package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	cmd := new(Lrem)
	commands[cmd.Name()] = cmd
}

//Lrem is the LREM command
type Lrem struct{}

//Name implements Name of Command interface
func (c *Lrem) Name() string {
	return "LREM"
}

//Help implements Help of Command interface
func (c *Lrem) Help() string {
	return `Usage: LREM key value
Removes all occurrences of elements equal to value from the list stored at key.`
}

//Execute implements Execute of Command interface
func (c *Lrem) Execute(strg storage.Storage, args ...string) Result {
	if len(args) != 2 {
		return ErrResult{Value: ErrWrongArgsNumber}
	}

	strg.Lock()
	defer strg.Unlock()

	key := storage.Key(args[0])

	val, err := strg.Get(key)
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return OkResult{}
		}
		return ErrResult{Value: err}
	}

	if val.Type() != storage.ListDataType {
		return ErrResult{Value: ErrWrongTypeOp}
	}

	list := val.Data().([]string)
	newList := list[:0]

	for _, val := range list {
		if val != args[1] {
			newList = append(newList, val)
		}
	}

	if len(newList) == 0 {
		if err = strg.Del(key); err != nil {
			return ErrResult{Value: err}
		}
		return OkResult{}
	}

	if err = strg.Put(key, storage.NewListValue(newList)); err != nil {
		return ErrResult{Value: err}
	}

	return OkResult{}
}
