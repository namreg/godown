package server

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "serverClock" can be found in github.com/namreg/godown/internal/server
*/
import (
	"sync/atomic"
	time "time"

	"github.com/gojuno/minimock"
)

//serverClockMock implements github.com/namreg/godown/internal/server.serverClock
type serverClockMock struct {
	t minimock.Tester

	NowFunc       func() (r time.Time)
	NowCounter    uint64
	NowPreCounter uint64
	NowMock       mserverClockMockNow
}

//NewserverClockMock returns a mock for github.com/namreg/godown/internal/server.serverClock
func NewserverClockMock(t minimock.Tester) *serverClockMock {
	m := &serverClockMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NowMock = mserverClockMockNow{mock: m}

	return m
}

type mserverClockMockNow struct {
	mock *serverClockMock
}

//Return sets up a mock for serverClock.Now to return Return's arguments
func (m *mserverClockMockNow) Return(r time.Time) *serverClockMock {
	m.mock.NowFunc = func() time.Time {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of serverClock.Now method
func (m *mserverClockMockNow) Set(f func() (r time.Time)) *serverClockMock {
	m.mock.NowFunc = f

	return m.mock
}

//Now implements github.com/namreg/godown/internal/server.serverClock interface
func (m *serverClockMock) Now() (r time.Time) {
	atomic.AddUint64(&m.NowPreCounter, 1)
	defer atomic.AddUint64(&m.NowCounter, 1)

	if m.NowFunc == nil {
		m.t.Fatal("Unexpected call to serverClockMock.Now")
		return
	}

	return m.NowFunc()
}

//NowMinimockCounter returns a count of serverClockMock.NowFunc invocations
func (m *serverClockMock) NowMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.NowCounter)
}

//NowMinimockPreCounter returns the value of serverClockMock.Now invocations
func (m *serverClockMock) NowMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.NowPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *serverClockMock) ValidateCallCounters() {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		m.t.Fatal("Expected call to serverClockMock.Now")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *serverClockMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *serverClockMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *serverClockMock) MinimockFinish() {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		m.t.Fatal("Expected call to serverClockMock.Now")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *serverClockMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *serverClockMock) MinimockWait(timeout time.Duration) {
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
				m.t.Error("Expected call to serverClockMock.Now")
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
func (m *serverClockMock) AllMocksCalled() bool {

	if m.NowFunc != nil && atomic.LoadUint64(&m.NowCounter) == 0 {
		return false
	}

	return true
}
