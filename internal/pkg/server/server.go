package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/namreg/godown-v2/internal/pkg/command"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/pkg/clock"
)

type commandParser interface {
	Parse(str string) (cmd command.Command, args []string, err error)
}

//go:generate minimock -i github.com/namreg/godown-v2/internal/pkg/server.serverStorage -o ./

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

//go:generate minimock -i github.com/namreg/godown-v2/internal/pkg/server.serverClock -o ./

type serverClock interface {
	//Now returns current time
	Now() time.Time
}

const defaultGCInterval = 500 * time.Millisecond

//Server represents a server that handles user requests and executes commands
type Server struct {
	parser     commandParser
	strg       serverStorage
	clck       serverClock
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

//Run runs the server on the given host and port
func (s *Server) Run(hostPort string) error {
	s.logger.Printf("[INFO] running on %s\n", hostPort)

	// starting a garbage collector
	go func() {
		gc := newGc(s.strg, s.logger, s.clck, s.gcInterval)
		gc.start()
	}()

	l, err := net.Listen("tcp", hostPort)
	if err != nil {
		return fmt.Errorf("server: could not listen %s: %v", hostPort, err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			s.logger.Printf("[WARN] could not accept connection: %v\n", err)
			continue
		}

		go s.handleConn(newConn(conn))
	}
}

func (s *Server) handleConn(conn *conn) {
	defer conn.Close()

	conn.writeWelcomeMessage()

	scanner := bufio.NewScanner(conn.conn)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			conn.writePrompt()
			continue
		}

		cmd, args, err := s.parser.Parse(input)

		if err != nil {
			switch err {
			case command.ErrCommandNotFound:
				conn.writeError(fmt.Errorf("command %q not found", input))
			default:
				conn.writeError(err)
			}
			continue
		}

		res := cmd.Execute(args...)

		conn.writeCommandResult(res)
	}

	if err := scanner.Err(); err != nil {
		s.logger.Printf("[WARN] scanner error: %v", err)
	}
}
