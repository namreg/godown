package server

import (
	"log"
	"time"

	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/pkg/clock"
)

//gc is the garbage collector that collects expired values
type gc struct {
	clck     clock.Clock
	strg     storage.Storage
	interval time.Duration
}

func newGc(strg storage.Storage, clck clock.Clock, interval time.Duration) *gc {
	return &gc{
		clck:     clck,
		strg:     strg,
		interval: interval,
	}
}

func (g *gc) start() {
	ticker := time.NewTicker(g.interval)
	defer ticker.Stop()

	for range ticker.C {
		items, err := g.strg.AllWithTTL()
		if err != nil {
			log.Printf("[WARN] gc: could not retrieve values: %v", err)
		}
		for k, v := range items {
			if v.IsExpired(g.clck.Now()) {
				g.strg.Del(k)
			}
		}
	}
}
