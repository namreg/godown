package command

import "github.com/namreg/godown/internal/storage"

// Rpop is a RPOP command.
type Rpop struct {
	strg dataStore
}

// Name returns a command name.
// Implements Name method of Command interface.
func (c *Rpop) Name() string {
	return "RPOP"
}

// Help returns a help message for the command.
// Implements Help method of Command interface.
func (c *Rpop) Help() string {
	return `Usage: RPOP key
Removes and returns the last element of the list stored at key.`
}

// Execute excutes a command with the given arguments.
// Implements Execute method of Command interface.
func (c *Rpop) Execute(args ...string) Reply {
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
		popped, list = list[len(list)-1], list[:len(list)-1]

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
