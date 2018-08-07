package server

import (
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/pkg/command"
	"github.com/stretchr/testify/suite"
)

type connTestSuite struct {
	suite.Suite
	writedBytes []byte
	conn        *conn
}

func (ts *connTestSuite) SetupTest() {
	mock := NewConnMock(ts.T())
	mock.WriteFunc = func(p []byte) (r int, r1 error) {
		ts.writedBytes = append(ts.writedBytes, p...)
		return len(p), nil
	}
	ts.conn = newConn(mock)
	ts.writedBytes = make([]byte, 0)
}

func (ts *connTestSuite) TestClose() {
	mock := NewConnMock(ts.T())
	mock.CloseMock.Return(nil)

	conn := newConn(mock)
	conn.Close()
}

func (ts *connTestSuite) Test_writeWelcomeMessage() {
	expected := []byte("\nWelcome to godown. Version is 000\ngodown > ")

	ts.conn.writeWelcomeMessage()

	ts.Equal(expected, ts.writedBytes)
}

func (ts *connTestSuite) Test_writeNil() {
	expected := []byte("(nil)\ngodown > ")

	ts.conn.writeNil()

	ts.Equal(expected, ts.writedBytes)
}

func (ts *connTestSuite) Test_writeString() {
	expected := []byte("(string): writed string\ngodown > ")

	ts.conn.writeString("writed string")

	ts.Equal(expected, ts.writedBytes)
}

func (ts *connTestSuite) Test_writeInt() {
	expected := []byte("(integer): 25\ngodown > ")

	ts.conn.writeInt(25)

	ts.Equal(expected, ts.writedBytes)
}

func (ts *connTestSuite) Test_writeError() {
	expected := []byte("(error): writed error\ngodown > ")

	ts.conn.writeError(errors.New("writed error"))

	ts.Equal(expected, ts.writedBytes)
}

func (ts *connTestSuite) Test_writeCommandResult() {
	cmdMock := command.NewCommandMock(ts.T())
	cmdMock.HelpMock.Return("help message")

	testCases := []struct {
		name string
		res  command.Result
		want []byte
	}{
		{"ok", command.OkResult{}, []byte("(string): OK\ngodown > ")},
		{"error", command.ErrResult{Err: errors.New("test error")}, []byte("(error): test error\ngodown > ")},
		{"nil", command.NilResult{}, []byte("(nil)\ngodown > ")},
		{"string", command.StringResult{Str: "test string"}, []byte("(string): test string\ngodown > ")},
		{"int", command.IntResult{Value: 10}, []byte("(integer): 10\ngodown > ")},
		{"help", command.HelpResult{Cmd: cmdMock}, []byte("help message\ngodown > ")},
		{"slice", command.SliceResult{Value: []string{"value 1", "value 2"}}, []byte("1) \"value 1\"\n2) \"value 2\"\ngodown > ")},
		{"empty_slice", command.SliceResult{Value: []string{}}, []byte("(nil)\ngodown > ")},
		{"unknown", nil, []byte("(error): could not recognize result\ngodown > ")},
	}

	for _, tc := range testCases {
		ts.writedBytes = []byte{}

		ts.conn.writeCommandResult(tc.res)

		ts.Equal(tc.want, ts.writedBytes)
	}
}

func (ts *connTestSuite) Test_write() {
	expected := []byte("random string")

	ts.conn.write("random string")

	ts.Equal(expected, ts.writedBytes)
}

func (ts *connTestSuite) Test_writePrompt() {
	expected := []byte("\ngodown > ")

	ts.conn.writePrompt()

	ts.Equal(expected, ts.writedBytes)
}

func Test_connTestSuite(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	suite.Run(t, new(connTestSuite))
}
