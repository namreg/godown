package command

import (
	"unicode/utf8"

	"github.com/namreg/godown/internal/storage"
)

//Strlen is the Strlen command
type Strlen struct {
	strg dataStore
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
func (c *Strlen) Execute(args ...string) Reply {
	if len(args) != 1 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	value, err := c.strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return IntReply{Value: 0}
		}
		return ErrReply{Value: err}
	}
	if value.Type() != storage.StringDataType {
		return ErrReply{Value: ErrWrongTypeOp}
	}
	cnt := utf8.RuneCountInString(value.Data().(string))
	return IntReply{Value: int64(cnt)}
}
