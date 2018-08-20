package command

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestOkResult_Val(t *testing.T) {
	res := OkResult{}
	assert.Nil(t, res.Val())
}

func TestErrResult_Val(t *testing.T) {
	err := errors.New("error")
	res := ErrResult{Value: err}

	assert.Equal(t, err, res.Val())
}

func TestNilResult_Val(t *testing.T) {
	res := NilResult{}
	assert.Nil(t, res.Val())
}

func TestRawStringResult_Val(t *testing.T) {
	res := RawStringResult{Value: "help message"}
	assert.Equal(t, "help message", res.Val())
}

func TestStringResult_Val(t *testing.T) {
	res := StringResult{Value: "string"}
	assert.Equal(t, "string", res.Val())
}

func TestIntResult_Val(t *testing.T) {
	res := IntResult{Value: 100}
	assert.Equal(t, int64(100), res.Val())
}

func TestSliceResult_Val(t *testing.T) {
	res := SliceResult{Value: []string{"val1", "val2"}}
	assert.Equal(t, []string{"val1", "val2"}, res.Val())
}
