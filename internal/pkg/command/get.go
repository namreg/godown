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

//ValidateArgs implements ValidateArgs of Command interface
func (c *Get) ValidateArgs(args ...string) error {
	if len(args) != 1 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Get) Execute(strg storage.Storage, args ...string) Result {
	if err := c.ValidateArgs(args...); err != nil {
		return ErrResult{err}
	}
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
