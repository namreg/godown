package command

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/pkg/errors"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
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
		"string": storage.NewStringValue("value"),
		"list":   storage.NewListValue("val1"),
		"map":    storage.NewMapValue(map[string]string{"field": "values"}),
		"bitmap": storage.NewBitMapValue([]uint64{1}),
	})

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"string", []string{"string"}, StringResult{"string"}},
		{"list", []string{"list"}, StringResult{"list"}},
		{"map", []string{"map"}, StringResult{"map"}},
		{"bitmap", []string{"bitmap"}, StringResult{"bitmap"}},
		{"not_existing_key", []string{"not_existing_key"}, NilResult{}},
		{"wrong_number_of_args/1", []string{}, ErrResult{ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"key", "arg1"}, ErrResult{ErrWrongArgsNumber}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Type)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestType_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := storage.NewStorageMock(t)
	strg.GetMock.Return(nil, err)

	cmd := new(Type)
	res := cmd.Execute(strg, "key")

	assert.Equal(t, ErrResult{err}, res)
}
