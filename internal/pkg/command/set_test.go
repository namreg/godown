package command

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/namreg/godown-v2/test"
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

func TestSet_ValidateArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  error
	}{
		{"valid_number_of_agrs", []string{"1", "2"}, nil},
		{"not_valid_number_of_agrs/1", []string{}, ErrWrongArgsNumber},
		{"not_valid_number_of_agrs/2", []string{"1"}, ErrWrongArgsNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Set)
			err := cmd.ValidateArgs(tt.args...)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestSet_Execute(t *testing.T) {
	strg := memory.New(nil)
	tests := []struct {
		name   string
		args   []string
		result Result
	}{
		{"ok", []string{"key", "value"}, OkResult{}},
		{"wrong_args_number", []string{}, ErrResult{ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Set)
			assert.Equal(t, tt.result, cmd.Execute(strg, tt.args...))
		})
	}
}

func TestSet_Execute_StorageErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	strg := test.NewMockStorage(ctrl)
	err := errors.New("error")
	strg.EXPECT().Put(storage.Key("key"), gomock.Any()).DoAndReturn(func(_ storage.Key, _ storage.ValueSetter) error {
		return err
	})

	cmd := new(Set)

	assert.Equal(t, ErrResult{err}, cmd.Execute(strg, "key", "value"))
}
