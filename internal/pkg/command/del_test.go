package command

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
	"github.com/namreg/godown-v2/test"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDel_Name(t *testing.T) {
	cmd := new(Del)
	assert.Equal(t, "DEL", cmd.Name())
}

func TestDel_Help(t *testing.T) {
	cmd := new(Del)
	expexted := `Usage: DEL key
Del the given key.`

	assert.Equal(t, expexted, cmd.Help())
}

func TestDel_ValidateArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  error
	}{
		{"valid_number_of_agrs", []string{"1"}, nil},
		{"not_valid_number_of_agrs/1", []string{}, ErrWrongArgsNumber},
		{"not_valid_number_of_agrs/2", []string{"1", "2"}, ErrWrongArgsNumber},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Del)
			err := cmd.ValidateArgs(tt.args...)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestDel_Execute(t *testing.T) {
	strg := memory.New(
		map[storage.Key]*storage.Value{
			"key": storage.NewStringValue("value"),
		},
	)
	tests := []struct {
		name   string
		args   []string
		result Result
	}{
		{"ok", []string{"key"}, OkResult{}},
		{"del_not_existing_key", []string{"not_existing_key"}, OkResult{}},
		{"wrong_number_of_args", []string{}, ErrResult{ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Del)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.result, res)
		})
	}
}

func TestDel_Execute_StorageErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("error")
	strg := test.NewMockStorage(ctrl)
	strg.EXPECT().Del(storage.Key("key")).DoAndReturn(func(_ storage.Key) error {
		return err
	})

	cmd := new(Del)
	res := cmd.Execute(strg, "key")

	assert.Equal(t, ErrResult{err}, res)
}
