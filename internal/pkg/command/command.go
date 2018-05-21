package command

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/pkg/errors"
)

var (
	//ErrCommandNotFound means that command could not be parsed. Returns by Parse
	ErrCommandNotFound = errors.New("command: not found")
	//ErrWrongArgsNumber means that given arguments not acceptable by Command. Returns by Parse
	ErrWrongArgsNumber = errors.New("command: wrong args number")
	//ErrWrongTypeOp means that operation is not acceptable for the given key
	ErrWrongTypeOp = errors.New("command: wrong type operation")
)

//commands is the all available commands
var commands = make(map[string]Command)

//Command represents a command thats server can execute
type Command interface {
	//Name returns the command name
	Name() string
	//Help returns information about the command. Description, usage and etc.
	Help() string
	//ArgsNumber returns the number of arguments that command cat accept
	ArgsNumber() int
	//Execute executes the command in the context of Storage with the given arguments
	Execute(strg storage.Storage, args ...string) Result
}

//Parse parses string to Command with args
func Parse(value string) (Command, []string, error) {
	args := extractArgs(value)

	cmd, ok := commands[strings.ToUpper(args[0])]
	if !ok {
		return nil, nil, ErrCommandNotFound
	}

	args = args[1:]

	if len(args) != cmd.ArgsNumber() {
		return nil, nil, ErrWrongArgsNumber
	}

	return cmd, args, nil
}

func extractArgs(val string) []string {
	args := make([]string, 0)
	var inQuote bool
	var buf bytes.Buffer
	for _, r := range []rune(val) {
		switch {
		case r == '"':
			inQuote = !inQuote
		case unicode.IsSpace(r):
			if !inQuote && buf.Len() > 0 {
				args = append(args, buf.String())
				buf.Reset()
			} else {
				buf.WriteRune(r)
			}
		default:
			buf.WriteRune(r)
		}
	}
	if buf.Len() > 0 {
		args = append(args, buf.String())
	}
	return args
}
