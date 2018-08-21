package command

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "commandStorage" can be found in github.com/namreg/godown-v2/internal/command
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	storage "github.com/namreg/godown-v2/internal/storage"

	testify_assert "github.com/stretchr/testify/assert"
)

//commandStorageMock implements github.com/namreg/godown-v2/internal/command.commandStorage
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

	PutFunc       func(p storage.Key, p1 storage.ValueSetter) (r error)
	PutCounter    uint64
	PutPreCounter uint64
	PutMock       mcommandStorageMockPut
}

//NewcommandStorageMock returns a mock for github.com/namreg/godown-v2/internal/command.commandStorage
func NewcommandStorageMock(t minimock.Tester) *commandStorageMock {
	m := &commandStorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DelMock = mcommandStorageMockDel{mock: m}
	m.GetMock = mcommandStorageMockGet{mock: m}
	m.KeysMock = mcommandStorageMockKeys{mock: m}
	m.PutMock = mcommandStorageMockPut{mock: m}

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
	m.mockExpectations = nil
	return m.mock
}

//Del implements github.com/namreg/godown-v2/internal/command.commandStorage interface
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
	m.mockExpectations = nil
	return m.mock
}

//Get implements github.com/namreg/godown-v2/internal/command.commandStorage interface
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

//Keys implements github.com/namreg/godown-v2/internal/command.commandStorage interface
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

type mcommandStorageMockPut struct {
	mock             *commandStorageMock
	mockExpectations *commandStorageMockPutParams
}

//commandStorageMockPutParams represents input parameters of the commandStorage.Put
type commandStorageMockPutParams struct {
	p  storage.Key
	p1 storage.ValueSetter
}

//Expect sets up expected params for the commandStorage.Put
func (m *mcommandStorageMockPut) Expect(p storage.Key, p1 storage.ValueSetter) *mcommandStorageMockPut {
	m.mockExpectations = &commandStorageMockPutParams{p, p1}
	return m
}

//Return sets up a mock for commandStorage.Put to return Return's arguments
func (m *mcommandStorageMockPut) Return(r error) *commandStorageMock {
	m.mock.PutFunc = func(p storage.Key, p1 storage.ValueSetter) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of commandStorage.Put method
func (m *mcommandStorageMockPut) Set(f func(p storage.Key, p1 storage.ValueSetter) (r error)) *commandStorageMock {
	m.mock.PutFunc = f
	m.mockExpectations = nil
	return m.mock
}

//Put implements github.com/namreg/godown-v2/internal/command.commandStorage interface
func (m *commandStorageMock) Put(p storage.Key, p1 storage.ValueSetter) (r error) {
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

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Put")
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

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		m.t.Fatal("Expected call to commandStorageMock.Put")
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
		ok = ok && (m.PutFunc == nil || atomic.LoadUint64(&m.PutCounter) > 0)

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

			if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
				m.t.Error("Expected call to commandStorageMock.Put")
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

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		return false
	}

	return true
}
