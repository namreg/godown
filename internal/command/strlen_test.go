package command

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/pkg/errors"

	"github.com/namreg/godown-v2/internal/storage"
	"github.com/namreg/godown-v2/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestStrlen_Name(t *testing.T) {
	cmd := new(Strlen)
	assert.Equal(t, "STRLEN", cmd.Name())
}

func TestStrlen_Help(t *testing.T) {
	cmd := new(Strlen)
	expected := `Usage: STRLEN key
Returns length of the given key.
If key does not exists, 0 will be returned.`
	assert.Equal(t, expected, cmd.Help())
}

func TestStrlen_Execute(t *testing.T) {
	expired := storage.NewStringValue("val")
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":       storage.NewStringValue("value"),
		"empty_string": storage.NewStringValue(""),
		"expired":      expired,
		"list":         storage.NewListValue([]string{"val1", "val2"}),
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"not_empty_string", []string{"string"}, IntResult{Value: 5}},
		{"empty_string", []string{"empty_string"}, IntResult{Value: 0}},
		{"expired_key", []string{"expired"}, IntResult{Value: 0}},
		{"not_existing_key", []string{"not_existing_key"}, IntResult{Value: 0}},
		{"wrong_type_op", []string{"list"}, ErrResult{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrResult{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"string", "list"}, ErrResult{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Strlen{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestStrlen_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewStorageMock(t)
	strg.GetMock.Return(nil, err)
	strg.RLockMock.Return()
	strg.RUnlockMock.Return()

	cmd := Strlen{strg: strg}
	res := cmd.Execute("key")

	assert.Equal(t, ErrResult{Value: err}, res)
}
