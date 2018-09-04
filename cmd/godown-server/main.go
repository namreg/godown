package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/namreg/godown/internal/clock"
	"github.com/namreg/godown/internal/command"
	"github.com/namreg/godown/internal/server"
	"github.com/namreg/godown/internal/storage/memory"
)

//Values populated by the Go linker.
var (
	version = "unknown"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	id := flag.String("id", "", "Server unique id.")
	dir := flag.String("dir", "", "Directory where data is stored.")
	listenAddr := flag.String("listen", "", "Server address to listen.")
	raftAddr := flag.String("raft", "", "Raft protocol listen address.")
	joinAddr := flag.String("join", "", "Server address to join.")
	gcInterval := flag.Duration("gc", 0, "Garbage collector interval.")
	showVersion := flag.Bool("version", false, "Show version.")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\ncommit: %s\nbuildtime: %s", version, commit, date)
		os.Exit(0)
	}

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
