package command

import (
	"errors"
	"sort"
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
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
	expired := storage.NewStringValue("value")
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":  storage.NewStringValue("value"),
		"string2": storage.NewStringValue("value2"),
		"map":     storage.NewMapValue(map[string]string{"field1": "val1", "field2": "val2"}),
		"expired": expired,
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"all_keys", []string{"*"}, SliceResult{[]string{"string", "string2", "map"}}},
		{"partial_match", []string{"str*"}, SliceResult{[]string{"string", "string2"}}},
		{"invalid_pattern", []string{"str++"}, ErrResult{errors.New("invalid pattern syntax")}},
		{"wrong_args_number/1", []string{}, ErrResult{ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"*", "*"}, ErrResult{ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Keys)
			res := cmd.Execute(strg, tt.args...)
			if sr, ok := res.(SliceResult); ok {
				expected := tt.want.(SliceResult).val
				sort.Strings(expected)

				actual := sr.val
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

	strg := NewStorageMock(t)
	strg.KeysMock.Return(nil, err)

	cmd := new(Keys)
	res := cmd.Execute(strg, "*")

	assert.Equal(t, ErrResult{err}, res)
}
