// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1 (interfaces: SlasherClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ethereum_beacon_rpc_v1 "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	eth "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	grpc "google.golang.org/grpc"
)

// MockSlasherClient is a mock of SlasherClient interface.
type MockSlasherClient struct {
	ctrl     *gomock.Controller
	recorder *MockSlasherClientMockRecorder
}

// MockSlasherClientMockRecorder is the mock recorder for MockSlasherClient.
type MockSlasherClientMockRecorder struct {
	mock *MockSlasherClient
}

// NewMockSlasherClient creates a new mock instance.
func NewMockSlasherClient(ctrl *gomock.Controller) *MockSlasherClient {
	mock := &MockSlasherClient{ctrl: ctrl}
	mock.recorder = &MockSlasherClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlasherClient) EXPECT() *MockSlasherClientMockRecorder {
	return m.recorder
}

// HighestAttestations mocks base method.
func (m *MockSlasherClient) HighestAttestations(arg0 context.Context, arg1 *ethereum_beacon_rpc_v1.HighestAttestationRequest, arg2 ...grpc.CallOption) (*ethereum_beacon_rpc_v1.HighestAttestationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HighestAttestations", varargs...)
	ret0, _ := ret[0].(*ethereum_beacon_rpc_v1.HighestAttestationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HighestAttestations indicates an expected call of HighestAttestations.
func (mr *MockSlasherClientMockRecorder) HighestAttestations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HighestAttestations", reflect.TypeOf((*MockSlasherClient)(nil).HighestAttestations), varargs...)
}

// IsSlashableAttestation mocks base method.
func (m *MockSlasherClient) IsSlashableAttestation(arg0 context.Context, arg1 *eth.IndexedAttestation, arg2 ...grpc.CallOption) (*ethereum_beacon_rpc_v1.AttesterSlashingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsSlashableAttestation", varargs...)
	ret0, _ := ret[0].(*ethereum_beacon_rpc_v1.AttesterSlashingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSlashableAttestation indicates an expected call of IsSlashableAttestation.
func (mr *MockSlasherClientMockRecorder) IsSlashableAttestation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSlashableAttestation", reflect.TypeOf((*MockSlasherClient)(nil).IsSlashableAttestation), varargs...)
}

// IsSlashableBlock mocks base method.
func (m *MockSlasherClient) IsSlashableBlock(arg0 context.Context, arg1 *eth.SignedBeaconBlockHeader, arg2 ...grpc.CallOption) (*ethereum_beacon_rpc_v1.ProposerSlashingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsSlashableBlock", varargs...)
	ret0, _ := ret[0].(*ethereum_beacon_rpc_v1.ProposerSlashingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSlashableBlock indicates an expected call of IsSlashableBlock.
func (mr *MockSlasherClientMockRecorder) IsSlashableBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSlashableBlock", reflect.TypeOf((*MockSlasherClient)(nil).IsSlashableBlock), varargs...)
}
