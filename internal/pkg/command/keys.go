package command

import (
	"regexp"
	"strings"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/pkg/errors"
)

func init() {
	cmd := new(Keys)
	commands[cmd.Name()] = cmd
}

//Keys is the Keys command
type Keys struct{}

//Name implements Name of Command interface
func (c *Keys) Name() string {
	return "KEYS"
}

//Help implements Help of Command interface
func (c *Keys) Help() string {
	return `Usage: KEYS pattern
Find all keys matching the given pattern.`
}

//ValidateArgs implements ValidateArgs of Command interface
func (c *Keys) ValidateArgs(args ...string) error {
	if len(args) != 1 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Keys) Execute(strg storage.Storage, args ...string) Result {
	if err := c.ValidateArgs(args...); err != nil {
		return ErrResult{err}
	}

	keys, err := strg.Keys()
	if err != nil {
		return ErrResult{err}
	}

	re, err := c.compilePattern(args[0])
	if err != nil {
		return ErrResult{err}
	}

	keyNames := make([]string, 0, len(keys))
	for _, k := range keys {
		sk := string(k)
		if re.MatchString(sk) {
			keyNames = append(keyNames, sk)
		}
	}
	return SliceResult{keyNames}
}

func (c *Keys) compilePattern(input string) (*regexp.Regexp, error) {
	pattern := strings.Replace(input, "*", ".+?", -1)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("invalid pattern syntax")
	}
	return re, nil
}
