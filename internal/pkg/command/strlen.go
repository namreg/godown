package command

import (
	"unicode/utf8"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["STRLEN"] = new(Strlen)
}

//Strlen is the Strlen command
type Strlen struct{}

//Name implements Name of Command interface
func (c *Strlen) Name() string {
	return "STRLEN"
}

//Help implements Help of Command interface
func (c *Strlen) Help() string {
	return `Usage: STRLEN key
Returns length of the given key.
If key does not exists, 0 will be returned.`
}

//ArgsNumber implements ArgsNumber of Command interface
func (c *Strlen) ArgsNumber() int {
	return 1
}

//Execute implements Execute of Command interface
func (c *Strlen) Execute(strg storage.Storage, args ...string) Result {
	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return IntResult{0}
		}
		return ErrResult{err}
	}
	if value.Type() != storage.StringDataType {
		return ErrResult{ErrWrongTypeOp}
	}
	cnt := utf8.RuneCountInString(value.Data().(string))
	return IntResult{int64(cnt)}
}
