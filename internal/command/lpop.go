package command

import (
	"github.com/namreg/godown-v2/internal/storage"
)

//Lpop is the LPOP command
type Lpop struct {
	strg dataStore
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
func (c *Lpop) Execute(args ...string) Reply {
	if len(args) != 1 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	var popped string

	setter := func(old *storage.Value) (*storage.Value, error) {
		if old == nil {
			return nil, nil
		}

		if old.Type() != storage.ListDataType {
			return nil, ErrWrongTypeOp
		}

		list := old.Data().([]string)
		popped, list = list[0], list[1:]

		if len(list) == 0 {
			return nil, nil
		}

		return storage.NewList(list), nil
	}

	if err := c.strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrReply{Value: err}
	}

	if popped == "" {
		return NilReply{}
	}
	return StringReply{Value: popped}
}
