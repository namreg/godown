package command

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/storage"
	"github.com/namreg/godown/internal/storage/memory"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGet_Name(t *testing.T) {
	cmd := new(Get)
	assert.Equal(t, "GET", cmd.Name())
}

func TestGet_Help(t *testing.T) {
	cmd := new(Get)
	expected := `Usage: GET key
Get the value by key.
If provided key does not exist NIL will be returned.`
	assert.Equal(t, expected, cmd.Help())
}

func TestGet_Execute(t *testing.T) {
	expired := storage.NewString("expired_value")
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(
		map[storage.Key]*storage.Value{
			"key_string": storage.NewString("string_value"),
			"key_list":   storage.NewList([]string{"list_value_1", "list_value_2"}),
			"expired":    expired,
		},
	)
	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"ok", []string{"key_string"}, StringReply{Value: "string_value"}},
		{"not_existing_key", []string{"not_existing_key"}, NilReply{}},
		{"expired_key", []string{"expired"}, NilReply{}},
		{"wrong_type_op", []string{"key_list"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_number_of_args/1", []string{"key1", "key2"}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Get{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestGet_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.GetMock.Return(nil, err)

	cmd := Get{strg: strg}
	res := cmd.Execute("key")

	assert.Equal(t, ErrReply{Value: err}, res)
}
