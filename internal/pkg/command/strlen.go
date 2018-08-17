package command

import (
	"unicode/utf8"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

//Strlen is the Strlen command
type Strlen struct {
	strg commandStorage
}

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
func (c *Strlen) Execute(args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	c.strg.RLock()
	value, err := c.strg.Get(storage.Key(args[0]))
	c.strg.RUnlock()
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
