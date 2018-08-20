package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Get is the GET command
type Get struct {
	strg commandStorage
}

//Name implements Name of Command interface
func (c *Get) Name() string {
	return "GET"
}

//Help implements Help of Command interface
func (c *Get) Help() string {
	return `Usage: GET key
Get the value by key.
If provided key does not exist NIL will be returned.`
}

//Execute implements Execute of Command interface
func (c *Get) Execute(args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}
	c.strg.RLock()
	value, err := c.strg.Get(storage.Key(args[0]))
	c.strg.RUnlock()
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{Value: err}
	}
	if value.Type() != storage.StringDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}
	return StringResult{Value: value.Data().(string)}
}
