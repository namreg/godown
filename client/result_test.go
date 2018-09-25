package client

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScalarResult_Err(t *testing.T) {
	res := ScalarResult{err: errors.New("error")}
	assert.Equal(t, errors.New("error"), res.Err())
}

func TestScalarResult_IsNil(t *testing.T) {
	tests := []struct {
		name   string
		result ScalarResult
		want   bool
	}{
		{"value_nil_and_err_nil", ScalarResult{}, true},
		{"value_nil_and_err_not_nil", ScalarResult{err: errors.New("error")}, false},
		{"value_not_nil_and_err_not_nil", ScalarResult{val: stringToPtr("vale"), err: errors.New("error")}, false},
		{"value_not_nil_and_err_nil", ScalarResult{val: stringToPtr("vale")}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.result.IsNil())
	}
}

func TestScalarResult_Val(t *testing.T) {
	tests := []struct {
		name    string
		result  ScalarResult
		wantVal string
		wantErr error
	}{
		{"val_nil_and_err_not_nil", ScalarResult{err: errors.New("error")}, "", errors.New("error")},
		{"val_nil_and_err_nil", ScalarResult{}, "", nil},
		{"val_not_nil_and_err_nil", ScalarResult{val: stringToPtr("val")}, "val", nil},
	}

	for _, tt := range tests {
		val, err := tt.result.Val()
		assert.Equal(t, tt.wantVal, val)
		assert.Equal(t, tt.wantErr, err)
	}
}

func TestScalarResult_Int64(t *testing.T) {
	tests := []struct {
		name    string
		result  ScalarResult
		wantVal int64
		wantErr error
	}{
		{"val_nil_and_err_not_nil", ScalarResult{err: errors.New("error")}, 0, errors.New("error")},
		{"val_nil_and_err_nil", ScalarResult{}, 0, nil},
		{"val_not_nil_and_err_nil", ScalarResult{val: stringToPtr("10")}, 10, nil},
		{"val_is_not_int64", ScalarResult{val: stringToPtr("string")}, 0, &strconv.NumError{
			Func: "ParseInt",
			Num:  "string",
			Err:  strconv.ErrSyntax,
		}},
	}

	for _, tt := range tests {
		ival, err := tt.result.Int64()
		assert.Equal(t, tt.wantVal, ival)
		assert.Equal(t, tt.wantErr, err)
	}
}

func TestStatusResult_Err(t *testing.T) {
	res := StatusResult{err: errors.New("error")}
	assert.Equal(t, errors.New("error"), res.Err())
}

func TestListResult_Err(t *testing.T) {
	res := ListResult{err: errors.New("error")}
	assert.Equal(t, errors.New("error"), res.Err())
}

func TestListResult_IsNil(t *testing.T) {
	tests := []struct {
		name   string
		result ListResult
		want   bool
	}{
		{"value_nil_and_err_nil", ListResult{}, true},
		{"value_nil_and_err_not_nil", ListResult{err: errors.New("error")}, false},
		{"value_not_nil_and_err_not_nil", ListResult{val: []string{"val"}, err: errors.New("error")}, false},
		{"value_not_nil_and_err_nil", ListResult{val: []string{"val"}}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.result.IsNil())
	}
}

func TestListResult_Val(t *testing.T) {
	tests := []struct {
		name    string
		result  ListResult
		wantVal []string
		wantErr error
	}{
		{"val_nil_and_err_not_nil", ListResult{err: errors.New("error")}, nil, errors.New("error")},
		{"val_nil_and_err_nil", ListResult{}, nil, nil},
		{"val_not_nil_and_err_nil", ListResult{val: []string{"val"}}, []string{"val"}, nil},
	}

	for _, tt := range tests {
		val, err := tt.result.Val()
		assert.Equal(t, tt.wantVal, val)
		assert.Equal(t, tt.wantErr, err)
	}
}
