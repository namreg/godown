package command

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Storage" can be found in github.com/namreg/godown-v2/internal/pkg/storage
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	storage "github.com/namreg/godown-v2/internal/pkg/storage"

	testify_assert "github.com/stretchr/testify/assert"
)

//StorageMock implements github.com/namreg/godown-v2/internal/pkg/storage.Storage
type StorageMock struct {
	t minimock.Tester

	AllFunc       func() (r map[storage.Key]*storage.Value, r1 error)
	AllCounter    uint64
	AllPreCounter uint64
	AllMock       mStorageMockAll

	AllWithTTLFunc       func() (r map[storage.Key]*storage.Value, r1 error)
	AllWithTTLCounter    uint64
	AllWithTTLPreCounter uint64
	AllWithTTLMock       mStorageMockAllWithTTL

	DelFunc       func(p storage.Key) (r error)
	DelCounter    uint64
	DelPreCounter uint64
	DelMock       mStorageMockDel

	GetFunc       func(p storage.Key) (r *storage.Value, r1 error)
	GetCounter    uint64
	GetPreCounter uint64
	GetMock       mStorageMockGet

	KeysFunc       func() (r []storage.Key, r1 error)
	KeysCounter    uint64
	KeysPreCounter uint64
	KeysMock       mStorageMockKeys

	PutFunc       func(p storage.Key, p1 storage.ValueSetter) (r error)
	PutCounter    uint64
	PutPreCounter uint64
	PutMock       mStorageMockPut
}

//NewStorageMock returns a mock for github.com/namreg/godown-v2/internal/pkg/storage.Storage
func NewStorageMock(t minimock.Tester) *StorageMock {
	m := &StorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AllMock = mStorageMockAll{mock: m}
	m.AllWithTTLMock = mStorageMockAllWithTTL{mock: m}
	m.DelMock = mStorageMockDel{mock: m}
	m.GetMock = mStorageMockGet{mock: m}
	m.KeysMock = mStorageMockKeys{mock: m}
	m.PutMock = mStorageMockPut{mock: m}

	return m
}

type mStorageMockAll struct {
	mock *StorageMock
}

//Return sets up a mock for Storage.All to return Return's arguments
func (m *mStorageMockAll) Return(r map[storage.Key]*storage.Value, r1 error) *StorageMock {
	m.mock.AllFunc = func() (map[storage.Key]*storage.Value, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.All method
func (m *mStorageMockAll) Set(f func() (r map[storage.Key]*storage.Value, r1 error)) *StorageMock {
	m.mock.AllFunc = f
	return m.mock
}

//All implements github.com/namreg/godown-v2/internal/pkg/storage.Storage interface
func (m *StorageMock) All() (r map[storage.Key]*storage.Value, r1 error) {
	atomic.AddUint64(&m.AllPreCounter, 1)
	defer atomic.AddUint64(&m.AllCounter, 1)

	if m.AllFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.All")
		return
	}

	return m.AllFunc()
}

//AllMinimockCounter returns a count of StorageMock.AllFunc invocations
func (m *StorageMock) AllMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.AllCounter)
}

//AllMinimockPreCounter returns the value of StorageMock.All invocations
func (m *StorageMock) AllMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.AllPreCounter)
}

type mStorageMockAllWithTTL struct {
	mock *StorageMock
}

//Return sets up a mock for Storage.AllWithTTL to return Return's arguments
func (m *mStorageMockAllWithTTL) Return(r map[storage.Key]*storage.Value, r1 error) *StorageMock {
	m.mock.AllWithTTLFunc = func() (map[storage.Key]*storage.Value, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.AllWithTTL method
func (m *mStorageMockAllWithTTL) Set(f func() (r map[storage.Key]*storage.Value, r1 error)) *StorageMock {
	m.mock.AllWithTTLFunc = f
	return m.mock
}

//AllWithTTL implements github.com/namreg/godown-v2/internal/pkg/storage.Storage interface
func (m *StorageMock) AllWithTTL() (r map[storage.Key]*storage.Value, r1 error) {
	atomic.AddUint64(&m.AllWithTTLPreCounter, 1)
	defer atomic.AddUint64(&m.AllWithTTLCounter, 1)

	if m.AllWithTTLFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.AllWithTTL")
		return
	}

	return m.AllWithTTLFunc()
}

//AllWithTTLMinimockCounter returns a count of StorageMock.AllWithTTLFunc invocations
func (m *StorageMock) AllWithTTLMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.AllWithTTLCounter)
}

//AllWithTTLMinimockPreCounter returns the value of StorageMock.AllWithTTL invocations
func (m *StorageMock) AllWithTTLMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.AllWithTTLPreCounter)
}

type mStorageMockDel struct {
	mock             *StorageMock
	mockExpectations *StorageMockDelParams
}

//StorageMockDelParams represents input parameters of the Storage.Del
type StorageMockDelParams struct {
	p storage.Key
}

//Expect sets up expected params for the Storage.Del
func (m *mStorageMockDel) Expect(p storage.Key) *mStorageMockDel {
	m.mockExpectations = &StorageMockDelParams{p}
	return m
}

//Return sets up a mock for Storage.Del to return Return's arguments
func (m *mStorageMockDel) Return(r error) *StorageMock {
	m.mock.DelFunc = func(p storage.Key) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.Del method
func (m *mStorageMockDel) Set(f func(p storage.Key) (r error)) *StorageMock {
	m.mock.DelFunc = f
	return m.mock
}

//Del implements github.com/namreg/godown-v2/internal/pkg/storage.Storage interface
func (m *StorageMock) Del(p storage.Key) (r error) {
	atomic.AddUint64(&m.DelPreCounter, 1)
	defer atomic.AddUint64(&m.DelCounter, 1)

	if m.DelMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.DelMock.mockExpectations, StorageMockDelParams{p},
			"Storage.Del got unexpected parameters")

		if m.DelFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.Del")

			return
		}
	}

	if m.DelFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.Del")
		return
	}

	return m.DelFunc(p)
}

//DelMinimockCounter returns a count of StorageMock.DelFunc invocations
func (m *StorageMock) DelMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DelCounter)
}

//DelMinimockPreCounter returns the value of StorageMock.Del invocations
func (m *StorageMock) DelMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.DelPreCounter)
}

type mStorageMockGet struct {
	mock             *StorageMock
	mockExpectations *StorageMockGetParams
}

//StorageMockGetParams represents input parameters of the Storage.Get
type StorageMockGetParams struct {
	p storage.Key
}

//Expect sets up expected params for the Storage.Get
func (m *mStorageMockGet) Expect(p storage.Key) *mStorageMockGet {
	m.mockExpectations = &StorageMockGetParams{p}
	return m
}

//Return sets up a mock for Storage.Get to return Return's arguments
func (m *mStorageMockGet) Return(r *storage.Value, r1 error) *StorageMock {
	m.mock.GetFunc = func(p storage.Key) (*storage.Value, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.Get method
func (m *mStorageMockGet) Set(f func(p storage.Key) (r *storage.Value, r1 error)) *StorageMock {
	m.mock.GetFunc = f
	return m.mock
}

//Get implements github.com/namreg/godown-v2/internal/pkg/storage.Storage interface
func (m *StorageMock) Get(p storage.Key) (r *storage.Value, r1 error) {
	atomic.AddUint64(&m.GetPreCounter, 1)
	defer atomic.AddUint64(&m.GetCounter, 1)

	if m.GetMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.GetMock.mockExpectations, StorageMockGetParams{p},
			"Storage.Get got unexpected parameters")

		if m.GetFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.Get")

			return
		}
	}

	if m.GetFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.Get")
		return
	}

	return m.GetFunc(p)
}

//GetMinimockCounter returns a count of StorageMock.GetFunc invocations
func (m *StorageMock) GetMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetCounter)
}

//GetMinimockPreCounter returns the value of StorageMock.Get invocations
func (m *StorageMock) GetMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetPreCounter)
}

type mStorageMockKeys struct {
	mock *StorageMock
}

//Return sets up a mock for Storage.Keys to return Return's arguments
func (m *mStorageMockKeys) Return(r []storage.Key, r1 error) *StorageMock {
	m.mock.KeysFunc = func() ([]storage.Key, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.Keys method
func (m *mStorageMockKeys) Set(f func() (r []storage.Key, r1 error)) *StorageMock {
	m.mock.KeysFunc = f
	return m.mock
}

//Keys implements github.com/namreg/godown-v2/internal/pkg/storage.Storage interface
func (m *StorageMock) Keys() (r []storage.Key, r1 error) {
	atomic.AddUint64(&m.KeysPreCounter, 1)
	defer atomic.AddUint64(&m.KeysCounter, 1)

	if m.KeysFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.Keys")
		return
	}

	return m.KeysFunc()
}

//KeysMinimockCounter returns a count of StorageMock.KeysFunc invocations
func (m *StorageMock) KeysMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.KeysCounter)
}

//KeysMinimockPreCounter returns the value of StorageMock.Keys invocations
func (m *StorageMock) KeysMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.KeysPreCounter)
}

type mStorageMockPut struct {
	mock             *StorageMock
	mockExpectations *StorageMockPutParams
}

//StorageMockPutParams represents input parameters of the Storage.Put
type StorageMockPutParams struct {
	p  storage.Key
	p1 storage.ValueSetter
}

//Expect sets up expected params for the Storage.Put
func (m *mStorageMockPut) Expect(p storage.Key, p1 storage.ValueSetter) *mStorageMockPut {
	m.mockExpectations = &StorageMockPutParams{p, p1}
	return m
}

//Return sets up a mock for Storage.Put to return Return's arguments
func (m *mStorageMockPut) Return(r error) *StorageMock {
	m.mock.PutFunc = func(p storage.Key, p1 storage.ValueSetter) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Storage.Put method
func (m *mStorageMockPut) Set(f func(p storage.Key, p1 storage.ValueSetter) (r error)) *StorageMock {
	m.mock.PutFunc = f
	return m.mock
}

//Put implements github.com/namreg/godown-v2/internal/pkg/storage.Storage interface
func (m *StorageMock) Put(p storage.Key, p1 storage.ValueSetter) (r error) {
	atomic.AddUint64(&m.PutPreCounter, 1)
	defer atomic.AddUint64(&m.PutCounter, 1)

	if m.PutMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PutMock.mockExpectations, StorageMockPutParams{p, p1},
			"Storage.Put got unexpected parameters")

		if m.PutFunc == nil {

			m.t.Fatal("No results are set for the StorageMock.Put")

			return
		}
	}

	if m.PutFunc == nil {
		m.t.Fatal("Unexpected call to StorageMock.Put")
		return
	}

	return m.PutFunc(p, p1)
}

//PutMinimockCounter returns a count of StorageMock.PutFunc invocations
func (m *StorageMock) PutMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PutCounter)
}

//PutMinimockPreCounter returns the value of StorageMock.Put invocations
func (m *StorageMock) PutMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.PutPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *StorageMock) ValidateCallCounters() {

	if m.AllFunc != nil && atomic.LoadUint64(&m.AllCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.All")
	}

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.AllWithTTL")
	}

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Del")
	}

	if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Get")
	}

	if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Keys")
	}

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Put")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *StorageMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *StorageMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *StorageMock) MinimockFinish() {

	if m.AllFunc != nil && atomic.LoadUint64(&m.AllCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.All")
	}

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.AllWithTTL")
	}

	if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Del")
	}

	if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Get")
	}

	if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Keys")
	}

	if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
		m.t.Fatal("Expected call to StorageMock.Put")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *StorageMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *StorageMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.AllFunc == nil || atomic.LoadUint64(&m.AllCounter) > 0)
		ok = ok && (m.AllWithTTLFunc == nil || atomic.LoadUint64(&m.AllWithTTLCounter) > 0)
		ok = ok && (m.DelFunc == nil || atomic.LoadUint64(&m.DelCounter) > 0)
		ok = ok && (m.GetFunc == nil || atomic.LoadUint64(&m.GetCounter) > 0)
		ok = ok && (m.KeysFunc == nil || atomic.LoadUint64(&m.KeysCounter) > 0)
		ok = ok && (m.PutFunc == nil || atomic.LoadUint64(&m.PutCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.AllFunc != nil && atomic.LoadUint64(&m.AllCounter) == 0 {
				m.t.Error("Expected call to StorageMock.All")
			}

			if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
				m.t.Error("Expected call to StorageMock.AllWithTTL")
			}

			if m.DelFunc != nil && atomic.LoadUint64(&m.DelCounter) == 0 {
				m.t.Error("Expected call to StorageMock.Del")
			}

			if m.GetFunc != nil && atomic.LoadUint64(&m.GetCounter) == 0 {
				m.t.Error("Expected call to StorageMock.Get")
			}

			if m.KeysFunc != nil && atomic.LoadUint64(&m.KeysCounter) == 0 {
				m.t.Error("Expected call to StorageMock.Keys")
			}

			if m.PutFunc != nil && atomic.LoadUint64(&m.PutCounter) == 0 {
				m.t.Error("Expected call to StorageMock.Put")
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
func (m *StorageMock) AllMocksCalled() bool {

	if m.AllFunc != nil && atomic.LoadUint64(&m.AllCounter) == 0 {
		return false
	}

	if m.AllWithTTLFunc != nil && atomic.LoadUint64(&m.AllWithTTLCounter) == 0 {
		return false
	}

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
