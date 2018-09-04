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

func TestLindex_Name(t *testing.T) {
	cmd := new(Lindex)
	assert.Equal(t, "LINDEX", cmd.Name())
}

func TestLindex_Help(t *testing.T) {
	cmd := new(Lindex)
	expected := `LINDEX key index
Returns the element at index index in the list stored at key. 
The index is zero-based, so 0 means the first element, 1 the second element and so on. 
Negative indices can be used to designate elements starting at the tail of the list.`
	assert.Equal(t, expected, cmd.Help())
}

func TestLindex_Execute(t *testing.T) {
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
		{"ok", []string{"list", "0"}, StringReply{Value: "val1"}},
		{"negative_index/1", []string{"list", "-1"}, StringReply{Value: "val2"}},
		{"negative_index/2", []string{"list", "-2"}, StringReply{Value: "val1"}},
		{"expired_key", []string{"expired", "0"}, NilReply{}},
		{"not_existing_key", []string{"not_existing_key", "0"}, NilReply{}},
		{"not_existing_index/1", []string{"list", "2"}, NilReply{}},
		{"not_existing_index/2", []string{"list", "-3"}, NilReply{}},
		{"wrong_type_op", []string{"string", "0"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"list", "0", "1"}, ErrReply{Value: ErrWrongArgsNumber}},
		{"index_not_integer", []string{"list", "string"}, ErrReply{Value: errors.New("index should be an integer")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Lindex{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLindex_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.GetMock.Return(nil, err)

	cmd := Lindex{strg: strg}
	res := cmd.Execute("key", "0")

	assert.Equal(t, ErrReply{Value: err}, res)
}
