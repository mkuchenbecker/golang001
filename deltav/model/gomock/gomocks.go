// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/golang001/deltav/model/gomodel (interfaces: StorageTankClient,ReactorClient)

// Package mock_gomodel is a generated GoMock package.
package mock_gomodel

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	gomodel "github.com/golang001/deltav/model/gomodel"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockStorageTankClient is a mock of StorageTankClient interface
type MockStorageTankClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageTankClientMockRecorder
}

// MockStorageTankClientMockRecorder is the mock recorder for MockStorageTankClient
type MockStorageTankClientMockRecorder struct {
	mock *MockStorageTankClient
}

// NewMockStorageTankClient creates a new mock instance
func NewMockStorageTankClient(ctrl *gomock.Controller) *MockStorageTankClient {
	mock := &MockStorageTankClient{ctrl: ctrl}
	mock.recorder = &MockStorageTankClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageTankClient) EXPECT() *MockStorageTankClientMockRecorder {
	return m.recorder
}

// AddStorage mocks base method
func (m *MockStorageTankClient) AddStorage(arg0 context.Context, arg1 *gomodel.AddStorageRequest, arg2 ...grpc.CallOption) (*gomodel.AddStorageResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddStorage", varargs...)
	ret0, _ := ret[0].(*gomodel.AddStorageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddStorage indicates an expected call of AddStorage
func (mr *MockStorageTankClientMockRecorder) AddStorage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStorage", reflect.TypeOf((*MockStorageTankClient)(nil).AddStorage), varargs...)
}

// Status mocks base method
func (m *MockStorageTankClient) Status(arg0 context.Context, arg1 *gomodel.StorageStatusRequest, arg2 ...grpc.CallOption) (*gomodel.StorageStatusResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*gomodel.StorageStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockStorageTankClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockStorageTankClient)(nil).Status), varargs...)
}

// WithdrawStorage mocks base method
func (m *MockStorageTankClient) WithdrawStorage(arg0 context.Context, arg1 *gomodel.WithdrawStorageRequest, arg2 ...grpc.CallOption) (*gomodel.WithdrawStorageResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithdrawStorage", varargs...)
	ret0, _ := ret[0].(*gomodel.WithdrawStorageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithdrawStorage indicates an expected call of WithdrawStorage
func (mr *MockStorageTankClientMockRecorder) WithdrawStorage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawStorage", reflect.TypeOf((*MockStorageTankClient)(nil).WithdrawStorage), varargs...)
}

// MockReactorClient is a mock of ReactorClient interface
type MockReactorClient struct {
	ctrl     *gomock.Controller
	recorder *MockReactorClientMockRecorder
}

// MockReactorClientMockRecorder is the mock recorder for MockReactorClient
type MockReactorClientMockRecorder struct {
	mock *MockReactorClient
}

// NewMockReactorClient creates a new mock instance
func NewMockReactorClient(ctrl *gomock.Controller) *MockReactorClient {
	mock := &MockReactorClient{ctrl: ctrl}
	mock.recorder = &MockReactorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReactorClient) EXPECT() *MockReactorClientMockRecorder {
	return m.recorder
}

// React mocks base method
func (m *MockReactorClient) React(arg0 context.Context, arg1 *gomodel.ReactRequest, arg2 ...grpc.CallOption) (*gomodel.ReactResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "React", varargs...)
	ret0, _ := ret[0].(*gomodel.ReactResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// React indicates an expected call of React
func (mr *MockReactorClientMockRecorder) React(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "React", reflect.TypeOf((*MockReactorClient)(nil).React), varargs...)
}
