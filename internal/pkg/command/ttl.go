package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["TTL"] = new(TTL)
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

//ValidateArgs implements ValidateArgs of Command interface
func (c *TTL) ValidateArgs(args ...string) error {
	if len(args) != 1 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *TTL) Execute(strg storage.Storage, args ...string) Result {
	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{err}
	}
	return IntResult{value.TTL()}
}
