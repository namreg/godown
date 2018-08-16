package command

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.10
The original interface "commandStorage" can be found in github.com/namreg/godown-v2/internal/pkg/command
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	storage "github.com/namreg/godown-v2/internal/pkg/storage"

	testify_assert "github.com/stretchr/testify/assert"
)

//commandStorageMock implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage
type commandStorageMock struct {
	t minimock.Tester

	DelFunc       func(p storage.Key) (r error)
	DelCounter    uint64
	DelPreCounter uint64
	DelMock       mcommandStorageMockDel

	GetFunc       func(p storage.Key) (r *storage.Value, r1 error)
	GetCounter    uint64
	GetPreCounter uint64
	GetMock       mcommandStorageMockGet

	KeysFunc       func() (r []storage.Key, r1 error)
	KeysCounter    uint64
	KeysPreCounter uint64
	KeysMock       mcommandStorageMockKeys

	LockFunc       func()
	LockCounter    uint64
	LockPreCounter uint64
	LockMock       mcommandStorageMockLock

	PutFunc       func(p storage.Key, p1 *storage.Value) (r error)
	PutCounter    uint64
	PutPreCounter uint64
	PutMock       mcommandStorageMockPut

	RLockFunc       func()
	RLockCounter    uint64
	RLockPreCounter uint64
	RLockMock       mcommandStorageMockRLock

	RUnlockFunc       func()
	RUnlockCounter    uint64
	RUnlockPreCounter uint64
	RUnlockMock       mcommandStorageMockRUnlock

	UnlockFunc       func()
	UnlockCounter    uint64
	UnlockPreCounter uint64
	UnlockMock       mcommandStorageMockUnlock
}

//NewcommandStorageMock returns a mock for github.com/namreg/godown-v2/internal/pkg/command.commandStorage
func NewcommandStorageMock(t minimock.Tester) *commandStorageMock {
	m := &commandStorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DelMock = mcommandStorageMockDel{mock: m}
	m.GetMock = mcommandStorageMockGet{mock: m}
	m.KeysMock = mcommandStorageMockKeys{mock: m}
	m.LockMock = mcommandStorageMockLock{mock: m}
	m.PutMock = mcommandStorageMockPut{mock: m}
	m.RLockMock = mcommandStorageMockRLock{mock: m}
	m.RUnlockMock = mcommandStorageMockRUnlock{mock: m}
	m.UnlockMock = mcommandStorageMockUnlock{mock: m}

	return m
}

type mcommandStorageMockDel struct {
	mock             *commandStorageMock
	mockExpectations *commandStorageMockDelParams
}

//commandStorageMockDelParams represents input parameters of the commandStorage.Del
type commandStorageMockDelParams struct {
	p storage.Key
}

//Expect sets up expected params for the commandStorage.Del
func (m *mcommandStorageMockDel) Expect(p storage.Key) *mcommandStorageMockDel {
	m.mockExpectations = &commandStorageMockDelParams{p}
	return m
}

//Return sets up a mock for commandStorage.Del to return Return's arguments
func (m *mcommandStorageMockDel) Return(r error) *commandStorageMock {
	m.mock.DelFunc = func(p storage.Key) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.Del method
func (m *mcommandStorageMockDel) Set(f func(p storage.Key) (r error)) *commandStorageMock {
	m.mock.DelFunc = f
	return m.mock
}

//Del implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage interface
func (m *commandStorageMock) Del(p storage.Key) (r error) {
	atomic.AddUint64(&m.DelPreCounter, 1)
	defer atomic.AddUint64(&m.DelCounter, 1)

	if m.DelMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.DelMock.mockExpectations, commandStorageMockDelParams{p},
			"commandStorage.Del got unexpected parameters")

		if m.DelFunc == nil {

			m.t.Fatal("No results are set for the commandStorageMock.Del")

			return
		}
	}

	if m.DelFunc == nil {
		m.t.Fatal("Unexpected call to commandStorageMock.Del")
		return
	}

	return m.DelFunc(p)
}

//DelMinimockCounter returns a count of commandStorageMock.DelFunc invocations
func (m *commandStorageMock) DelMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DelCounter)
}

//DelMinimockPreCounter returns the value of commandStorageMock.Del invocations
func (m *commandStorageMock) DelMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.DelPreCounter)
}

type mcommandStorageMockGet struct {
	mock             *commandStorageMock
	mockExpectations *commandStorageMockGetParams
}

//commandStorageMockGetParams represents input parameters of the commandStorage.Get
type commandStorageMockGetParams struct {
	p storage.Key
}

//Expect sets up expected params for the commandStorage.Get
func (m *mcommandStorageMockGet) Expect(p storage.Key) *mcommandStorageMockGet {
	m.mockExpectations = &commandStorageMockGetParams{p}
	return m
}

//Return sets up a mock for commandStorage.Get to return Return's arguments
func (m *mcommandStorageMockGet) Return(r *storage.Value, r1 error) *commandStorageMock {
	m.mock.GetFunc = func(p storage.Key) (*storage.Value, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.Get method
func (m *mcommandStorageMockGet) Set(f func(p storage.Key) (r *storage.Value, r1 error)) *commandStorageMock {
	m.mock.GetFunc = f
	return m.mock
}

//Get implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage interface
func (m *commandStorageMock) Get(p storage.Key) (r *storage.Value, r1 error) {
	atomic.AddUint64(&m.GetPreCounter, 1)
	defer atomic.AddUint64(&m.GetCounter, 1)

	if m.GetMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.GetMock.mockExpectations, commandStorageMockGetParams{p},
			"commandStorage.Get got unexpected parameters")

		if m.GetFunc == nil {

			m.t.Fatal("No results are set for the commandStorageMock.Get")

			return
		}
	}

	if m.GetFunc == nil {
		m.t.Fatal("Unexpected call to commandStorageMock.Get")
		return
	}

	return m.GetFunc(p)
}

//GetMinimockCounter returns a count of commandStorageMock.GetFunc invocations
func (m *commandStorageMock) GetMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetCounter)
}

//GetMinimockPreCounter returns the value of commandStorageMock.Get invocations
func (m *commandStorageMock) GetMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetPreCounter)
}

type mcommandStorageMockKeys struct {
	mock *commandStorageMock
}

//Return sets up a mock for commandStorage.Keys to return Return's arguments
func (m *mcommandStorageMockKeys) Return(r []storage.Key, r1 error) *commandStorageMock {
	m.mock.KeysFunc = func() ([]storage.Key, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.Keys method
func (m *mcommandStorageMockKeys) Set(f func() (r []storage.Key, r1 error)) *commandStorageMock {
	m.mock.KeysFunc = f
	return m.mock
}

//Keys implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage interface
func (m *commandStorageMock) Keys() (r []storage.Key, r1 error) {
	atomic.AddUint64(&m.KeysPreCounter, 1)
	defer atomic.AddUint64(&m.KeysCounter, 1)

	if m.KeysFunc == nil {
		m.t.Fatal("Unexpected call to commandStorageMock.Keys")
		return
	}

	return m.KeysFunc()
}

//KeysMinimockCounter returns a count of commandStorageMock.KeysFunc invocations
func (m *commandStorageMock) KeysMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.KeysCounter)
}

//KeysMinimockPreCounter returns the value of commandStorageMock.Keys invocations
func (m *commandStorageMock) KeysMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.KeysPreCounter)
}

type mcommandStorageMockLock struct {
	mock *commandStorageMock
}

//Return sets up a mock for commandStorage.Lock to return Return's arguments
func (m *mcommandStorageMockLock) Return() *commandStorageMock {
	m.mock.LockFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.Lock method
func (m *mcommandStorageMockLock) Set(f func()) *commandStorageMock {
	m.mock.LockFunc = f
	return m.mock
}

//Lock implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage interface
func (m *commandStorageMock) Lock() {
	atomic.AddUint64(&m.LockPreCounter, 1)
	defer atomic.AddUint64(&m.LockCounter, 1)

	if m.LockFunc == nil {
		m.t.Fatal("Unexpected call to commandStorageMock.Lock")
		return
	}

	m.LockFunc()
}

//LockMinimockCounter returns a count of commandStorageMock.LockFunc invocations
func (m *commandStorageMock) LockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.LockCounter)
}

//LockMinimockPreCounter returns the value of commandStorageMock.Lock invocations
func (m *commandStorageMock) LockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.LockPreCounter)
}

type mcommandStorageMockPut struct {
	mock             *commandStorageMock
	mockExpectations *commandStorageMockPutParams
}

//commandStorageMockPutParams represents input parameters of the commandStorage.Put
type commandStorageMockPutParams struct {
	p  storage.Key
	p1 *storage.Value
}

//Expect sets up expected params for the commandStorage.Put
func (m *mcommandStorageMockPut) Expect(p storage.Key, p1 *storage.Value) *mcommandStorageMockPut {
	m.mockExpectations = &commandStorageMockPutParams{p, p1}
	return m
}

//Return sets up a mock for commandStorage.Put to return Return's arguments
func (m *mcommandStorageMockPut) Return(r error) *commandStorageMock {
	m.mock.PutFunc = func(p storage.Key, p1 *storage.Value) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.Put method
func (m *mcommandStorageMockPut) Set(f func(p storage.Key, p1 *storage.Value) (r error)) *commandStorageMock {
	m.mock.PutFunc = f
	return m.mock
}

//Put implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage interface
func (m *commandStorageMock) Put(p storage.Key, p1 *storage.Value) (r error) {
	atomic.AddUint64(&m.PutPreCounter, 1)
	defer atomic.AddUint64(&m.PutCounter, 1)

	if m.PutMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PutMock.mockExpectations, commandStorageMockPutParams{p, p1},
			"commandStorage.Put got unexpected parameters")

		if m.PutFunc == nil {

			m.t.Fatal("No results are set for the commandStorageMock.Put")

			return
		}
	}

	if m.PutFunc == nil {
		m.t.Fatal("Unexpected call to commandStorageMock.Put")
		return
	}

	return m.PutFunc(p, p1)
}

//PutMinimockCounter returns a count of commandStorageMock.PutFunc invocations
func (m *commandStorageMock) PutMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PutCounter)
}

//PutMinimockPreCounter returns the value of commandStorageMock.Put invocations
func (m *commandStorageMock) PutMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.PutPreCounter)
}

type mcommandStorageMockRLock struct {
	mock *commandStorageMock
}

//Return sets up a mock for commandStorage.RLock to return Return's arguments
func (m *mcommandStorageMockRLock) Return() *commandStorageMock {
	m.mock.RLockFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.RLock method
func (m *mcommandStorageMockRLock) Set(f func()) *commandStorageMock {
	m.mock.RLockFunc = f
	return m.mock
}

//RLock implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage interface
func (m *commandStorageMock) RLock() {
	atomic.AddUint64(&m.RLockPreCounter, 1)
	defer atomic.AddUint64(&m.RLockCounter, 1)

	if m.RLockFunc == nil {
		m.t.Fatal("Unexpected call to commandStorageMock.RLock")
		return
	}

	m.RLockFunc()
}

//RLockMinimockCounter returns a count of commandStorageMock.RLockFunc invocations
func (m *commandStorageMock) RLockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RLockCounter)
}

//RLockMinimockPreCounter returns the value of commandStorageMock.RLock invocations
func (m *commandStorageMock) RLockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RLockPreCounter)
}

type mcommandStorageMockRUnlock struct {
	mock *commandStorageMock
}

//Return sets up a mock for commandStorage.RUnlock to return Return's arguments
func (m *mcommandStorageMockRUnlock) Return() *commandStorageMock {
	m.mock.RUnlockFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.RUnlock method
func (m *mcommandStorageMockRUnlock) Set(f func()) *commandStorageMock {
	m.mock.RUnlockFunc = f
	return m.mock
}

//RUnlock implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage interface
func (m *commandStorageMock) RUnlock() {
	atomic.AddUint64(&m.RUnlockPreCounter, 1)
	defer atomic.AddUint64(&m.RUnlockCounter, 1)

	if m.RUnlockFunc == nil {
		m.t.Fatal("Unexpected call to commandStorageMock.RUnlock")
		return
	}

	m.RUnlockFunc()
}

//RUnlockMinimockCounter returns a count of commandStorageMock.RUnlockFunc invocations
func (m *commandStorageMock) RUnlockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RUnlockCounter)
}

//RUnlockMinimockPreCounter returns the value of commandStorageMock.RUnlock invocations
func (m *commandStorageMock) RUnlockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RUnlockPreCounter)
}

type mcommandStorageMockUnlock struct {
	mock *commandStorageMock
}

//Return sets up a mock for commandStorage.Unlock to return Return's arguments
func (m *mcommandStorageMockUnlock) Return() *commandStorageMock {
	m.mock.UnlockFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.Unlock method
func (m *mcommandStorageMockUnlock) Set(f func()) *commandStorageMock {
	m.mock.UnlockFunc = f
	return m.mock
}

//Unlock implements github.com/namreg/godown-v2/internal/pkg/command.commandStorage interface
func (m *commandStorageMock) Unlock() {
	atomic.AddUint64(&m.UnlockPreCounter, 1)
	defer atomic.AddUint64(&m.UnlockCounter, 1)

	if m.UnlockFunc == nil {
		m.t.Fatal("Unexpected call to commandStorageMock.Unlock")
		return
	}

	m.UnlockFunc()
}

//UnlockMinimockCounter returns a count of commandStorageMock.UnlockFunc invocations
func (m *commandStorageMock) UnlockMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.UnlockCounter)
}

//UnlockMinimockPreCounter returns the value of commandStorageMock.Unlock invocations
func (m *commandStorageMock) UnlockMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.UnlockPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *commandStorageMock) ValidateCallCounters() {

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Del")
	}

	if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Get")
	}

	if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Keys")
	}

	if m.LockFunc != nil && atomic.LoadUint64(&m.LockCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Lock")
	}

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Put")
	}

	if m.RLockFunc != nil && atomic.LoadUint64(&m.RLockCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.RLock")
	}

	if m.RUnlockFunc != nil && atomic.LoadUint64(&m.RUnlockCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.RUnlock")
	}

	if m.UnlockFunc != nil && atomic.LoadUint64(&m.UnlockCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Unlock")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *commandStorageMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *commandStorageMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *commandStorageMock) MinimockFinish() {

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Del")
	}

	if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Get")
	}

	if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Keys")
	}

	if m.LockFunc != nil && atomic.LoadUint64(&m.LockCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Lock")
	}

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Put")
	}

	if m.RLockFunc != nil && atomic.LoadUint64(&m.RLockCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.RLock")
	}

	if m.RUnlockFunc != nil && atomic.LoadUint64(&m.RUnlockCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.RUnlock")
	}

	if m.UnlockFunc != nil && atomic.LoadUint64(&m.UnlockCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Unlock")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *commandStorageMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *commandStorageMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.DelFunc == nil || atomic.LoadUint64(&m.DelCounter) > 0)
		ok = ok && (m.GetFunc == nil || atomic.LoadUint64(&m.GetCounter) > 0)
		ok = ok && (m.KeysFunc == nil || atomic.LoadUint64(&m.KeysCounter) > 0)
		ok = ok && (m.LockFunc == nil || atomic.LoadUint64(&m.LockCounter) > 0)
		ok = ok && (m.PutFunc == nil || atomic.LoadUint64(&m.PutCounter) > 0)
		ok = ok && (m.RLockFunc == nil || atomic.LoadUint64(&m.RLockCounter) > 0)
		ok = ok && (m.RUnlockFunc == nil || atomic.LoadUint64(&m.RUnlockCounter) > 0)
		ok = ok && (m.UnlockFunc == nil || atomic.LoadUint64(&m.UnlockCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.Del")
			}

			if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.Get")
			}

			if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.Keys")
			}

			if m.LockFunc != nil && atomic.LoadUint64(&m.LockCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.Lock")
			}

			if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.Put")
			}

			if m.RLockFunc != nil && atomic.LoadUint64(&m.RLockCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.RLock")
			}

			if m.RUnlockFunc != nil && atomic.LoadUint64(&m.RUnlockCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.RUnlock")
			}

			if m.UnlockFunc != nil && atomic.LoadUint64(&m.UnlockCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.Unlock")
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
func (m *commandStorageMock) AllMocksCalled() bool {

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		return false
	}

	if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
		return false
	}

	if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
		return false
	}

	if m.LockFunc != nil && atomic.LoadUint64(&m.LockCounter) == 0 {
		return false
	}

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
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
