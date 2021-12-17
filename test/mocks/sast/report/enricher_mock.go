// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/sast/report (interfaces: Enricher)

// Package mock_report is a generated GoMock package.
package mock_report

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEnricher is a mock of Enricher interface.
type MockEnricher struct {
	ctrl     *gomock.Controller
	recorder *MockEnricherMockRecorder
}

// MockEnricherMockRecorder is the mock recorder for MockEnricher.
type MockEnricherMockRecorder struct {
	mock *MockEnricher
}

// NewMockEnricher creates a new mock instance.
func NewMockEnricher(ctrl *gomock.Controller) *MockEnricher {
	mock := &MockEnricher{ctrl: ctrl}
	mock.recorder = &MockEnricherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnricher) EXPECT() *MockEnricherMockRecorder {
	return m.recorder
}

// AddSimilarity mocks base method.
func (m *MockEnricher) AddSimilarity() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSimilarity")
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSimilarity indicates an expected call of AddSimilarity.
func (mr *MockEnricherMockRecorder) AddSimilarity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSimilarity", reflect.TypeOf((*MockEnricher)(nil).AddSimilarity))
}

// Marshal mocks base method.
func (m *MockEnricher) Marshal() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal.
func (mr *MockEnricherMockRecorder) Marshal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockEnricher)(nil).Marshal))
}

// Parse mocks base method.
func (m *MockEnricher) Parse(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Parse indicates an expected call of Parse.
func (mr *MockEnricherMockRecorder) Parse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockEnricher)(nil).Parse), arg0)
}
