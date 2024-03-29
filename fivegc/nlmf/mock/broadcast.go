// Code generated by MockGen. DO NOT EDIT.
// Source: ../broadcast.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	fivegc "github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	nlmf "github.com/5GCoreNet/5GCoreNetSDK/fivegc/nlmf"
	openapi_CommonData "github.com/5GCoreNet/openapi/openapi_CommonData"
	openapi_Nlmf_Broadcast "github.com/5GCoreNet/openapi/openapi_Nlmf_Broadcast"
	gomock "github.com/golang/mock/gomock"
)

// MockBroadcast is a mock of Broadcast interface.
type MockBroadcast struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcastMockRecorder
}

// MockBroadcastMockRecorder is the mock recorder for MockBroadcast.
type MockBroadcastMockRecorder struct {
	mock *MockBroadcast
}

// NewMockBroadcast creates a new mock instance.
func NewMockBroadcast(ctrl *gomock.Controller) *MockBroadcast {
	mock := &MockBroadcast{ctrl: ctrl}
	mock.recorder = &MockBroadcastMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroadcast) EXPECT() *MockBroadcastMockRecorder {
	return m.recorder
}

// CipherKeyData mocks base method.
func (m *MockBroadcast) CipherKeyData(arg0 context.Context, arg1 openapi_Nlmf_Broadcast.CipherRequestData) (openapi_Nlmf_Broadcast.CipherResponseData, openapi_CommonData.ProblemDetails, fivegc.RedirectResponse, nlmf.CypherResponseStatusCode) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CipherKeyData", arg0, arg1)
	ret0, _ := ret[0].(openapi_Nlmf_Broadcast.CipherResponseData)
	ret1, _ := ret[1].(openapi_CommonData.ProblemDetails)
	ret2, _ := ret[2].(fivegc.RedirectResponse)
	ret3, _ := ret[3].(nlmf.CypherResponseStatusCode)
	return ret0, ret1, ret2, ret3
}

// CipherKeyData indicates an expected call of CipherKeyData.
func (mr *MockBroadcastMockRecorder) CipherKeyData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CipherKeyData", reflect.TypeOf((*MockBroadcast)(nil).CipherKeyData), arg0, arg1)
}

// Error mocks base method.
func (m *MockBroadcast) Error(ctx context.Context, err error) openapi_CommonData.ProblemDetails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", ctx, err)
	ret0, _ := ret[0].(openapi_CommonData.ProblemDetails)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockBroadcastMockRecorder) Error(ctx, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockBroadcast)(nil).Error), ctx, err)
}
