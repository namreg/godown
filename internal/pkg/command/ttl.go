package command

import (
	"time"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(TTL)
	commands[cmd.Name()] = cmd
}

//TTL is the TTL command
type TTL struct{}

//Name implements Name of Command interface
func (c *TTL) Name() string {
	return "TTL"
}

//Help implements Help of Command interface
func (c *TTL) Help() string {
	return `Usage: TTL key
Ttl the given key.`
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
	return IntResult{int64(time.Until(time.Unix(value.TTL(), 0)).Seconds())}
}
