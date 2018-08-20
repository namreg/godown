package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/namreg/godown-v2/internal/api"

	"github.com/Bowery/prompt"
	"google.golang.org/grpc"
)

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
	term, err := prompt.NewTerminal()
	if err != nil {
		return fmt.Errorf("could not create a terminal: %v", err)
	}
	conn, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not dial the server: %v", err)
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
	c.printer.println("Welcome to the godown CLI!")
	for {
		input, err := c.term.GetPrompt(prefix)
		if err != nil {
			if err == prompt.ErrCTRLC || err == prompt.ErrEOF {
				break
			}
			c.printer.printError(err)
			continue
		}
		req := &api.Request{Command: input}
		if resp, err := c.client.ExecuteCommand(context.Background(), req); err != nil {
			c.printer.printError(err)
		} else {
			c.printer.printResponse(resp)
		}
	}
	c.printer.println("Bye!")
}
