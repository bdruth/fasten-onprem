// Code generated by MockGen. DO NOT EDIT.
// SourceCredential: interface.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	models "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockDatabaseRepository is a mock of DatabaseRepository interface.
type MockDatabaseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRepositoryMockRecorder
}

// MockDatabaseRepositoryMockRecorder is the mock recorder for MockDatabaseRepository.
type MockDatabaseRepositoryMockRecorder struct {
	mock *MockDatabaseRepository
}

// NewMockDatabaseRepository creates a new mock instance.
func NewMockDatabaseRepository(ctrl *gomock.Controller) *MockDatabaseRepository {
	mock := &MockDatabaseRepository{ctrl: ctrl}
	mock.recorder = &MockDatabaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRepository) EXPECT() *MockDatabaseRepositoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDatabaseRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatabaseRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDatabaseRepository)(nil).Close))
}

// CreateUser mocks base method.
func (m *MockDatabaseRepository) CreateUser(arg0 context.Context, arg1 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDatabaseRepositoryMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabaseRepository)(nil).CreateUser), arg0, arg1)
}
