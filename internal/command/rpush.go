package command

import "github.com/namreg/godown/internal/storage"

// Rpush is a RPUSH command.
type Rpush struct {
	strg dataStore
}

// Name returns a command name.
// Implements Name method of Command interface.
func (c *Rpush) Name() string {
	return "RPUSH"
}

// Help returns a help message for the command.
// Implements Help method of Command interface.
func (c *Rpush) Help() string {
	return `Usage: RPUSH key value [value ...]
Append one or multiple values to a list.`
}

// Execute excutes a command with the given arguments.
// Implements Execute method of Command interface.
func (c *Rpush) Execute(args ...string) Reply {
	if len(args) < 2 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	setter := func(old *storage.Value) (*storage.Value, error) {
		vals := args[1:]
		if old == nil {
			return storage.NewList(vals), nil
		}

		if old.Type() != storage.ListDataType {
			return nil, ErrWrongTypeOp
		}

		oldList := old.Data().([]string)
		newList := append(oldList, vals...)

		return storage.NewList(newList), nil
	}
	if err := c.strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrReply{Value: err}
	}
	return OkReply{}
}
