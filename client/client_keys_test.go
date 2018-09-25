package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_Keys(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	tests := []struct {
		name          string
		arg           string
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ListResult
	}{
		{
			name:          "could_not_execute_command",
			arg:           "*",
			expectCommand: "KEYS *",
			mockErr:       errors.New("something went wrong"),
			wantResult: ListResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			arg:           "*",
			expectCommand: "KEYS *",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ListResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_slice",
			arg:           "*",
			expectCommand: "KEYS *",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"key1"},
			},
			wantResult: ListResult{val: []string{"key1"}},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			arg:           "*",
			expectCommand: "KEYS *",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "1",
			},
			wantResult: ListResult{err: errors.New("unexpected reply: INT")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc)
			mock.ExecuteCommandMock.
				Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.Keys(tt.arg)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_KeysWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx     context.Context
		pattern string
	}

	tests := []struct {
		name          string
		args          args
		expectCtx     context.Context
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ListResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{ctx: context.Background(), pattern: "*"},
			expectCtx:     context.Background(),
			expectCommand: "KEYS *",
			mockErr:       errors.New("something went wrong"),
			wantResult: ListResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), pattern: "*"},
			expectCtx:     context.Background(),
			expectCommand: "KEYS *",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ListResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_slice",
			args:          args{ctx: context.Background(), pattern: "*"},
			expectCtx:     context.Background(),
			expectCommand: "KEYS *",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"key1"},
			},
			wantResult: ListResult{val: []string{"key1"}},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "value"), pattern: "*"},
			expectCtx:     contextWithValue("key", "value"),
			expectCommand: "KEYS *",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"key1"},
			},
			wantResult: ListResult{val: []string{"key1"}},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), pattern: "*"},
			expectCtx:     context.Background(),
			expectCommand: "KEYS *",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "1",
			},
			wantResult: ListResult{err: errors.New("unexpected reply: INT")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc)
			mock.ExecuteCommandMock.
				Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.KeysWithContext(tt.args.ctx, tt.args.pattern)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
