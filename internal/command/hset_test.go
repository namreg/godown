package command

import (
	"errors"
	"testing"

	"github.com/gojuno/minimock"

	"github.com/namreg/godown/internal/storage"
	"github.com/namreg/godown/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestHset_Name(t *testing.T) {
	cmd := new(Hset)
	assert.Equal(t, "HSET", cmd.Name())
}

func TestHset_Help(t *testing.T) {
	cmd := new(Hset)
	expected := `Usage: HSET key field value
Sets field in the hash stored at key to value.`
	assert.Equal(t, expected, cmd.Help())
}

func TestHset_Execute(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"string": storage.NewString("value"),
	})

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"ok", []string{"key", "field", "value"}, OkReply{}},
		{"wrong_type_op", []string{"string", "field", "value"}, ErrReply{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key", "field"}, ErrReply{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Hset{strg: strg}
			assert.Equal(t, tt.want, cmd.Execute(tt.args...))
		})
	}
}

func TestHset_Execute_WhiteBox(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"map": storage.NewMap(map[string]string{"field": "value"}),
	})

	tests := []struct {
		name   string
		args   []string
		verify func(t *testing.T, items map[storage.Key]*storage.Value)
	}{
		{
			"add_new_field_to_existing_key",
			[]string{"map", "field2", "value2"},
			func(t *testing.T, items map[storage.Key]*storage.Value) {
				val, ok := items["map"]
				assert.True(t, ok)

				m := val.Data().(map[string]string)

				fval, ok := m["field"]
				assert.True(t, ok)
				assert.Equal(t, "value", fval)

				fval, ok = m["field2"]
				assert.True(t, ok)
				assert.Equal(t, "value2", fval)
			},
		},

		{
			"replace_field_in_existing_key",
			[]string{"map", "field", "value2"},
			func(t *testing.T, items map[storage.Key]*storage.Value) {
				val, ok := items["map"]
				assert.True(t, ok)

				m := val.Data().(map[string]string)

				fval, ok := m["field"]
				assert.True(t, ok)
				assert.Equal(t, "value2", fval)
			},
		},
		{
			"add_new_field_to_not_existing_key",
			[]string{"map2", "field", "value"},
			func(t *testing.T, items map[storage.Key]*storage.Value) {
				val, ok := items["map2"]
				assert.True(t, ok)

				m := val.Data().(map[string]string)

				fval, ok := m["field"]
				assert.True(t, ok)
				assert.Equal(t, "value", fval)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Hset{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, OkReply{}, res)

			items, err := strg.All()
			assert.NoError(t, err)

			tt.verify(t, items)
		})
	}
}

func TestHset_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.PutMock.Return(err)

	cmd := Hset{strg: strg}
	res := cmd.Execute("key", "field", "value")

	assert.Equal(t, ErrReply{Value: err}, res)
}
