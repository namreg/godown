package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Hvals is the HVALS command
type Hvals struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Hvals) Name() string {
	return "HVALS"
}

//Help implements Help of Command interface
func (c *Hvals) Help() string {
	return `Usage: HVALS key
Returns all values in the hash stored at key`
}

//Execute implements Execute of Command interface
func (c *Hvals) Execute(args ...string) Reply {
	if len(args) != 1 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	value, err := c.strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilReply{}
		}
		return ErrReply{Value: err}
	}
	if value.Type() != storage.MapDataType {
		return ErrReply{Value: ErrWrongTypeOp}
	}
	m := value.Data().(map[string]string)
	vals := make([]string, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return SliceReply{Value: vals}
}
