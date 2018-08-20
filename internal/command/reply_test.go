package command

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestOkReply_Val(t *testing.T) {
	res := OkReply{}
	assert.Nil(t, res.Val())
}

func TestErrReply_Val(t *testing.T) {
	err := errors.New("error")
	res := ErrReply{Value: err}

	assert.Equal(t, err, res.Val())
}

func TestNilReply_Val(t *testing.T) {
	res := NilReply{}
	assert.Nil(t, res.Val())
}

func TestRawStringReply_Val(t *testing.T) {
	res := RawStringReply{Value: "help message"}
	assert.Equal(t, "help message", res.Val())
}

func TestStringReply_Val(t *testing.T) {
	res := StringReply{Value: "string"}
	assert.Equal(t, "string", res.Val())
}

func TestIntReply_Val(t *testing.T) {
	res := IntReply{Value: 100}
	assert.Equal(t, int64(100), res.Val())
}

func TestSliceReply_Val(t *testing.T) {
	res := SliceReply{Value: []string{"val1", "val2"}}
	assert.Equal(t, []string{"val1", "val2"}, res.Val())
}
