package command

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/pkg/errors"

	"github.com/namreg/godown-v2/internal/storage"
	"github.com/namreg/godown-v2/internal/storage/memory"
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
			"key": storage.NewString("value"),
		},
	)
	tests := []struct {
		name   string
		args   []string
		result Reply
	}{
		{"ok", []string{"key"}, OkReply{}},
		{"not_existing_key", []string{"not_existing_key"}, OkReply{}},
		{"wrong_number_of_args/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"key1", "key2"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &Del{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.result, res)
		})
	}
}

func TestDel_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.DelMock.Return(err)

	cmd := Del{strg: strg}
	res := cmd.Execute("key")

	assert.Equal(t, ErrReply{Value: err}, res)
}
