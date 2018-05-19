package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["KEYS"] = new(Keys)
}

//Keys is the Keys command
type Keys struct{}

//Name implements Name of Command interface
func (c *Keys) Name() string {
	return "KEYS"
}

//Help implements Help of Command interface
func (c *Keys) Help() string {
	return `Usage: KEYS pattern
Find all keys matching the given pattern.`
}

//ArgsNumber implements ArgsNumber of Command interface
func (c *Keys) ArgsNumber() int {
	return 1
}

//Execute implements Execute of Command interface
func (c *Keys) Execute(strg storage.Storage, args ...string) Result {
	keys, err := strg.Keys()
	if err != nil {
		return ErrResult{err}
	}

	keyNames := make([]string, 0, len(keys))
	for _, k := range keys {
		keyNames = append(keyNames, k.Val())
	}
	return SliceResult{keyNames}
}
