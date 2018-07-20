package command

import (
	"errors"
	"testing"

	"github.com/gojuno/minimock"
	"github.com/namreg/godown-v2/internal/pkg/storage/memory"

	"github.com/stretchr/testify/assert"
)

func TestHelp_Name(t *testing.T) {
	cmd := new(Help)
	assert.Equal(t, "HELP", cmd.Name())
}

func TestHelp_Help(t *testing.T) {
	cmd := new(Help)
	expected := `Usage: HELP command
Show the usage of the given command`
	assert.Equal(t, expected, cmd.Help())
}

func TestHelp_Execute(t *testing.T) {
	strg := memory.New(nil)

	mc := minimock.NewController(t)
	defer mc.Finish()

	mock := NewCommandMock(t)
	mock.HelpMock.Return("help message")
	commands["MOCK"] = mock

	tests := []struct {
		name string
		args []string
		want Result
	}{
		{"existing_command", []string{"mock"}, HelpResult{mock}},
		{"not_existing_command", []string{"not_existing_command"}, ErrResult{errors.New(`command "not_existing_command" not found`)}},
		{"wrong_number_of_args/1", []string{}, ErrResult{ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"mock", "mock"}, ErrResult{ErrWrongArgsNumber}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := new(Help)
			res := cmd.Execute(strg, tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}
