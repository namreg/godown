package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/namreg/godown-v2/internal/api"
	"github.com/namreg/godown-v2/internal/clock"
	"github.com/namreg/godown-v2/internal/command"
	"github.com/namreg/godown-v2/internal/storage"
)

type commandParser interface {
	Parse(str string) (cmd command.Command, args []string, err error)
}

//go:generate minimock -i github.com/namreg/godown-v2/internal/server.serverStorage -o ./

//serverStorage describes storage interface that Server works with.
type serverStorage interface {
	sync.Locker
	//RLock locks storage for reading.
	RLock()
	//RUnlock undoes single RLock call.
	RUnlock()
	//Del deletes the given key.
	Del(storage.Key) error
	//AllWithTTL returns all values that have TTL.
	AllWithTTL() (map[storage.Key]*storage.Value, error)
}

//go:generate minimock -i github.com/namreg/godown-v2/internal/server.serverClock -o ./

type serverClock interface {
	//Now returns current time
	Now() time.Time
}

const defaultGCInterval = 500 * time.Millisecond

//Server represents a server that handles user requests and executes commands
type Server struct {
	parser commandParser
	strg   serverStorage
	clck   serverClock

	srv *grpc.Server

	logger     *log.Logger
	gcInterval time.Duration
}

//WithLogger sets the logger
func WithLogger(logger *log.Logger) func(*Server) {
	return func(srv *Server) {
		srv.logger = logger
	}
}

//WithGCInterval sets the GC interval for garbage collector
func WithGCInterval(interval time.Duration) func(*Server) {
	return func(srv *Server) {
		srv.gcInterval = interval
	}
}

//WithClock sets the clock
func WithClock(clck serverClock) func(*Server) {
	return func(srv *Server) {
		srv.clck = clck
	}
}

//New creates a server with given storage and options
func New(strg serverStorage, parser commandParser, opts ...func(*Server)) *Server {
	srv := &Server{parser: parser, strg: strg}

	for _, f := range opts {
		f(srv)
	}

	if srv.logger == nil {
		srv.logger = log.New(os.Stdout, "[godown-server]: ", log.LstdFlags)
	}

	if srv.gcInterval == 0 {
		srv.gcInterval = defaultGCInterval
	}

	if srv.clck == nil {
		srv.clck = clock.New()
	}

	return srv
}

//Start starts a server
func (s *Server) Start(hostPort string) error {
	l, err := net.Listen("tcp", hostPort)
	if err != nil {
		return fmt.Errorf("could not listen on %s: %v", hostPort, err)
	}

	s.srv = grpc.NewServer()
	api.RegisterGodownServer(s.srv, s)

	// starting a garbage collector
	go func() {
		gc := newGc(s.strg, s.logger, s.clck, s.gcInterval)
		gc.start()
	}()

	return s.srv.Serve(l)
}

//Stop stops a grpc server
func (s *Server) Stop() {
	s.srv.Stop()
}

//ExecuteCommand executes a command that placed into the request.
func (s *Server) ExecuteCommand(ctx context.Context, req *api.Request) (*api.Response, error) {
	cmd, args, err := s.parser.Parse(req.Command)
	if err != nil {
		if err == command.ErrCommandNotFound {
			return &api.Response{
				Result: &api.Response_Result{
					Type: api.Response_ERR,
					Item: fmt.Sprintf("command %q not found", req.Command),
				},
			}, nil
		}
		return nil, fmt.Errorf("could not parse command: %v", err)
	}
	res := cmd.Execute(args...)

	apiRes := new(api.Response_Result)

	switch t := res.(type) {
	case command.NilResult:
		apiRes.Type = api.Response_NIL
	case command.OkResult:
		apiRes.Type = api.Response_OK
	case command.StringResult:
		apiRes.Type = api.Response_STRING
		apiRes.Item = t.Value
	case command.IntResult:
		apiRes.Type = api.Response_INT
		apiRes.Item = string(t.Value)
	case command.SliceResult:
		apiRes.Type = api.Response_SLICE
		apiRes.Items = t.Value
	case command.HelpResult:
		apiRes.Type = api.Response_HELP
		apiRes.Item = t.Value
	case command.ErrResult:
		apiRes.Type = api.Response_ERR
		apiRes.Item = t.Value.Error()
	default:
		return nil, fmt.Errorf("unsupported type %T", res)
	}

	return &api.Response{Result: apiRes}, nil
}
