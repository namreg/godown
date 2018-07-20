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
	expired := storage.NewListValue("val1", "val2")
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":  storage.NewStringValue("string"),
		"list":    storage.NewListValue("val1", "val2"),
		"expired": expired,
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"ok", []string{"list", "0"}, StringResult{"val1"}},
		{"negative_index/1", []string{"list", "-1"}, StringResult{"val2"}},
		{"negative_index/2", []string{"list", "-2"}, StringResult{"val1"}},
		{"expired_key", []string{"expired", "0"}, NilResult{}},
		{"not_existing_key", []string{"not_existing_key", "0"}, NilResult{}},
		{"not_existing_index/1", []string{"list", "2"}, NilResult{}},
		{"not_existing_index/2", []string{"list", "-3"}, NilResult{}},
		{"wrong_type_op", []string{"string", "0"}, ErrResult{ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrResult{ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"list", "0", "1"}, ErrResult{ErrWrongArgsNumber}},
		{"index_not_integer", []string{"list", "string"}, ErrResult{errors.New("index should be an integer")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Lindex)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLindex_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewStorageMock(t)
	strg.GetMock.Return(nil, err)

	cmd := new(Lindex)
	res := cmd.Execute(strg, "key", "0")

	assert.Equal(t, ErrResult{err}, res)
}
