package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

//Hset is the HSET command
type Hset struct {
	strg commandStorage
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
func (c *Hset) Execute(args ...string) Result {
	if len(args) != 3 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	key := storage.Key(args[0])
	mfield, mvalue := args[1], args[2]

	c.strg.Lock()
	defer c.strg.Unlock()

	old, err := c.strg.Get(key)
	if err != nil && err != storage.ErrKeyNotExists {
		return ErrResult{Value: err}
	}

	if old == nil {
		return c.put(key, storage.NewMapValue(map[string]string{mfield: mvalue}))
	}

	if old.Type() != storage.MapDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}

	m := old.Data().(map[string]string)
	m[mfield] = mvalue

	return c.put(key, storage.NewMapValue(m))
}

func (c *Hset) put(key storage.Key, value *storage.Value) Result {
	if err := c.strg.Put(key, value); err != nil {
		return ErrResult{Value: err}
	}
	return OkResult{}
}
