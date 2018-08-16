package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

//Llen is the LLEN command
type Llen struct {
	strg commandStorage
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
func (c *Llen) Execute(args ...string) Result {
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
	if value.Type() != storage.ListDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}
	l := len(value.Data().([]string))
	return IntResult{Value: int64(l)}
}
