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

func TestLpop_Name(t *testing.T) {
	cmd := new(Lpop)
	assert.Equal(t, "LPOP", cmd.Name())
}

func TestLpop_Help(t *testing.T) {
	cmd := new(Lpop)
	expected := `Usage: LPOP key
Removes and returns the first element of the list stored at key.`
	assert.Equal(t, expected, cmd.Help())
}

func TestLpop_Execute(t *testing.T) {
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
		{"ok", []string{"list"}, StringReply{Value: "val1"}},
		{"expired_key", []string{"expired"}, NilReply{}},
		{"not_existing_key", []string{"not_existing_key"}, NilReply{}},
		{"wrong_type_op", []string{"string"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"list", "0"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Lpop{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLpop_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.PutMock.Return(err)

	cmd := Lpop{strg: strg}

	res := cmd.Execute("list")
	assert.Equal(t, ErrReply{Value: err}, res)
}

func TestLpop_Execute_DelEmptyList(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"list": storage.NewList([]string{"val1"}),
	})

	cmd := Lpop{strg: strg}
	_ = cmd.Execute("list")

	items, err := strg.All()
	assert.NoError(t, err)

	value, ok := items["list"]
	assert.Nil(t, value)
	assert.False(t, ok)
}
