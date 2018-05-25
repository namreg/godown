package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["HGET"] = new(Hget)
}

//Hget is the HGET command
type Hget struct{}

//Name implements Name of Command interface
func (c *Hget) Name() string {
	return "HGET"
}

//Help implements Help of Command interface
func (c *Hget) Help() string {
	return `Usage: HGET key field
Returns the value associated with field in the hash stored at key.`
}

//ValidateArgs implements ValidateArgs of Command interface
func (c *Hget) ValidateArgs(args ...string) error {
	if len(args) != 2 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Hget) Execute(strg storage.Storage, args ...string) Result {
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
	if v, ok := m[args[1]]; ok {
		return StringResult{v}
	}
	return NilResult{}
}
