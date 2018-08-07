package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Get)
	commands[cmd.Name()] = cmd
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

//Execute implements Execute of Command interface
func (c *Get) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Err: ErrWrongArgsNumber}
	}
	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{Err: err}
	}
	if value.Type() != storage.StringDataType {
		return ErrResult{Err: ErrWrongTypeOp}
	}
	return StringResult{Str: value.Data().(string)}
}
