package command

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/pkg/errors"

	"github.com/namreg/godown-v2/internal/storage"
	"github.com/namreg/godown-v2/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestType_Name(t *testing.T) {
	cmd := new(Type)
	assert.Equal(t, "TYPE", cmd.Name())
}

func TestType_Help(t *testing.T) {
	cmd := new(Type)
	expected := `Usage: TYPE key
Returns the type stored at key.`
	assert.Equal(t, expected, cmd.Help())
}

func TestType_Execute(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"string": storage.NewString("value"),
		"list":   storage.NewList([]string{"val1"}),
		"map":    storage.NewMap(map[string]string{"field": "values"}),
		"bitmap": storage.NewBitMap([]uint64{1}),
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"string", []string{"string"}, StringResult{Value: "string"}},
		{"list", []string{"list"}, StringResult{Value: "list"}},
		{"map", []string{"map"}, StringResult{Value: "map"}},
		{"bitmap", []string{"bitmap"}, StringResult{Value: "bitmap"}},
		{"not_existing_key", []string{"not_existing_key"}, NilResult{}},
		{"wrong_number_of_args/1", []string{}, ErrResult{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"key", "arg1"}, ErrResult{Value: ErrWrongArgsNumber}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Type{strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestType_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewcommandStorageMock(mc)
	strg.GetMock.Return(nil, err)

	cmd := Type{strg: strg}
	res := cmd.Execute("key")

	assert.Equal(t, ErrResult{Value: err}, res)
}
