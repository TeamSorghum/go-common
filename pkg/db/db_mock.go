// Code generated by MockGen. DO NOT EDIT.
// Source: db.go
//
// Generated by this command:
//
//	mockgen -write_package_comment=false -source=db.go -destination=db_mock.go -package db
//

package db

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder[T]
	isgomock struct{}
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder[T any] struct {
	mock *MockRepo[T]
}

// NewMockRepo creates a new mock instance.
func NewMockRepo[T any](ctrl *gomock.Controller) *MockRepo[T] {
	mock := &MockRepo[T]{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo[T]) EXPECT() *MockRepoMockRecorder[T] {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRepo[T]) Delete(ctx context.Context, do *T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, do)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepoMockRecorder[T]) Delete(ctx, do any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepo[T])(nil).Delete), ctx, do)
}

// Insert mocks base method.
func (m *MockRepo[T]) Insert(ctx context.Context, do *T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, do)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRepoMockRecorder[T]) Insert(ctx, do any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepo[T])(nil).Insert), ctx, do)
}

// QueryByID mocks base method.
func (m *MockRepo[T]) QueryByID(ctx context.Context, id int64) (*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryByID", ctx, id)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryByID indicates an expected call of QueryByID.
func (mr *MockRepoMockRecorder[T]) QueryByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryByID", reflect.TypeOf((*MockRepo[T])(nil).QueryByID), ctx, id)
}

// Update mocks base method.
func (m *MockRepo[T]) Update(ctx context.Context, do *T) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, do)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepoMockRecorder[T]) Update(ctx, do any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepo[T])(nil).Update), ctx, do)
}
