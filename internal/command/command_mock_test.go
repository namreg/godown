package command

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Command" can be found in github.com/namreg/godown/internal/command
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	testify_assert "github.com/stretchr/testify/assert"
)

//CommandMock implements github.com/namreg/godown/internal/command.Command
type CommandMock struct {
	t minimock.Tester

	ExecuteFunc       func(p ...string) (r Reply)
	ExecuteCounter    uint64
	ExecutePreCounter uint64
	ExecuteMock       mCommandMockExecute

	HelpFunc       func() (r string)
	HelpCounter    uint64
	HelpPreCounter uint64
	HelpMock       mCommandMockHelp

	NameFunc       func() (r string)
	NameCounter    uint64
	NamePreCounter uint64
	NameMock       mCommandMockName
}

//NewCommandMock returns a mock for github.com/namreg/godown/internal/command.Command
func NewCommandMock(t minimock.Tester) *CommandMock {
	m := &CommandMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ExecuteMock = mCommandMockExecute{mock: m}
	m.HelpMock = mCommandMockHelp{mock: m}
	m.NameMock = mCommandMockName{mock: m}

	return m
}

type mCommandMockExecute struct {
	mock             *CommandMock
	mockExpectations *CommandMockExecuteParams
}

//CommandMockExecuteParams represents input parameters of the Command.Execute
type CommandMockExecuteParams struct {
	p []string
}

//Expect sets up expected params for the Command.Execute
func (m *mCommandMockExecute) Expect(p ...string) *mCommandMockExecute {
	m.mockExpectations = &CommandMockExecuteParams{p}
	return m
}

//Return sets up a mock for Command.Execute to return Return's arguments
func (m *mCommandMockExecute) Return(r Reply) *CommandMock {
	m.mock.ExecuteFunc = func(p ...string) Reply {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Command.Execute method
func (m *mCommandMockExecute) Set(f func(p ...string) (r Reply)) *CommandMock {
	m.mock.ExecuteFunc = f
	m.mockExpectations = nil
	return m.mock
}

//Execute implements github.com/namreg/godown/internal/command.Command interface
func (m *CommandMock) Execute(p ...string) (r Reply) {
	atomic.AddUint64(&m.ExecutePreCounter, 1)
	defer atomic.AddUint64(&m.ExecuteCounter, 1)

	if m.ExecuteMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ExecuteMock.mockExpectations, CommandMockExecuteParams{p},
			"Command.Execute got unexpected parameters")

		if m.ExecuteFunc == nil {

			m.t.Fatal("No results are set for the CommandMock.Execute")

			return
		}
	}

	if m.ExecuteFunc == nil {
		m.t.Fatal("Unexpected call to CommandMock.Execute")
		return
	}

	return m.ExecuteFunc(p...)
}

//ExecuteMinimockCounter returns a count of CommandMock.ExecuteFunc invocations
func (m *CommandMock) ExecuteMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ExecuteCounter)
}

//ExecuteMinimockPreCounter returns the value of CommandMock.Execute invocations
func (m *CommandMock) ExecuteMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ExecutePreCounter)
}

type mCommandMockHelp struct {
	mock *CommandMock
}

//Return sets up a mock for Command.Help to return Return's arguments
func (m *mCommandMockHelp) Return(r string) *CommandMock {
	m.mock.HelpFunc = func() string {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Command.Help method
func (m *mCommandMockHelp) Set(f func() (r string)) *CommandMock {
	m.mock.HelpFunc = f

	return m.mock
}

//Help implements github.com/namreg/godown/internal/command.Command interface
func (m *CommandMock) Help() (r string) {
	atomic.AddUint64(&m.HelpPreCounter, 1)
	defer atomic.AddUint64(&m.HelpCounter, 1)

	if m.HelpFunc == nil {
		m.t.Fatal("Unexpected call to CommandMock.Help")
		return
	}

	return m.HelpFunc()
}

//HelpMinimockCounter returns a count of CommandMock.HelpFunc invocations
func (m *CommandMock) HelpMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.HelpCounter)
}

//HelpMinimockPreCounter returns the value of CommandMock.Help invocations
func (m *CommandMock) HelpMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.HelpPreCounter)
}

type mCommandMockName struct {
	mock *CommandMock
}

//Return sets up a mock for Command.Name to return Return's arguments
func (m *mCommandMockName) Return(r string) *CommandMock {
	m.mock.NameFunc = func() string {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Command.Name method
func (m *mCommandMockName) Set(f func() (r string)) *CommandMock {
	m.mock.NameFunc = f

	return m.mock
}

//Name implements github.com/namreg/godown/internal/command.Command interface
func (m *CommandMock) Name() (r string) {
	atomic.AddUint64(&m.NamePreCounter, 1)
	defer atomic.AddUint64(&m.NameCounter, 1)

	if m.NameFunc == nil {
		m.t.Fatal("Unexpected call to CommandMock.Name")
		return
	}

	return m.NameFunc()
}

//NameMinimockCounter returns a count of CommandMock.NameFunc invocations
func (m *CommandMock) NameMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.NameCounter)
}

//NameMinimockPreCounter returns the value of CommandMock.Name invocations
func (m *CommandMock) NameMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.NamePreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *CommandMock) ValidateCallCounters() {

	if m.ExecuteFunc != nil && atomic.LoadUint64(&m.ExecuteCounter) == 0 {
		m.t.Fatal("Expected call to CommandMock.Execute")
	}

	if m.HelpFunc != nil && atomic.LoadUint64(&m.HelpCounter) == 0 {
		m.t.Fatal("Expected call to CommandMock.Help")
	}

	if m.NameFunc != nil && atomic.LoadUint64(&m.NameCounter) == 0 {
		m.t.Fatal("Expected call to CommandMock.Name")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *CommandMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *CommandMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *CommandMock) MinimockFinish() {

	if m.ExecuteFunc != nil && atomic.LoadUint64(&m.ExecuteCounter) == 0 {
		m.t.Fatal("Expected call to CommandMock.Execute")
	}

	if m.HelpFunc != nil && atomic.LoadUint64(&m.HelpCounter) == 0 {
		m.t.Fatal("Expected call to CommandMock.Help")
	}

	if m.NameFunc != nil && atomic.LoadUint64(&m.NameCounter) == 0 {
		m.t.Fatal("Expected call to CommandMock.Name")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *CommandMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *CommandMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.ExecuteFunc == nil || atomic.LoadUint64(&m.ExecuteCounter) > 0)
		ok = ok && (m.HelpFunc == nil || atomic.LoadUint64(&m.HelpCounter) > 0)
		ok = ok && (m.NameFunc == nil || atomic.LoadUint64(&m.NameCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.ExecuteFunc != nil && atomic.LoadUint64(&m.ExecuteCounter) == 0 {
				m.t.Error("Expected call to CommandMock.Execute")
			}

			if m.HelpFunc != nil && atomic.LoadUint64(&m.HelpCounter) == 0 {
				m.t.Error("Expected call to CommandMock.Help")
			}

			if m.NameFunc != nil && atomic.LoadUint64(&m.NameCounter) == 0 {
				m.t.Error("Expected call to CommandMock.Name")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *CommandMock) AllMocksCalled() bool {

	if m.ExecuteFunc != nil && atomic.LoadUint64(&m.ExecuteCounter) == 0 {
		return false
	}

	if m.HelpFunc != nil && atomic.LoadUint64(&m.HelpCounter) == 0 {
		return false
	}

	if m.NameFunc != nil && atomic.LoadUint64(&m.NameCounter) == 0 {
		return false
	}

	return true
}
