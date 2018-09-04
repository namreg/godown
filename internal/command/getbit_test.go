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

func TestGetBit_Name(t *testing.T) {
	cmd := new(GetBit)
	assert.Equal(t, "GETBIT", cmd.Name())
}

func TestGetBit_Help(t *testing.T) {
	cmd := new(GetBit)
	expected := `Usage: GETBIT key offset
Returns the bit value at offset in the string value stored at key.`
	assert.Equal(t, expected, cmd.Help())
}

func TestGetBit_Execute(t *testing.T) {
	expired := storage.NewBitMap([]uint64{1 << 10})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":                 storage.NewString("string"),
		"bitmap":                 storage.NewBitMap([]uint64{1 << 5}),
		"bitmap_with_big_offset": storage.NewBitMap([]uint64{0, 3}),
		"expired_bitmap":         expired,
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"set_bit", []string{"bitmap", "5"}, IntReply{Value: 1}},
		{"unset_bit", []string{"bitmap", "10"}, IntReply{Value: 0}},
		{"big_offset/1", []string{"bitmap_with_big_offset", "64"}, IntReply{Value: 1}},
		{"big_offset/2", []string{"bitmap_with_big_offset", "65"}, IntReply{Value: 1}},
		{"big_offset/3", []string{"bitmap_with_big_offset", "1000"}, IntReply{Value: 0}},
		{"key_not_exists", []string{"key_not_exists", "0"}, IntReply{Value: 0}},
		{"key_not_exists", []string{"key_not_exists", "0"}, IntReply{Value: 0}},
		{"expired_key", []string{"expired_bitmap", "10"}, IntReply{Value: 0}},
		{"wrong_type_op", []string{"string", "1"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_number_of_args/1", []string{"key1"}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"negative_offset", []string{"bitmap", "-1"}, ErrReply{Value: errors.New("invalid offset")}},
		{"offset_not_integer", []string{"bitmap", "string"}, ErrReply{Value: errors.New("invalid offset")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetBit{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestGetBit_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.GetMock.Return(nil, err)

	cmd := GetBit{strg: strg}
	res := cmd.Execute("key", "10")

	assert.Equal(t, ErrReply{Value: err}, res)
}
