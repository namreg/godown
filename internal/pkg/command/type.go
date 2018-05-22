package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["TYPE"] = new(Type)
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

//ValidateArgs implements ValidateArgs of Command interface
func (c *Type) ValidateArgs(args ...string) error {
	if len(args) != 1 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Type) Execute(strg storage.Storage, args ...string) Result {
	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{err}
	}
	return StringResult{string(value.Type())}
}
