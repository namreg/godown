package client

import (
	"context"
	"errors"
	"testing"

	"github.com/namreg/godown/internal/api"

	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/assert"
)

func TestClient_Get(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	test := []struct {
		name          string
		arg           string
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ScalarResult
	}{
		{
			name:          "could_not_execute_command",
			arg:           "test_key",
			expectCommand: "GET test_key",
			mockErr:       errors.New("server error"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: server error"),
			},
		},
		{
			name:          "server_responds_with_error",
			arg:           "test_key",
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "something went wrong",
			},
			wantResult: ScalarResult{
				err: errors.New("something went wrong"),
			},
		},
		{
			name:          "server_responds_with_nil",
			arg:           "test_key",
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ScalarResult{},
		},
		{
			name:          "server_responds_with_string",
			arg:           "test_key",
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "value",
			},
			wantResult: ScalarResult{val: stringToPtr("value")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			arg:           "test_key",
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
	}

	for _, tt := range test {
		mock := NewexecutorMock(mc)
		mock.ExecuteCommandMock.
			Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.Get(tt.arg)
		assert.Equal(t, tt.wantResult, res)
	}
}

func TestClient_GetWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx context.Context
		key string
	}

	test := []struct {
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
			args:          args{key: "test_key", ctx: context.Background()},
			expectCtx:     context.Background(),
			expectCommand: "GET test_key",
			mockErr:       errors.New("server error"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: server error"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{key: "test_key", ctx: context.Background()},
			expectCtx:     context.Background(),
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "something went wrong",
			},
			wantResult: ScalarResult{
				err: errors.New("something went wrong"),
			},
		},
		{
			name:          "server_responds_with_nil",
			args:          args{key: "test_key", ctx: context.Background()},
			expectCtx:     context.Background(),
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ScalarResult{},
		},
		{
			name:          "server_responds_with_string",
			args:          args{key: "test_key", ctx: context.Background()},
			expectCtx:     context.Background(),
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "value",
			},
			wantResult: ScalarResult{val: stringToPtr("value")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{key: "test_key", ctx: context.Background()},
			expectCtx:     context.Background(),
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("ctx_key", "ctx_value"), key: "test_key"},
			expectCtx:     contextWithValue("ctx_key", "ctx_value"),
			expectCommand: "GET test_key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
	}

	for _, tt := range test {
		mock := NewexecutorMock(mc)
		mock.ExecuteCommandMock.
			Expect(tt.args.ctx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.GetWithContext(tt.args.ctx, tt.args.key)
		assert.Equal(t, tt.wantResult, res)
	}
}
