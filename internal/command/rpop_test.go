package command

import (
	"errors"
	"testing"
	"time"

	"github.com/gojuno/minimock"

	"github.com/namreg/godown/internal/storage"
	"github.com/namreg/godown/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestRpop_Name(t *testing.T) {
	cmd := new(Rpop)
	assert.Equal(t, "RPOP", cmd.Name())
}

func TestRpop_Help(t *testing.T) {
	cmd := new(Rpop)
	expected := `Usage: RPOP key
Removes and returns the last element of the list stored at key.`
	assert.Equal(t, expected, cmd.Help())
}

func TestRpop_Execute(t *testing.T) {
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
		{"ok", []string{"list"}, StringReply{Value: "val2"}},
		{"expired_key", []string{"expired"}, NilReply{}},
		{"not_existing_key", []string{"not_existing_key"}, NilReply{}},
		{"wrong_type_op", []string{"string"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"list", "0"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Rpop{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestRpop_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.PutMock.Return(err)

	cmd := Rpop{strg: strg}

	res := cmd.Execute("list")
	assert.Equal(t, ErrReply{Value: err}, res)
}

func TestRpop_Execute_DelEmptyList(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"list": storage.NewList([]string{"val1"}),
	})

	cmd := Rpop{strg: strg}
	_ = cmd.Execute("list")

	items, err := strg.All()
	assert.NoError(t, err)

	value, ok := items["list"]
	assert.Nil(t, value)
	assert.False(t, ok)
}
