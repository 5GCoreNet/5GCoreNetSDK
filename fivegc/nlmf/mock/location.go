// Code generated by MockGen. DO NOT EDIT.
// Source: ../location.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	fivegc "github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	nlmf "github.com/5GCoreNet/5GCoreNetSDK/fivegc/nlmf"
	openapi_CommonData "github.com/5GCoreNet/openapi/openapi_CommonData"
	openapi_Nlmf_Location "github.com/5GCoreNet/openapi/openapi_Nlmf_Location"
	gomock "github.com/golang/mock/gomock"
)

// MockLocation is a mock of Location interface.
type MockLocation struct {
	ctrl     *gomock.Controller
	recorder *MockLocationMockRecorder
}

// MockLocationMockRecorder is the mock recorder for MockLocation.
type MockLocationMockRecorder struct {
	mock *MockLocation
}

// NewMockLocation creates a new mock instance.
func NewMockLocation(ctrl *gomock.Controller) *MockLocation {
	mock := &MockLocation{ctrl: ctrl}
	mock.recorder = &MockLocationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocation) EXPECT() *MockLocationMockRecorder {
	return m.recorder
}

// CancelLocation mocks base method.
func (m *MockLocation) CancelLocation(arg0 context.Context, arg1 openapi_Nlmf_Location.CancelLocData) (openapi_CommonData.ProblemDetails, fivegc.RedirectResponse, nlmf.CancelLocationStatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelLocation", arg0, arg1)
	ret0, _ := ret[0].(openapi_CommonData.ProblemDetails)
	ret1, _ := ret[1].(fivegc.RedirectResponse)
	ret2, _ := ret[2].(nlmf.CancelLocationStatusCode)
	return ret0, ret1, ret2
}

// CancelLocation indicates an expected call of CancelLocation.
func (mr *MockLocationMockRecorder) CancelLocation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelLocation", reflect.TypeOf((*MockLocation)(nil).CancelLocation), arg0, arg1)
}

// DetermineLocation mocks base method.
func (m *MockLocation) DetermineLocation(arg0 context.Context, arg1 openapi_Nlmf_Location.InputData) (openapi_Nlmf_Location.LocationData, openapi_CommonData.ProblemDetails, fivegc.RedirectResponse, nlmf.DetermineLocationStatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetermineLocation", arg0, arg1)
	ret0, _ := ret[0].(openapi_Nlmf_Location.LocationData)
	ret1, _ := ret[1].(openapi_CommonData.ProblemDetails)
	ret2, _ := ret[2].(fivegc.RedirectResponse)
	ret3, _ := ret[3].(nlmf.DetermineLocationStatusCode)
	return ret0, ret1, ret2, ret3
}

// DetermineLocation indicates an expected call of DetermineLocation.
func (mr *MockLocationMockRecorder) DetermineLocation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetermineLocation", reflect.TypeOf((*MockLocation)(nil).DetermineLocation), arg0, arg1)
}

// Error mocks base method.
func (m *MockLocation) Error(ctx context.Context, err error) openapi_CommonData.ProblemDetails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", ctx, err)
	ret0, _ := ret[0].(openapi_CommonData.ProblemDetails)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockLocationMockRecorder) Error(ctx, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLocation)(nil).Error), ctx, err)
}

// LocationContextTransfer mocks base method.
func (m *MockLocation) LocationContextTransfer(arg0 context.Context, arg1 openapi_Nlmf_Location.LocContextData) (openapi_CommonData.ProblemDetails, fivegc.RedirectResponse, nlmf.LocationContextTransferStatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationContextTransfer", arg0, arg1)
	ret0, _ := ret[0].(openapi_CommonData.ProblemDetails)
	ret1, _ := ret[1].(fivegc.RedirectResponse)
	ret2, _ := ret[2].(nlmf.LocationContextTransferStatusCode)
	return ret0, ret1, ret2
}

// LocationContextTransfer indicates an expected call of LocationContextTransfer.
func (mr *MockLocationMockRecorder) LocationContextTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationContextTransfer", reflect.TypeOf((*MockLocation)(nil).LocationContextTransfer), arg0, arg1)
}
