package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

//Hget is the HGET command
type Hget struct {
	strg commandStorage
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
func (c *Hget) Execute(args ...string) Result {
	if len(args) != 2 {
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
	if v, ok := m[args[1]]; ok {
		return StringResult{Value: v}
	}
	return NilResult{}
}
