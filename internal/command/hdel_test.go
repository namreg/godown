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

func TestHdel_Name(t *testing.T) {
	cmd := new(Hdel)
	assert.Equal(t, "HDEL", cmd.Name())
}

func TestHdel_Help(t *testing.T) {
	cmd := new(Hdel)
	expected := `Usage: HDEL key field [field ...]
Removes the specified fields from the hash stored at key.
Returns the number of fields that were removed.`
	assert.Equal(t, expected, cmd.Help())
}

func TestHdel_Execute(t *testing.T) {
	expired := storage.NewMap(map[string]string{"field": "value"})
	expired.SetTTL(time.Now().Add(-1 * time.Second))

	strg := memory.New(map[storage.Key]*storage.Value{
		"string":  storage.NewString("value"),
		"map":     storage.NewMap(map[string]string{"field1": "value1", "field2": "value2", "field3": "value3"}),
		"expired": expired,
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"ok", []string{"map", "field1", "field2", "bar"}, IntReply{Value: 2}},
		{"not_existing_key", []string{"not_existing_key", "field1"}, IntReply{Value: 0}},
		{"expired_key", []string{"expired", "field"}, IntReply{Value: 0}},
		{"wrong_type_op", []string{"string", "field"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Hdel{strg: strg}
			assert.Equal(t, tt.want, cmd.Execute(tt.args...))
		})
	}
}

func TestHdel_Execute_DeleteKeyWhenAllFieldsDeleted(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"map": storage.NewMap(map[string]string{"field": "value"}),
	})

	cmd := Hdel{strg: strg}
	res := cmd.Execute("map", "field")
	assert.Equal(t, IntReply{Value: 1}, res)

	items, err := strg.All()
	assert.NoError(t, err)

	_, ok := items["map"]
	assert.False(t, ok)
}

func TestHdel_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.PutMock.Return(err)

	cmd := Hdel{strg: strg}
	res := cmd.Execute("key", "field")

	assert.Equal(t, ErrReply{Value: err}, res)
}
