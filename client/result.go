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

//IsNil returns true if the underlying value and error is empty.
func (sr ScalarResult) IsNil() bool {
	return sr.val == nil && sr.err == nil
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
	case api.StringCommandReply, api.IntCommandReply, api.RawStringCommandReply:
		res.val = &resp.Item
	default:
		res.err = fmt.Errorf("unexpected reply: %v", resp.GetReply())
	}
	return res
}

//StatusResult is used when a command can respond only with success or with error.
type StatusResult struct {
	err error
}

//Err returns an error.
func (sr StatusResult) Err() error {
	return sr.err
}

func newStatusResult(resp *api.ExecuteCommandResponse) StatusResult {
	res := StatusResult{}
	switch resp.GetReply() {
	case api.OkCommandReply:
	case api.ErrCommandReply:
		res.err = errors.New(resp.Item)
	default:
		res.err = fmt.Errorf("unexpected reply: %v", resp.GetReply())
	}
	return res
}

//ListResult is used when command respond with slice.
type ListResult struct {
	val []string
	err error
}

//Err returns an error.
func (lr ListResult) Err() error {
	return lr.err
}

//IsNil returns true if the underlaying value and error is empty.
func (lr ListResult) IsNil() bool {
	return lr.err == nil && lr.val == nil
}

//Val returns the underlying slice if no error. Otherwise, the error will be returned.
func (lr ListResult) Val() ([]string, error) {
	if err := lr.Err(); err != nil {
		return nil, err
	}
	return lr.val, nil
}

func newListResult(resp *api.ExecuteCommandResponse) ListResult {
	res := ListResult{}
	switch resp.GetReply() {
	case api.NilCommandReply:
	case api.SliceCommandReply:
		res.val = resp.GetItems()
	case api.ErrCommandReply:
		res.err = errors.New(resp.GetItem())
	default:
		res.err = fmt.Errorf("unexpected reply: %v", resp.GetReply())
	}
	return res
}
