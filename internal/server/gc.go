package server

import (
	"fmt"
	"log"
	"time"

	"github.com/namreg/godown-v2/internal/api"
	context "golang.org/x/net/context"
)

//gc is the garbage collector that collects expired values
type gc struct {
	srv      *Server
	logger   *log.Logger
	clck     serverClock
	data     dataStore
	interval time.Duration
	ticker   *time.Ticker
}

func newGc(
	srv *Server,
	strg dataStore,
	logger *log.Logger,
	clck serverClock,
	interval time.Duration,
) *gc {
	return &gc{
		srv:      srv,
		logger:   logger,
		clck:     clck,
		data:     strg,
		interval: interval,
	}
}

func (g *gc) start() {
	g.ticker = time.NewTicker(g.interval)

	for range g.ticker.C {
		g.deleteExpired()
	}
}

func (g *gc) stop() {
	if g.ticker != nil {
		g.ticker.Stop()
	}
}

func (g *gc) deleteExpired() {
	items, err := g.data.AllWithTTL()
	if err != nil {
		g.logger.Printf("[WARN] gc: could not retrieve values: %v", err)
	}
	for k, v := range items {
		if v.IsExpired(g.clck.Now()) {
			var err error
			req := &api.ExecuteCommandRequest{Command: fmt.Sprintf("DEL %s", k)}

			if g.srv.isLeader() {
				_, err = g.srv.handleExecuteCommandRequest(req)
			} else {
				_, err = g.srv.ExecuteCommand(context.Background(), req)
			}

			if err != nil {
				g.logger.Printf("[WANR] gc: could not delete item: %v", err)
			}
		}
	}
}
