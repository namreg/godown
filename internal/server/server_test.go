package server

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/namreg/godown-v2/internal/clock"
	"github.com/namreg/godown-v2/internal/command"
	"github.com/namreg/godown-v2/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestNew_WithoutOptions(t *testing.T) {
	defaultLogger := log.New(os.Stdout, "[godown-server]: ", log.LstdFlags)

	strg := memory.New(nil)
	parser := command.NewParser(strg, NewserverClockMock(t))
	srv := New(strg, parser)

	assert.Equal(t, defaultLogger, srv.logger)
	assert.Equal(t, defaultGCInterval, srv.gcInterval)
	assert.Equal(t, clock.New(), srv.clck)
}

func TestNew_WithOptions(t *testing.T) {
	logger := log.New(os.Stdout, "[test server]: ", log.LUTC)
	clck := NewserverClockMock(t)
	strg := memory.New(nil)
	parser := command.NewParser(strg, clck)
	srv := New(
		strg,
		parser,
		WithLogger(logger),
		WithClock(NewserverClockMock(t)),
		WithGCInterval(10*time.Second),
	)

	assert.Equal(t, logger, srv.logger)
	assert.Equal(t, 10*time.Second, srv.gcInterval)
	assert.Equal(t, clck, srv.clck)
}
