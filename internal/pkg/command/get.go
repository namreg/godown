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
		return ErrResult{Value: ErrWrongArgsNumber}
	}
	strg.RLock()
	value, err := strg.Get(storage.Key(args[0]))
	strg.RUnlock()
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
