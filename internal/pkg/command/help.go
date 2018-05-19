package command

import (
	"fmt"
	"strings"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["HELP"] = new(Help)
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

//ArgsNumber implements ArgsNumber of Command interface
func (c *Help) ArgsNumber() int {
	return 1
}

//Execute implements Execute of Command interface
func (c *Help) Execute(strg storage.Storage, args ...string) Result {
	cmdName := args[0]
	if cmd, ok := commands[strings.ToUpper(cmdName)]; ok {
		return HelpResult{cmd}
	}
	return ErrResult{fmt.Errorf("command %q not found", cmdName)}
}
