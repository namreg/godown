package command

import (
	"strconv"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/pkg/errors"
)

var errListOutOfRange = errors.New("lrange: out of range")

func init() {
	commands["LRANGE"] = new(Lrange)
}

//Lrange is the LRANGE command
type Lrange struct{}

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

//ValidateArgs implements ValidateArgs of Command interface
func (c *Lrange) ValidateArgs(args ...string) error {
	if len(args) != 3 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Lrange) Execute(strg storage.Storage, args ...string) Result {
	value, err := strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilResult{}
		}
		return ErrResult{err}
	}

	if value.Type() != storage.ListDataType {
		return ErrResult{ErrWrongTypeOp}
	}

	list := value.Data().([]string)

	start, stop, err := c.extractStartStopIndexes(list, args)
	if err != nil {
		if err == errListOutOfRange {
			return NilResult{}
		}
		return ErrResult{err}
	}

	rng := list[start:stop]
	if len(rng) == 0 {
		return NilResult{}
	}

	return SliceResult{rng}
}

func (c *Lrange) extractStartStopIndexes(list, args []string) (int, int, error) {
	start, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, errors.New("start should be an integer")
	}

	stop, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, errors.New("stop should be an integer")
	}

	if start < 0 {
		if start = len(list) + start; start < 0 {
			start = 0
		}
	} else if start > len(list)-1 {
		return 0, 0, errListOutOfRange
	}

	if stop < 0 {
		stop = len(list) + stop
	} else if stop > len(list) {
		stop = len(list)
	}

	return start, stop + 1, nil // +1 due to the range operator of slice
}
