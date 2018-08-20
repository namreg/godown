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

func TestHvals_Name(t *testing.T) {
	cmd := new(Hvals)
	assert.Equal(t, "HVALS", cmd.Name())
}

func TestHvals_Help(t *testing.T) {
	cmd := new(Hvals)
	expected := `Usage: HVALS key
Returns all values in the hash stored at key`
	assert.Equal(t, expected, cmd.Help())
}

func TestHvals_Execute(t *testing.T) {
	expired := storage.NewMap(map[string]string{"field1": "val1", "field2": "val2"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":  storage.NewString("value"),
		"map":     storage.NewMap(map[string]string{"field1": "val1", "field2": "val2"}),
		"expired": expired,
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"existing_key", []string{"map"}, SliceReply{Value: []string{"val1", "val2"}}},
		{"expired_key", []string{"expired"}, NilReply{}},
		{"not_existing_key", []string{"not_existing_key"}, NilReply{}},
		{"wrong_type_op", []string{"string"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key", "field"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Hvals{strg: strg}
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

func TestHvals_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.GetMock.Return(nil, err)

	cmd := Hvals{strg: strg}
	res := cmd.Execute("key")

	assert.Equal(t, ErrReply{Value: err}, res)
}
