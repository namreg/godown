package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Hset is the HSET command
type Hset struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Hset) Name() string {
	return "HSET"
}

//Help implements Help of Command interface
func (c *Hset) Help() string {
	return `Usage: HSET key field value
Sets field in the hash stored at key to value.`
}

//Execute implements Execute of Command interface
func (c *Hset) Execute(args ...string) Reply {
	if len(args) != 3 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	setter := func(old *storage.Value) (*storage.Value, error) {
		mfield, mvalue := args[1], args[2]
		if old == nil {
			return storage.NewMap(map[string]string{mfield: mvalue}), nil
		}

		if old.Type() != storage.MapDataType {
			return nil, ErrWrongTypeOp
		}

		m := old.Data().(map[string]string)

		m[mfield] = mvalue

		return storage.NewMap(m), nil
	}
	if err := c.strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrReply{Value: err}
	}
	return OkReply{}
}
