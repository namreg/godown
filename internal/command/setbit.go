package command

import (
	"errors"
	"strconv"

	"github.com/namreg/godown-v2/internal/storage"
)

//SetBit is the SetBit command
type SetBit struct {
	strg commandStorage
}

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
func (c *SetBit) Execute(args ...string) Result {
	if len(args) != 3 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	offset, err := c.parseOffset(args)
	if err != nil {
		return ErrResult{Value: err}
	}

	bitValue, err := c.parseValue(args)
	if err != nil {
		return ErrResult{Value: err}
	}

	key := storage.Key(args[0])

	c.strg.Lock()
	defer c.strg.Unlock()

	var value []uint64

	old, err := c.strg.Get(key)
	if err != nil && err != storage.ErrKeyNotExists {
		return ErrResult{Value: err}
	}

	if old != nil {
		if old.Type() != storage.BitMapDataType {
			return ErrResult{Value: ErrWrongTypeOp}
		}
		value = old.Data().([]uint64)
	}

	value = c.growSlice(value, offset)
	idx := c.resolveIndex(offset)

	if bitValue == 1 {
		value[idx] = value[idx] | 1<<(offset%64)
	} else {
		value[idx] = value[idx] & ^(1 << (offset % 64))
	}
	if c.isZeroSlice(value) {
		if err = c.strg.Del(key); err != nil {
			return ErrResult{Value: err}
		}
		return OkResult{}
	}
	if err = c.strg.Put(key, storage.NewBitMapValue(value)); err != nil {
		return ErrResult{Value: err}
	}
	return OkResult{}
}

func (c *SetBit) resolveIndex(offset uint64) uint64 {
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

func (c *SetBit) growSlice(sl []uint64, offset uint64) []uint64 {
	if sl == nil {
		sl = make([]uint64, 1)
	}

	maxIdx := uint64(len(sl) - 1)
	idx := c.resolveIndex(offset)

	if maxIdx >= idx {
		return sl
	}

	gsl := make([]uint64, idx+1)
	copy(gsl, sl)

	return gsl
}

func (c *SetBit) isZeroSlice(sl []uint64) bool {
	var sum uint64
	for _, v := range sl {
		sum += v
	}
	return sum == 0
}

func (c *SetBit) parseOffset(args []string) (uint64, error) {
	offset, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return 0, errors.New("invalid offset")
	}
	return offset, nil
}

func (c *SetBit) parseValue(args []string) (uint64, error) {
	bitValue, err := strconv.ParseUint(args[2], 10, 1)
	if err != nil {
		return 0, errors.New("value should be 0 or 1")
	}
	return bitValue, nil
}
