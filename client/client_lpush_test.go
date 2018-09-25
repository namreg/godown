package client

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_LPush(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key, value string
	}

	tests := []struct {
		name          string
		args          args
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    StatusResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{key: "key", value: "val"},
			expectCommand: "LPUSH key val",
			mockErr:       errors.New("something went wrong"),
			wantResult: StatusResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{key: "key", value: "val"},
			expectCommand: "LPUSH key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_ok",
			args:          args{key: "key", value: "val"},
			expectCommand: "LPUSH key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{key: "key", value: "val"},
			expectCommand: "LPUSH key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "string",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: STRING")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc)
			mock.ExecuteCommandMock.
				Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.LPush(tt.args.key, tt.args.value)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_LPushWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key   string
		value string
		ctx   context.Context
	}

	tests := []struct {
		name          string
		args          args
		expectCtx     context.Context
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    StatusResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{ctx: context.Background(), key: "key", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "LPUSH key val",
			mockErr:       errors.New("something went wrong"),
			wantResult: StatusResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "LPUSH key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_ok",
			args:          args{ctx: context.Background(), key: "key", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "LPUSH key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "value"), key: "key", value: "val"},
			expectCtx:     contextWithValue("key", "value"),
			expectCommand: "LPUSH key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "LPUSH key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "string",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: STRING")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc)
			mock.ExecuteCommandMock.
				Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.LPushWithContext(tt.expectCtx, tt.args.key, tt.args.value)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
