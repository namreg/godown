package command

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/pkg/errors"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
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
	expired := storage.NewListValue([]string{"val"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":  storage.NewStringValue("value"),
		"list":    storage.NewListValue([]string{"val1", "val2"}),
		"expired": expired,
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"expired_key", []string{"expired", "val"}, OkResult{}},
		{"not_existing_key", []string{"not_existing_key", "val"}, OkResult{}},
		{"wrong_type_op", []string{"string", "val"}, ErrResult{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrResult{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key", "val", "val"}, ErrResult{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Lrem)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLrem_Execute_WhiteBox(t *testing.T) {
	expired := storage.NewListValue([]string{"val"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"list1": storage.NewListValue([]string{"val1", "val2", "val1"}),
		"list2": storage.NewListValue([]string{"val1", "val1"}),
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
			cmd := new(Lrem)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, OkResult{}, res)

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

	cmd := new(Lrem)

	//storage.Get err
	strg1 := storage.NewStorageMock(t)
	strg1.GetMock.Return(nil, err)
	strg1.LockMock.Return()
	strg1.UnlockMock.Return()

	res1 := cmd.Execute(strg1, "key", "val")
	assert.Equal(t, ErrResult{Value: err}, res1)

	//storage.Put err
	strg2 := storage.NewStorageMock(t)
	strg2.GetMock.Return(storage.NewListValue([]string{"val", "val2"}), nil)
	strg2.PutMock.Return(err)
	strg2.LockMock.Return()
	strg2.UnlockMock.Return()

	res2 := cmd.Execute(strg2, "key", "val")
	assert.Equal(t, ErrResult{Value: err}, res2)

	//storage.Del err
	strg3 := storage.NewStorageMock(t)
	strg3.GetMock.Return(storage.NewListValue([]string{"val"}), nil)
	strg3.DelMock.Return(err)
	strg3.LockMock.Return()
	strg3.UnlockMock.Return()

	res3 := cmd.Execute(strg3, "key", "val")
	assert.Equal(t, ErrResult{Value: err}, res3)
}
