package command

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/storage"
	"github.com/namreg/godown-v2/internal/storage/memory"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestLlen_Name(t *testing.T) {
	cmd := new(Llen)
	assert.Equal(t, "LLEN", cmd.Name())
}

func TestLlen_Help(t *testing.T) {
	cmd := new(Llen)
	expected := `Usage: LLEN key
Returns the length of the list stored at key. 
If key does not exist, it is interpreted as an empty list and 0 is returned.`
	assert.Equal(t, expected, cmd.Help())
}

func TestLlen_Execute(t *testing.T) {
	expired := storage.NewList([]string{"val1", "val2"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":  storage.NewString("string"),
		"list":    storage.NewList([]string{"val1", "val2"}),
		"expired": expired,
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"ok", []string{"list"}, IntReply{Value: 2}},
		{"expired_key", []string{"expired"}, IntReply{Value: 0}},
		{"not_existing_key", []string{"not_existing_key"}, IntReply{Value: 0}},
		{"wrong_type_op", []string{"string"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"list", "0", "1"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Llen{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLlen_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(t)
	strg.GetMock.Return(nil, err)

	cmd := Llen{strg: strg}
	res := cmd.Execute("key")

	assert.Equal(t, ErrReply{Value: err}, res)
}
