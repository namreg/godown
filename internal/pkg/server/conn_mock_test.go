package server

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.10
The original interface "Conn" can be found in net
*/
import (
	net "net"
	"sync/atomic"
	time "time"

	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

//ConnMock implements net.Conn
type ConnMock struct {
	t minimock.Tester

	CloseFunc       func() (r error)
	CloseCounter    uint64
	ClosePreCounter uint64
	CloseMock       mConnMockClose

	LocalAddrFunc       func() (r net.Addr)
	LocalAddrCounter    uint64
	LocalAddrPreCounter uint64
	LocalAddrMock       mConnMockLocalAddr

	ReadFunc       func(p []byte) (r int, r1 error)
	ReadCounter    uint64
	ReadPreCounter uint64
	ReadMock       mConnMockRead

	RemoteAddrFunc       func() (r net.Addr)
	RemoteAddrCounter    uint64
	RemoteAddrPreCounter uint64
	RemoteAddrMock       mConnMockRemoteAddr

	SetDeadlineFunc       func(p time.Time) (r error)
	SetDeadlineCounter    uint64
	SetDeadlinePreCounter uint64
	SetDeadlineMock       mConnMockSetDeadline

	SetReadDeadlineFunc       func(p time.Time) (r error)
	SetReadDeadlineCounter    uint64
	SetReadDeadlinePreCounter uint64
	SetReadDeadlineMock       mConnMockSetReadDeadline

	SetWriteDeadlineFunc       func(p time.Time) (r error)
	SetWriteDeadlineCounter    uint64
	SetWriteDeadlinePreCounter uint64
	SetWriteDeadlineMock       mConnMockSetWriteDeadline

	WriteFunc       func(p []byte) (r int, r1 error)
	WriteCounter    uint64
	WritePreCounter uint64
	WriteMock       mConnMockWrite
}

//NewConnMock returns a mock for net.Conn
func NewConnMock(t minimock.Tester) *ConnMock {
	m := &ConnMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mConnMockClose{mock: m}
	m.LocalAddrMock = mConnMockLocalAddr{mock: m}
	m.ReadMock = mConnMockRead{mock: m}
	m.RemoteAddrMock = mConnMockRemoteAddr{mock: m}
	m.SetDeadlineMock = mConnMockSetDeadline{mock: m}
	m.SetReadDeadlineMock = mConnMockSetReadDeadline{mock: m}
	m.SetWriteDeadlineMock = mConnMockSetWriteDeadline{mock: m}
	m.WriteMock = mConnMockWrite{mock: m}

	return m
}

type mConnMockClose struct {
	mock *ConnMock
}

//Return sets up a mock for Conn.Close to return Return's arguments
func (m *mConnMockClose) Return(r error) *ConnMock {
	m.mock.CloseFunc = func() error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Conn.Close method
func (m *mConnMockClose) Set(f func() (r error)) *ConnMock {
	m.mock.CloseFunc = f
	return m.mock
}

//Close implements net.Conn interface
func (m *ConnMock) Close() (r error) {
	atomic.AddUint64(&m.ClosePreCounter, 1)
	defer atomic.AddUint64(&m.CloseCounter, 1)

	if m.CloseFunc == nil {
		m.t.Fatal("Unexpected call to ConnMock.Close")
		return
	}

	return m.CloseFunc()
}

//CloseMinimockCounter returns a count of ConnMock.CloseFunc invocations
func (m *ConnMock) CloseMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CloseCounter)
}

//CloseMinimockPreCounter returns the value of ConnMock.Close invocations
func (m *ConnMock) CloseMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ClosePreCounter)
}

type mConnMockLocalAddr struct {
	mock *ConnMock
}

//Return sets up a mock for Conn.LocalAddr to return Return's arguments
func (m *mConnMockLocalAddr) Return(r net.Addr) *ConnMock {
	m.mock.LocalAddrFunc = func() net.Addr {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Conn.LocalAddr method
func (m *mConnMockLocalAddr) Set(f func() (r net.Addr)) *ConnMock {
	m.mock.LocalAddrFunc = f
	return m.mock
}

//LocalAddr implements net.Conn interface
func (m *ConnMock) LocalAddr() (r net.Addr) {
	atomic.AddUint64(&m.LocalAddrPreCounter, 1)
	defer atomic.AddUint64(&m.LocalAddrCounter, 1)

	if m.LocalAddrFunc == nil {
		m.t.Fatal("Unexpected call to ConnMock.LocalAddr")
		return
	}

	return m.LocalAddrFunc()
}

//LocalAddrMinimockCounter returns a count of ConnMock.LocalAddrFunc invocations
func (m *ConnMock) LocalAddrMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.LocalAddrCounter)
}

//LocalAddrMinimockPreCounter returns the value of ConnMock.LocalAddr invocations
func (m *ConnMock) LocalAddrMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.LocalAddrPreCounter)
}

type mConnMockRead struct {
	mock             *ConnMock
	mockExpectations *ConnMockReadParams
}

//ConnMockReadParams represents input parameters of the Conn.Read
type ConnMockReadParams struct {
	p []byte
}

//Expect sets up expected params for the Conn.Read
func (m *mConnMockRead) Expect(p []byte) *mConnMockRead {
	m.mockExpectations = &ConnMockReadParams{p}
	return m
}

//Return sets up a mock for Conn.Read to return Return's arguments
func (m *mConnMockRead) Return(r int, r1 error) *ConnMock {
	m.mock.ReadFunc = func(p []byte) (int, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Conn.Read method
func (m *mConnMockRead) Set(f func(p []byte) (r int, r1 error)) *ConnMock {
	m.mock.ReadFunc = f
	return m.mock
}

//Read implements net.Conn interface
func (m *ConnMock) Read(p []byte) (r int, r1 error) {
	atomic.AddUint64(&m.ReadPreCounter, 1)
	defer atomic.AddUint64(&m.ReadCounter, 1)

	if m.ReadMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ReadMock.mockExpectations, ConnMockReadParams{p},
			"Conn.Read got unexpected parameters")

		if m.ReadFunc == nil {

			m.t.Fatal("No results are set for the ConnMock.Read")

			return
		}
	}

	if m.ReadFunc == nil {
		m.t.Fatal("Unexpected call to ConnMock.Read")
		return
	}

	return m.ReadFunc(p)
}

//ReadMinimockCounter returns a count of ConnMock.ReadFunc invocations
func (m *ConnMock) ReadMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ReadCounter)
}

//ReadMinimockPreCounter returns the value of ConnMock.Read invocations
func (m *ConnMock) ReadMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ReadPreCounter)
}

type mConnMockRemoteAddr struct {
	mock *ConnMock
}

//Return sets up a mock for Conn.RemoteAddr to return Return's arguments
func (m *mConnMockRemoteAddr) Return(r net.Addr) *ConnMock {
	m.mock.RemoteAddrFunc = func() net.Addr {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Conn.RemoteAddr method
func (m *mConnMockRemoteAddr) Set(f func() (r net.Addr)) *ConnMock {
	m.mock.RemoteAddrFunc = f
	return m.mock
}

//RemoteAddr implements net.Conn interface
func (m *ConnMock) RemoteAddr() (r net.Addr) {
	atomic.AddUint64(&m.RemoteAddrPreCounter, 1)
	defer atomic.AddUint64(&m.RemoteAddrCounter, 1)

	if m.RemoteAddrFunc == nil {
		m.t.Fatal("Unexpected call to ConnMock.RemoteAddr")
		return
	}

	return m.RemoteAddrFunc()
}

//RemoteAddrMinimockCounter returns a count of ConnMock.RemoteAddrFunc invocations
func (m *ConnMock) RemoteAddrMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RemoteAddrCounter)
}

//RemoteAddrMinimockPreCounter returns the value of ConnMock.RemoteAddr invocations
func (m *ConnMock) RemoteAddrMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.RemoteAddrPreCounter)
}

type mConnMockSetDeadline struct {
	mock             *ConnMock
	mockExpectations *ConnMockSetDeadlineParams
}

//ConnMockSetDeadlineParams represents input parameters of the Conn.SetDeadline
type ConnMockSetDeadlineParams struct {
	p time.Time
}

//Expect sets up expected params for the Conn.SetDeadline
func (m *mConnMockSetDeadline) Expect(p time.Time) *mConnMockSetDeadline {
	m.mockExpectations = &ConnMockSetDeadlineParams{p}
	return m
}

//Return sets up a mock for Conn.SetDeadline to return Return's arguments
func (m *mConnMockSetDeadline) Return(r error) *ConnMock {
	m.mock.SetDeadlineFunc = func(p time.Time) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Conn.SetDeadline method
func (m *mConnMockSetDeadline) Set(f func(p time.Time) (r error)) *ConnMock {
	m.mock.SetDeadlineFunc = f
	return m.mock
}

//SetDeadline implements net.Conn interface
func (m *ConnMock) SetDeadline(p time.Time) (r error) {
	atomic.AddUint64(&m.SetDeadlinePreCounter, 1)
	defer atomic.AddUint64(&m.SetDeadlineCounter, 1)

	if m.SetDeadlineMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetDeadlineMock.mockExpectations, ConnMockSetDeadlineParams{p},
			"Conn.SetDeadline got unexpected parameters")

		if m.SetDeadlineFunc == nil {

			m.t.Fatal("No results are set for the ConnMock.SetDeadline")

			return
		}
	}

	if m.SetDeadlineFunc == nil {
		m.t.Fatal("Unexpected call to ConnMock.SetDeadline")
		return
	}

	return m.SetDeadlineFunc(p)
}

//SetDeadlineMinimockCounter returns a count of ConnMock.SetDeadlineFunc invocations
func (m *ConnMock) SetDeadlineMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetDeadlineCounter)
}

//SetDeadlineMinimockPreCounter returns the value of ConnMock.SetDeadline invocations
func (m *ConnMock) SetDeadlineMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SetDeadlinePreCounter)
}

type mConnMockSetReadDeadline struct {
	mock             *ConnMock
	mockExpectations *ConnMockSetReadDeadlineParams
}

//ConnMockSetReadDeadlineParams represents input parameters of the Conn.SetReadDeadline
type ConnMockSetReadDeadlineParams struct {
	p time.Time
}

//Expect sets up expected params for the Conn.SetReadDeadline
func (m *mConnMockSetReadDeadline) Expect(p time.Time) *mConnMockSetReadDeadline {
	m.mockExpectations = &ConnMockSetReadDeadlineParams{p}
	return m
}

//Return sets up a mock for Conn.SetReadDeadline to return Return's arguments
func (m *mConnMockSetReadDeadline) Return(r error) *ConnMock {
	m.mock.SetReadDeadlineFunc = func(p time.Time) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Conn.SetReadDeadline method
func (m *mConnMockSetReadDeadline) Set(f func(p time.Time) (r error)) *ConnMock {
	m.mock.SetReadDeadlineFunc = f
	return m.mock
}

//SetReadDeadline implements net.Conn interface
func (m *ConnMock) SetReadDeadline(p time.Time) (r error) {
	atomic.AddUint64(&m.SetReadDeadlinePreCounter, 1)
	defer atomic.AddUint64(&m.SetReadDeadlineCounter, 1)

	if m.SetReadDeadlineMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetReadDeadlineMock.mockExpectations, ConnMockSetReadDeadlineParams{p},
			"Conn.SetReadDeadline got unexpected parameters")

		if m.SetReadDeadlineFunc == nil {

			m.t.Fatal("No results are set for the ConnMock.SetReadDeadline")

			return
		}
	}

	if m.SetReadDeadlineFunc == nil {
		m.t.Fatal("Unexpected call to ConnMock.SetReadDeadline")
		return
	}

	return m.SetReadDeadlineFunc(p)
}

//SetReadDeadlineMinimockCounter returns a count of ConnMock.SetReadDeadlineFunc invocations
func (m *ConnMock) SetReadDeadlineMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetReadDeadlineCounter)
}

//SetReadDeadlineMinimockPreCounter returns the value of ConnMock.SetReadDeadline invocations
func (m *ConnMock) SetReadDeadlineMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SetReadDeadlinePreCounter)
}

type mConnMockSetWriteDeadline struct {
	mock             *ConnMock
	mockExpectations *ConnMockSetWriteDeadlineParams
}

//ConnMockSetWriteDeadlineParams represents input parameters of the Conn.SetWriteDeadline
type ConnMockSetWriteDeadlineParams struct {
	p time.Time
}

//Expect sets up expected params for the Conn.SetWriteDeadline
func (m *mConnMockSetWriteDeadline) Expect(p time.Time) *mConnMockSetWriteDeadline {
	m.mockExpectations = &ConnMockSetWriteDeadlineParams{p}
	return m
}

//Return sets up a mock for Conn.SetWriteDeadline to return Return's arguments
func (m *mConnMockSetWriteDeadline) Return(r error) *ConnMock {
	m.mock.SetWriteDeadlineFunc = func(p time.Time) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Conn.SetWriteDeadline method
func (m *mConnMockSetWriteDeadline) Set(f func(p time.Time) (r error)) *ConnMock {
	m.mock.SetWriteDeadlineFunc = f
	return m.mock
}

//SetWriteDeadline implements net.Conn interface
func (m *ConnMock) SetWriteDeadline(p time.Time) (r error) {
	atomic.AddUint64(&m.SetWriteDeadlinePreCounter, 1)
	defer atomic.AddUint64(&m.SetWriteDeadlineCounter, 1)

	if m.SetWriteDeadlineMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetWriteDeadlineMock.mockExpectations, ConnMockSetWriteDeadlineParams{p},
			"Conn.SetWriteDeadline got unexpected parameters")

		if m.SetWriteDeadlineFunc == nil {

			m.t.Fatal("No results are set for the ConnMock.SetWriteDeadline")

			return
		}
	}

	if m.SetWriteDeadlineFunc == nil {
		m.t.Fatal("Unexpected call to ConnMock.SetWriteDeadline")
		return
	}

	return m.SetWriteDeadlineFunc(p)
}

//SetWriteDeadlineMinimockCounter returns a count of ConnMock.SetWriteDeadlineFunc invocations
func (m *ConnMock) SetWriteDeadlineMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetWriteDeadlineCounter)
}

//SetWriteDeadlineMinimockPreCounter returns the value of ConnMock.SetWriteDeadline invocations
func (m *ConnMock) SetWriteDeadlineMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SetWriteDeadlinePreCounter)
}

type mConnMockWrite struct {
	mock             *ConnMock
	mockExpectations *ConnMockWriteParams
}

//ConnMockWriteParams represents input parameters of the Conn.Write
type ConnMockWriteParams struct {
	p []byte
}

//Expect sets up expected params for the Conn.Write
func (m *mConnMockWrite) Expect(p []byte) *mConnMockWrite {
	m.mockExpectations = &ConnMockWriteParams{p}
	return m
}

//Return sets up a mock for Conn.Write to return Return's arguments
func (m *mConnMockWrite) Return(r int, r1 error) *ConnMock {
	m.mock.WriteFunc = func(p []byte) (int, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Conn.Write method
func (m *mConnMockWrite) Set(f func(p []byte) (r int, r1 error)) *ConnMock {
	m.mock.WriteFunc = f
	return m.mock
}

//Write implements net.Conn interface
func (m *ConnMock) Write(p []byte) (r int, r1 error) {
	atomic.AddUint64(&m.WritePreCounter, 1)
	defer atomic.AddUint64(&m.WriteCounter, 1)

	if m.WriteMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.WriteMock.mockExpectations, ConnMockWriteParams{p},
			"Conn.Write got unexpected parameters")

		if m.WriteFunc == nil {

			m.t.Fatal("No results are set for the ConnMock.Write")

			return
		}
	}

	if m.WriteFunc == nil {
		m.t.Fatal("Unexpected call to ConnMock.Write")
		return
	}

	return m.WriteFunc(p)
}

//WriteMinimockCounter returns a count of ConnMock.WriteFunc invocations
func (m *ConnMock) WriteMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.WriteCounter)
}

//WriteMinimockPreCounter returns the value of ConnMock.Write invocations
func (m *ConnMock) WriteMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.WritePreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ConnMock) ValidateCallCounters() {

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.Close")
	}

	if m.LocalAddrFunc != nil && atomic.LoadUint64(&m.LocalAddrCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.LocalAddr")
	}

	if m.ReadFunc != nil && atomic.LoadUint64(&m.ReadCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.Read")
	}

	if m.RemoteAddrFunc != nil && atomic.LoadUint64(&m.RemoteAddrCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.RemoteAddr")
	}

	if m.SetDeadlineFunc != nil && atomic.LoadUint64(&m.SetDeadlineCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.SetDeadline")
	}

	if m.SetReadDeadlineFunc != nil && atomic.LoadUint64(&m.SetReadDeadlineCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.SetReadDeadline")
	}

	if m.SetWriteDeadlineFunc != nil && atomic.LoadUint64(&m.SetWriteDeadlineCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.SetWriteDeadline")
	}

	if m.WriteFunc != nil && atomic.LoadUint64(&m.WriteCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.Write")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ConnMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *ConnMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *ConnMock) MinimockFinish() {

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.Close")
	}

	if m.LocalAddrFunc != nil && atomic.LoadUint64(&m.LocalAddrCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.LocalAddr")
	}

	if m.ReadFunc != nil && atomic.LoadUint64(&m.ReadCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.Read")
	}

	if m.RemoteAddrFunc != nil && atomic.LoadUint64(&m.RemoteAddrCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.RemoteAddr")
	}

	if m.SetDeadlineFunc != nil && atomic.LoadUint64(&m.SetDeadlineCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.SetDeadline")
	}

	if m.SetReadDeadlineFunc != nil && atomic.LoadUint64(&m.SetReadDeadlineCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.SetReadDeadline")
	}

	if m.SetWriteDeadlineFunc != nil && atomic.LoadUint64(&m.SetWriteDeadlineCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.SetWriteDeadline")
	}

	if m.WriteFunc != nil && atomic.LoadUint64(&m.WriteCounter) == 0 {
		m.t.Fatal("Expected call to ConnMock.Write")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *ConnMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *ConnMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.CloseFunc == nil || atomic.LoadUint64(&m.CloseCounter) > 0)
		ok = ok && (m.LocalAddrFunc == nil || atomic.LoadUint64(&m.LocalAddrCounter) > 0)
		ok = ok && (m.ReadFunc == nil || atomic.LoadUint64(&m.ReadCounter) > 0)
		ok = ok && (m.RemoteAddrFunc == nil || atomic.LoadUint64(&m.RemoteAddrCounter) > 0)
		ok = ok && (m.SetDeadlineFunc == nil || atomic.LoadUint64(&m.SetDeadlineCounter) > 0)
		ok = ok && (m.SetReadDeadlineFunc == nil || atomic.LoadUint64(&m.SetReadDeadlineCounter) > 0)
		ok = ok && (m.SetWriteDeadlineFunc == nil || atomic.LoadUint64(&m.SetWriteDeadlineCounter) > 0)
		ok = ok && (m.WriteFunc == nil || atomic.LoadUint64(&m.WriteCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
				m.t.Error("Expected call to ConnMock.Close")
			}

			if m.LocalAddrFunc != nil && atomic.LoadUint64(&m.LocalAddrCounter) == 0 {
				m.t.Error("Expected call to ConnMock.LocalAddr")
			}

			if m.ReadFunc != nil && atomic.LoadUint64(&m.ReadCounter) == 0 {
				m.t.Error("Expected call to ConnMock.Read")
			}

			if m.RemoteAddrFunc != nil && atomic.LoadUint64(&m.RemoteAddrCounter) == 0 {
				m.t.Error("Expected call to ConnMock.RemoteAddr")
			}

			if m.SetDeadlineFunc != nil && atomic.LoadUint64(&m.SetDeadlineCounter) == 0 {
				m.t.Error("Expected call to ConnMock.SetDeadline")
			}

			if m.SetReadDeadlineFunc != nil && atomic.LoadUint64(&m.SetReadDeadlineCounter) == 0 {
				m.t.Error("Expected call to ConnMock.SetReadDeadline")
			}

			if m.SetWriteDeadlineFunc != nil && atomic.LoadUint64(&m.SetWriteDeadlineCounter) == 0 {
				m.t.Error("Expected call to ConnMock.SetWriteDeadline")
			}

			if m.WriteFunc != nil && atomic.LoadUint64(&m.WriteCounter) == 0 {
				m.t.Error("Expected call to ConnMock.Write")
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
func (m *ConnMock) AllMocksCalled() bool {

	if m.CloseFunc != nil && atomic.LoadUint64(&m.CloseCounter) == 0 {
		return false
	}

	if m.LocalAddrFunc != nil && atomic.LoadUint64(&m.LocalAddrCounter) == 0 {
		return false
	}

	if m.ReadFunc != nil && atomic.LoadUint64(&m.ReadCounter) == 0 {
		return false
	}

	if m.RemoteAddrFunc != nil && atomic.LoadUint64(&m.RemoteAddrCounter) == 0 {
		return false
	}

	if m.SetDeadlineFunc != nil && atomic.LoadUint64(&m.SetDeadlineCounter) == 0 {
		return false
	}

	if m.SetReadDeadlineFunc != nil && atomic.LoadUint64(&m.SetReadDeadlineCounter) == 0 {
		return false
	}

	if m.SetWriteDeadlineFunc != nil && atomic.LoadUint64(&m.SetWriteDeadlineCounter) == 0 {
		return false
	}

	if m.WriteFunc != nil && atomic.LoadUint64(&m.WriteCounter) == 0 {
		return false
	}

	return true
}
