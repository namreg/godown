package command

import (
	"github.com/namreg/godown/internal/storage"
)

//Get is the GET command
type Get struct {
	strg dataStore
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
func (c *Get) Execute(args ...string) Reply {
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
	if value.Type() != storage.StringDataType {
		return ErrReply{Value: ErrWrongTypeOp}
	}
	return StringReply{Value: value.Data().(string)}
}
