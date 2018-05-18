package command

import (
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

func init() {
	commands["GET"] = new(Get)
}

//Get is the GET command
type Get struct{}

//Name implements Name of Command interface
func (c *Get) Name() string {
	return "GET"
}

//Help implements Help of Command interface
func (c *Get) Help() string {
	return `Usage: GET key
Get the value by key. 
If provided key does not exist NIL will be returned.`
}

//ArgsNumber implements ArgsNumber of Command interface
func (c *Get) ArgsNumber() int {
	return 1
}

//Execute implements Execute of Command interface
func (c *Get) Execute(strg storage.Storage, args ...string) Resulter {
	key := storage.NewStringKey(args[0])
	val, err := strg.Get(key)
	if err != nil {
		return Result{err: err}
	}

	if val == nil {
		return EmptyResult{}
	}
	return Result{value: val}
}
