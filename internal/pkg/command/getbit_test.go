package command

import (
	"errors"
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
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
	expired := storage.NewBitMapValue(1 << 10)
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":         storage.NewStringValue("string"),
		"bitmap":         storage.NewBitMapValue(1 << 5),
		"expired_bitmap": expired,
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"set_bit", []string{"bitmap", "5"}, IntResult{1}},
		{"not_set_bit", []string{"bitmap", "10"}, IntResult{0}},
		{"big_offset", []string{"bitmap", "100"}, ErrResult{errors.New("invalid offset")}},
		{"key_not_exists", []string{"key_not_exists", "0"}, IntResult{0}},
		{"key_not_exists", []string{"key_not_exists", "0"}, IntResult{0}},
		{"expired_key", []string{"expired_bitmap", "10"}, IntResult{0}},
		{"wrong_type_op", []string{"string", "1"}, ErrResult{ErrWrongTypeOp}},
		{"wrong_number_of_args/1", []string{"key1"}, ErrResult{ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{}, ErrResult{ErrWrongArgsNumber}},
		{"negative_offset", []string{"bitmap", "-1"}, ErrResult{errors.New("invalid offset")}},
		{"offset_not_integer", []string{"bitmap", "string"}, ErrResult{errors.New("invalid offset")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(GetBit)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestGetBit_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	strg := NewStorageMock(t)

	err := errors.New("error")

	strg.GetMock.Return(nil, err)

	cmd := new(GetBit)
	res := cmd.Execute(strg, "key", "10")

	assert.Equal(t, ErrResult{err}, res)
}
