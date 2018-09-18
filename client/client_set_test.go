package client

import (
	"context"
	"errors"
	"testing"

	"github.com/namreg/godown/internal/api"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/assert"
)

func TestClient_Set(t *testing.T) {
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
			expectCommand: "SET key val",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_respond_with_error",
			args:          args{key: "key", value: "val"},
			expectCommand: "SET key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_respond_with_ok",
			args:          args{key: "key", value: "val"},
			expectCommand: "SET key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_respond_with_unexpected_reply",
			args:          args{key: "key", value: "val"},
			expectCommand: "SET key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: SLICE")},
		},
	}

	for _, tt := range tests {
		mock := NewexecutorMock(mc).ExecuteCommandMock.
			Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.Set(tt.args.key, tt.args.value)
		assert.Equal(t, tt.wantResult, res)
	}
}

func TestClient_SetWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx        context.Context
		key, value string
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
			expectCommand: "SET key val",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_respond_with_error",
			args:          args{ctx: context.Background(), key: "key", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "SET key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_respond_with_ok",
			args:          args{ctx: context.Background(), key: "key", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "SET key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_respond_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "SET key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: SLICE")},
		},
		{
			name:          "with_custom_context",
			args:          args{ctx: contextWithValue("ctx_key", "ctx_value"), key: "key", value: "val"},
			expectCtx:     contextWithValue("ctx_key", "ctx_value"),
			expectCommand: "SET key val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
	}

	for _, tt := range tests {
		mock := NewexecutorMock(mc).ExecuteCommandMock.
			Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.SetWithContext(tt.args.ctx, tt.args.key, tt.args.value)
		assert.Equal(t, tt.wantResult, res)
	}
}
