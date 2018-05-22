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

//ValidateArgs implements ValidateArgs of Command interface
func (c *Set) ValidateArgs(args ...string) error {
	if len(args) != 2 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Set) Execute(strg storage.Storage, args ...string) Result {
	value := strings.Join(args[1:], " ")

	setter := func(old *storage.Value) (*storage.Value, error) {
		return storage.NewStringValue(value), nil
	}

	if err := strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}
