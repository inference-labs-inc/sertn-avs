// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/inference-labs-inc/sertn-avs/operator (interfaces: AggregatorRpcClienter)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/rpc_client.go -package=mocks github.com/inference-labs-inc/sertn-avs/operator AggregatorRpcClienter
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	aggregator "github.com/inference-labs-inc/sertn-avs/aggregator"
	gomock "go.uber.org/mock/gomock"
)

// MockAggregatorRpcClienter is a mock of AggregatorRpcClienter interface.
type MockAggregatorRpcClienter struct {
	ctrl     *gomock.Controller
	recorder *MockAggregatorRpcClienterMockRecorder
}

// MockAggregatorRpcClienterMockRecorder is the mock recorder for MockAggregatorRpcClienter.
type MockAggregatorRpcClienterMockRecorder struct {
	mock *MockAggregatorRpcClienter
}

// NewMockAggregatorRpcClienter creates a new mock instance.
func NewMockAggregatorRpcClienter(ctrl *gomock.Controller) *MockAggregatorRpcClienter {
	mock := &MockAggregatorRpcClienter{ctrl: ctrl}
	mock.recorder = &MockAggregatorRpcClienterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregatorRpcClienter) EXPECT() *MockAggregatorRpcClienterMockRecorder {
	return m.recorder
}

// SendSignedTaskResponseToAggregator mocks base method.
func (m *MockAggregatorRpcClienter) SendSignedTaskResponseToAggregator(arg0 *aggregator.SignedTaskResponse) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendSignedTaskResponseToAggregator", arg0)
}

// SendSignedTaskResponseToAggregator indicates an expected call of SendSignedTaskResponseToAggregator.
func (mr *MockAggregatorRpcClienterMockRecorder) SendSignedTaskResponseToAggregator(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSignedTaskResponseToAggregator", reflect.TypeOf((*MockAggregatorRpcClienter)(nil).SendSignedTaskResponseToAggregator), arg0)
}
