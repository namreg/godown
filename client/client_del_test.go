package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_Del(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	tests := []struct {
		name          string
		arg           string
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    StatusResult
	}{
		{
			name:          "could_not_execute_command",
			arg:           "key",
			expectCommand: "DEL key",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_respond_with_error",
			arg:           "key",
			expectCommand: "DEL key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_respond_with_ok",
			arg:           "key",
			expectCommand: "DEL key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_respond_with_unexpected_reply",
			arg:           "key",
			expectCommand: "DEL key",
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

		res := cl.Del(tt.arg)
		assert.Equal(t, tt.wantResult, res)
	}
}

func TestClient_DelWithContext(t *testing.T) {
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
		wantResult    StatusResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "DEL key",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_respond_with_error",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "DEL key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_respond_with_ok",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "DEL key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_respond_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "DEL key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.RawStringCommandReply,
				Item:  "raw string",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: RAW_STRING")},
		},
		{
			name:          "with_custom_context",
			args:          args{ctx: contextWithValue("ctx_key", "ctx_value"), key: "key"},
			expectCtx:     contextWithValue("ctx_key", "ctx_value"),
			expectCommand: "DEL key",
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

		res := cl.DelWithContext(tt.args.ctx, tt.args.key)
		assert.Equal(t, tt.wantResult, res)
	}
}
