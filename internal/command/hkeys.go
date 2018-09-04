package command

import (
	"github.com/namreg/godown/internal/storage"
)

//Hkeys is the HKEYS command
type Hkeys struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Hkeys) Name() string {
	return "HKEYS"
}

//Help implements Help of Command interface
func (c *Hkeys) Help() string {
	return `Usage: HKEYS key
Returns all field names in the hash stored at key. Order of fields is not guaranteed`
}

//Execute implements Execute of Command interface
func (c *Hkeys) Execute(args ...string) Reply {
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

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return SliceReply{Value: keys}
}
