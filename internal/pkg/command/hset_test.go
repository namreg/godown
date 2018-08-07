package command

import (
	"errors"
	"testing"

	"github.com/gojuno/minimock"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
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
		"string": storage.NewStringValue("value"),
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"ok", []string{"key", "field", "value"}, OkResult{}},
		{"wrong_type_op", []string{"string", "field", "value"}, ErrResult{Value: ErrWrongTypeOp}},
		{"wrong_args_number/1", []string{}, ErrResult{Value: ErrWrongArgsNumber}},
		{"wrong_args_number/2", []string{"key", "field"}, ErrResult{Value: ErrWrongArgsNumber}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Hset)
			assert.Equal(t, tt.want, cmd.Execute(strg, tt.args...))
		})
	}
}

func TestHset_Execute_Setter(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"map": storage.NewMapValue(map[string]string{"field": "value"}),
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
			cmd := new(Hset)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, OkResult{}, res)

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

	strg := storage.NewStorageMock(t)
	strg.PutMock.Return(err)

	cmd := new(Hset)
	res := cmd.Execute(strg, "key", "field", "value")

	assert.Equal(t, ErrResult{Value: err}, res)
}
