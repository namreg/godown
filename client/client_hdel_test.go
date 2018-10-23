package client

import (
	context "context"
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestClient_HDel(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		key    string
		fields []string
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
			args:          args{key: "key", fields: []string{"field1", "field2"}},
			expectCommand: "HDEL key field1 field2",
			mockErr:       errors.New("something went wrong"),
			wantResult:    ScalarResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_responds_with_error",
			args:          args{key: "key", fields: []string{"field1", "field2"}},
			expectCommand: "HDEL key field1 field2",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_int",
			args:          args{key: "key", fields: []string{"field1", "field2"}},
			expectCommand: "HDEL key field1 field2",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "2",
			},
			wantResult: ScalarResult{val: stringToPtr("2")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{key: "key", fields: []string{"field1", "field2"}},
			expectCommand: "HDEL key field1 field2",
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

			res := cl.HDel(tt.args.key, tt.args.fields[0], tt.args.fields[1:]...)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}

func TestClient_HDelWithContext(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	type args struct {
		ctx    context.Context
		key    string
		fields []string
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
			args:          args{ctx: context.Background(), key: "key", fields: []string{"field1", "field2"}},
			expectCtx:     context.Background(),
			expectCommand: "HDEL key field1 field2",
			mockErr:       errors.New("something went wrong"),
			wantResult:    ScalarResult{err: errors.New("could not execute command: something went wrong")},
		},
		{
			name:          "server_responds_with_error",
			args:          args{ctx: context.Background(), key: "key", fields: []string{"field1", "field2"}},
			expectCtx:     context.Background(),
			expectCommand: "HDEL key field1 field2",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.ErrCommandReply,
				Item:  "internal server error",
			},
			wantResult: ScalarResult{err: errors.New("internal server error")},
		},
		{
			name:          "server_responds_with_int",
			args:          args{ctx: context.Background(), key: "key", fields: []string{"field1", "field2"}},
			expectCtx:     context.Background(),
			expectCommand: "HDEL key field1 field2",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "2",
			},
			wantResult: ScalarResult{val: stringToPtr("2")},
		},
		{
			name:          "custom_context",
			args:          args{ctx: contextWithValue("key", "value"), key: "key", fields: []string{"field1", "field2"}},
			expectCtx:     contextWithValue("key", "value"),
			expectCommand: "HDEL key field1 field2",
			mockResponse: &api.ExecuteCommandResponse{
				Reply: api.IntCommandReply,
				Item:  "2",
			},
			wantResult: ScalarResult{val: stringToPtr("2")},
		},
		{
			name:          "server_responds_with_unexpected_reply",
			args:          args{ctx: context.Background(), key: "key", fields: []string{"field1", "field2"}},
			expectCtx:     context.Background(),
			expectCommand: "HDEL key field1 field2",
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

			res := cl.HDelWithContext(tt.args.ctx, tt.args.key, tt.args.fields[0], tt.args.fields[1:]...)
			assert.Equal(t, tt.wantResult, res)
		})
	}
}
