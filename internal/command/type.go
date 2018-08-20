package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Type is the Type command
type Type struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Type) Name() string {
	return "TYPE"
}

//Help implements Help of Command interface
func (c *Type) Help() string {
	return `Usage: TYPE key
Returns the type stored at key.`
}

//Execute implements Execute of Command interface
func (c *Type) Execute(args ...string) Reply {
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
	return StringReply{Value: string(value.Type())}
}
