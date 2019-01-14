// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/golang001/brewery (interfaces: Controller)

// Package mock_brewery is a generated GoMock package.
package mock_brewery

import (
	gomock "github.com/golang/mock/gomock"
	brewery "github.com/golang001/brewery"
	reflect "reflect"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// PowerPin mocks base method
func (m *MockController) PowerPin(arg0 int, arg1 bool) error {
	ret := m.ctrl.Call(m, "PowerPin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerPin indicates an expected call of PowerPin
func (mr *MockControllerMockRecorder) PowerPin(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerPin", reflect.TypeOf((*MockController)(nil).PowerPin), arg0, arg1)
}

// ReadTemperature mocks base method
func (m *MockController) ReadTemperature(arg0 brewery.TemperatureAddress) (float64, error) {
	ret := m.ctrl.Call(m, "ReadTemperature", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTemperature indicates an expected call of ReadTemperature
func (mr *MockControllerMockRecorder) ReadTemperature(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTemperature", reflect.TypeOf((*MockController)(nil).ReadTemperature), arg0)
}