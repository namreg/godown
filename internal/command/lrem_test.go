package command

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/pkg/errors"

	"github.com/namreg/godown/internal/storage"
	"github.com/namreg/godown/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestLrem_Name(t *testing.T) {
	cmd := new(Lrem)
	assert.Equal(t, "LREM", cmd.Name())
}

func TestLrem_Help(t *testing.T) {
	cmd := new(Lrem)
	expected := `Usage: LREM key value
Removes all occurrences of elements equal to value from the list stored at key.`
	assert.Equal(t, expected, cmd.Help())
}

func TestLrem_Execute(t *testing.T) {
	expired := storage.NewList([]string{"val"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":  storage.NewString("value"),
		"list":    storage.NewList([]string{"val1", "val2"}),
		"expired": expired,
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"expired_key", []string{"expired", "val"}, OkReply{}},
		{"not_existing_key", []string{"not_existing_key", "val"}, OkReply{}},
		{"wrong_type_op", []string{"string", "val"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key", "val", "val"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Lrem{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLrem_Execute_WhiteBox(t *testing.T) {
	expired := storage.NewList([]string{"val"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"list1": storage.NewList([]string{"val1", "val2", "val1"}),
		"list2": storage.NewList([]string{"val1", "val1"}),
	})
	tests := []struct {
		name   string
		args   []string
		verify func(t *testing.T, items map[storage.Key]*storage.Value)
	}{
		{
			"remove_from_existing_key",
			[]string{"list1", "val1"},
			func(t *testing.T, items map[storage.Key]*storage.Value) {
				val, ok := items["list1"]
				assert.True(t, ok)

				expected := []string{"val2"}
				actual := val.Data().([]string)

				assert.Equal(t, expected, actual)
			},
		},
		{
			"empty_list_should_be_deleted",
			[]string{"list2", "val1"},
			func(t *testing.T, items map[storage.Key]*storage.Value) {
				_, ok := items["list2"]
				assert.False(t, ok)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Lrem{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, OkReply{}, res)

			items, err := strg.All()
			assert.NoError(t, err)

			tt.verify(t, items)
		})
	}
}

func TestLrem_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.PutMock.Return(err)

	cmd := Lrem{strg: strg}
	res := cmd.Execute("key", "val")

	assert.Equal(t, ErrReply{Value: err}, res)
}
