package client

import (
	"context"
	"fmt"
	"time"

	"github.com/namreg/godown/internal/api"

	"github.com/hashicorp/go-multierror"
	"google.golang.org/grpc"
)

const connectTimeout = 100 * time.Millisecond

//Client is a client that communicates with a server.
type Client struct {
	addrs  []string
	conn   *grpc.ClientConn
	client api.GodownClient
}

//New creates a new client with the given servet addresses.
func New(addr string, addrs ...string) (*Client, error) {
	c := &Client{addrs: append([]string{addr}, addrs...)}
	if err := c.tryConnect(); err != nil {
		return nil, fmt.Errorf("could not connect to server: %v", err)
	}
	return c, nil
}

//Close closes the client.
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) tryConnect() error {
	var (
		result *multierror.Error
		err    error
		conn   *grpc.ClientConn
	)

	for addrs := c.addrs; len(addrs) > 0; addrs = addrs[1:] {
		ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
		defer cancel()

		if conn, err = grpc.DialContext(ctx, addrs[0], grpc.WithInsecure(), grpc.WithBlock()); err == nil {
			c.conn = conn
			c.client = api.NewGodownClient(c.conn)
			return nil
		}
		result = multierror.Append(result, err)
	}
	return result.ErrorOrNil()
}
