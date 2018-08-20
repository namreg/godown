package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Hget is the HGET command
type Hget struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Hget) Name() string {
	return "HGET"
}

//Help implements Help of Command interface
func (c *Hget) Help() string {
	return `Usage: HGET key field
Returns the value associated with field in the hash stored at key.`
}

//Execute implements Execute of Command interface
func (c *Hget) Execute(args ...string) Reply {
	if len(args) != 2 {
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
	if v, ok := m[args[1]]; ok {
		return StringReply{Value: v}
	}
	return NilReply{}
}
