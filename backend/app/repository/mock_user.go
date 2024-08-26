// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -source=user.go -destination=mock_user.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"
	model "sudoku/model"

	mo "github.com/samber/mo"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// FindByGitHubUserID mocks base method.
func (m *MockIUserRepository) FindByGitHubUserID(githubUserID uint) (mo.Option[*model.User], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGitHubUserID", githubUserID)
	ret0, _ := ret[0].(mo.Option[*model.User])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGitHubUserID indicates an expected call of FindByGitHubUserID.
func (mr *MockIUserRepositoryMockRecorder) FindByGitHubUserID(githubUserID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGitHubUserID", reflect.TypeOf((*MockIUserRepository)(nil).FindByGitHubUserID), githubUserID)
}

// Save mocks base method.
func (m *MockIUserRepository) Save(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIUserRepositoryMockRecorder) Save(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIUserRepository)(nil).Save), user)
}