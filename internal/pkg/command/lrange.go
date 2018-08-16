package command

import (
	"errors"
	"strconv"

	"github.com/namreg/godown-v2/internal/pkg/storage"
)

var errListOutOfRange = errors.New("lrange: out of range")

//Lrange is the LRANGE command
type Lrange struct {
	strg commandStorage
}

//Name implements Name of Command interface
func (c *Lrange) Name() string {
	return "LRANGE"
}

//Help implements Help of Command interface
func (c *Lrange) Help() string {
	return `Usage: LRANGE key start stop
Returns the specified elements of the list stored at key. 
The offsets start and stop are zero-based indexes, 
with 0 being the first element of the list (the head of the list), 1 being the next element and so on.`
}

//Execute implements Execute of Command interface
func (c *Lrange) Execute(args ...string) Result {
	if len(args) != 3 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	c.strg.RLock()
	value, err := c.strg.Get(storage.Key(args[0]))
	c.strg.RUnlock()
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{Value: err}
	}

	if value.Type() != storage.ListDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}

	list := value.Data().([]string)

	start, stop, err := c.extractStartStopIndexes(len(list), args)
	if err != nil {
		if err == errListOutOfRange {
			return NilResult{}
		}
		return ErrResult{Value: err}
	}

	rng := list[start:stop]
	if len(rng) == 0 {
		return NilResult{}
	}

	return SliceResult{Value: rng}
}

func (c *Lrange) extractStartStopIndexes(len int, args []string) (int, int, error) {
	start, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, errors.New("start should be an integer")
	}

	stop, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, errors.New("stop should be an integer")
	}

	if start > len-1 {
		return 0, 0, errListOutOfRange
	}

	if stop > len {
		return start, len, nil
	}

	if start < 0 {
		if start = len - 1 + start; start < 0 {
			start = 0
		}
	}

	if stop < 0 {
		if stop = len + stop; stop < 0 {
			stop = len - 1
		}
	}

	return start, stop + 1, nil // +1 due to the range operator of slice
}
