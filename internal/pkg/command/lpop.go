package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["LPOP"] = new(Lpop)
}

//Lpop is the LPOP command
type Lpop struct{}

//Name implements Name of Command interface
func (c *Lpop) Name() string {
	return "LPOP"
}

//Help implements Help of Command interface
func (c *Lpop) Help() string {
	return `Usage: LPOP key
Removes and returns the first element of the list stored at key.`
}

//ValidateArgs implements ValidateArgs of Command interface
func (c *Lpop) ValidateArgs(args ...string) error {
	if len(args) < 1 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Lpop) Execute(strg storage.Storage, args ...string) Result {
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

		return storage.NewListValue(list...), nil
	}

	if err := strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{err}
	}

	if popped == "" {
		return NilResult{}
	}
	return StringResult{popped}
}
