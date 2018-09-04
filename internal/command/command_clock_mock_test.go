package command

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "commandClock" can be found in github.com/namreg/godown/internal/command
*/
import (
	"sync/atomic"
	time "time"

	"github.com/gojuno/minimock"
)

//commandClockMock implements github.com/namreg/godown/internal/command.commandClock
type commandClockMock struct {
	t minimock.Tester

	NowFunc       func() (r time.Time)
	NowCounter    uint64
	NowPreCounter uint64
	NowMock       mcommandClockMockNow
}

//NewcommandClockMock returns a mock for github.com/namreg/godown/internal/command.commandClock
func NewcommandClockMock(t minimock.Tester) *commandClockMock {
	m := &commandClockMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NowMock = mcommandClockMockNow{mock: m}

	return m
}

type mcommandClockMockNow struct {
	mock *commandClockMock
}

//Return sets up a mock for commandClock.Now to return Return's arguments
func (m *mcommandClockMockNow) Return(r time.Time) *commandClockMock {
	m.mock.NowFunc = func() time.Time {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of commandClock.Now method
func (m *mcommandClockMockNow) Set(f func() (r time.Time)) *commandClockMock {
	m.mock.NowFunc = f

	return m.mock
}

//Now implements github.com/namreg/godown/internal/command.commandClock interface
func (m *commandClockMock) Now() (r time.Time) {
	atomic.AddUint64(&m.NowPreCounter, 1)
	defer atomic.AddUint64(&m.NowCounter, 1)

	if m.NowFunc == nil {
		m.t.Fatal("Unexpected call to commandClockMock.Now")
		return
	}

	return m.NowFunc()
}

//NowMinimockCounter returns a count of commandClockMock.NowFunc invocations
func (m *commandClockMock) NowMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.NowCounter)
}

//NowMinimockPreCounter returns the value of commandClockMock.Now invocations
func (m *commandClockMock) NowMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.NowPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *commandClockMock) ValidateCallCounters() {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		m.t.Fatal("Expected call to commandClockMock.Now")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *commandClockMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *commandClockMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *commandClockMock) MinimockFinish() {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		m.t.Fatal("Expected call to commandClockMock.Now")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *commandClockMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *commandClockMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.NowFunc == nil || atomic.LoadUint64(&m.NowCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
				m.t.Error("Expected call to commandClockMock.Now")
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
func (m *commandClockMock) AllMocksCalled() bool {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		return false
	}

	return true
}
