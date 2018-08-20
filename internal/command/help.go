package command

import (
	"fmt"
)

//Help is the Help command
type Help struct {
	parser commandParser
}

//Name implements Name of Command interface
func (c *Help) Name() string {
	return "HELP"
}

//Help implements Help of Command interface
func (c *Help) Help() string {
	return `Usage: HELP command
Show the usage of the given command`
}

//Execute implements Execute of Command interface
func (c *Help) Execute(args ...string) Result {
	if len(args) != 1 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	cmdName := args[0]

	cmd, _, err := c.parser.Parse(cmdName)
	if err != nil {
		if err == ErrCommandNotFound {
			return ErrResult{Value: fmt.Errorf("command %q not found", cmdName)}
		}
		return ErrResult{Value: err}
	}
	return RawStringResult{Value: cmd.Help()}
}
