package command

import (
	"bytes"
	"errors"
	"strings"
	"unicode"
)

//ErrCommandNotFound means that command could not be parsed.
var ErrCommandNotFound = errors.New("command: not found")

//Parser is a parser that parses user input and creates the appropriate command.
type Parser struct {
	strg dataStore
	clck commandClock
}

//NewParser creates a new parser
func NewParser(strg dataStore, clck commandClock) *Parser {
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
	case "RPUSH":
		cmd = &Rpush{strg: p.strg}
	case "RPOP":
		cmd = &Rpop{strg: p.strg}
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
