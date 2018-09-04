package command

import (
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/storage"
	"github.com/namreg/godown/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestSetBit_Name(t *testing.T) {
	cmd := new(SetBit)
	assert.Equal(t, "SETBIT", cmd.Name())
}

func TestSetBit_Help(t *testing.T) {
	cmd := new(SetBit)
	expected := `Usage: SETBIT key offset value
Sets or clears the bit at offset in the string value stored at key.`
	assert.Equal(t, expected, cmd.Help())
}

func TestSetBit_Execute(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"string": storage.NewString("value"),
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"ok/1", []string{"key", "1", "1"}, OkReply{}},
		{"ok/2", []string{"key", "0", "0"}, OkReply{}},
		{"big_offset", []string{"key", "100", "1"}, OkReply{}},
		{"negative_offset", []string{"key", "-1", "1"}, ErrReply{Value: errors.New("invalid offset")}},
		{"invalid_value/1", []string{"key", "1", "-1"}, ErrReply{Value: errors.New("value should be 0 or 1")}},
		{"invalid_value/2", []string{"key", "1", "2"}, ErrReply{Value: errors.New("value should be 0 or 1")}},
		{"wrong_type_op", []string{"string", "1", "1"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key", "field"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := SetBit{strg: strg}
			assert.Equal(t, tt.want, cmd.Execute(tt.args...))
		})
	}
}

func TestSetBit_Execute_WhiteBox(t *testing.T) {
	items := map[storage.Key]*storage.Value{
		"bitmap2":                storage.NewBitMap([]uint64{2}),
		"bitmap3":                storage.NewBitMap([]uint64{2}),
		"bitmap_with_big_offset": storage.NewBitMap([]uint64{0, 1}),
	}
	strg := memory.New(items)

	tests := []struct {
		name   string
		args   []string
		verify func(t *testing.T)
	}{
		{
			"set_bit_in_not_existing_key",
			[]string{"bitmap", "1", "1"},
			func(t *testing.T) {
				val, ok := items["bitmap"]
				assert.True(t, ok)

				v, ok := val.Data().([]uint64)
				assert.True(t, ok)
				assert.Equal(t, uint64(2), v[0])
			},
		},
		{
			"set_bit_in_existing_key",
			[]string{"bitmap2", "2", "1"},
			func(t *testing.T) {
				val, ok := items["bitmap2"]
				assert.True(t, ok)

				v, ok := val.Data().([]uint64)
				assert.True(t, ok)
				assert.Equal(t, uint64(6), v[0])
			},
		},
		{
			"delete_key_when_all_bits_not_set",
			[]string{"bitmap3", "1", "0"},
			func(t *testing.T) {
				_, ok := items["bitmap3"]
				assert.False(t, ok)
			},
		},
		{
			"delete_key_when_all_bits_not_set/big_offset",
			[]string{"bitmap_with_big_offset", "64", "0"},
			func(t *testing.T) {
				_, ok := items["bitmap_with_big_offset"]
				assert.False(t, ok)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := SetBit{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, OkReply{}, res)

			tt.verify(t)
		})
	}
}

func TestSetBit_growSlice(t *testing.T) {
	type args struct {
		sl     []uint64
		offset uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{"offset < 64", args{sl: []uint64{1}, offset: 60}, []uint64{1}},
		{"offset > 63", args{sl: []uint64{1}, offset: 65}, []uint64{1, 0}},
		{"offset == 63", args{sl: []uint64{1}, offset: 63}, []uint64{1}},
		{"offset == 64", args{sl: []uint64{1}, offset: 64}, []uint64{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SetBit{}
			got := c.growSlice(tt.args.sl, tt.args.offset)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSetBit_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.PutMock.Return(err)

	cmd := SetBit{strg: strg}
	res := cmd.Execute("key", "1", "1")

	assert.Equal(t, ErrReply{Value: err}, res)
}
