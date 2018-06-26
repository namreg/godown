package command

import (
	"errors"
	"testing"
	"time"

	"github.com/gojuno/minimock"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"

	"github.com/stretchr/testify/assert"
)

func TestExpire_Name(t *testing.T) {
	cmd := new(Expire)
	assert.Equal(t, "EXPIRE", cmd.Name())
}

func TestExpire_Help(t *testing.T) {
	cmd := new(Expire)
	expexted := `Usage: EXPIRE key seconds
Set a timeout on key. After the timeout has expired, the key will automatically be deleted.`

	assert.Equal(t, expexted, cmd.Help())
}

func TestExpire_Execute(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"key": storage.NewStringValue("value"),
	})
	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"wrong_arg_number/1", []string{}, ErrResult{ErrWrongArgsNumber}},
		{"wrong_arg_number/2", []string{"key", "10", "20"}, ErrResult{ErrWrongArgsNumber}},
		{"existing_key", []string{"key", "10"}, OkResult{}},
		{"not_existing_key", []string{"not_existing_key", "10"}, OkResult{}},
		{"secs_as_string", []string{"not_existing_key", "seconds"}, ErrResult{errors.New("seconds should be integer")}},
		{"secs_negative", []string{"not_existing_key", "-10"}, ErrResult{errors.New("seconds should be positive")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Expire)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestExpire_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	strg := NewStorageMock(t)
	err := errors.New("error")
	strg.PutMock.Return(err)

	cmd := new(Expire)

	expectedRes := ErrResult{err}
	actualRes := cmd.Execute(strg, []string{"key", "10"}...)

	assert.Equal(t, expectedRes, actualRes)
}

func TestExpire_Execute_Setter(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"key": storage.NewStringValue("value"),
	})

	cmd := new(Expire)

	expectedRes := OkResult{}
	actualRes := cmd.Execute(strg, []string{"key", "10"}...)
	assert.Equal(t, expectedRes, actualRes)

	items, err := strg.All()
	assert.NoError(t, err)

	value, ok := items["key"]
	assert.True(t, ok)
	assert.True(t, time.Unix(value.TTL(), 0).After(time.Now()))
}
