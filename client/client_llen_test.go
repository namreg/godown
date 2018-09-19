package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_LLen(t *testing.T) {
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
			expectCommand: "LLEN key",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			arg:           "key",
			expectCommand: "LLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_nil",
			arg:           "key",
			expectCommand: "LLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ScalarResult{},
		},
		{
			name:          "server_responds_with_integer",
			arg:           "key",
			expectCommand: "LLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "10",
			},
			wantResult: ScalarResult{val: stringToPtr("10")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			arg:           "key",
			expectCommand: "LLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"10"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
	}

	for _, tt := range tests {
		mock := NewexecutorMock(mc)
		mock.ExecuteCommandMock.
			Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.LLen(tt.arg)
		assert.Equal(t, tt.wantResult, res)
	}
}

func TestClient_LLenWithContext(t *testing.T) {
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
			expectCommand: "LLEN key",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "LLEN key",
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
			expectCommand: "LLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ScalarResult{},
		},
		{
			name:          "server_responds_with_integer",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "LLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "10",
			},
			wantResult: ScalarResult{val: stringToPtr("10")},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "value"), key: "key"},
			expectCtx:     contextWithValue("key", "value"),
			expectCommand: "LLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "10",
			},
			wantResult: ScalarResult{val: stringToPtr("10")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "LLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"10"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
	}

	for _, tt := range tests {
		mock := NewexecutorMock(mc)
		mock.ExecuteCommandMock.
			Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.LLenWithContext(tt.args.ctx, tt.args.key)
		assert.Equal(t, tt.wantResult, res)
	}
}
