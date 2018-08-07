package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Del)
	commands[cmd.Name()] = cmd
}

//Del is the DEL command
type Del struct{}

//Name implements Name of Command i`nterface
func (c *Del) Name() string {
	return "DEL"
}

//Help implements Help of Command interface
func (c *Del) Help() string {
	return `Usage: DEL key
Del the given key.`
}

//Execute implements Execute of Command interface
func (c *Del) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}
	err := strg.Del(storage.Key(args[0]))
	if err != nil {
		return ErrResult{Value: err}
	}
	return OkResult{}
}
