package command

import (
	"strconv"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/pkg/errors"
)

func init() {
	commands["LINDEX"] = new(Lindex)
}

//Lindex is the LINDEX command
type Lindex struct{}

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

//ValidateArgs implements ValidateArgs of Command interface
func (c *Lindex) ValidateArgs(args ...string) error {
	if len(args) != 2 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Lindex) Execute(strg storage.Storage, args ...string) Result {
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

	index, err := c.parseIndex(list, args[1])
	if err != nil {
		return ErrResult{err}
	}

	if index < 0 || index > len(list)-1 {
		return NilResult{}
	}
	return StringResult{list[index]}
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
