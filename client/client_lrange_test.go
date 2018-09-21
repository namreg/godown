package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_LRange(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key         string
		start, stop int
	}

	tests := []struct {
		name          string
		args          args
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ListResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{key: "key", start: 1, stop: 10},
			expectCommand: "LRANGE key 1 10",
			mockErr:       errors.New("something went wrong"),
			wantResult: ListResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{key: "key", start: 1, stop: 10},
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ListResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_nil",
			args:          args{key: "key", start: 1, stop: 10},
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ListResult{},
		},
		{
			name:          "server_responds_with_slice",
			args:          args{key: "key", start: 1, stop: 10},
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"value 1", "value 2"},
			},
			wantResult: ListResult{val: []string{"value 1", "value 2"}},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{key: "key", start: 1, stop: 10},
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "2",
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

			res := cl.LRange(tt.args.key, tt.args.start, tt.args.stop)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_LRangeWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx         context.Context
		key         string
		start, stop int
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
			args:          args{ctx: context.Background(), key: "key", start: 1, stop: 10},
			expectCtx:     context.Background(),
			expectCommand: "LRANGE key 1 10",
			mockErr:       errors.New("something went wrong"),
			wantResult: ListResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key", start: 1, stop: 10},
			expectCtx:     context.Background(),
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ListResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_nil",
			args:          args{ctx: context.Background(), key: "key", start: 1, stop: 10},
			expectCtx:     context.Background(),
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ListResult{},
		},
		{
			name:          "server_responds_with_slice",
			args:          args{ctx: context.Background(), key: "key", start: 1, stop: 10},
			expectCtx:     context.Background(),
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"value 1", "value 2"},
			},
			wantResult: ListResult{val: []string{"value 1", "value 2"}},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("ctx_key", "value"), key: "key", start: 1, stop: 10},
			expectCtx:     contextWithValue("ctx_key", "value"),
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"value 1", "value 2"},
			},
			wantResult: ListResult{val: []string{"value 1", "value 2"}},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", start: 1, stop: 10},
			expectCtx:     context.Background(),
			expectCommand: "LRANGE key 1 10",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "2",
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

			res := cl.LRangeWithContext(tt.args.ctx, tt.args.key, tt.args.start, tt.args.stop)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
