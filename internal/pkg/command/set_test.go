package command

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
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
		result Result
	}{
		{"ok", []string{"key", "value"}, OkResult{}},
		{"wrong_args_number/1", []string{}, ErrResult{Err: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key", "value1", "value2"}, ErrResult{Err: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Set)
			assert.Equal(t, tt.result, cmd.Execute(strg, tt.args...))
		})
	}
}

func TestSet_Execute_Setter(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"string": storage.NewStringValue("value"),
	})

	cmd := new(Set)
	_ = cmd.Execute(strg, "string", "new_value")

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

	strg := storage.NewStorageMock(t)
	strg.PutMock.Return(err)

	cmd := new(Set)
	res := cmd.Execute(strg, "key", "value")

	assert.Equal(t, ErrResult{Err: err}, res)
}
