package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_SetBit(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key           string
		offset, value uint64
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
			args:          args{key: "key", offset: 1024, value: 1},
			expectCommand: "SETBIT key 1024 1",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_responds_with_error",
			args:          args{key: "key", offset: 1024, value: 1},
			expectCommand: "SETBIT key 1024 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_ok",
			args:          args{key: "key", offset: 1024, value: 1},
			expectCommand: "SETBIT key 1024 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{key: "key", offset: 1024, value: 1},
			expectCommand: "SETBIT key 1024 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "10",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: INT")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc).ExecuteCommandMock.
				Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.SetBit(tt.args.key, tt.args.offset, tt.args.value)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_SetBitWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx           context.Context
		key           string
		offset, value uint64
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
			args:          args{ctx: context.Background(), key: "key", offset: 1024, value: 1},
			expectCtx:     context.Background(),
			expectCommand: "SETBIT key 1024 1",
			mockErr:       errors.New("something went wrong"),
			wantResult:    StatusResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key", offset: 1024, value: 1},
			expectCtx:     context.Background(),
			expectCommand: "SETBIT key 1024 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: StatusResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_ok",
			args:          args{ctx: context.Background(), key: "key", offset: 1024, value: 1},
			expectCtx:     context.Background(),
			expectCommand: "SETBIT key 1024 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "val"), key: "key", offset: 1024, value: 1},
			expectCtx:     contextWithValue("key", "val"),
			expectCommand: "SETBIT key 1024 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: StatusResult{},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", offset: 1024, value: 1},
			expectCtx:     context.Background(),
			expectCommand: "SETBIT key 1024 1",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "10",
			},
			wantResult: StatusResult{err: errors.New("unexpected reply: INT")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc).ExecuteCommandMock.
				Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.SetBitWithContext(tt.args.ctx, tt.args.key, tt.args.offset, tt.args.value)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
