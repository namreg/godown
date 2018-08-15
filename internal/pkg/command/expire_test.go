package command

import (
	"errors"
	"testing"
	"time"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
	"github.com/namreg/godown-v2/pkg/clock"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/assert"
)

func TestExpire_Name(t *testing.T) {
	cmd := Expire{clck: clock.TimeClock{}}
	assert.Equal(t, "EXPIRE", cmd.Name())
}

func TestExpire_Help(t *testing.T) {
	cmd := Expire{clck: clock.TimeClock{}}
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
		{"wrong_arg_number/1", []string{}, ErrResult{Value: ErrWrongArgsNumber}},
		{"wrong_arg_number/2", []string{"key", "10", "20"}, ErrResult{Value: ErrWrongArgsNumber}},
		{"existing_key", []string{"key", "10"}, OkResult{}},
		{"not_existing_key", []string{"not_existing_key", "10"}, OkResult{}},
		{"secs_as_string", []string{"not_existing_key", "seconds"}, ErrResult{Value: errors.New("seconds should be integer")}},
		{"secs_negative", []string{"not_existing_key", "-10"}, ErrResult{Value: errors.New("seconds should be positive")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Expire{clck: clock.TimeClock{}}
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestExpire_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg1 := storage.NewStorageMock(t)
	strg1.GetMock.Return(nil, err)
	strg1.LockMock.Return()
	strg1.UnlockMock.Return()

	strg2 := storage.NewStorageMock(t)
	strg2.GetMock.Return(storage.NewStringValue("value"), nil)
	strg2.PutMock.Return(err)
	strg2.LockMock.Return()
	strg2.UnlockMock.Return()

	cmd := Expire{clck: clock.TimeClock{}}

	expectedRes := ErrResult{Value: err}

	actualRes1 := cmd.Execute(strg1, []string{"key", "10"}...)
	assert.Equal(t, expectedRes, actualRes1)

	actualRes2 := cmd.Execute(strg2, []string{"key", "10"}...)
	assert.Equal(t, expectedRes, actualRes2)
}

func TestExpire_Execute_WhiteBox(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	now, _ := time.Parse("2006-01-02 15:04:05", "2018-01-01 11:11:11")

	clck := clock.NewClockMock(t)
	clck.NowMock.Return(now)

	strg := memory.New(map[storage.Key]*storage.Value{
		"key": storage.NewStringValue("value"),
	})

	cmd := Expire{clck: clck}

	res := cmd.Execute(strg, []string{"key", "10"}...)
	assert.Equal(t, OkResult{}, res)

	items, err := strg.All()
	assert.NoError(t, err)

	value, ok := items["key"]
	assert.True(t, ok)
	assert.Equal(t, now.Add(10*time.Second).Unix(), value.TTL())
}
