package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["GET"] = new(Get)
}

//Get is the GET command
type Get struct{}

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

//ArgsNumber implements ArgsNumber of Command interface
func (c *Get) ArgsNumber() int {
	return 1
}

//Execute implements Execute of Command interface
func (c *Get) Execute(strg storage.Storage, args ...string) Result {
	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{err}
	}
	if value.Type() != storage.StringDataType {
		return ErrResult{ErrWrongTypeOp}
	}
	return StringResult{value.Data().(string)}
}
