package command

import (
	"errors"
	"testing"
	"time"

	"github.com/namreg/godown-v2/internal/clock"
	"github.com/namreg/godown-v2/internal/storage"
	"github.com/namreg/godown-v2/internal/storage/memory"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/assert"
)

func TestExpire_Name(t *testing.T) {
	cmd := Expire{clck: clock.New()}
	assert.Equal(t, "EXPIRE", cmd.Name())
}

func TestExpire_Help(t *testing.T) {
	cmd := Expire{clck: clock.New()}
	expexted := `Usage: EXPIRE key seconds
Set a timeout on key. After the timeout has expired, the key will automatically be deleted.`

	assert.Equal(t, expexted, cmd.Help())
}

func TestExpire_Execute(t *testing.T) {
	strg := memory.New(map[storage.Key]*storage.Value{
		"key": storage.NewString("value"),
	})
	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"wrong_arg_number/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_arg_number/2", []string{"key", "10", "20"}, ErrReply{Value: ErrWrongArgsNumber}},
		{"existing_key", []string{"key", "10"}, OkReply{}},
		{"not_existing_key", []string{"not_existing_key", "10"}, OkReply{}},
		{"secs_as_string", []string{"not_existing_key", "seconds"}, ErrReply{Value: errors.New("seconds should be integer")}},
		{"secs_negative", []string{"not_existing_key", "-10"}, ErrReply{Value: errors.New("seconds should be positive")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Expire{strg: strg, clck: clock.New()}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestExpire_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.PutMock.Return(err)

	cmd := Expire{strg: strg}

	expectedRes := ErrReply{Value: err}
	actualRes := cmd.Execute("key", "10")

	assert.Equal(t, expectedRes, actualRes)
}

func TestExpire_Execute_WhiteBox(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	now, _ := time.Parse("2006-01-02 15:04:05", "2018-01-01 11:11:11")

	clck := NewcommandClockMock(mc)
	clck.NowMock.Return(now)

	strg := memory.New(map[storage.Key]*storage.Value{
		"key": storage.NewString("value"),
	})

	cmd := Expire{strg: strg, clck: clck}

	res := cmd.Execute([]string{"key", "10"}...)
	assert.Equal(t, OkReply{}, res)

	items, err := strg.All()
	assert.NoError(t, err)

	value, ok := items["key"]
	assert.True(t, ok)
	assert.Equal(t, now.Add(10*time.Second).Unix(), value.TTL())
}
