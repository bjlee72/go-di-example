// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package component is a generated GoMock package.
package component

import (
	reflect "reflect"

	"go.uber.org/mock/gomock"
)

// MockDataProvider is a mock of DataProvider interface.
type MockDataProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDataProviderMockRecorder
}

// MockDataProviderMockRecorder is the mock recorder for MockDataProvider.
type MockDataProviderMockRecorder struct {
	mock *MockDataProvider
}

// NewMockDataProvider creates a new mock instance.
func NewMockDataProvider(ctrl *gomock.Controller) *MockDataProvider {
	mock := &MockDataProvider{ctrl: ctrl}
	mock.recorder = &MockDataProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataProvider) EXPECT() *MockDataProviderMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockDataProvider) Query() []*Record {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query")
	ret0, _ := ret[0].([]*Record)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockDataProviderMockRecorder) Query() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDataProvider)(nil).Query))
}
