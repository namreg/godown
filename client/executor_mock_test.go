package client

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "executor" can be found in github.com/namreg/godown/client
*/
import (
	context "context"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	api "github.com/namreg/godown/internal/api"
	grpc "google.golang.org/grpc"

	testify_assert "github.com/stretchr/testify/assert"
)

//executorMock implements github.com/namreg/godown/client.executor
type executorMock struct {
	t minimock.Tester

	ExecuteCommandFunc       func(p context.Context, p1 *api.ExecuteCommandRequest, p2 ...grpc.CallOption) (r *api.ExecuteCommandResponse, r1 error)
	ExecuteCommandCounter    uint64
	ExecuteCommandPreCounter uint64
	ExecuteCommandMock       mexecutorMockExecuteCommand
}

//NewexecutorMock returns a mock for github.com/namreg/godown/client.executor
func NewexecutorMock(t minimock.Tester) *executorMock {
	m := &executorMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ExecuteCommandMock = mexecutorMockExecuteCommand{mock: m}

	return m
}

type mexecutorMockExecuteCommand struct {
	mock             *executorMock
	mockExpectations *executorMockExecuteCommandParams
}

//executorMockExecuteCommandParams represents input parameters of the executor.ExecuteCommand
type executorMockExecuteCommandParams struct {
	p  context.Context
	p1 *api.ExecuteCommandRequest
	p2 []grpc.CallOption
}

//Expect sets up expected params for the executor.ExecuteCommand
func (m *mexecutorMockExecuteCommand) Expect(p context.Context, p1 *api.ExecuteCommandRequest, p2 ...grpc.CallOption) *mexecutorMockExecuteCommand {
	m.mockExpectations = &executorMockExecuteCommandParams{p, p1, p2}
	return m
}

//Return sets up a mock for executor.ExecuteCommand to return Return's arguments
func (m *mexecutorMockExecuteCommand) Return(r *api.ExecuteCommandResponse, r1 error) *executorMock {
	m.mock.ExecuteCommandFunc = func(p context.Context, p1 *api.ExecuteCommandRequest, p2 ...grpc.CallOption) (*api.ExecuteCommandResponse, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of executor.ExecuteCommand method
func (m *mexecutorMockExecuteCommand) Set(f func(p context.Context, p1 *api.ExecuteCommandRequest, p2 ...grpc.CallOption) (r *api.ExecuteCommandResponse, r1 error)) *executorMock {
	m.mock.ExecuteCommandFunc = f
	m.mockExpectations = nil
	return m.mock
}

//ExecuteCommand implements github.com/namreg/godown/client.executor interface
func (m *executorMock) ExecuteCommand(p context.Context, p1 *api.ExecuteCommandRequest, p2 ...grpc.CallOption) (r *api.ExecuteCommandResponse, r1 error) {
	atomic.AddUint64(&m.ExecuteCommandPreCounter, 1)
	defer atomic.AddUint64(&m.ExecuteCommandCounter, 1)

	if m.ExecuteCommandMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ExecuteCommandMock.mockExpectations, executorMockExecuteCommandParams{p, p1, p2},
			"executor.ExecuteCommand got unexpected parameters")

		if m.ExecuteCommandFunc == nil {

			m.t.Fatal("No results are set for the executorMock.ExecuteCommand")

			return
		}
	}

	if m.ExecuteCommandFunc == nil {
		m.t.Fatal("Unexpected call to executorMock.ExecuteCommand")
		return
	}

	return m.ExecuteCommandFunc(p, p1, p2...)
}

//ExecuteCommandMinimockCounter returns a count of executorMock.ExecuteCommandFunc invocations
func (m *executorMock) ExecuteCommandMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ExecuteCommandCounter)
}

//ExecuteCommandMinimockPreCounter returns the value of executorMock.ExecuteCommand invocations
func (m *executorMock) ExecuteCommandMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.ExecuteCommandPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *executorMock) ValidateCallCounters() {

	if m.ExecuteCommandFunc != nil && atomic.LoadUint64(&m.ExecuteCommandCounter) == 0 {
		m.t.Fatal("Expected call to executorMock.ExecuteCommand")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *executorMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *executorMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *executorMock) MinimockFinish() {

	if m.ExecuteCommandFunc != nil && atomic.LoadUint64(&m.ExecuteCommandCounter) == 0 {
		m.t.Fatal("Expected call to executorMock.ExecuteCommand")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *executorMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *executorMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.ExecuteCommandFunc == nil || atomic.LoadUint64(&m.ExecuteCommandCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.ExecuteCommandFunc != nil && atomic.LoadUint64(&m.ExecuteCommandCounter) == 0 {
				m.t.Error("Expected call to executorMock.ExecuteCommand")
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
func (m *executorMock) AllMocksCalled() bool {

	if m.ExecuteCommandFunc != nil && atomic.LoadUint64(&m.ExecuteCommandCounter) == 0 {
		return false
	}

	return true
}
