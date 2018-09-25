package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_Strlen(t *testing.T) {
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
			expectCommand: "STRLEN key",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			arg:           "key",
			expectCommand: "STRLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_int",
			arg:           "key",
			expectCommand: "STRLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "5",
			},
			wantResult: ScalarResult{val: stringToPtr("5")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			arg:           "key",
			expectCommand: "STRLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: OK")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc).ExecuteCommandMock.
				Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.Strlen(tt.arg)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_StrlenWithContext(t *testing.T) {
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
			expectCommand: "STRLEN key",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "STRLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_int",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "STRLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "5",
			},
			wantResult: ScalarResult{val: stringToPtr("5")},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "val"), key: "key"},
			expectCtx:     contextWithValue("key", "val"),
			expectCommand: "STRLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "5",
			},
			wantResult: ScalarResult{val: stringToPtr("5")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "STRLEN key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: OK")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc).ExecuteCommandMock.
				Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.StrlenWithContext(tt.args.ctx, tt.args.key)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
