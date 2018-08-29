package command

import (
	"errors"
	"testing"

	"github.com/gojuno/minimock"

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
	mc := minimock.NewController(t)
	defer mc.Finish()

	parserMock := NewcommandParserMock(mc)
	parserMock.ParseFunc = func(str string) (Command, []string, error) {
		if str == "mock" {
			return NewCommandMock(mc).HelpMock.Return("help message"), nil, nil
		}
		return nil, nil, ErrCommandNotFound
	}

	tests := []struct {
		name string
		args []string
		want Reply
	}{
		{"existing_command", []string{"mock"}, RawStringReply{Value: "help message"}},
		{"not_existing_command", []string{"not_existing_command"}, ErrReply{Value: errors.New(`command "not_existing_command" not found`)}},
		{"wrong_number_of_args/1", []string{}, ErrReply{Value: ErrWrongArgsNumber}},
		{"wrong_number_of_args/2", []string{"mock", "mock"}, ErrReply{Value: ErrWrongArgsNumber}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Help{parser: parserMock}
			res := cmd.Execute(tt.args...)
			assert.Equal(t, tt.want, res)
		})
	}
}
