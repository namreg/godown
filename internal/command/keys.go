package command

import (
	"errors"
	"regexp"
	"strings"
)

//Keys is the Keys command
type Keys struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Keys) Name() string {
	return "KEYS"
}

//Help implements Help of Command interface
func (c *Keys) Help() string {
	return `Usage: KEYS pattern
Find all keys matching the given pattern.`
}

//Execute implements Execute of Command interface
func (c *Keys) Execute(args ...string) Reply {
	if len(args) != 1 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	keys, err := c.strg.Keys()
	if err != nil {
		return ErrReply{Value: err}
	}

	re, err := c.compilePattern(args[0])
	if err != nil {
		return ErrReply{Value: err}
	}

	keyNames := make([]string, 0)
	for _, k := range keys {
		sk := string(k)
		if re.MatchString(sk) {
			keyNames = append(keyNames, sk)
		}
	}
	return SliceReply{Value: keyNames}
}

func (c *Keys) compilePattern(input string) (*regexp.Regexp, error) {
	pattern := strings.Replace(input, "*", ".+?", -1)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("invalid pattern syntax")
	}
	return re, nil
}
