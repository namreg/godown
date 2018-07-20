package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/namreg/godown-v2/pkg/clock"

	"github.com/namreg/godown-v2/internal/pkg/command"
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

const defaultGCInterval = 500 * time.Millisecond

//Server represents a server
type Server struct {
	logger     *log.Logger
	strg       storage.Storage
	gcInterval time.Duration
	clock      clock.Clock
}

//WithLogger sets the logger
func WithLogger(logger *log.Logger) func(*Server) {
	return func(srv *Server) {
		srv.logger = logger
	}
}

//WithGCInterval sets GC interval for garbage collector
func WithGCInterval(interval time.Duration) func(*Server) {
	return func(srv *Server) {
		srv.gcInterval = interval
	}
}

//WithClock sets clock
func WithClock(clck clock.Clock) func(*Server) {
	return func(srv *Server) {
		srv.clock = clck
	}
}

//New creates a server with given storage and options
func New(strg storage.Storage, opts ...func(*Server)) *Server {
	srv := &Server{strg: strg}

	for _, f := range opts {
		f(srv)
	}

	if srv.logger == nil {
		srv.logger = log.New(os.Stdout, "[godown-server]: ", log.LstdFlags)
	}

	if srv.gcInterval == 0 {
		srv.gcInterval = defaultGCInterval
	}

	if srv.clock == nil {
		srv.clock = clock.TimeClock{}
	}

	return srv
}

//Run runs the server on the given host and port
func (s *Server) Run(hostPort string) error {
	s.logger.Printf("[INFO] running on %s\n", hostPort)

	// starting a garbage collector
	go func() {
		gc := newGc(s.strg, s.logger, s.clock, s.gcInterval)
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

		cmd, args, err := command.Parse(input)

		if err != nil {
			switch err {
			case command.ErrCommandNotFound:
				conn.writeError(fmt.Errorf("command %q not found", input))
			case command.ErrWrongArgsNumber:
				conn.writeError(fmt.Errorf("wrong number of arguments"))
			default:
				conn.writeError(err)
			}
			continue
		}

		res := cmd.Execute(s.strg, args...)

		conn.writeCommandResult(res)
	}

	if err := scanner.Err(); err != nil {
		s.logger.Printf("[WARN] scanner error: %v", err)
	}
}
