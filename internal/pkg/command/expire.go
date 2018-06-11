package command

import (
	"strconv"
	"time"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/pkg/errors"
)

func init() {
	cmd := new(Expire)
	commands[cmd.Name()] = cmd
}

//Expire is the Expire command
type Expire struct{}

//Name implements Name of Command interface
func (c *Expire) Name() string {
	return "EXPIRE"
}

//Help implements Help of Command interface
func (c *Expire) Help() string {
	return `Usage: EXPIRE key seconds
Set a timeout on key. After the timeout has expired, the key will automatically be deleted.`
}

//ValidateArgs implements ValidateArgs of Command interface
func (c *Expire) ValidateArgs(args ...string) error {
	if len(args) != 2 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Expire) Execute(strg storage.Storage, args ...string) Result {
	if err := c.ValidateArgs(args...); err != nil {
		return ErrResult{err}
	}
	secs, err := strconv.Atoi(args[1])
	if err != nil {
		return ErrResult{errors.New("seconds should be integer")}
	}
	if secs < 0 {
		return ErrResult{errors.New("seconds should be positive")}
	}
	setter := func(old *storage.Value) (*storage.Value, error) {
		if old == nil {
			return nil, nil
		}
		old.SetTTL(time.Now().Add(time.Duration(secs) * time.Second))
		return old, nil
	}
	if err := strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}
