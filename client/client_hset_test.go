package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_HSet(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key, field, value string
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
			args:          args{key: "key", field: "field", value: "val"},
			expectCommand: "HSET key field val",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_responds_with_error",
			args:          args{key: "key", field: "field", value: "val"},
			expectCommand: "HSET key field val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_ok",
			args:          args{key: "key", field: "field", value: "val"},
			expectCommand: "HSET key field val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{key: "key", field: "field", value: "val"},
			expectCommand: "HSET key field val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "5",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: INT")},
		},
	}

	for _, tt := range tests {
		mock := NewexecutorMock(mc).ExecuteCommandMock.
			Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.HSet(tt.args.key, tt.args.field, tt.args.value)
		assert.Equal(t, tt.wantResult, res)
	}
}

func TestClient_HSetWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx   context.Context
		key   string
		field string
		value string
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
			args:          args{ctx: context.Background(), key: "key", field: "field", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "HSET key field val",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key", field: "field", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "HSET key field val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_ok",
			args:          args{ctx: context.Background(), key: "key", field: "field", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "HSET key field val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "value"), key: "key", field: "field", value: "val"},
			expectCtx:     contextWithValue("key", "value"),
			expectCommand: "HSET key field val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", field: "field", value: "val"},
			expectCtx:     context.Background(),
			expectCommand: "HSET key field val",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "5",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: INT")},
		},
	}

	for _, tt := range tests {
		mock := NewexecutorMock(mc).ExecuteCommandMock.
			Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.HSetWithContext(tt.args.ctx, tt.args.key, tt.args.field, tt.args.value)
		assert.Equal(t, tt.wantResult, res)
	}
}
