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

func TestGet_Name(t *testing.T) {
	cmd := new(Get)
	assert.Equal(t, "GET", cmd.Name())
}

func TestGet_Help(t *testing.T) {
	cmd := new(Get)
	expected := `Usage: GET key
Get the value by key.
If provided key does not exist NIL will be returned.`
	assert.Equal(t, expected, cmd.Help())
}

func TestGet_Execute(t *testing.T) {
	expired := storage.NewStringValue("expired_value")
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(
		map[storage.Key]*storage.Value{
			"key_string": storage.NewStringValue("string_value"),
			"key_list":   storage.NewListValue([]string{"list_value_1", "list_value_2"}),
			"expired":    expired,
		},
	)
	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"ok", []string{"key_string"}, StringResult{Value: "string_value"}},
		{"not_existing_key", []string{"not_existing_key"}, NilResult{}},
		{"expired_key", []string{"expired"}, NilResult{}},
		{"wrong_type_op", []string{"key_list"}, ErrResult{Value: ErrWrongTypeOp}},
		{"wrong_number_of_args/1", []string{"key1", "key2"}, ErrResult{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{}, ErrResult{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Get{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestGet_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewStorageMock(t)
	strg.GetMock.Return(nil, err)
	strg.RLockMock.Return()
	strg.RUnlockMock.Return()

	cmd := Get{strg: strg}
	res := cmd.Execute("key")

	assert.Equal(t, ErrResult{Value: err}, res)
}
