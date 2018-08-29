package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/namreg/godown-v2/internal/clock"
	"github.com/namreg/godown-v2/internal/command"
	"github.com/namreg/godown-v2/internal/server"
	"github.com/namreg/godown-v2/internal/storage/memory"
)

func main() {
	id := flag.String("id", "", "Server unique id.")
	dir := flag.String("dir", "", "Directory where data is stored.")
	listenAddr := flag.String("listen", "", "Server address to listen.")
	raftAddr := flag.String("raft", "", "Raft protocol listen address.")
	joinAddr := flag.String("join", "", "Server address to join.")
	gcInterval := flag.Duration("gc", 0, "Garbage collector interval.")

	flag.Parse()

	clck := clock.New()
	strg := memory.New(nil)
	parser := command.NewParser(strg, clck)

	opts := server.DefaultOptions()
	opts.ID = *id

	if *listenAddr != "" {
		opts.ListenAddr = *listenAddr
	}
	if *raftAddr != "" {
		opts.RaftAddr = *raftAddr
	}

	if *dir != "" {
		opts.Dir = *dir
	}

	if *gcInterval != 0 {
		opts.GCInterval = *gcInterval
	}

	srv := server.New(strg, strg, parser, opts)

	var err error

	if *joinAddr != "" {
		err = srv.JoinCluster(*joinAddr)
	} else {
		err = srv.BootstrapCluster()
	}

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}
