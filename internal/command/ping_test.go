package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing_Name(t *testing.T) {
	cmd := new(Ping)
	assert.Equal(t, "PING", cmd.Name())
}

func TestPing_Help(t *testing.T) {
	cmd := new(Ping)
	expected := `Usage: PING [message]
Returns PONG if no argument is provided, otherwise return a copy of the argument as a bulk.`
	assert.Equal(t, expected, cmd.Help())
}

func TestPing_Execute(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"no_args", []string{}, RawStringReply{Value: "PONG"}},
		{"with_args", []string{"hello", "world"}, RawStringReply{Value: "hello world"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Ping)
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}
