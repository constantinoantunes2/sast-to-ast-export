// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/integration/rest (interfaces: Client)

// Package mock_integration_rest is a generated GoMock package.
package mock_integration_rest

import (
	io "io"
	reflect "reflect"

	rest "github.com/checkmarxDev/ast-sast-export/internal/integration/rest"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockClient) Authenticate(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockClientMockRecorder) Authenticate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockClient)(nil).Authenticate), arg0, arg1)
}

// CreateScanReport mocks base method.
func (m *MockClient) CreateScanReport(arg0 int, arg1 string, arg2 rest.Retry) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScanReport", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScanReport indicates an expected call of CreateScanReport.
func (mr *MockClientMockRecorder) CreateScanReport(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScanReport", reflect.TypeOf((*MockClient)(nil).CreateScanReport), arg0, arg1, arg2)
}

// GetLdapRoleMappings mocks base method.
func (m *MockClient) GetLdapRoleMappings() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLdapRoleMappings")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLdapRoleMappings indicates an expected call of GetLdapRoleMappings.
func (mr *MockClientMockRecorder) GetLdapRoleMappings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLdapRoleMappings", reflect.TypeOf((*MockClient)(nil).GetLdapRoleMappings))
}

// GetLdapServers mocks base method.
func (m *MockClient) GetLdapServers() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLdapServers")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLdapServers indicates an expected call of GetLdapServers.
func (mr *MockClientMockRecorder) GetLdapServers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLdapServers", reflect.TypeOf((*MockClient)(nil).GetLdapServers))
}

// GetLdapTeamMappings mocks base method.
func (m *MockClient) GetLdapTeamMappings() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLdapTeamMappings")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLdapTeamMappings indicates an expected call of GetLdapTeamMappings.
func (mr *MockClientMockRecorder) GetLdapTeamMappings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLdapTeamMappings", reflect.TypeOf((*MockClient)(nil).GetLdapTeamMappings))
}

// GetProjectsWithLastScanID mocks base method.
func (m *MockClient) GetProjectsWithLastScanID(arg0 string, arg1, arg2 int) (*[]rest.ProjectWithLastScanID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsWithLastScanID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*[]rest.ProjectWithLastScanID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsWithLastScanID indicates an expected call of GetProjectsWithLastScanID.
func (mr *MockClientMockRecorder) GetProjectsWithLastScanID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsWithLastScanID", reflect.TypeOf((*MockClient)(nil).GetProjectsWithLastScanID), arg0, arg1, arg2)
}

// GetRoles mocks base method.
func (m *MockClient) GetRoles() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles.
func (mr *MockClientMockRecorder) GetRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockClient)(nil).GetRoles))
}

// GetSamlIdentityProviders mocks base method.
func (m *MockClient) GetSamlIdentityProviders() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSamlIdentityProviders")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamlIdentityProviders indicates an expected call of GetSamlIdentityProviders.
func (mr *MockClientMockRecorder) GetSamlIdentityProviders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamlIdentityProviders", reflect.TypeOf((*MockClient)(nil).GetSamlIdentityProviders))
}

// GetSamlRoleMappings mocks base method.
func (m *MockClient) GetSamlRoleMappings() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSamlRoleMappings")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamlRoleMappings indicates an expected call of GetSamlRoleMappings.
func (mr *MockClientMockRecorder) GetSamlRoleMappings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamlRoleMappings", reflect.TypeOf((*MockClient)(nil).GetSamlRoleMappings))
}

// GetSamlTeamMappings mocks base method.
func (m *MockClient) GetSamlTeamMappings() ([]*rest.SamlTeamMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSamlTeamMappings")
	ret0, _ := ret[0].([]*rest.SamlTeamMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamlTeamMappings indicates an expected call of GetSamlTeamMappings.
func (mr *MockClientMockRecorder) GetSamlTeamMappings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamlTeamMappings", reflect.TypeOf((*MockClient)(nil).GetSamlTeamMappings))
}

// GetTeams mocks base method.
func (m *MockClient) GetTeams() ([]*rest.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeams")
	ret0, _ := ret[0].([]*rest.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeams indicates an expected call of GetTeams.
func (mr *MockClientMockRecorder) GetTeams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeams", reflect.TypeOf((*MockClient)(nil).GetTeams))
}

// GetTriagedResultsByScanID mocks base method.
func (m *MockClient) GetTriagedResultsByScanID(arg0 int) (*[]rest.TriagedScanResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTriagedResultsByScanID", arg0)
	ret0, _ := ret[0].(*[]rest.TriagedScanResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTriagedResultsByScanID indicates an expected call of GetTriagedResultsByScanID.
func (mr *MockClientMockRecorder) GetTriagedResultsByScanID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTriagedResultsByScanID", reflect.TypeOf((*MockClient)(nil).GetTriagedResultsByScanID), arg0)
}

// GetUsers mocks base method.
func (m *MockClient) GetUsers() ([]*rest.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]*rest.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockClientMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockClient)(nil).GetUsers))
}

// PostResponseBody mocks base method.
func (m *MockClient) PostResponseBody(arg0 string, arg1 io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostResponseBody", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostResponseBody indicates an expected call of PostResponseBody.
func (mr *MockClientMockRecorder) PostResponseBody(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostResponseBody", reflect.TypeOf((*MockClient)(nil).PostResponseBody), arg0, arg1)
}
