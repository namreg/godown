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

	LockFunc       func()
	LockCounter    uint64
	LockPreCounter uint64
	LockMock       mserverStorageMockLock

	RLockFunc       func()
	RLockCounter    uint64
	RLockPreCounter uint64
	RLockMock       mserverStorageMockRLock

	RUnlockFunc       func()
	RUnlockCounter    uint64
	RUnlockPreCounter uint64
	RUnlockMock       mserverStorageMockRUnlock

	UnlockFunc       func()
	UnlockCounter    uint64
	UnlockPreCounter uint64
	UnlockMock       mserverStorageMockUnlock
}

//NewserverStorageMock returns a mock for github.com/namreg/godown-v2/internal/server.serverStorage
func NewserverStorageMock(t minimock.Tester) *serverStorageMock {
	m := &serverStorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AllWithTTLMock = mserverStorageMockAllWithTTL{mock: m}
	m.DelMock = mserverStorageMockDel{mock: m}
	m.LockMock = mserverStorageMockLock{mock: m}
	m.RLockMock = mserverStorageMockRLock{mock: m}
	m.RUnlockMock = mserverStorageMockRUnlock{mock: m}
	m.UnlockMock = mserverStorageMockUnlock{mock: m}

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

type mserverStorageMockLock struct {
	mock *serverStorageMock
}

//Return sets up a mock for serverStorage.Lock to return Return's arguments
func (m *mserverStorageMockLock) Return() *serverStorageMock {
	m.mock.LockFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of serverStorage.Lock method
func (m *mserverStorageMockLock) Set(f func()) *serverStorageMock {
	m.mock.LockFunc = f

	return m.mock
}

//Lock implements github.com/namreg/godown-v2/internal/server.serverStorage interface
func (m *serverStorageMock) Lock() {
	atomic.AddUint64(&m.LockPreCounter, 1)
	defer atomic.AddUint64(&m.LockCounter, 1)

	if m.LockFunc == nil {
		m.t.Fatal("Unexpected call to serverStorageMock.Lock")
		return
	}

	m.LockFunc()
}

//LockMinimockCounter returns a count of serverStorageMock.LockFunc invocations
func (m *serverStorageMock) LockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.LockCounter)
}

//LockMinimockPreCounter returns the value of serverStorageMock.Lock invocations
func (m *serverStorageMock) LockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.LockPreCounter)
}

type mserverStorageMockRLock struct {
	mock *serverStorageMock
}

//Return sets up a mock for serverStorage.RLock to return Return's arguments
func (m *mserverStorageMockRLock) Return() *serverStorageMock {
	m.mock.RLockFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of serverStorage.RLock method
func (m *mserverStorageMockRLock) Set(f func()) *serverStorageMock {
	m.mock.RLockFunc = f

	return m.mock
}

//RLock implements github.com/namreg/godown-v2/internal/server.serverStorage interface
func (m *serverStorageMock) RLock() {
	atomic.AddUint64(&m.RLockPreCounter, 1)
	defer atomic.AddUint64(&m.RLockCounter, 1)

	if m.RLockFunc == nil {
		m.t.Fatal("Unexpected call to serverStorageMock.RLock")
		return
	}

	m.RLockFunc()
}

//RLockMinimockCounter returns a count of serverStorageMock.RLockFunc invocations
func (m *serverStorageMock) RLockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RLockCounter)
}

//RLockMinimockPreCounter returns the value of serverStorageMock.RLock invocations
func (m *serverStorageMock) RLockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RLockPreCounter)
}

type mserverStorageMockRUnlock struct {
	mock *serverStorageMock
}

//Return sets up a mock for serverStorage.RUnlock to return Return's arguments
func (m *mserverStorageMockRUnlock) Return() *serverStorageMock {
	m.mock.RUnlockFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of serverStorage.RUnlock method
func (m *mserverStorageMockRUnlock) Set(f func()) *serverStorageMock {
	m.mock.RUnlockFunc = f

	return m.mock
}

//RUnlock implements github.com/namreg/godown-v2/internal/server.serverStorage interface
func (m *serverStorageMock) RUnlock() {
	atomic.AddUint64(&m.RUnlockPreCounter, 1)
	defer atomic.AddUint64(&m.RUnlockCounter, 1)

	if m.RUnlockFunc == nil {
		m.t.Fatal("Unexpected call to serverStorageMock.RUnlock")
		return
	}

	m.RUnlockFunc()
}

//RUnlockMinimockCounter returns a count of serverStorageMock.RUnlockFunc invocations
func (m *serverStorageMock) RUnlockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RUnlockCounter)
}

//RUnlockMinimockPreCounter returns the value of serverStorageMock.RUnlock invocations
func (m *serverStorageMock) RUnlockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RUnlockPreCounter)
}

type mserverStorageMockUnlock struct {
	mock *serverStorageMock
}

//Return sets up a mock for serverStorage.Unlock to return Return's arguments
func (m *mserverStorageMockUnlock) Return() *serverStorageMock {
	m.mock.UnlockFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of serverStorage.Unlock method
func (m *mserverStorageMockUnlock) Set(f func()) *serverStorageMock {
	m.mock.UnlockFunc = f

	return m.mock
}

//Unlock implements github.com/namreg/godown-v2/internal/server.serverStorage interface
func (m *serverStorageMock) Unlock() {
	atomic.AddUint64(&m.UnlockPreCounter, 1)
	defer atomic.AddUint64(&m.UnlockCounter, 1)

	if m.UnlockFunc == nil {
		m.t.Fatal("Unexpected call to serverStorageMock.Unlock")
		return
	}

	m.UnlockFunc()
}

//UnlockMinimockCounter returns a count of serverStorageMock.UnlockFunc invocations
func (m *serverStorageMock) UnlockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.UnlockCounter)
}

//UnlockMinimockPreCounter returns the value of serverStorageMock.Unlock invocations
func (m *serverStorageMock) UnlockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.UnlockPreCounter)
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

	if m.LockFunc != nil && atomic.LoadUint64(&m.LockCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.Lock")
	}

	if m.RLockFunc != nil && atomic.LoadUint64(&m.RLockCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.RLock")
	}

	if m.RUnlockFunc != nil && atomic.LoadUint64(&m.RUnlockCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.RUnlock")
	}

	if m.UnlockFunc != nil && atomic.LoadUint64(&m.UnlockCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.Unlock")
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

	if m.LockFunc != nil && atomic.LoadUint64(&m.LockCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.Lock")
	}

	if m.RLockFunc != nil && atomic.LoadUint64(&m.RLockCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.RLock")
	}

	if m.RUnlockFunc != nil && atomic.LoadUint64(&m.RUnlockCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.RUnlock")
	}

	if m.UnlockFunc != nil && atomic.LoadUint64(&m.UnlockCounter) == 0 {
		m.t.Fatal("Expected call to serverStorageMock.Unlock")
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
		ok = ok && (m.LockFunc == nil || atomic.LoadUint64(&m.LockCounter) > 0)
		ok = ok && (m.RLockFunc == nil || atomic.LoadUint64(&m.RLockCounter) > 0)
		ok = ok && (m.RUnlockFunc == nil || atomic.LoadUint64(&m.RUnlockCounter) > 0)
		ok = ok && (m.UnlockFunc == nil || atomic.LoadUint64(&m.UnlockCounter) > 0)

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

			if m.LockFunc != nil && atomic.LoadUint64(&m.LockCounter) == 0 {
				m.t.Error("Expected call to serverStorageMock.Lock")
			}

			if m.RLockFunc != nil && atomic.LoadUint64(&m.RLockCounter) == 0 {
				m.t.Error("Expected call to serverStorageMock.RLock")
			}

			if m.RUnlockFunc != nil && atomic.LoadUint64(&m.RUnlockCounter) == 0 {
				m.t.Error("Expected call to serverStorageMock.RUnlock")
			}

			if m.UnlockFunc != nil && atomic.LoadUint64(&m.UnlockCounter) == 0 {
				m.t.Error("Expected call to serverStorageMock.Unlock")
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

	if m.LockFunc != nil && atomic.LoadUint64(&m.LockCounter) == 0 {
		return false
	}

	if m.RLockFunc != nil && atomic.LoadUint64(&m.RLockCounter) == 0 {
		return false
	}

	if m.RUnlockFunc != nil && atomic.LoadUint64(&m.RUnlockCounter) == 0 {
		return false
	}

	if m.UnlockFunc != nil && atomic.LoadUint64(&m.UnlockCounter) == 0 {
		return false
	}

	return true
}
