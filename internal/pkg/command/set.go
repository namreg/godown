package command

import (
	"strings"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["SET"] = new(Set)
}

//Set is the SET command
type Set struct{}

//Name implements Name of Command interface
func (c *Set) Name() string {
	return "SET"
}

//Help implements Help of Command interface
func (c *Set) Help() string {
	return `Usage: SET key value
Set key to hold the string value.
If key already holds a value, it is overwritten.`
}

//ArgsNumber implements ArgsNumber of Command interface
func (c *Set) ArgsNumber() int {
	return 2
}

//Execute implements Execute of Command interface
func (c *Set) Execute(strg storage.Storage, args ...string) Result {
	key := storage.NewStringKey(args[0])
	err := strg.Put(key, storage.Value(strings.Join(args[1:], " ")))
	if err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}
