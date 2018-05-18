package main

import (
	"flag"
	"log"
	"net"

	"github.com/namreg/godown-v2/internal/pkg/server"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
)

var host = flag.String("host", "127.0.0.1", "Server host")
var port = flag.String("port", "4000", "Server port")

func main() {
	flag.Parse()

	if *host == "" {
		log.Fatalf("host can not be empty")
	}
	if *port == "" {
		log.Fatalf("port can not be empty")
	}

	strg := memory.New()
	srv := server.New(strg)

	log.Fatal(srv.Run(net.JoinHostPort(*host, *port)))
}
