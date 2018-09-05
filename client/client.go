package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/namreg/godown/internal/api"

	"github.com/hashicorp/go-multierror"
	"google.golang.org/grpc"
)

const connectTimeout = 100 * time.Millisecond

//go:generate minimock -i github.com/namreg/godown/client.executor -o ./
type executor interface {
	ExecuteCommand(context.Context, *api.ExecuteCommandRequest, ...grpc.CallOption) (*api.ExecuteCommandResponse, error)
}

//Client is a client that communicates with a server.
type Client struct {
	addrs    []string
	conn     *grpc.ClientConn
	executor executor
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
			c.executor = api.NewGodownClient(c.conn)
			return nil
		}
		result = multierror.Append(result, err)
	}
	return result.ErrorOrNil()
}

func (c *Client) newExecuteRequest(cmd string, args ...string) *api.ExecuteCommandRequest {
	args = append([]string{cmd}, args...)
	return &api.ExecuteCommandRequest{
		Command: strings.Join(args, " "),
	}
}

//Get gets a value at the given key.
func (c *Client) Get(key string) ScalarResult {
	return c.get(context.Background(), key)
}

//GetWithContext similar to Get but with the context.
func (c *Client) GetWithContext(ctx context.Context, key string) ScalarResult {
	return c.get(ctx, key)
}

func (c *Client) get(ctx context.Context, key string) ScalarResult {
	req := c.newExecuteRequest("GET", key)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return ScalarResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newScalarResult(resp)
}
