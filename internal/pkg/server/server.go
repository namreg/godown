package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/namreg/godown-v2/internal/pkg/command"
	"github.com/namreg/godown-v2/internal/pkg/storage"
)

//Server represents a server
type Server struct {
	strg storage.Storage
}

//New creates a server with given storage
func New(strg storage.Storage) *Server {
	return &Server{strg}
}

//Run runs the server on the given host and port
func (s *Server) Run(hostPort string) error {
	log.Printf("[INFO] server: runing on %s\n", hostPort)
	l, err := net.Listen("tcp", hostPort)
	if err != nil {
		return fmt.Errorf("server: could not listen %s: %v", hostPort, err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("[WARN] server: could not accept connection: %v\n", err)
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
			log.Printf("[INFO] server: could not parse command %q: %v", input, err)
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

		log.Printf("[INFO] server: recieved command: %s", cmd.Name())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("[WARN] server: scanner error: %v", err)
	}
}
