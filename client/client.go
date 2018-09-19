package client

import (
	"context"
	"fmt"
	"strconv"
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

//Set sets a new value at the given key.
func (c *Client) Set(key, value string) StatusResult {
	return c.set(context.Background(), key, value)
}

//SetWithContext similar to Set but with the context.
func (c *Client) SetWithContext(ctx context.Context, key, value string) StatusResult {
	return c.set(ctx, key, value)
}

func (c *Client) set(ctx context.Context, key, value string) StatusResult {
	req := c.newExecuteRequest("SET", key, value)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return StatusResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newStatusResult(resp)
}

//Del deletes the given key.
func (c *Client) Del(key string) StatusResult {
	return c.del(context.Background(), key)
}

//DelWithContext similar to Del but with context.
func (c *Client) DelWithContext(ctx context.Context, key string) StatusResult {
	return c.del(ctx, key)
}

func (c *Client) del(ctx context.Context, key string) StatusResult {
	req := c.newExecuteRequest("DEL", key)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return StatusResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newStatusResult(resp)
}

//Expire sets expiration of the given key as `now + secs`.
func (c *Client) Expire(key string, secs int) StatusResult {
	return c.expire(context.Background(), key, secs)
}

//ExpireWithContext similar to Expire but with context.
func (c *Client) ExpireWithContext(ctx context.Context, key string, secs int) StatusResult {
	return c.expire(ctx, key, secs)
}

func (c *Client) expire(ctx context.Context, key string, secs int) StatusResult {
	req := c.newExecuteRequest("EXPIRE", key, strconv.Itoa(secs))
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return StatusResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newStatusResult(resp)
}

//GetBit returns the bit value at the offset in the string stored at key.
func (c *Client) GetBit(key string, offset uint64) ScalarResult {
	return c.getBit(context.Background(), key, offset)
}

//GetBitWithContext similar to GetBit but with context.
func (c *Client) GetBitWithContext(ctx context.Context, key string, offset uint64) ScalarResult {
	return c.getBit(ctx, key, offset)
}

func (c *Client) getBit(ctx context.Context, key string, offset uint64) ScalarResult {
	req := c.newExecuteRequest("GETBIT", key, strconv.FormatUint(offset, 10))
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return ScalarResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newScalarResult(resp)
}

//HGet returns the value associated with field in the hash stored at key.
func (c *Client) HGet(key, field string) ScalarResult {
	return c.hget(context.Background(), key, field)
}

//HGetWithContext similar to HGet but with context.
func (c *Client) HGetWithContext(ctx context.Context, key, field string) ScalarResult {
	return c.hget(ctx, key, field)
}

func (c *Client) hget(ctx context.Context, key, field string) ScalarResult {
	req := c.newExecuteRequest("HGET", key, field)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return ScalarResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newScalarResult(resp)
}

//HKeys returns all keys of the map stored at the given key.
func (c *Client) HKeys(key string) ListResult {
	return c.hkeys(context.Background(), key)
}

//HKeysWithContext similar to Hkeys by with context.
func (c *Client) HKeysWithContext(ctx context.Context, key string) ListResult {
	return c.hkeys(ctx, key)
}

func (c *Client) hkeys(ctx context.Context, key string) ListResult {
	req := c.newExecuteRequest("HKEYS", key)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return ListResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newListResult(resp)
}

//HSet sets a field in hash at the given key.
func (c *Client) HSet(key, field, value string) StatusResult {
	return c.hset(context.Background(), key, field, value)
}

//HSetWithContext similar to HSet but with context.
func (c *Client) HSetWithContext(ctx context.Context, key, field, value string) StatusResult {
	return c.hset(ctx, key, field, value)
}

func (c *Client) hset(ctx context.Context, key, field, value string) StatusResult {
	req := c.newExecuteRequest("HSET", key, field, value)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return StatusResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newStatusResult(resp)
}

//HVals returns values of a hash stored at the given key.
func (c *Client) HVals(key string) ListResult {
	return c.hvals(context.Background(), key)
}

//HValsWithContext similar to HVals but with context.
func (c *Client) HValsWithContext(ctx context.Context, key string) ListResult {
	return c.hvals(ctx, key)
}

func (c *Client) hvals(ctx context.Context, key string) ListResult {
	req := c.newExecuteRequest("HVALS", key)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return ListResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newListResult(resp)
}

//Keys returns all keys that matched to the given pattern.
func (c *Client) Keys(pattern string) ListResult {
	return c.keys(context.Background(), pattern)
}

//KeysWithContext similar to Keys but with context.
func (c *Client) KeysWithContext(ctx context.Context, pattern string) ListResult {
	return c.keys(ctx, pattern)
}

func (c *Client) keys(ctx context.Context, pattern string) ListResult {
	req := c.newExecuteRequest("KEYS", pattern)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return ListResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newListResult(resp)
}

//LIndex returns a value at the index in the list stored at the given key.
func (c *Client) LIndex(key string, index int) ScalarResult {
	return c.lindex(context.Background(), key, index)
}

//LIndexWithContext similar to LIndex but with context.
func (c *Client) LIndexWithContext(ctx context.Context, key string, index int) ScalarResult {
	return c.lindex(ctx, key, index)
}

func (c *Client) lindex(ctx context.Context, key string, index int) ScalarResult {
	req := c.newExecuteRequest("LINDEX", key, strconv.Itoa(index))
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return ScalarResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newScalarResult(resp)
}

//LLen returns a number of the elements in the list stored at the given key.
func (c *Client) LLen(key string) ScalarResult {
	return c.llen(context.Background(), key)
}

//LLenWithContext similar to LLen but with context.
func (c *Client) LLenWithContext(ctx context.Context, key string) ScalarResult {
	return c.llen(ctx, key)
}

func (c *Client) llen(ctx context.Context, key string) ScalarResult {
	req := c.newExecuteRequest("LLEN", key)
	resp, err := c.executor.ExecuteCommand(ctx, req)
	if err != nil {
		return ScalarResult{err: fmt.Errorf("could not execute command: %v", err)}
	}
	return newScalarResult(resp)
}
