package command

import (
	"github.com/namreg/godown/internal/storage"
)

//Hdel is the Hdel command
type Hdel struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Hdel) Name() string {
	return "HDEL"
}

//Help implements Help of Command interface
func (c *Hdel) Help() string {
	return `Usage: HDEL key field [field ...]
Removes the specified fields from the hash stored at key.
Returns the number of fields that were removed.`
}

//Execute implements Execute of Command interface
func (c *Hdel) Execute(args ...string) Reply {
	if len(args) < 2 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	var deleted int

	setter := func(old *storage.Value) (*storage.Value, error) {
		if old == nil {
			return nil, nil
		}

		if old.Type() != storage.MapDataType {
			return nil, ErrWrongTypeOp
		}

		m := old.Data().(map[string]string)

		for _, field := range args[1:] {
			if _, ok := m[field]; ok {
				deleted++
			}
			delete(m, field)
		}

		if len(m) == 0 {
			return nil, nil
		}

		return storage.NewMap(m), nil
	}
	if err := c.strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrReply{Value: err}
	}
	return IntReply{Value: int64(deleted)}
}
