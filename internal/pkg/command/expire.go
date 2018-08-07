package command

import (
	"errors"
	"strconv"
	"time"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/pkg/clock"
)

func init() {
	cmd := &Expire{clck: clock.TimeClock{}}
	commands[cmd.Name()] = cmd
}

//Expire is the Expire command
type Expire struct {
	clck clock.Clock
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
func (c *Expire) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 2 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}
	secs, err := strconv.Atoi(args[1])
	if err != nil {
		return ErrResult{Value: errors.New("seconds should be integer")}
	}
	if secs < 0 {
		return ErrResult{Value: errors.New("seconds should be positive")}
	}
	setter := func(old *storage.Value) (*storage.Value, error) {
		if old == nil {
			return nil, nil
		}
		now := c.clck.Now()
		old.SetTTL(now.Add(time.Duration(secs) * time.Second))
		return old, nil
	}
	if err := strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{Value: err}
	}
	return OkResult{}
}
