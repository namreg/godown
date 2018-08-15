package command

import (
	"strings"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Set)
	commands[cmd.Name()] = cmd
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

//Execute implements Execute of Command interface
func (c *Set) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 2 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	key := storage.Key(args[0])
	value := strings.Join(args[1:], " ")

	if err := strg.Put(key, storage.NewStringValue(value)); err != nil {
		return ErrResult{Value: err}
	}

	return OkResult{}
}
