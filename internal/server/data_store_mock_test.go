package server

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "dataStore" can be found in github.com/namreg/godown/internal/server
*/
import (
	"sync/atomic"
	time "time"

	"github.com/gojuno/minimock"
	storage "github.com/namreg/godown/internal/storage"

	testify_assert "github.com/stretchr/testify/assert"
)

//dataStoreMock implements github.com/namreg/godown/internal/server.dataStore
type dataStoreMock struct {
	t minimock.Tester

	AllFunc       func() (r map[storage.Key]*storage.Value, r1 error)
	AllCounter    uint64
	AllPreCounter uint64
	AllMock       mdataStoreMockAll

	AllWithTTLFunc       func() (r map[storage.Key]*storage.Value, r1 error)
	AllWithTTLCounter    uint64
	AllWithTTLPreCounter uint64
	AllWithTTLMock       mdataStoreMockAllWithTTL

	DelFunc       func(p storage.Key) (r error)
	DelCounter    uint64
	DelPreCounter uint64
	DelMock       mdataStoreMockDel

	RestoreFunc       func(p map[storage.Key]*storage.Value) (r error)
	RestoreCounter    uint64
	RestorePreCounter uint64
	RestoreMock       mdataStoreMockRestore
}

//NewdataStoreMock returns a mock for github.com/namreg/godown/internal/server.dataStore
func NewdataStoreMock(t minimock.Tester) *dataStoreMock {
	m := &dataStoreMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AllMock = mdataStoreMockAll{mock: m}
	m.AllWithTTLMock = mdataStoreMockAllWithTTL{mock: m}
	m.DelMock = mdataStoreMockDel{mock: m}
	m.RestoreMock = mdataStoreMockRestore{mock: m}

	return m
}

type mdataStoreMockAll struct {
	mock *dataStoreMock
}

//Return sets up a mock for dataStore.All to return Return's arguments
func (m *mdataStoreMockAll) Return(r map[storage.Key]*storage.Value, r1 error) *dataStoreMock {
	m.mock.AllFunc = func() (map[storage.Key]*storage.Value, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of dataStore.All method
func (m *mdataStoreMockAll) Set(f func() (r map[storage.Key]*storage.Value, r1 error)) *dataStoreMock {
	m.mock.AllFunc = f

	return m.mock
}

//All implements github.com/namreg/godown/internal/server.dataStore interface
func (m *dataStoreMock) All() (r map[storage.Key]*storage.Value, r1 error) {
	atomic.AddUint64(&m.AllPreCounter, 1)
	defer atomic.AddUint64(&m.AllCounter, 1)

	if m.AllFunc == nil {
		m.t.Fatal("Unexpected call to dataStoreMock.All")
		return
	}

	return m.AllFunc()
}

//AllMinimockCounter returns a count of dataStoreMock.AllFunc invocations
func (m *dataStoreMock) AllMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.AllCounter)
}

//AllMinimockPreCounter returns the value of dataStoreMock.All invocations
func (m *dataStoreMock) AllMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.AllPreCounter)
}

type mdataStoreMockAllWithTTL struct {
	mock *dataStoreMock
}

//Return sets up a mock for dataStore.AllWithTTL to return Return's arguments
func (m *mdataStoreMockAllWithTTL) Return(r map[storage.Key]*storage.Value, r1 error) *dataStoreMock {
	m.mock.AllWithTTLFunc = func() (map[storage.Key]*storage.Value, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of dataStore.AllWithTTL method
func (m *mdataStoreMockAllWithTTL) Set(f func() (r map[storage.Key]*storage.Value, r1 error)) *dataStoreMock {
	m.mock.AllWithTTLFunc = f

	return m.mock
}

//AllWithTTL implements github.com/namreg/godown/internal/server.dataStore interface
func (m *dataStoreMock) AllWithTTL() (r map[storage.Key]*storage.Value, r1 error) {
	atomic.AddUint64(&m.AllWithTTLPreCounter, 1)
	defer atomic.AddUint64(&m.AllWithTTLCounter, 1)

	if m.AllWithTTLFunc == nil {
		m.t.Fatal("Unexpected call to dataStoreMock.AllWithTTL")
		return
	}

	return m.AllWithTTLFunc()
}

//AllWithTTLMinimockCounter returns a count of dataStoreMock.AllWithTTLFunc invocations
func (m *dataStoreMock) AllWithTTLMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.AllWithTTLCounter)
}

//AllWithTTLMinimockPreCounter returns the value of dataStoreMock.AllWithTTL invocations
func (m *dataStoreMock) AllWithTTLMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.AllWithTTLPreCounter)
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

//Del implements github.com/namreg/godown/internal/server.dataStore interface
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

type mdataStoreMockRestore struct {
	mock             *dataStoreMock
	mockExpectations *dataStoreMockRestoreParams
}

//dataStoreMockRestoreParams represents input parameters of the dataStore.Restore
type dataStoreMockRestoreParams struct {
	p map[storage.Key]*storage.Value
}

//Expect sets up expected params for the dataStore.Restore
func (m *mdataStoreMockRestore) Expect(p map[storage.Key]*storage.Value) *mdataStoreMockRestore {
	m.mockExpectations = &dataStoreMockRestoreParams{p}
	return m
}

//Return sets up a mock for dataStore.Restore to return Return's arguments
func (m *mdataStoreMockRestore) Return(r error) *dataStoreMock {
	m.mock.RestoreFunc = func(p map[storage.Key]*storage.Value) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of dataStore.Restore method
func (m *mdataStoreMockRestore) Set(f func(p map[storage.Key]*storage.Value) (r error)) *dataStoreMock {
	m.mock.RestoreFunc = f
	m.mockExpectations = nil
	return m.mock
}

//Restore implements github.com/namreg/godown/internal/server.dataStore interface
func (m *dataStoreMock) Restore(p map[storage.Key]*storage.Value) (r error) {
	atomic.AddUint64(&m.RestorePreCounter, 1)
	defer atomic.AddUint64(&m.RestoreCounter, 1)

	if m.RestoreMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.RestoreMock.mockExpectations, dataStoreMockRestoreParams{p},
			"dataStore.Restore got unexpected parameters")

		if m.RestoreFunc == nil {

			m.t.Fatal("No results are set for the dataStoreMock.Restore")

			return
		}
	}

	if m.RestoreFunc == nil {
		m.t.Fatal("Unexpected call to dataStoreMock.Restore")
		return
	}

	return m.RestoreFunc(p)
}

//RestoreMinimockCounter returns a count of dataStoreMock.RestoreFunc invocations
func (m *dataStoreMock) RestoreMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RestoreCounter)
}

//RestoreMinimockPreCounter returns the value of dataStoreMock.Restore invocations
func (m *dataStoreMock) RestoreMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RestorePreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *dataStoreMock) ValidateCallCounters() {

	if m.AllFunc != nil && atomic.LoadUint64(&m.AllCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.All")
	}

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.AllWithTTL")
	}

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Del")
	}

	if m.RestoreFunc != nil && atomic.LoadUint64(&m.RestoreCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Restore")
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

	if m.AllFunc != nil && atomic.LoadUint64(&m.AllCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.All")
	}

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.AllWithTTL")
	}

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Del")
	}

	if m.RestoreFunc != nil && atomic.LoadUint64(&m.RestoreCounter) == 0 {
		m.t.Fatal("Expected call to dataStoreMock.Restore")
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
		ok = ok && (m.AllFunc == nil || atomic.LoadUint64(&m.AllCounter) > 0)
		ok = ok && (m.AllWithTTLFunc == nil || atomic.LoadUint64(&m.AllWithTTLCounter) > 0)
		ok = ok && (m.DelFunc == nil || atomic.LoadUint64(&m.DelCounter) > 0)
		ok = ok && (m.RestoreFunc == nil || atomic.LoadUint64(&m.RestoreCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.AllFunc != nil && atomic.LoadUint64(&m.AllCounter) == 0 {
				m.t.Error("Expected call to dataStoreMock.All")
			}

			if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
				m.t.Error("Expected call to dataStoreMock.AllWithTTL")
			}

			if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
				m.t.Error("Expected call to dataStoreMock.Del")
			}

			if m.RestoreFunc != nil && atomic.LoadUint64(&m.RestoreCounter) == 0 {
				m.t.Error("Expected call to dataStoreMock.Restore")
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

	if m.AllFunc != nil && atomic.LoadUint64(&m.AllCounter) == 0 {
		return false
	}

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		return false
	}

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		return false
	}

	if m.RestoreFunc != nil && atomic.LoadUint64(&m.RestoreCounter) == 0 {
		return false
	}

	return true
}
