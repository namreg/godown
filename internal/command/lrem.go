package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Lrem is the LREM command
type Lrem struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Lrem) Name() string {
	return "LREM"
}

//Help implements Help of Command interface
func (c *Lrem) Help() string {
	return `Usage: LREM key value
Removes all occurrences of elements equal to value from the list stored at key.`
}

//Execute implements Execute of Command interface
func (c *Lrem) Execute(args ...string) Reply {
	if len(args) != 2 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	setter := func(old *storage.Value) (*storage.Value, error) {
		if old == nil {
			return nil, nil
		}
		if old.Type() != storage.ListDataType {
			return nil, ErrWrongTypeOp
		}

		list := old.Data().([]string)
		newList := list[:0]

		for _, val := range list {
			if val != args[1] {
				newList = append(newList, val)
			}
		}

		if len(newList) == 0 {
			return nil, nil
		}
		return storage.NewList(newList), nil
	}
	if err := c.strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrReply{Value: err}
	}
	return OkReply{}
}
