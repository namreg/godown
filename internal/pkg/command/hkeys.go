package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Hkeys)
	commands[cmd.Name()] = cmd
}

//Hkeys is the HKEYS command
type Hkeys struct{}

//Name implements Name of Command interface
func (c *Hkeys) Name() string {
	return "HKEYS"
}

//Help implements Help of Command interface
func (c *Hkeys) Help() string {
	return `Usage: HKEYS key
Returns all field names in the hash stored at key. Order of fields is not guaranteed`
}

//Execute implements Execute of Command interface
func (c *Hkeys) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	strg.RLock()
	value, err := strg.Get(storage.Key(args[0]))
	strg.RUnlock()
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{Value: err}
	}
	if value.Type() != storage.MapDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}
	m := value.Data().(map[string]string)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return SliceResult{Value: keys}
}
