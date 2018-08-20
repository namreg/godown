package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantCmd  Command
		wantArgs []string
		wantErr  error
	}{
		{"uppercase", "GET key", new(Get), []string{"key"}, nil},
		{"lowercase", "get key", new(Get), []string{"key"}, nil},
		{"not_existing_command", "command_not_exists arg", nil, nil, ErrCommandNotFound},
		{"args_with_space", `SET key "value with space"`, new(Set), []string{"key", "value with space"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{strg: NewStorageMock(t)}
			cmd, args, err := p.Parse(tt.input)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.IsType(t, tt.wantCmd, cmd)
			assert.Equal(t, tt.wantArgs, args)
		})
	}
}
