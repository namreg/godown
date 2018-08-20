package command

import (
	"errors"
	"testing"
	"time"

	"github.com/gojuno/minimock"

	"github.com/namreg/godown-v2/internal/storage"
	"github.com/namreg/godown-v2/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestHget_Name(t *testing.T) {
	cmd := new(Hget)
	assert.Equal(t, "HGET", cmd.Name())
}

func TestHget_Help(t *testing.T) {
	cmd := new(Hget)
	expected := `Usage: HGET key field
Returns the value associated with field in the hash stored at key.`
	assert.Equal(t, expected, cmd.Help())
}

func TestHget_Execute(t *testing.T) {
	expired := storage.NewMapValue(map[string]string{"field": "value"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string_key":  storage.NewStringValue("string"),
		"expired_key": expired,
		"key":         storage.NewMapValue(map[string]string{"field": "value"}),
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"ok", []string{"key", "field"}, StringResult{Value: "value"}},
		{"not_existing_key", []string{"not_existing_key", "field"}, NilResult{}},
		{"not_existing_field", []string{"key", "not_existing_field"}, NilResult{}},
		{"expired_key", []string{"expired_key", "field"}, NilResult{}},
		{"wront_type_op", []string{"string_key", "field"}, ErrResult{Value: ErrWrongTypeOp}},
		{"wrong_number_of_args/1", []string{}, ErrResult{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"key", "arg1", "arg2"}, ErrResult{Value: ErrWrongArgsNumber}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Hget{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestHget_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewStorageMock(t)
	strg.GetMock.Return(nil, err)
	strg.RLockMock.Return()
	strg.RUnlockMock.Return()

	cmd := Hget{strg: strg}
	res := cmd.Execute("key", "field")

	assert.Equal(t, ErrResult{Value: err}, res)
}
