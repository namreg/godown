package command

import (
	"testing"

	"github.com/gojuno/minimock"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestOkResult_Val(t *testing.T) {
	res := OkResult{}
	assert.Nil(t, res.Val())
}

func TestErrResult_Val(t *testing.T) {
	err := errors.New("error")
	res := ErrResult{err}

	assert.Equal(t, err, res.Val())
}

func TestNilResult_Val(t *testing.T) {
	res := NilResult{}
	assert.Nil(t, res.Val())
}

func TestHelpResult_Val(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	cmd := NewCommandMock(t)
	cmd.HelpMock.Return("help message")

	res := HelpResult{cmd}
	assert.Equal(t, "help message", res.Val())
}

func TestStringResult_Val(t *testing.T) {
	res := StringResult{"string"}
	assert.Equal(t, "string", res.Val())
}

func TestIntResult_Val(t *testing.T) {
	res := IntResult{100}
	assert.Equal(t, int64(100), res.Val())
}

func TestSliceResult_Val(t *testing.T) {
	res := SliceResult{[]string{"val1", "val2"}}
	assert.Equal(t, []string{"val1", "val2"}, res.Val())
}
