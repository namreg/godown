package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Type)
	commands[cmd.Name()] = cmd
}

//Type is the Type command
type Type struct{}

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
func (c *Type) Execute(strg storage.Storage, args ...string) Result {
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
	return StringResult{Value: string(value.Type())}
}
