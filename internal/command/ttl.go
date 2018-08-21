package command

import (
	"time"

	"github.com/namreg/godown-v2/internal/storage"
)

//TTL is the TTL command
type TTL struct {
	strg commandStorage
	clck commandClock
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
func (c *TTL) Execute(args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	value, err := c.strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{Value: err}
	}
	if value.TTL() < 0 {
		return IntResult{Value: -1}
	}
	secs := time.Unix(value.TTL(), 0).Sub(c.clck.Now()).Seconds()
	return IntResult{Value: int64(secs)}
}
