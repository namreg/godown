package server

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	storage "github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
	"github.com/namreg/godown-v2/pkg/clock"

	"github.com/gojuno/minimock"
)

func Test_gc_deleteExpired(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	testTime, _ := time.Parse("2006-01-02 15:04:05", "2018-01-01 11:11:11")

	clck := NewClockMock(t)
	clck.NowMock.Return(testTime)

	now := clck.Now()

	expired := storage.NewStringValue("value")
	expired.SetTTL(now.Add(-1 * time.Second))

	willExpire := storage.NewStringValue("value")
	willExpire.SetTTL(now.Add(10 * time.Second))

	strg := memory.New(
		map[storage.Key]*storage.Value{
			"no_timeout":  storage.NewStringValue("value"),
			"expired":     expired,
			"will_expire": willExpire,
		},
		memory.WithClock(clck),
	)

	gc := newGc(strg, log.New(os.Stdout, "", 0), clck, 1*time.Millisecond)
	go gc.start()
	defer gc.stop()

	time.Sleep(2 * time.Millisecond)

	items, err := strg.All()
	assert.NoError(t, err)

	_, ok := items["no_timeout"]
	assert.True(t, ok)

	_, ok = items["expired"]
	assert.False(t, ok)

	_, ok = items["will_expire"]
	assert.True(t, ok)
}

func Test_gc_deleteExpired_StorageErr(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	err := errors.New("error")

	strg := NewStorageMock(mc)
	strg.AllWithTTLMock.Return(nil, err)

	loggerOutput := new(bytes.Buffer)

	gc := newGc(strg, log.New(loggerOutput, "", 0), clock.TimeClock{}, 1*time.Millisecond)
	gc.deleteExpired()

	assert.Equal(t, "[WARN] gc: could not retrieve values: error\n", loggerOutput.String())
}
