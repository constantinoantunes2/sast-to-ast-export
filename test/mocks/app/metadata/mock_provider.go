// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/app/metadata (interfaces: MetadataProvider)

// Package mock_app_metadata is a generated GoMock package.
package mock_app_metadata

import (
	reflect "reflect"

	metadata "github.com/checkmarxDev/ast-sast-export/internal/app/metadata"
	gomock "github.com/golang/mock/gomock"
)

// MockMetadataProvider is a mock of MetadataProvider interface.
type MockMetadataProvider struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataProviderMockRecorder
}

// MockMetadataProviderMockRecorder is the mock recorder for MockMetadataProvider.
type MockMetadataProviderMockRecorder struct {
	mock *MockMetadataProvider
}

// NewMockMetadataProvider creates a new mock instance.
func NewMockMetadataProvider(ctrl *gomock.Controller) *MockMetadataProvider {
	mock := &MockMetadataProvider{ctrl: ctrl}
	mock.recorder = &MockMetadataProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataProvider) EXPECT() *MockMetadataProviderMockRecorder {
	return m.recorder
}

// GetMetadataRecord mocks base method.
func (m *MockMetadataProvider) GetMetadataRecord(arg0 string, arg1 []*metadata.Query) (*metadata.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadataRecord", arg0, arg1)
	ret0, _ := ret[0].(*metadata.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadataRecord indicates an expected call of GetMetadataRecord.
func (mr *MockMetadataProviderMockRecorder) GetMetadataRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataRecord", reflect.TypeOf((*MockMetadataProvider)(nil).GetMetadataRecord), arg0, arg1)
}
