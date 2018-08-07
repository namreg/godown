package command

import (
	"errors"
	"strconv"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(GetBit)
	commands[cmd.Name()] = cmd
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

//Execute implements Execute of Command interface
func (c *GetBit) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 2 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	offset, err := c.parseOffset(args)
	if err != nil {
		return ErrResult{Value: err}
	}

	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return IntResult{Value: 0}
		}
		return ErrResult{Value: err}
	}

	if value.Type() != storage.BitMapDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}

	vals := value.Data().([]uint64)
	idx := c.resolveIndex(offset)

	if idx > uint64(len(vals)-1) {
		return IntResult{Value: 0}
	}

	if vals[idx]&(1<<(offset%64)) != 0 {
		return IntResult{Value: 1}
	}
	return IntResult{Value: 0}
}

func (c *GetBit) parseOffset(args []string) (uint64, error) {
	offset, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return 0, errors.New("invalid offset")
	}
	return offset, nil
}

func (c *GetBit) resolveIndex(offset uint64) uint64 {
	var idx uint64
	if offset > 63 {
		if offset == 64 {
			idx = 1
		} else {
			idx = offset % 64
		}
	}
	return idx
}
