package command

import (
	"errors"
	"strconv"

	"github.com/namreg/godown/internal/storage"
)

//Lindex is the LINDEX command
type Lindex struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Lindex) Name() string {
	return "LINDEX"
}

//Help implements Help of Command interface
func (c *Lindex) Help() string {
	return `LINDEX key index
Returns the element at index index in the list stored at key. 
The index is zero-based, so 0 means the first element, 1 the second element and so on. 
Negative indices can be used to designate elements starting at the tail of the list.`
}

//Execute implements Execute of Command interface
func (c *Lindex) Execute(args ...string) Reply {
	if len(args) != 2 {
		return ErrReply{Value: ErrWrongArgsNumber}
	}

	value, err := c.strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilReply{}
		}
		return ErrReply{Value: err}
	}

	if value.Type() != storage.ListDataType {
		return ErrReply{Value: ErrWrongTypeOp}
	}

	list := value.Data().([]string)

	index, err := c.parseIndex(list, args[1])
	if err != nil {
		return ErrReply{Value: err}
	}

	if index < 0 || index > len(list)-1 {
		return NilReply{}
	}
	return StringReply{Value: list[index]}
}

func (c *Lindex) parseIndex(list []string, index string) (int, error) {
	i, err := strconv.Atoi(index)
	if err != nil {
		return 0, errors.New("index should be an integer")
	}
	if i < 0 {
		return len(list) + i, nil
	}
	return i, nil
}
