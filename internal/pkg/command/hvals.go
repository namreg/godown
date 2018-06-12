package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Hvals)
	commands[cmd.Name()] = cmd
}

//Hvals is the HVALS command
type Hvals struct{}

//Name implements Name of Command interface
func (c *Hvals) Name() string {
	return "HVALS"
}

//Help implements Help of Command interface
func (c *Hvals) Help() string {
	return `Usage: HVALS key
Returns all values in the hash stored at key`
}

//Execute implements Execute of Command interface
func (c *Hvals) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 1 {
		return ErrResult{ErrWrongArgsNumber}
	}

	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{err}
	}
	if value.Type() != storage.MapDataType {
		return ErrResult{ErrWrongTypeOp}
	}
	m := value.Data().(map[string]string)
	vals := make([]string, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return SliceResult{vals}
}
