package client

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/namreg/godown/internal/api"
)

//ScalarResult is a scalar result like string, int, etc.
type ScalarResult struct {
	val *string
	err error
}

//Err returns an error.
func (sr ScalarResult) Err() error {
	return sr.err
}

//IsNil returns true if the underlying value is nil.
func (sr ScalarResult) IsNil() bool {
	return sr.val == nil
}

//Val returns a string representation.
func (sr ScalarResult) Val() (string, error) {
	if sr.err != nil {
		return "", sr.err
	}
	if sr.IsNil() {
		return "", nil
	}
	return *sr.val, nil
}

//Int64 converts scalar value to the int64.
func (sr ScalarResult) Int64() (int64, error) {
	if sr.err != nil {
		return 0, sr.err
	}
	if sr.IsNil() {
		return 0, nil
	}
	return strconv.ParseInt(*sr.val, 10, 64)
}

func newScalarResult(resp *api.ExecuteCommandResponse) ScalarResult {
	res := ScalarResult{}
	switch resp.GetReply() {
	case api.NilCommandReply:
	case api.ErrCommandReply:
		res.err = errors.New(resp.Item)
	case api.StringCommandReply:
		res.val = &resp.Item
	default:
		res.err = fmt.Errorf("unexpected reply: %v", resp.Reply)
	}
	return res
}
