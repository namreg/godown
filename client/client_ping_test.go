package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_Ping(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	tests := []struct {
		name          string
		args          []string
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ScalarResult
	}{
		{
			name:          "could_not_execute_command",
			args:          []string{"hello", "world"},
			expectCommand: "PING hello world",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          []string{"hello", "world"},
			expectCommand: "PING hello world",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_raw_string",
			args:          []string{"hello", "world"},
			expectCommand: "PING hello world",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.RawStringCommandReply,
				Item:  "PONG hello world",
			},
			wantResult: ScalarResult{val: stringToPtr("PONG hello world")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          []string{"hello", "world"},
			expectCommand: "PING hello world",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: OK")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc).ExecuteCommandMock.
				Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.Ping(tt.args...)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_PingWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type arguments struct {
		ctx  context.Context
		args []string
	}

	tests := []struct {
		name          string
		args          arguments
		expectCtx     context.Context
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ScalarResult
	}{
		{
			name:          "could_not_execute_command",
			args:          arguments{ctx: context.Background(), args: []string{"hello", "world"}},
			expectCtx:     context.Background(),
			expectCommand: "PING hello world",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          arguments{ctx: context.Background(), args: []string{"hello", "world"}},
			expectCtx:     context.Background(),
			expectCommand: "PING hello world",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_raw_string",
			args:          arguments{ctx: context.Background(), args: []string{"hello", "world"}},
			expectCtx:     context.Background(),
			expectCommand: "PING hello world",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.RawStringCommandReply,
				Item:  "PONG hello world",
			},
			wantResult: ScalarResult{val: stringToPtr("PONG hello world")},
		},
		{
			name:          "custom_context",
			args:          arguments{ctx: contextWithValue("key", "value"), args: []string{"hello", "world"}},
			expectCtx:     contextWithValue("key", "value"),
			expectCommand: "PING hello world",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.RawStringCommandReply,
				Item:  "PONG hello world",
			},
			wantResult: ScalarResult{val: stringToPtr("PONG hello world")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          arguments{ctx: context.Background(), args: []string{"hello", "world"}},
			expectCtx:     context.Background(),
			expectCommand: "PING hello world",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: OK")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc).ExecuteCommandMock.
				Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.PingWithContext(tt.args.ctx, tt.args.args...)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
