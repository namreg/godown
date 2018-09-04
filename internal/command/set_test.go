package command

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/storage"
	"github.com/namreg/godown/internal/storage/memory"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSet_Name(t *testing.T) {
	cmd := new(Set)
	assert.Equal(t, "SET", cmd.Name())
}

func TestSet_Help(t *testing.T) {
	cmd := new(Set)
	expected := `Usage: SET key value
Set key to hold the string value.
If key already holds a value, it is overwritten.`
	assert.Equal(t, expected, cmd.Help())
}

func TestSet_Execute(t *testing.T) {
	strg := memory.New(nil)
	tests := []struct {
		name   string
		args   []string
		result Reply
	}{
		{"ok", []string{"key", "value"}, OkReply{}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key", "value1", "value2"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Set{strg: strg}
			assert.Equal(t, tt.result, cmd.Execute(tt.args...))
		})
	}
}

func TestSet_Execute_WhiteBox(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"string": storage.NewString("value"),
	})

	cmd := Set{strg: strg}
	_ = cmd.Execute("string", "new_value")

	items, err := strg.All()
	assert.NoError(t, err)

	val, ok := items["string"]
	assert.True(t, ok)
	assert.Equal(t, "new_value", val.Data())
}

func TestSet_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.PutMock.Return(err)

	cmd := Set{strg: strg}
	res := cmd.Execute("key", "value")

	assert.Equal(t, ErrReply{Value: err}, res)
}
