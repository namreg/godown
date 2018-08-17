package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/namreg/godown-v2/internal/pkg/command"

	"github.com/namreg/godown-v2/internal/pkg/storage/memory"

	"github.com/namreg/godown-v2/internal/pkg/server"
	"github.com/namreg/godown-v2/pkg/clock"
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

	clck := clock.New()
	strg := memory.New(nil, memory.WithClock(clck))
	parser := command.NewParser(strg, clck)

	srv := server.New(strg, parser, server.WithClock(clck), server.WithGCInterval(1*time.Second))

	log.Fatal(srv.Run(net.JoinHostPort(*host, *port)))
}
