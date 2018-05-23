package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["LPUSH"] = new(Lpush)
}

//Lpush is the LPUSH command
type Lpush struct{}

//Name implements Name of Command interface
func (c *Lpush) Name() string {
	return "LPUSH"
}

//Help implements Help of Command interface
func (c *Lpush) Help() string {
	return `Usage: LPUSH key value [value ...]
Prepend one or multiple values to a list.`
}

//ValidateArgs implements ValidateArgs of Command interface
func (c *Lpush) ValidateArgs(args ...string) error {
	if len(args) < 2 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Lpush) Execute(strg storage.Storage, args ...string) Result {
	setter := func(old *storage.Value) (*storage.Value, error) {
		vals := args[1:]
		// reverse vals
		for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
			vals[i], vals[j] = vals[j], vals[i]
		}
		if old == nil {
			return storage.NewListValue(vals...), nil
		}

		if old.Type() != storage.ListDataType {
			return nil, ErrWrongTypeOp
		}

		oldList := old.Data().([]string)

		newList := make([]string, 0, len(oldList)+len(vals))
		newList = append(newList, vals...)
		newList = append(newList, oldList...)

		return storage.NewListValue(newList...), nil
	}
	if err := strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}
