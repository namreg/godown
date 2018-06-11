package command

import (
	"fmt"
	"strings"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Help)
	commands[cmd.Name()] = cmd
}

//Help is the Help command
type Help struct{}

//Name implements Name of Command interface
func (c *Help) Name() string {
	return "HELP"
}

//Help implements Help of Command interface
func (c *Help) Help() string {
	return `Usage: HELP command
Show the usage of the given command
`
}

//ValidateArgs implements ValidateArgs of Command interface
func (c *Help) ValidateArgs(args ...string) error {
	if len(args) != 1 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Help) Execute(strg storage.Storage, args ...string) Result {
	if err := c.ValidateArgs(args...); err != nil {
		return ErrResult{err}
	}

	cmdName := args[0]
	if cmd, ok := commands[strings.ToUpper(cmdName)]; ok {
		return HelpResult{cmd}
	}
	return ErrResult{fmt.Errorf("command %q not found", cmdName)}
}
