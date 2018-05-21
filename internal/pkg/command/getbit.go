package command

import (
	"strconv"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/pkg/errors"
)

func init() {
	commands["GETBIT"] = new(GetBit)
}

//GetBit is the GetBit command
type GetBit struct{}

//Name implements Name of Command interface
func (c *GetBit) Name() string {
	return "GETBIT"
}

//Help implements Help of Command interface
func (c *GetBit) Help() string {
	return `Usage: GETBIT key offset
Returns the bit value at offset in the string value stored at key.`
}

//ArgsNumber implements ArgsNumber of Command interface
func (c *GetBit) ArgsNumber() int {
	return 2
}

//Execute implements Execute of Command interface
func (c *GetBit) Execute(strg storage.Storage, args ...string) Result {
	offset, err := c.parseOffset(args)
	if err != nil {
		return ErrResult{err}
	}

	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return IntResult{0}
		}
		return ErrResult{err}
	}

	intValue := value.Data().(int64)

	if intValue&(1<<offset) != 0 {
		return IntResult{1}
	}
	return IntResult{0}
}

func (c *GetBit) parseOffset(args []string) (uint64, error) {
	offset, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return 0, err
	}
	if offset < 0 {
		return 0, errors.New("offset should not be negative")
	}
	return offset, nil
}
