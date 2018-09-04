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
	expired := storage.NewMap(map[string]string{"field": "value"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string_key":  storage.NewString("string"),
		"expired_key": expired,
		"key":         storage.NewMap(map[string]string{"field": "value"}),
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"ok", []string{"key", "field"}, StringReply{Value: "value"}},
		{"not_existing_key", []string{"not_existing_key", "field"}, NilReply{}},
		{"not_existing_field", []string{"key", "not_existing_field"}, NilReply{}},
		{"expired_key", []string{"expired_key", "field"}, NilReply{}},
		{"wront_type_op", []string{"string_key", "field"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_number_of_args/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"key", "arg1", "arg2"}, ErrReply{Value: ErrWrongArgsNumber}},
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

	strg := NewdataStoreMock(mc)
	strg.GetMock.Return(nil, err)

	cmd := Hget{strg: strg}
	res := cmd.Execute("key", "field")

	assert.Equal(t, ErrReply{Value: err}, res)
}
