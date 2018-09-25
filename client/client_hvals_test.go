package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_HVals(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	tests := []struct {
		name          string
		arg           string
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ListResult
	}{
		{
			name:          "could_not_execute_command",
			arg:           "key",
			expectCommand: "HVALS key",
			mockErr:       errors.New("something went wrong"),
			wantResult: ListResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			arg:           "key",
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ListResult{
				err: errors.New("internal server error"),
			},
		},
		{
			name:          "server_responds_with_nil",
			arg:           "key",
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ListResult{},
		},
		{
			name:          "server_responds_with_slice",
			arg:           "key",
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"string1", "string2"},
			},
			wantResult: ListResult{val: []string{"string1", "string2"}},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			arg:           "key",
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: ListResult{err: errors.New("unexpected reply: OK")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc)
			mock.ExecuteCommandMock.
				Expect(context.Background(), &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.HVals(tt.arg)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_HValsWithContext(t *testing.T) {
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
		wantResult    ListResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "HVALS key",
			mockErr:       errors.New("something went wrong"),
			wantResult: ListResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ListResult{
				err: errors.New("internal server error"),
			},
		},
		{
			name:          "server_responds_with_nil",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ListResult{},
		},
		{
			name:          "server_responds_with_slice",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"string1", "string2"},
			},
			wantResult: ListResult{val: []string{"string1", "string2"}},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "value"), key: "key"},
			expectCtx:     contextWithValue("key", "value"),
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"string1", "string2"},
			},
			wantResult: ListResult{val: []string{"string1", "string2"}},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key"},
			expectCtx:     context.Background(),
			expectCommand: "HVALS key",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.OkCommandReply,
			},
			wantResult: ListResult{err: errors.New("unexpected reply: OK")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewexecutorMock(mc)
			mock.ExecuteCommandMock.
				Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
				Return(tt.mockResponse, tt.mockErr)

			cl := Client{executor: mock}

			res := cl.HValsWithContext(tt.args.ctx, tt.args.key)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
