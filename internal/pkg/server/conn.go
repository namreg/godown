package server

import (
	"fmt"
	"net"

	"github.com/namreg/godown-v2/internal/pkg/command"
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

const (
	nilString     = "(nil)"
	okString      = "OK"
	newLineString = "\n> "
)

type conn struct {
	conn net.Conn
}

func newConn(c net.Conn) *conn {
	return &conn{c}
}

func (c *conn) Close() {
	c.conn.Close()
}

func (c *conn) writeWelcomeMessage() {
	c.writeMessage("\nWelcome to godown. Version is 000")
}

func (c *conn) writeMessage(msg string) {
	fmt.Fprint(c.conn, msg)
	c.writeNewLine()
}

func (c *conn) writeType(str string) {
	fmt.Fprintf(c.conn, "(%s)", str)
	c.writeNewLine()
}

func (c *conn) writeString(str string) {
	fmt.Fprintf(c.conn, "(string): %s", str)
	c.writeNewLine()
}

func (c *conn) writeError(err error) {
	fmt.Fprintf(c.conn, "(error): %s", err.Error())
	c.writeNewLine()
}

func (c *conn) writeCommandResult(res command.Resulter) {
	if _, ok := res.(command.EmptyResult); ok {
		c.writeMessage(nilString)
		return
	}
	if res, ok := res.(command.UsageResult); ok {
		c.writeMessage(res.Value().(string))
		return
	}
	switch {
	case res.Err() == nil && res.Value() == nil:
		c.writeMessage(okString)
	case res.Err() != nil:
		c.writeError(res.Err())
	default:
		switch res.Value().(type) {
		case string:
			c.writeString(res.Value().(string))
		case storage.Key:
			c.writeType(res.Value().(storage.Key).DataType())
		}
	}
}

func (c *conn) writeNewLine() {
	fmt.Fprintf(c.conn, newLineString)
}
