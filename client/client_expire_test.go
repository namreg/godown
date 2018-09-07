package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_Expire(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key  string
		secs int
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
			args:          args{key: "key", secs: 10},
			expectCommand: "EXPIRE key 10",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_respond_with_error",
			args:          args{key: "key", secs: 10},
			expectCommand: "EXPIRE key 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_respond_with_ok",
			args:          args{key: "key", secs: 10},
			expectCommand: "EXPIRE key 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_respond_with_unexpected_reply",
			args:          args{key: "key", secs: 10},
			expectCommand: "EXPIRE key 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.RawStringCommandReply,
				Item:  "raw string",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: RAW_STRING")},
		},
	}

	for _, tt := range tests {
		mock := NewexecutorMock(mc)
		mock.ExecuteCommandMock.
			Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.Expire(tt.args.key, tt.args.secs)
		assert.Equal(t, tt.wantResult, res)
	}
}

func TestClient_ExpireWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx  context.Context
		key  string
		secs int
	}

	type ctxKey string

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
			args:          args{ctx: context.Background(), key: "key", secs: 10},
			expectCtx:     context.Background(),
			expectCommand: "EXPIRE key 10",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_respond_with_error",
			args:          args{ctx: context.Background(), key: "key", secs: 10},
			expectCtx:     context.Background(),
			expectCommand: "EXPIRE key 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_respond_with_ok",
			args:          args{ctx: context.Background(), key: "key", secs: 10},
			expectCtx:     context.Background(),
			expectCommand: "EXPIRE key 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_respond_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", secs: 10},
			expectCtx:     context.Background(),
			expectCommand: "EXPIRE key 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.RawStringCommandReply,
				Item:  "raw string",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: RAW_STRING")},
		},
		{
			name:          "with_custom_context",
			args:          args{ctx: context.WithValue(context.Background(), ctxKey("ctx_key"), "ctx_value"), key: "key", secs: 10},
			expectCtx:     context.WithValue(context.Background(), ctxKey("ctx_key"), "ctx_value"),
			expectCommand: "EXPIRE key 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
	}

	for _, tt := range tests {
		mock := NewexecutorMock(mc)
		mock.ExecuteCommandMock.
			Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.ExpireWithContext(tt.args.ctx, tt.args.key, tt.args.secs)
		assert.Equal(t, tt.wantResult, res)
	}
}
