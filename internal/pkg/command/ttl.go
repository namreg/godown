package command

import (
	"time"

	"github.com/namreg/godown-v2/pkg/clock"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := &TTL{clock.TimeClock{}}
	commands[cmd.Name()] = cmd
}

//TTL is the TTL command
type TTL struct {
	clck clock.Clock
}

//Name implements Name of Command interface
func (c *TTL) Name() string {
	return "TTL"
}

//Help implements Help of Command interface
func (c *TTL) Help() string {
	return `Usage: TTL key
Returns the remaining time to live of a key. -1 returns if key does not have timeout.`
}

//Execute implements Execute of Command interface
func (c *TTL) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 1 {
		return ErrResult{ErrWrongArgsNumber}
	}

	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{err}
	}
	if value.TTL() < 0 {
		return IntResult{-1}
	}
	secs := time.Unix(value.TTL(), 0).Sub(c.clck.Now()).Seconds()
	return IntResult{int64(secs)}
}
