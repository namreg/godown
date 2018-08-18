package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Lpop is the LPOP command
type Lpop struct {
	strg commandStorage
}

//Name implements Name of Command interface
func (c *Lpop) Name() string {
	return "LPOP"
}

//Help implements Help of Command interface
func (c *Lpop) Help() string {
	return `Usage: LPOP key
Removes and returns the first element of the list stored at key.`
}

//Execute implements Execute of Command interface
func (c *Lpop) Execute(args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	c.strg.Lock()
	defer c.strg.Unlock()

	key := storage.Key(args[0])

	val, err := c.strg.Get(key)
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{Value: err}
	}

	if val.Type() != storage.ListDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}

	list := val.Data().([]string)
	popped, list := list[0], list[1:]

	if len(list) == 0 {
		if err = c.strg.Del(key); err != nil {
			return ErrResult{Value: err}
		}
		return StringResult{Value: popped}
	}

	if err = c.strg.Put(key, storage.NewListValue(list)); err != nil {
		return ErrResult{Value: err}
	}

	return StringResult{Value: popped}
}
