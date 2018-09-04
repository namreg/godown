package command

import (
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/storage"
	"github.com/namreg/godown/internal/storage/memory"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestTTL_Name(t *testing.T) {
	cmd := new(TTL)
	assert.Equal(t, "TTL", cmd.Name())
}

func TestTTL_Help(t *testing.T) {
	cmd := new(TTL)
	expected := `Usage: TTL key
Returns the remaining time to live of a key. -1 returns if key does not have timeout.`
	assert.Equal(t, expected, cmd.Help())
}

func TestTTL_Execute(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	testTime, _ := time.Parse("2006-01-02 15:04:05", "2018-01-01 11:11:11")

	clck := NewcommandClockMock(t)
	clck.NowMock.Return(testTime)

	now := clck.Now()

	expired := storage.NewString("value")
	expired.SetTTL(now.Add(-1 * time.Second))

	willExpire := storage.NewString("value")
	willExpire.SetTTL(now.Add(10 * time.Second))

	strg := memory.New(
		map[storage.Key]*storage.Value{
			"no_timeout":  storage.NewString("value"),
			"expired":     expired,
			"will_expire": willExpire,
		},
		memory.WithClock(clck),
	)

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"no_timeout", []string{"no_timeout"}, IntReply{Value: -1}},
		{"expired", []string{"expired"}, NilReply{}},
		{"will_expire", []string{"will_expire"}, IntReply{Value: now.Add(10*time.Second).Unix() - now.Unix()}},
		{"not_existing_key", []string{"not_existing_key"}, NilReply{}},
		{"wrong_number_of_args/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"key", "arg1"}, ErrReply{Value: ErrWrongArgsNumber}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := TTL{clck: clck, strg: strg}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestTTL_Execute_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewdataStoreMock(mc)
	strg.GetMock.Return(nil, err)

	cmd := TTL{strg: strg}
	res := cmd.Execute("key")

	assert.Equal(t, ErrReply{Value: err}, res)
}
