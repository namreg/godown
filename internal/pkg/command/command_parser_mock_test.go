package command

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.10
The original interface "commandParser" can be found in github.com/namreg/godown-v2/internal/pkg/command
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	testify_assert "github.com/stretchr/testify/assert"
)

//commandParserMock implements github.com/namreg/godown-v2/internal/pkg/command.commandParser
type commandParserMock struct {
	t minimock.Tester

	ParseFunc       func(p string) (r Command, r1 []string, r2 error)
	ParseCounter    uint64
	ParsePreCounter uint64
	ParseMock       mcommandParserMockParse
}

//NewcommandParserMock returns a mock for github.com/namreg/godown-v2/internal/pkg/command.commandParser
func NewcommandParserMock(t minimock.Tester) *commandParserMock {
	m := &commandParserMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ParseMock = mcommandParserMockParse{mock: m}

	return m
}

type mcommandParserMockParse struct {
	mock             *commandParserMock
	mockExpectations *commandParserMockParseParams
}

//commandParserMockParseParams represents input parameters of the commandParser.Parse
type commandParserMockParseParams struct {
	p string
}

//Expect sets up expected params for the commandParser.Parse
func (m *mcommandParserMockParse) Expect(p string) *mcommandParserMockParse {
	m.mockExpectations = &commandParserMockParseParams{p}
	return m
}

//Return sets up a mock for commandParser.Parse to return Return's arguments
func (m *mcommandParserMockParse) Return(r Command, r1 []string, r2 error) *commandParserMock {
	m.mock.ParseFunc = func(p string) (Command, []string, error) {
		return r, r1, r2
	}
	return m.mock
}

//Set uses given function f as a mock of commandParser.Parse method
func (m *mcommandParserMockParse) Set(f func(p string) (r Command, r1 []string, r2 error)) *commandParserMock {
	m.mock.ParseFunc = f
	return m.mock
}

//Parse implements github.com/namreg/godown-v2/internal/pkg/command.commandParser interface
func (m *commandParserMock) Parse(p string) (r Command, r1 []string, r2 error) {
	atomic.AddUint64(&m.ParsePreCounter, 1)
	defer atomic.AddUint64(&m.ParseCounter, 1)

	if m.ParseMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ParseMock.mockExpectations, commandParserMockParseParams{p},
			"commandParser.Parse got unexpected parameters")

		if m.ParseFunc == nil {

			m.t.Fatal("No results are set for the commandParserMock.Parse")

			return
		}
	}

	if m.ParseFunc == nil {
		m.t.Fatal("Unexpected call to commandParserMock.Parse")
		return
	}

	return m.ParseFunc(p)
}

//ParseMinimockCounter returns a count of commandParserMock.ParseFunc invocations
func (m *commandParserMock) ParseMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ParseCounter)
}

//ParseMinimockPreCounter returns the value of commandParserMock.Parse invocations
func (m *commandParserMock) ParseMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ParsePreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *commandParserMock) ValidateCallCounters() {

	if m.ParseFunc != nil && atomic.LoadUint64(&m.ParseCounter) == 0 {
		m.t.Fatal("Expected call to commandParserMock.Parse")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *commandParserMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *commandParserMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *commandParserMock) MinimockFinish() {

	if m.ParseFunc != nil && atomic.LoadUint64(&m.ParseCounter) == 0 {
		m.t.Fatal("Expected call to commandParserMock.Parse")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *commandParserMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *commandParserMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.ParseFunc == nil || atomic.LoadUint64(&m.ParseCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.ParseFunc != nil && atomic.LoadUint64(&m.ParseCounter) == 0 {
				m.t.Error("Expected call to commandParserMock.Parse")
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
func (m *commandParserMock) AllMocksCalled() bool {

	if m.ParseFunc != nil && atomic.LoadUint64(&m.ParseCounter) == 0 {
		return false
	}

	return true
}
