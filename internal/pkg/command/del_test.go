package command

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/pkg/errors"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
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
		{"not_existing_key", []string{"not_existing_key"}, OkResult{}},
		{"wrong_number_of_args/1", []string{}, ErrResult{ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"key1", "key2"}, ErrResult{ErrWrongArgsNumber}},
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
	mc := minimock.NewController(t)
	defer mc.Finish()

	strg := storage.NewStorageMock(t)

	err := errors.New("error")

	strg.DelMock.Return(err)

	cmd := new(Del)
	res := cmd.Execute(strg, "key")

	assert.Equal(t, ErrResult{err}, res)
}
