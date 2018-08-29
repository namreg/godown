package command

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "dataStore" can be found in github.com/namreg/godown-v2/internal/command
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	storage "github.com/namreg/godown-v2/internal/storage"

	testify_assert "github.com/stretchr/testify/assert"
)

//dataStoreMock implements github.com/namreg/godown-v2/internal/command.dataStore
type dataStoreMock struct {
	t minimock.Tester

	DelFunc       func(p storage.Key) (r error)
	DelCounter    uint64
	DelPreCounter uint64
	DelMock       mdataStoreMockDel

	GetFunc       func(p storage.Key) (r *storage.Value, r1 error)
	GetCounter    uint64
	GetPreCounter uint64
	GetMock       mdataStoreMockGet

	KeysFunc       func() (r []storage.Key, r1 error)
	KeysCounter    uint64
	KeysPreCounter uint64
	KeysMock       mdataStoreMockKeys

	PutFunc       func(p storage.Key, p1 storage.ValueSetter) (r error)
	PutCounter    uint64
	PutPreCounter uint64
	PutMock       mdataStoreMockPut
}

//NewdataStoreMock returns a mock for github.com/namreg/godown-v2/internal/command.dataStore
func NewdataStoreMock(t minimock.Tester) *dataStoreMock {
	m := &dataStoreMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DelMock = mdataStoreMockDel{mock: m}
	m.GetMock = mdataStoreMockGet{mock: m}
	m.KeysMock = mdataStoreMockKeys{mock: m}
	m.PutMock = mdataStoreMockPut{mock: m}

	return m
}

type mdataStoreMockDel struct {
	mock             *dataStoreMock
	mockExpectations *dataStoreMockDelParams
}

//dataStoreMockDelParams represents input parameters of the dataStore.Del
type dataStoreMockDelParams struct {
	p storage.Key
}

//Expect sets up expected params for the dataStore.Del
func (m *mdataStoreMockDel) Expect(p storage.Key) *mdataStoreMockDel {
	m.mockExpectations = &dataStoreMockDelParams{p}
	return m
}

//Return sets up a mock for dataStore.Del to return Return's arguments
func (m *mdataStoreMockDel) Return(r error) *dataStoreMock {
	m.mock.DelFunc = func(p storage.Key) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of dataStore.Del method
func (m *mdataStoreMockDel) Set(f func(p storage.Key) (r error)) *dataStoreMock {
	m.mock.DelFunc = f
	m.mockExpectations = nil
	return m.mock
}

//Del implements github.com/namreg/godown-v2/internal/command.dataStore interface
func (m *dataStoreMock) Del(p storage.Key) (r error) {
	atomic.AddUint64(&m.DelPreCounter, 1)
	defer atomic.AddUint64(&m.DelCounter, 1)

	if m.DelMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.DelMock.mockExpectations, dataStoreMockDelParams{p},
			"dataStore.Del got unexpected parameters")

		if m.DelFunc == nil {

			m.t.Fatal("No results are set for the dataStoreMock.Del")

			return
		}
	}

	if m.DelFunc == nil {
		m.t.Fatal("Unexpected call to dataStoreMock.Del")
		return
	}

	return m.DelFunc(p)
}

//DelMinimockCounter returns a count of dataStoreMock.DelFunc invocations
func (m *dataStoreMock) DelMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DelCounter)
}

//DelMinimockPreCounter returns the value of dataStoreMock.Del invocations
func (m *dataStoreMock) DelMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.DelPreCounter)
}

type mdataStoreMockGet struct {
	mock             *dataStoreMock
	mockExpectations *dataStoreMockGetParams
}

//dataStoreMockGetParams represents input parameters of the dataStore.Get
type dataStoreMockGetParams struct {
	p storage.Key
}

//Expect sets up expected params for the dataStore.Get
func (m *mdataStoreMockGet) Expect(p storage.Key) *mdataStoreMockGet {
	m.mockExpectations = &dataStoreMockGetParams{p}
	return m
}

//Return sets up a mock for dataStore.Get to return Return's arguments
func (m *mdataStoreMockGet) Return(r *storage.Value, r1 error) *dataStoreMock {
	m.mock.GetFunc = func(p storage.Key) (*storage.Value, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of dataStore.Get method
func (m *mdataStoreMockGet) Set(f func(p storage.Key) (r *storage.Value, r1 error)) *dataStoreMock {
	m.mock.GetFunc = f
	m.mockExpectations = nil
	return m.mock
}

//Get implements github.com/namreg/godown-v2/internal/command.dataStore interface
func (m *dataStoreMock) Get(p storage.Key) (r *storage.Value, r1 error) {
	atomic.AddUint64(&m.GetPreCounter, 1)
	defer atomic.AddUint64(&m.GetCounter, 1)

	if m.GetMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.GetMock.mockExpectations, dataStoreMockGetParams{p},
			"dataStore.Get got unexpected parameters")

		if m.GetFunc == nil {

			m.t.Fatal("No results are set for the dataStoreMock.Get")

			return
		}
	}

	if m.GetFunc == nil {
		m.t.Fatal("Unexpected call to dataStoreMock.Get")
		return
	}

	return m.GetFunc(p)
}

//GetMinimockCounter returns a count of dataStoreMock.GetFunc invocations
func (m *dataStoreMock) GetMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetCounter)
}

//GetMinimockPreCounter returns the value of dataStoreMock.Get invocations
func (m *dataStoreMock) GetMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetPreCounter)
}

type mdataStoreMockKeys struct {
	mock *dataStoreMock
}

//Return sets up a mock for dataStore.Keys to return Return's arguments
func (m *mdataStoreMockKeys) Return(r []storage.Key, r1 error) *dataStoreMock {
	m.mock.KeysFunc = func() ([]storage.Key, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of dataStore.Keys method
func (m *mdataStoreMockKeys) Set(f func() (r []storage.Key, r1 error)) *dataStoreMock {
	m.mock.KeysFunc = f

	return m.mock
}

//Keys implements github.com/namreg/godown-v2/internal/command.dataStore interface
func (m *dataStoreMock) Keys() (r []storage.Key, r1 error) {
	atomic.AddUint64(&m.KeysPreCounter, 1)
	defer atomic.AddUint64(&m.KeysCounter, 1)

	if m.KeysFunc == nil {
		m.t.Fatal("Unexpected call to dataStoreMock.Keys")
		return
	}

	return m.KeysFunc()
}

//KeysMinimockCounter returns a count of dataStoreMock.KeysFunc invocations
func (m *dataStoreMock) KeysMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.KeysCounter)
}

//KeysMinimockPreCounter returns the value of dataStoreMock.Keys invocations
func (m *dataStoreMock) KeysMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.KeysPreCounter)
}

type mdataStoreMockPut struct {
	mock             *dataStoreMock
	mockExpectations *dataStoreMockPutParams
}

//dataStoreMockPutParams represents input parameters of the dataStore.Put
type dataStoreMockPutParams struct {
	p  storage.Key
	p1 storage.ValueSetter
}

//Expect sets up expected params for the dataStore.Put
func (m *mdataStoreMockPut) Expect(p storage.Key, p1 storage.ValueSetter) *mdataStoreMockPut {
	m.mockExpectations = &dataStoreMockPutParams{p, p1}
	return m
}

//Return sets up a mock for dataStore.Put to return Return's arguments
func (m *mdataStoreMockPut) Return(r error) *dataStoreMock {
	m.mock.PutFunc = func(p storage.Key, p1 storage.ValueSetter) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of dataStore.Put method
func (m *mdataStoreMockPut) Set(f func(p storage.Key, p1 storage.ValueSetter) (r error)) *dataStoreMock {
	m.mock.PutFunc = f
	m.mockExpectations = nil
	return m.mock
}

//Put implements github.com/namreg/godown-v2/internal/command.dataStore interface
func (m *dataStoreMock) Put(p storage.Key, p1 storage.ValueSetter) (r error) {
	atomic.AddUint64(&m.PutPreCounter, 1)
	defer atomic.AddUint64(&m.PutCounter, 1)

	if m.PutMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PutMock.mockExpectations, dataStoreMockPutParams{p, p1},
			"dataStore.Put got unexpected parameters")

		if m.PutFunc == nil {

			m.t.Fatal("No results are set for the dataStoreMock.Put")

			return
		}
	}

	if m.PutFunc == nil {
		m.t.Fatal("Unexpected call to dataStoreMock.Put")
		return
	}

	return m.PutFunc(p, p1)
}

//PutMinimockCounter returns a count of dataStoreMock.PutFunc invocations
func (m *dataStoreMock) PutMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PutCounter)
}

//PutMinimockPreCounter returns the value of dataStoreMock.Put invocations
func (m *dataStoreMock) PutMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.PutPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *dataStoreMock) ValidateCallCounters() {

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Del")
	}

	if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Get")
	}

	if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Keys")
	}

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Put")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *dataStoreMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *dataStoreMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *dataStoreMock) MinimockFinish() {

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Del")
	}

	if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Get")
	}

	if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Keys")
	}

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Put")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *dataStoreMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *dataStoreMock) MinimockWait(timeout time.Duration) {
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
				m.t.Error("Expected call to dataStoreMock.Del")
			}

			if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
				m.t.Error("Expected call to dataStoreMock.Get")
			}

			if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
				m.t.Error("Expected call to dataStoreMock.Keys")
			}

			if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
				m.t.Error("Expected call to dataStoreMock.Put")
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
func (m *dataStoreMock) AllMocksCalled() bool {

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
