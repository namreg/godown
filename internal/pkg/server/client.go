package server

import (
	"fmt"
	"net"
)

type client struct {
	conn net.Conn
}

func newClient(conn net.Conn) *client {
	return &client{conn}
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) respondWithCommandWaiting() {
	fmt.Fprintf(c.conn, "\n> ")
}
