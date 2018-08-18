package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Del is the DEL command
type Del struct {
	strg commandStorage
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
func (c *Del) Execute(args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}
	c.strg.Lock()
	err := c.strg.Del(storage.Key(args[0]))
	c.strg.Unlock()
	if err != nil {
		return ErrResult{Value: err}
	}
	return OkResult{}
}
