package server

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/pkg/storage"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"
	"github.com/namreg/godown-v2/pkg/clock"
	"github.com/stretchr/testify/assert"
)

func TestNew_WithoutOptions(t *testing.T) {
	defaultLogger := log.New(os.Stdout, "[godown-server]: ", log.LstdFlags)
	defaultClock := clock.TimeClock{}
	defaultStorage := memory.New(nil, memory.WithClock(defaultClock))

	srv := New()

	assert.IsType(t, &Server{}, srv)
	assert.Equal(t, defaultLogger, srv.logger)
	assert.Equal(t, defaultGCInterval, srv.gcInterval)
	assert.Equal(t, clock.TimeClock{}, srv.clock)
	assert.Equal(t, defaultStorage, srv.strg)
}

func TestNew_WithStorage_And_WithClock(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	clck := clock.NewClockMock(mc)
	strg := memory.New(map[storage.Key]*storage.Value{"key": storage.NewStringValue("value")}, memory.WithClock(clck))

	srv := New(WithStorage(strg), WithClock(clck))

	assert.IsType(t, &Server{}, srv)
	assert.Equal(t, strg, srv.strg)
	assert.Equal(t, clck, srv.clock)
}

func TestNew_WithLogger(t *testing.T) {
	logger := log.New(os.Stdout, "[test server]: ", log.LUTC)
	srv := New(WithLogger(logger))

	assert.Equal(t, logger, srv.logger)
}

func TestNew_WithGCInterval(t *testing.T) {
	interval := 10 * time.Second
	srv := New(WithGCInterval(interval))

	assert.Equal(t, interval, srv.gcInterval)
}

func TestServer_handleConn(t *testing.T) {
	tests := []struct {
		name  string
		init  func() *Server
		input []byte
		want  []byte
	}{
		{
			"existing_command",
			func() *Server {
				strg := memory.New(map[storage.Key]*storage.Value{
					"key": storage.NewStringValue("value"),
				})
				return New(WithStorage(strg))
			},
			[]byte("GET key"),
			[]byte("(string): value"),
		},
		{
			"not_existing_command",
			func() *Server {
				return New()
			},
			[]byte("UNKNOWN arg"),
			[]byte("(error): command \"UNKNOWN arg\" not found"),
		},
		{
			"empty_input",
			func() *Server {
				return New()
			},
			[]byte(" "),
			[]byte("godown > "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			defer mc.Finish()

			srv := tt.init()
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
