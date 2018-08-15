package command

import (
	"unicode/utf8"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Strlen)
	commands[cmd.Name()] = cmd
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

//Execute implements Execute of Command interface
func (c *Strlen) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	strg.RLock()
	value, err := strg.Get(storage.Key(args[0]))
	strg.RUnlock()
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return IntResult{Value: 0}
		}
		return ErrResult{Value: err}
	}
	if value.Type() != storage.StringDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}
	cnt := utf8.RuneCountInString(value.Data().(string))
	return IntResult{Value: int64(cnt)}
}
