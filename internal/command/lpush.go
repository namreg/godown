package command

import (
	"github.com/namreg/godown/internal/storage"
)

//Lpush is the LPUSH command
type Lpush struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Lpush) Name() string {
	return "LPUSH"
}

//Help implements Help of Command interface
func (c *Lpush) Help() string {
	return `Usage: LPUSH key value [value ...]
Prepend one or multiple values to a list.`
}

//Execute implements Execute of Command interface
func (c *Lpush) Execute(args ...string) Reply {
	if len(args) < 2 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	setter := func(old *storage.Value) (*storage.Value, error) {
		vals := args[1:]
		// reverse vals
		for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
			vals[i], vals[j] = vals[j], vals[i]
		}
		if old == nil {
			return storage.NewList(vals), nil
		}

		if old.Type() != storage.ListDataType {
			return nil, ErrWrongTypeOp
		}

		oldList := old.Data().([]string)

		newList := make([]string, 0, len(oldList)+len(vals))
		newList = append(newList, vals...)
		newList = append(newList, oldList...)

		return storage.NewList(newList), nil
	}
	if err := c.strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrReply{Value: err}
	}
	return OkReply{}
}
