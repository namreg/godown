package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Del is the DEL command
type Del struct {
	strg dataStore
}

//Name implements Name of Command i`nterface
func (c *Del) Name() string {
	return "DEL"
}

//Help implements Help of Command interface
func (c *Del) Help() string {
	return `Usage: DEL key
Del the given key.`
}

//Execute implements Execute of Command interface
func (c *Del) Execute(args ...string) Reply {
	if len(args) != 1 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	if err := c.strg.Del(storage.Key(args[0])); err != nil {
		return ErrReply{Value: err}
	}
	return OkReply{}
}
