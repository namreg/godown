package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetBit(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key    string
		offset uint64
	}

	tests := []struct {
		name          string
		args          args
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ScalarResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{key: "key", offset: 1},
			expectCommand: "GETBIT key 1",
			mockErr:       errors.New("something went wrong"),
			wantResult:    ScalarResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_responds_with_error",
			args:          args{key: "key", offset: 1},
			expectCommand: "GETBIT key 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_int",
			args:          args{key: "key", offset: 1},
			expectCommand: "GETBIT key 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "1",
			},
			wantResult: ScalarResult{val: stringToPtr("1")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{key: "key", offset: 1},
			expectCommand: "GETBIT key 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc)
			mock.ExecuteCommandMock.
				Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.GetBit(tt.args.key, tt.args.offset)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_GetBitWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx    context.Context
		key    string
		offset uint64
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
			args:          args{ctx: context.Background(), key: "key", offset: 1},
			expectCtx:     context.Background(),
			expectCommand: "GETBIT key 1",
			mockErr:       errors.New("something went wrong"),
			wantResult:    ScalarResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key", offset: 1},
			expectCtx:     context.Background(),
			expectCommand: "GETBIT key 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_int",
			args:          args{ctx: context.Background(), key: "key", offset: 1},
			expectCtx:     context.Background(),
			expectCommand: "GETBIT key 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "1",
			},
			wantResult: ScalarResult{val: stringToPtr("1")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", offset: 1},
			expectCtx:     context.Background(),
			expectCommand: "GETBIT key 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
		{
			name:          "with_custom_context",
			args:          args{ctx: contextWithValue("ctx_key", "ctx_value"), key: "key", offset: 1},
			expectCtx:     contextWithValue("ctx_key", "ctx_value"),
			expectCommand: "GETBIT key 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "1",
			},
			wantResult: ScalarResult{val: stringToPtr("1")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc)
			mock.ExecuteCommandMock.
				Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.GetBitWithContext(tt.args.ctx, tt.args.key, tt.args.offset)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
