// Code generated by MockGen. DO NOT EDIT.
// Source: repository/todo_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mFsl16/go-restapi-example/model"
)

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// DeleteTodo mocks base method.
func (m *MockTodoRepository) DeleteTodo(ctx context.Context, db *sql.DB, id int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", ctx, db, id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockTodoRepositoryMockRecorder) DeleteTodo(ctx, db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockTodoRepository)(nil).DeleteTodo), ctx, db, id)
}

// FindAll mocks base method.
func (m *MockTodoRepository) FindAll(ctx context.Context, db *sql.DB) []model.Todo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, db)
	ret0, _ := ret[0].([]model.Todo)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockTodoRepositoryMockRecorder) FindAll(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTodoRepository)(nil).FindAll), ctx, db)
}

// FindById mocks base method.
func (m *MockTodoRepository) FindById(ctx context.Context, db *sql.DB, id int) (model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, db, id)
	ret0, _ := ret[0].(model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockTodoRepositoryMockRecorder) FindById(ctx, db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockTodoRepository)(nil).FindById), ctx, db, id)
}

// InsertTodo mocks base method.
func (m *MockTodoRepository) InsertTodo(ctx context.Context, db *sql.DB, todo model.Todo) (model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTodo", ctx, db, todo)
	ret0, _ := ret[0].(model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTodo indicates an expected call of InsertTodo.
func (mr *MockTodoRepositoryMockRecorder) InsertTodo(ctx, db, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTodo", reflect.TypeOf((*MockTodoRepository)(nil).InsertTodo), ctx, db, todo)
}

// UpdateTodo mocks base method.
func (m *MockTodoRepository) UpdateTodo(ctx context.Context, db *sql.DB, id int, todo model.Todo) (model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", ctx, db, id, todo)
	ret0, _ := ret[0].(model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockTodoRepositoryMockRecorder) UpdateTodo(ctx, db, id, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockTodoRepository)(nil).UpdateTodo), ctx, db, id, todo)
}
