package command

import (
	"errors"
	"time"

	"github.com/namreg/godown-v2/internal/storage"
)

var (
	//ErrWrongArgsNumber means that given arguments not acceptable by Command.
	ErrWrongArgsNumber = errors.New("command: wrong args number")
	//ErrWrongTypeOp means that operation is not acceptable for the given key.
	ErrWrongTypeOp = errors.New("command: wrong type operation")
)

//Command represents a command thats server can execute.
//go:generate minimock -i github.com/namreg/godown-v2/internal/command.Command -o ./
type Command interface {
	//Name returns the command name.
	Name() string
	//Help returns information about the command. Description, usage and etc.
	Help() string
	//Execute executes the command with the given arguments.
	Execute(args ...string) Result
}

//go:generate minimock -i github.com/namreg/godown-v2/internal/command.commandStorage -o ./
type commandStorage interface {
	//Put puts a new value at the given Key.
	Put(key storage.Key, sttr storage.ValueSetter) error
	//Get gets a value by the given key.
	Get(key storage.Key) (*storage.Value, error)
	//Del deletes a value by the given key.
	Del(key storage.Key) error
	//Keys returns all stored keys.
	Keys() ([]storage.Key, error)
}

//go:generate minimock -i github.com/namreg/godown-v2/internal/command.commandParser -o ./
type commandParser interface {
	Parse(str string) (cmd Command, args []string, err error)
}

//go:generate minimock -i github.com/namreg/godown-v2/internal/command.commandClock -o ./
type commandClock interface {
	//Now returns current time.
	Now() time.Time
}
