package command

import (
	"errors"
	"sort"
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/storage"
	"github.com/namreg/godown-v2/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestKeys_Name(t *testing.T) {
	cmd := new(Keys)
	assert.Equal(t, "KEYS", cmd.Name())
}

func TestKeys_Help(t *testing.T) {
	cmd := new(Keys)
	expected := `Usage: KEYS pattern
Find all keys matching the given pattern.`
	assert.Equal(t, expected, cmd.Help())
}

func TestKeys_Execute(t *testing.T) {
	expired := storage.NewString("value")
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":  storage.NewString("value"),
		"string2": storage.NewString("value2"),
		"map":     storage.NewMap(map[string]string{"field1": "val1", "field2": "val2"}),
		"expired": expired,
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"all_keys", []string{"*"}, SliceReply{Value: []string{"string", "string2", "map"}}},
		{"partial_match", []string{"str*"}, SliceReply{Value: []string{"string", "string2"}}},
		{"invalid_pattern", []string{"str++"}, ErrReply{Value: errors.New("invalid pattern syntax")}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"*", "*"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Keys{strg: strg}
			res := cmd.Execute(tt.args...)
			if sr, ok := res.(SliceReply); ok {
				expected := tt.want.(SliceReply).Value
				sort.Strings(expected)

				actual := sr.Value
				sort.Strings(actual)

				assert.Equal(t, tt.want, res)
			} else {
				assert.Equal(t, tt.want, res)
			}
		})
	}
}

func TestKeys_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.KeysMock.Return(nil, err)

	cmd := Keys{strg: strg}
	res := cmd.Execute("*")

	assert.Equal(t, ErrReply{Value: err}, res)
}
