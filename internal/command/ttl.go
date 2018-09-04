package command

import (
	"time"

	"github.com/namreg/godown/internal/storage"
)

//TTL is the TTL command
type TTL struct {
	strg dataStore
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
func (c *TTL) Execute(args ...string) Reply {
	if len(args) != 1 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	value, err := c.strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilReply{}
		}
		return ErrReply{Value: err}
	}
	if value.TTL() < 0 {
		return IntReply{Value: -1}
	}
	secs := time.Unix(value.TTL(), 0).Sub(c.clck.Now()).Seconds()
	return IntReply{Value: int64(secs)}
}
