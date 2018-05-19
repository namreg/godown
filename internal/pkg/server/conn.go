package server

import (
	"errors"
	"fmt"
	"net"

	"github.com/namreg/godown-v2/internal/pkg/command"
)

const (
	nilString     = "nil"
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

func (c *conn) writeCommandResult(res command.Result) {
	switch res.(type) {
	case command.OkResult:
		c.writeString(okString)
	case command.ErrResult:
		c.writeError(res.Val().(error))
	case command.NilResult:
		c.writeType(nilString)
	case command.StringResult:
		c.writeString(res.Val().(string))
	case command.HelpResult:
		c.writeMessage(res.Val().(string))
	case command.SliceResult:
		s := res.Val().([]string)
		if len(s) == 0 {
			c.writeType(nilString)
		} else {
			for i, v := range s {
				c.write(fmt.Sprintf("%d) %q", i+1, v))
				if i != len(s)-1 {
					c.write("\n")
				}
			}
			c.writeNewLine()
		}
	default:
		c.writeError(errors.New("could not recognize result"))
	}
}

func (c *conn) write(str string) {
	fmt.Fprintf(c.conn, str)
}

func (c *conn) writeNewLine() {
	fmt.Fprintf(c.conn, newLineString)
}
