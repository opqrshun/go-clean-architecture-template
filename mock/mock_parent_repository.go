// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/i_parent_repository.go

// Package mock_usecase is a generated GoMock package.
package mock

import (
	domain "gobackend/model"
	database "gobackend/repository/database"
	request "gobackend/model"
	response "gobackend/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockParentRepository is a mock of ParentRepository interface.
type MockParentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockParentRepositoryMockRecorder
}

// MockParentRepositoryMockRecorder is the mock recorder for MockParentRepository.
type MockParentRepositoryMockRecorder struct {
	mock *MockParentRepository
}

// NewMockParentRepository creates a new mock instance.
func NewMockParentRepository(ctrl *gomock.Controller) *MockParentRepository {
	mock := &MockParentRepository{ctrl: ctrl}
	mock.recorder = &MockParentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParentRepository) EXPECT() *MockParentRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockParentRepository) Delete(arg0 database.Base) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockParentRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockParentRepository)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockParentRepository) FindAll() ([]model.Parent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Parent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockParentRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockParentRepository)(nil).FindAll))
}

// FindAllFull mocks base method.
func (m *MockParentRepository) FindAllFull(arg0 *model.ParentQuery) ([]model.Parent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllFull", arg0)
	ret0, _ := ret[0].([]model.Parent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllFull indicates an expected call of FindAllFull.
func (mr *MockParentRepositoryMockRecorder) FindAllFull(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllFull", reflect.TypeOf((*MockParentRepository)(nil).FindAllFull), arg0)
}

// FindByID mocks base method.
func (m *MockParentRepository) FindByID(arg0 int) (model.Parent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(model.Parent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockParentRepositoryMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockParentRepository)(nil).FindByID), arg0)
}

// FindFullByID mocks base method.
func (m *MockParentRepository) FindFullByID(arg0 int) (model.Parent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFullByID", arg0)
	ret0, _ := ret[0].(model.Parent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFullByID indicates an expected call of FindFullByID.
func (mr *MockParentRepositoryMockRecorder) FindFullByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFullByID", reflect.TypeOf((*MockParentRepository)(nil).FindFullByID), arg0)
}

// FindUserIDByIDpID mocks base method.
func (m *MockParentRepository) FindUserIDByIDpID(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserIDByIDpID", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserIDByIDpID indicates an expected call of FindUserIDByIDpID.
func (mr *MockParentRepositoryMockRecorder) FindUserIDByIDpID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserIDByIDpID", reflect.TypeOf((*MockParentRepository)(nil).FindUserIDByIDpID), arg0)
}

// FindUserIDByUsername mocks base method.
func (m *MockParentRepository) FindUserIDByUsername(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserIDByUsername", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserIDByUsername indicates an expected call of FindUserIDByUsername.
func (mr *MockParentRepositoryMockRecorder) FindUserIDByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserIDByUsername", reflect.TypeOf((*MockParentRepository)(nil).FindUserIDByUsername), arg0)
}

// Store mocks base method.
func (m *MockParentRepository) Store(arg0 database.Base) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockParentRepositoryMockRecorder) Store(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockParentRepository)(nil).Store), arg0)
}

// Update mocks base method.
func (m *MockParentRepository) Update(arg0 database.Base) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockParentRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockParentRepository)(nil).Update), arg0)
}
