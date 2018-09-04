package command

import (
	"errors"
	"strconv"
	"time"

	"github.com/namreg/godown/internal/storage"
)

//Expire is the Expire command
type Expire struct {
	clck commandClock
	strg dataStore
}

//Name implements Name of Command interface
func (c *Expire) Name() string {
	return "EXPIRE"
}

//Help implements Help of Command interface
func (c *Expire) Help() string {
	return `Usage: EXPIRE key seconds
Set a timeout on key. After the timeout has expired, the key will automatically be deleted.`
}

//Execute implements Execute of Command interface
func (c *Expire) Execute(args ...string) Reply {
	if len(args) != 2 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}
	secs, err := strconv.Atoi(args[1])
	if err != nil {
		return ErrReply{Value: errors.New("seconds should be integer")}
	}
	if secs < 0 {
		return ErrReply{Value: errors.New("seconds should be positive")}
	}
	setter := func(old *storage.Value) (*storage.Value, error) {
		if old == nil {
			return nil, nil
		}

		now := c.clck.Now()

		old.SetTTL(now.Add(time.Duration(secs) * time.Second))

		return old, nil
	}
	if err := c.strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrReply{Value: err}
	}
	return OkReply{}
}
