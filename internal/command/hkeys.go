package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Hkeys is the HKEYS command
type Hkeys struct {
	strg commandStorage
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
func (c *Hkeys) Execute(args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	c.strg.RLock()
	value, err := c.strg.Get(storage.Key(args[0]))
	c.strg.RUnlock()
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{Value: err}
	}
	if value.Type() != storage.MapDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}
	m := value.Data().(map[string]string)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return SliceResult{Value: keys}
}
