package command

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
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
		{"ok", []string{"list"}, IntResult{Value: 2}},
		{"expired_key", []string{"expired"}, IntResult{Value: 0}},
		{"not_existing_key", []string{"not_existing_key"}, IntResult{Value: 0}},
		{"wrong_type_op", []string{"string"}, ErrResult{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrResult{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"list", "0", "1"}, ErrResult{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Llen)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestLlen_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := storage.NewStorageMock(t)
	strg.GetMock.Return(nil, err)

	cmd := new(Llen)
	res := cmd.Execute(strg, "key")

	assert.Equal(t, ErrResult{Value: err}, res)
}
