package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_Type(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	tests := []struct {
		name          string
		arg           string
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ScalarResult
	}{
		{
			name:          "could_not_execute_command",
			arg:           "key",
			expectCommand: "TYPE key",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			arg:           "key",
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_nil",
			arg:           "key",
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ScalarResult{},
		},
		{
			name:          "server_responds_with_string",
			arg:           "key",
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "list",
			},
			wantResult: ScalarResult{val: stringToPtr("list")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			arg:           "key",
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc).ExecuteCommandMock.
				Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.Type(tt.arg)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_TypeWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx context.Context
		key string
	}

	tests := []struct {
		name          string
		args          args
		expectCtx     context.Context
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ScalarResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "TYPE key",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_nil",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ScalarResult{},
		},
		{
			name:          "server_responds_with_string",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "list",
			},
			wantResult: ScalarResult{val: stringToPtr("list")},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "val"), key: "key"},
			expectCtx:     contextWithValue("key", "val"),
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "list",
			},
			wantResult: ScalarResult{val: stringToPtr("list")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "TYPE key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc).ExecuteCommandMock.
				Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.TypeWithContext(tt.args.ctx, tt.args.key)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
