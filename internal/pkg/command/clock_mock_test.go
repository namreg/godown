package command

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Clock" can be found in github.com/namreg/godown-v2/pkg/clock
*/
import (
	"sync/atomic"
	time "time"

	"github.com/gojuno/minimock"
)

//ClockMock implements github.com/namreg/godown-v2/pkg/clock.Clock
type ClockMock struct {
	t minimock.Tester

	NowFunc       func() (r time.Time)
	NowCounter    uint64
	NowPreCounter uint64
	NowMock       mClockMockNow
}

//NewClockMock returns a mock for github.com/namreg/godown-v2/pkg/clock.Clock
func NewClockMock(t minimock.Tester) *ClockMock {
	m := &ClockMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NowMock = mClockMockNow{mock: m}

	return m
}

type mClockMockNow struct {
	mock *ClockMock
}

//Return sets up a mock for Clock.Now to return Return's arguments
func (m *mClockMockNow) Return(r time.Time) *ClockMock {
	m.mock.NowFunc = func() time.Time {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Clock.Now method
func (m *mClockMockNow) Set(f func() (r time.Time)) *ClockMock {
	m.mock.NowFunc = f
	return m.mock
}

//Now implements github.com/namreg/godown-v2/pkg/clock.Clock interface
func (m *ClockMock) Now() (r time.Time) {
	atomic.AddUint64(&m.NowPreCounter, 1)
	defer atomic.AddUint64(&m.NowCounter, 1)

	if m.NowFunc == nil {
		m.t.Fatal("Unexpected call to ClockMock.Now")
		return
	}

	return m.NowFunc()
}

//NowMinimockCounter returns a count of ClockMock.NowFunc invocations
func (m *ClockMock) NowMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.NowCounter)
}

//NowMinimockPreCounter returns the value of ClockMock.Now invocations
func (m *ClockMock) NowMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.NowPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ClockMock) ValidateCallCounters() {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		m.t.Fatal("Expected call to ClockMock.Now")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ClockMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *ClockMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *ClockMock) MinimockFinish() {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		m.t.Fatal("Expected call to ClockMock.Now")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *ClockMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *ClockMock) MinimockWait(timeout time.Duration) {
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
				m.t.Error("Expected call to ClockMock.Now")
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
func (m *ClockMock) AllMocksCalled() bool {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		return false
	}

	return true
}
