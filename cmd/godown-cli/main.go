package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/namreg/godown-v2/internal/cli"
)

var host = flag.String("host", "127.0.0.1", "Host to connect to a server")
var port = flag.String("port", "4000", "Port to connect to a server")

func main() {
	flag.Parse()

	hostPort := net.JoinHostPort(*host, *port)

	if err := cli.Run(hostPort); err != nil {
		fmt.Fprintf(os.Stderr, "could not run CLI: %v", err)
	}
}
