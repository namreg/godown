package resp

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/tidwall/redcon"

	"github.com/namreg/godown/internal/api"
)

// commandExecutor executes commands.
type commandExecutor interface {
	// ExecuteCommand executes command request.
	ExecuteCommand(ctx context.Context, req *api.ExecuteCommandRequest) (*api.ExecuteCommandResponse, error)
}

// Server is a server that works over Redis Serialization Protocol.
type Server struct {
	e commandExecutor
}

// New creates a new RESP server.
func New(e commandExecutor) *Server {
	return &Server{e: e}
}

// Start starts a server.
func (s *Server) Start(hostPort string) error {
	return redcon.ListenAndServeNetwork("tcp", hostPort, s.handle, nil, nil)
}

func (s *Server) handle(conn redcon.Conn, cmd redcon.Command) {
	cmds := append([]redcon.Command{cmd}, conn.ReadPipeline()...)
	for _, cmd := range cmds {
		buf := bytes.Buffer{}
		for _, arg := range cmd.Args {
			buf.Write(arg)
			buf.WriteRune(' ')
		}

		req := &api.ExecuteCommandRequest{Command: buf.String()}
		resp, err := s.e.ExecuteCommand(context.Background(), req)
		if err != nil {
			conn.WriteError(fmt.Sprintf("could not execute command: %v", err.Error()))
		} else {
			s.writeResponse(conn, resp)
		}
	}
}

func (s *Server) writeResponse(conn redcon.Conn, resp *api.ExecuteCommandResponse) {
	switch resp.Reply {
	case api.OkCommandReply:
		conn.WriteString("OK")
	case api.NilCommandReply:
		conn.WriteNull()
	case api.RawStringCommandReply, api.StringCommandReply:
		conn.WriteBulkString(resp.Item)
	case api.IntCommandReply:
		n, err := strconv.Atoi(resp.Item)
		if err != nil {
			conn.WriteError(err.Error())
			break
		}
		conn.WriteInt(n)
	case api.ErrCommandReply:
		conn.WriteError(resp.Item)
	case api.SliceCommandReply:
		conn.WriteArray(len(resp.Items))
		for _, item := range resp.Items {
			conn.WriteString(item)
		}
	default:
		conn.WriteError(fmt.Sprintf("unexpected reply: %s", resp.Reply.String()))
	}
}
