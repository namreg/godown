package command

import (
	"strconv"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/pkg/errors"
)

func init() {
	cmd := new(SetBit)
	commands[cmd.Name()] = cmd
}

//SetBit is the SetBit command
type SetBit struct{}

//Name implements Name of Command interface
func (c *SetBit) Name() string {
	return "SETBIT"
}

//Help implements Help of Command interface
func (c *SetBit) Help() string {
	return `Usage: SETBIT key offset value
Sets or clears the bit at offset in the string value stored at key.`
}

//Execute implements Execute of Command interface
func (c *SetBit) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 3 {
		return ErrResult{ErrWrongArgsNumber}
	}

	offset, err := c.parseOffset(args)
	if err != nil {
		return ErrResult{err}
	}

	bitValue, err := c.parseValue(args)
	if err != nil {
		return ErrResult{err}
	}

	setter := func(old *storage.Value) (*storage.Value, error) {
		var value int64
		if old != nil {
			if old.Type() != storage.BitMapDataType {
				return nil, ErrWrongTypeOp
			}
			value = old.Data().(int64)
		}

		if bitValue == 1 {
			value = value | 1<<offset
		} else {
			value = value & ^(1 << offset)
		}
		return storage.NewBitMapValue(value), nil
	}

	if err := strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}

func (c *SetBit) parseOffset(args []string) (uint64, error) {
	offset, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return 0, err
	}
	if offset < 0 {
		return 0, errors.New("offset should not be negative")
	}
	return offset, nil
}

func (c *SetBit) parseValue(args []string) (uint64, error) {
	bitValue, err := strconv.ParseUint(args[2], 10, 1)
	if err != nil {
		return 0, errors.New("could not parse value")
	}
	if bitValue != 0 && bitValue != 1 {
		return 0, errors.New("value should be 0 or 1")
	}
	return bitValue, nil
}
