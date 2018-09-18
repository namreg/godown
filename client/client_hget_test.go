package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_HGet(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key, field string
	}

	test := []struct {
		name          string
		args          args
		expectCommand string
		mockResponse  *api.ExecuteCommandResponse
		mockErr       error
		wantResult    ScalarResult
	}{
		{
			name:          "could_not_execute_command",
			args:          args{key: "key", field: "field"},
			expectCommand: "HGET key field",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{key: "key", field: "field"},
			expectCommand: "HGET key field",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{
				err: errors.New("internal server error"),
			},
		},
		{
			name:          "server_responds_with_nil",
			args:          args{key: "key", field: "field"},
			expectCommand: "HGET key field",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ScalarResult{},
		},
		{
			name:          "server_responds_with_string",
			args:          args{key: "key", field: "field"},
			expectCommand: "HGET key field",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "string",
			},
			wantResult: ScalarResult{val: stringToPtr("string")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{key: "key", field: "field"},
			expectCommand: "HGET key field",
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

		res := cl.HGet(tt.args.key, tt.args.field)
		assert.Equal(t, tt.wantResult, res)
	}
}

func TestClient_HGetWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx        context.Context
		key, field string
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
			args:          args{ctx: context.Background(), key: "key", field: "field"},
			expectCtx:     context.Background(),
			expectCommand: "HGET key field",
			mockErr:       errors.New("something went wrong"),
			wantResult: ScalarResult{
				err: errors.New("could not execute command: something went wrong"),
			},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key", field: "field"},
			expectCtx:     context.Background(),
			expectCommand: "HGET key field",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{
				err: errors.New("internal server error"),
			},
		},
		{
			name:          "server_responds_with_nil",
			args:          args{ctx: context.Background(), key: "key", field: "field"},
			expectCtx:     context.Background(),
			expectCommand: "HGET key field",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.NilCommandReply,
			},
			wantResult: ScalarResult{},
		},
		{
			name:          "server_responds_with_string",
			args:          args{ctx: context.Background(), key: "key", field: "field"},
			expectCtx:     context.Background(),
			expectCommand: "HGET key field",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.StringCommandReply,
				Item:  "string",
			},
			wantResult: ScalarResult{val: stringToPtr("string")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", field: "field"},
			expectCtx:     context.Background(),
			expectCommand: "HGET key field",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.SliceCommandReply,
				Items: []string{"val"},
			},
			wantResult: ScalarResult{err: errors.New("unexpected reply: SLICE")},
		},
		{
			name:          "custon_context",
			args:          args{ctx: contextWithValue("custom_key", "custom_value"), key: "key", field: "field"},
			expectCtx:     contextWithValue("custom_key", "custom_value"),
			expectCommand: "HGET key field",
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
			Expect(tt.expectCtx, &api.ExecuteCommandRequest{Command: tt.expectCommand}).
			Return(tt.mockResponse, tt.mockErr)

		cl := Client{executor: mock}

		res := cl.HGetWithContext(tt.args.ctx, tt.args.key, tt.args.field)
		assert.Equal(t, tt.wantResult, res)
	}
}
