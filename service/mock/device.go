// Code generated by MockGen. DO NOT EDIT.
// Source: device.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	service "github.com/jakob-moeller-cloud/octi-sync-server/service"
)

// MockDevice is a mock of Device interface.
type MockDevice struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceMockRecorder
}

// MockDeviceMockRecorder is the mock recorder for MockDevice.
type MockDeviceMockRecorder struct {
	mock *MockDevice
}

// NewMockDevice creates a new mock instance.
func NewMockDevice(ctrl *gomock.Controller) *MockDevice {
	mock := &MockDevice{ctrl: ctrl}
	mock.recorder = &MockDeviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevice) EXPECT() *MockDeviceMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockDevice) ID() service.DeviceID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(service.DeviceID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockDeviceMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockDevice)(nil).ID))
}
