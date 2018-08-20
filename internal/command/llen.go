package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Llen is the LLEN command
type Llen struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Llen) Name() string {
	return "LLEN"
}

//Help implements Help of Command interface
func (c *Llen) Help() string {
	return `Usage: LLEN key
Returns the length of the list stored at key. 
If key does not exist, it is interpreted as an empty list and 0 is returned.`
}

//Execute implements Execute of Command interface
func (c *Llen) Execute(args ...string) Reply {
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
	if value.Type() != storage.ListDataType {
		return ErrReply{Value: ErrWrongTypeOp}
	}
	l := len(value.Data().([]string))
	return IntReply{Value: int64(l)}
}
