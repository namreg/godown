package server

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/pkg/command"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
	"github.com/namreg/godown-v2/pkg/clock"
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

func TestServer_handleConn(t *testing.T) {
	tests := []struct {
		name  string
		items map[storage.Key]*storage.Value
		input []byte
		want  []byte
	}{
		{
			"existing_command",
			map[storage.Key]*storage.Value{
				"key": storage.NewStringValue("value"),
			},
			[]byte("GET key"),
			[]byte("(string): value"),
		},
		{
			"not_existing_command",
			nil,
			[]byte("UNKNOWN arg"),
			[]byte("(error): command \"UNKNOWN arg\" not found"),
		},
		{
			"empty_input",
			nil,
			[]byte(" "),
			[]byte("godown > "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			defer mc.Finish()

			strg := memory.New(tt.items)
			parser := command.NewParser(strg, NewserverClockMock(t))
			srv := New(strg, parser)

			r := bytes.NewReader(tt.input)
			w := new(bytes.Buffer)

			netConn := NewConnMock(mc)
			netConn.CloseMock.Return(nil)
			netConn.WriteFunc = func(p []byte) (int, error) {
				return w.Write(p)
			}
			netConn.ReadFunc = func(p []byte) (int, error) {
				return r.Read(p)
			}

			conn := newConn(netConn)

			srv.handleConn(conn)

			assert.Truef(t, bytes.Contains(w.Bytes(), tt.want), "output does not contain %s. output %s", tt.want, w)
		})
	}
}
