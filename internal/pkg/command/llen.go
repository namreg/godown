package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Llen)
	commands[cmd.Name()] = cmd
}

//Llen is the LLEN command
type Llen struct{}

//Name implements Name of Command interface
func (c *Llen) Name() string {
	return "LLEN"
}

//Help implements Help of Command interface
func (c *Llen) Help() string {
	return `Usage: LLEN key
Returns the length of the list stored at key. 
If key does not exist, it is interpreted as an empty list and 0 is returned.`
}

//Execute implements Execute of Command interface
func (c *Llen) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Err: ErrWrongArgsNumber}
	}

	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return IntResult{Value: 0}
		}
		return ErrResult{Err: err}
	}
	if value.Type() != storage.ListDataType {
		return ErrResult{Err: ErrWrongTypeOp}
	}
	l := len(value.Data().([]string))
	return IntResult{Value: int64(l)}
}
