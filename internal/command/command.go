package command

import (
	"bytes"
	"errors"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/namreg/godown-v2/internal/storage"
)

var (
	//ErrCommandNotFound means that command could not be parsed. Returns by Parser.Parse.
	ErrCommandNotFound = errors.New("command: not found")
	//ErrWrongArgsNumber means that given arguments not acceptable by Command.
	ErrWrongArgsNumber = errors.New("command: wrong args number")
	//ErrWrongTypeOp means that operation is not acceptable for the given key.
	ErrWrongTypeOp = errors.New("command: wrong type operation")
)

//go:generate minimock -i github.com/namreg/godown-v2/internal/command.Command -o ./

//Command represents a command thats server can execute.
type Command interface {
	//Name returns the command name.
	Name() string
	//Help returns information about the command. Description, usage and etc.
	Help() string
	//Execute executes the command with the given arguments.
	Execute(args ...string) Result
}

type rwLocker interface {
	sync.Locker
	RLock()
	RUnlock()
}

//go:generate minimock -i github.com/namreg/godown-v2/internal/command.commandStorage -o ./

type commandStorage interface {
	rwLocker
	//Put puts a new value at the given Key.
	//Warning: method is not thread safe! You should call Lock mannually before calling
	Put(key storage.Key, value *storage.Value) error
	//Get gets a value of a storage by the given key.
	//Warning: method is not thread safe! You should call RLock mannually before calling.
	Get(key storage.Key) (*storage.Value, error)
	//Del deletes a value by the given key.
	//Warning: method is not thread safe! You should call Lock mannually before calling.
	Del(key storage.Key) error
	//Keys returns all stored keys.
	//Warning: method is not thread safe! You should call RLock mannually before calling.
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

//Parser is a parser that parses user input and creates the appropriate command.
type Parser struct {
	strg commandStorage
	clck commandClock
}

//NewParser creates a new parser
func NewParser(strg commandStorage, clck commandClock) *Parser {
	return &Parser{strg: strg, clck: clck}
}

//Parse parses string to Command with args
func (p *Parser) Parse(str string) (Command, []string, error) {
	var cmd Command
	args := p.extractArgs(str)

	switch strings.ToUpper(args[0]) {
	case "HELP":
		cmd = &Help{parser: p}
	case "DEL":
		cmd = &Del{strg: p.strg}
	case "EXPIRE":
		cmd = &Expire{clck: p.clck, strg: p.strg}
	case "GET":
		cmd = &Get{strg: p.strg}
	case "SET":
		cmd = &Set{strg: p.strg}
	case "STRLEN":
		cmd = &Strlen{strg: p.strg}
	case "GETBIT":
		cmd = &GetBit{strg: p.strg}
	case "SETBIT":
		cmd = &SetBit{strg: p.strg}
	case "HGET":
		cmd = &Hget{strg: p.strg}
	case "HKEYS":
		cmd = &Hkeys{strg: p.strg}
	case "HSET":
		cmd = &Hset{strg: p.strg}
	case "HVALS":
		cmd = &Hvals{strg: p.strg}
	case "KEYS":
		cmd = &Keys{strg: p.strg}
	case "LINDEX":
		cmd = &Lindex{strg: p.strg}
	case "LLEN":
		cmd = &Llen{strg: p.strg}
	case "LPOP":
		cmd = &Lpop{strg: p.strg}
	case "LPUSH":
		cmd = &Lpush{strg: p.strg}
	case "LRANGE":
		cmd = &Lrange{strg: p.strg}
	case "LREM":
		cmd = &Lrem{strg: p.strg}
	case "TTL":
		cmd = &TTL{strg: p.strg, clck: p.clck}
	case "TYPE":
		cmd = &Type{strg: p.strg}
	case "PING":
		cmd = &Ping{}
	default:
		return nil, nil, ErrCommandNotFound
	}

	return cmd, args[1:], nil
}

func (p *Parser) extractArgs(val string) []string {
	args := make([]string, 0)
	var inQuote bool
	var buf bytes.Buffer
	for _, r := range val {
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
