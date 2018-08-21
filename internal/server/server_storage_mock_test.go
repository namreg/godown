package server

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "serverStorage" can be found in github.com/namreg/godown-v2/internal/server
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	storage "github.com/namreg/godown-v2/internal/storage"

	testify_assert "github.com/stretchr/testify/assert"
)

//serverStorageMock implements github.com/namreg/godown-v2/internal/server.serverStorage
type serverStorageMock struct {
	t minimock.Tester

	AllWithTTLFunc       func() (r map[storage.Key]*storage.Value, r1 error)
	AllWithTTLCounter    uint64
	AllWithTTLPreCounter uint64
	AllWithTTLMock       mserverStorageMockAllWithTTL

	DelFunc       func(p storage.Key) (r error)
	DelCounter    uint64
	DelPreCounter uint64
	DelMock       mserverStorageMockDel
}

//NewserverStorageMock returns a mock for github.com/namreg/godown-v2/internal/server.serverStorage
func NewserverStorageMock(t minimock.Tester) *serverStorageMock {
	m := &serverStorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AllWithTTLMock = mserverStorageMockAllWithTTL{mock: m}
	m.DelMock = mserverStorageMockDel{mock: m}

	return m
}

type mserverStorageMockAllWithTTL struct {
	mock *serverStorageMock
}

//Return sets up a mock for serverStorage.AllWithTTL to return Return's arguments
func (m *mserverStorageMockAllWithTTL) Return(r map[storage.Key]*storage.Value, r1 error) *serverStorageMock {
	m.mock.AllWithTTLFunc = func() (map[storage.Key]*storage.Value, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of serverStorage.AllWithTTL method
func (m *mserverStorageMockAllWithTTL) Set(f func() (r map[storage.Key]*storage.Value, r1 error)) *serverStorageMock {
	m.mock.AllWithTTLFunc = f

	return m.mock
}

//AllWithTTL implements github.com/namreg/godown-v2/internal/server.serverStorage interface
func (m *serverStorageMock) AllWithTTL() (r map[storage.Key]*storage.Value, r1 error) {
	atomic.AddUint64(&m.AllWithTTLPreCounter, 1)
	defer atomic.AddUint64(&m.AllWithTTLCounter, 1)

	if m.AllWithTTLFunc == nil {
		m.t.Fatal("Unexpected call to serverStorageMock.AllWithTTL")
		return
	}

	return m.AllWithTTLFunc()
}

//AllWithTTLMinimockCounter returns a count of serverStorageMock.AllWithTTLFunc invocations
func (m *serverStorageMock) AllWithTTLMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.AllWithTTLCounter)
}

//AllWithTTLMinimockPreCounter returns the value of serverStorageMock.AllWithTTL invocations
func (m *serverStorageMock) AllWithTTLMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.AllWithTTLPreCounter)
}

type mserverStorageMockDel struct {
	mock             *serverStorageMock
	mockExpectations *serverStorageMockDelParams
}

//serverStorageMockDelParams represents input parameters of the serverStorage.Del
type serverStorageMockDelParams struct {
	p storage.Key
}

//Expect sets up expected params for the serverStorage.Del
func (m *mserverStorageMockDel) Expect(p storage.Key) *mserverStorageMockDel {
	m.mockExpectations = &serverStorageMockDelParams{p}
	return m
}

//Return sets up a mock for serverStorage.Del to return Return's arguments
func (m *mserverStorageMockDel) Return(r error) *serverStorageMock {
	m.mock.DelFunc = func(p storage.Key) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of serverStorage.Del method
func (m *mserverStorageMockDel) Set(f func(p storage.Key) (r error)) *serverStorageMock {
	m.mock.DelFunc = f
	m.mockExpectations = nil
	return m.mock
}

//Del implements github.com/namreg/godown-v2/internal/server.serverStorage interface
func (m *serverStorageMock) Del(p storage.Key) (r error) {
	atomic.AddUint64(&m.DelPreCounter, 1)
	defer atomic.AddUint64(&m.DelCounter, 1)

	if m.DelMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.DelMock.mockExpectations, serverStorageMockDelParams{p},
			"serverStorage.Del got unexpected parameters")

		if m.DelFunc == nil {

			m.t.Fatal("No results are set for the serverStorageMock.Del")

			return
		}
	}

	if m.DelFunc == nil {
		m.t.Fatal("Unexpected call to serverStorageMock.Del")
		return
	}

	return m.DelFunc(p)
}

//DelMinimockCounter returns a count of serverStorageMock.DelFunc invocations
func (m *serverStorageMock) DelMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DelCounter)
}

//DelMinimockPreCounter returns the value of serverStorageMock.Del invocations
func (m *serverStorageMock) DelMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.DelPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *serverStorageMock) ValidateCallCounters() {

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.AllWithTTL")
	}

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.Del")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *serverStorageMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *serverStorageMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *serverStorageMock) MinimockFinish() {

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.AllWithTTL")
	}

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.Del")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *serverStorageMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *serverStorageMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.AllWithTTLFunc == nil || atomic.LoadUint64(&m.AllWithTTLCounter) > 0)
		ok = ok && (m.DelFunc == nil || atomic.LoadUint64(&m.DelCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
				m.t.Error("Expected call to serverStorageMock.AllWithTTL")
			}

			if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
				m.t.Error("Expected call to serverStorageMock.Del")
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
func (m *serverStorageMock) AllMocksCalled() bool {

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		return false
	}

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		return false
	}

	return true
}
