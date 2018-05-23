package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["LREM"] = new(Lrem)
}

//Lrem is the LREM command
type Lrem struct{}

//Name implements Name of Command interface
func (c *Lrem) Name() string {
	return "LREM"
}

//Help implements Help of Command interface
func (c *Lrem) Help() string {
	return `Usage: LREM key value
Removes all occurrences of elements equal to value from the list stored at key.`
}

//ValidateArgs implements ValidateArgs of Command interface
func (c *Lrem) ValidateArgs(args ...string) error {
	if len(args) < 2 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Lrem) Execute(strg storage.Storage, args ...string) Result {
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
		return storage.NewListValue(newList...), nil
	}
	if err := strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}
