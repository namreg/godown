package cli

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/namreg/godown-v2/internal/api"

	"github.com/Bowery/prompt"
	"google.golang.org/grpc"
)

const connectTimeout = 200 * time.Millisecond

const prefix = "godown >"

//CLI allows users to interact with a server.
type CLI struct {
	printer *printer
	term    *prompt.Terminal
	conn    *grpc.ClientConn
	client  api.GodownClient
}

//Run runs a new CLI.
func Run(hostPort string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, hostPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("could not connect to %s", hostPort)
	}

	term, err := prompt.NewTerminal()
	if err != nil {
		return fmt.Errorf("could not create a terminal: %v", err)
	}

	c := &CLI{
		printer: newPrinter(os.Stdout),
		term:    term,
		client:  api.NewGodownClient(conn),
		conn:    conn,
	}

	defer func() {
		err = c.Close()
	}()

	c.run()

	return nil
}

//Close closes the CLI.
func (c *CLI) Close() error {
	if err := c.printer.Close(); err != nil {
		return err
	}
	if err := c.conn.Close(); err != nil {
		return err
	}
	return c.term.Close()
}

func (c *CLI) run() {
	c.printer.printLogo()
	for {
		input, err := c.term.GetPrompt(prefix)
		if err != nil {
			if err == prompt.ErrCTRLC || err == prompt.ErrEOF {
				break
			}
			c.printer.printError(err)
			continue
		}
		if input == "" {
			continue
		}
		req := &api.ExecuteCommandRequest{Command: input}
		if resp, err := c.client.ExecuteCommand(context.Background(), req); err != nil {
			c.printer.printError(err)
		} else {
			c.printer.printResponse(resp)
		}
	}
	c.printer.println("Bye!")
}
