// Code generated by MockGen. DO NOT EDIT.
// Source: api.go
//
// Generated by this command:
//
//	mockgen -source api.go -destination mocks/mocks.go -typed true service
//

// Package mock_shortener is a generated GoMock package.
package mock_shortener

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// Mockservice is a mock of service interface.
type Mockservice struct {
	ctrl     *gomock.Controller
	recorder *MockserviceMockRecorder
}

// MockserviceMockRecorder is the mock recorder for Mockservice.
type MockserviceMockRecorder struct {
	mock *Mockservice
}

// NewMockservice creates a new mock instance.
func NewMockservice(ctrl *gomock.Controller) *Mockservice {
	mock := &Mockservice{ctrl: ctrl}
	mock.recorder = &MockserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockservice) EXPECT() *MockserviceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *Mockservice) Create(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockserviceMockRecorder) Create(arg0 any) *MockserviceCreateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mockservice)(nil).Create), arg0)
	return &MockserviceCreateCall{Call: call}
}

// MockserviceCreateCall wrap *gomock.Call
type MockserviceCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockserviceCreateCall) Return(arg0 string, arg1 error) *MockserviceCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockserviceCreateCall) Do(f func(string) (string, error)) *MockserviceCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockserviceCreateCall) DoAndReturn(f func(string) (string, error)) *MockserviceCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *Mockservice) GetByID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockserviceMockRecorder) GetByID(arg0 any) *MockserviceGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*Mockservice)(nil).GetByID), arg0)
	return &MockserviceGetByIDCall{Call: call}
}

// MockserviceGetByIDCall wrap *gomock.Call
type MockserviceGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockserviceGetByIDCall) Return(arg0 string, arg1 error) *MockserviceGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockserviceGetByIDCall) Do(f func(string) (string, error)) *MockserviceGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockserviceGetByIDCall) DoAndReturn(f func(string) (string, error)) *MockserviceGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
