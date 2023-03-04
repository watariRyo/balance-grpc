// Code generated by MockGen. DO NOT EDIT.
// Source: group.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/watariRyo/balance/server/proto"
)

// MockGroupRepository is a mock of GroupRepository interface.
type MockGroupRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGroupRepositoryMockRecorder
}

// MockGroupRepositoryMockRecorder is the mock recorder for MockGroupRepository.
type MockGroupRepositoryMockRecorder struct {
	mock *MockGroupRepository
}

// NewMockGroupRepository creates a new mock instance.
func NewMockGroupRepository(ctrl *gomock.Controller) *MockGroupRepository {
	mock := &MockGroupRepository{ctrl: ctrl}
	mock.recorder = &MockGroupRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupRepository) EXPECT() *MockGroupRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockGroupRepository) Delete(ctx context.Context, input *proto.GroupID) (*proto.GroupID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, input)
	ret0, _ := ret[0].(*proto.GroupID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockGroupRepositoryMockRecorder) Delete(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGroupRepository)(nil).Delete), ctx, input)
}

// Get mocks base method.
func (m *MockGroupRepository) Get(ctx context.Context, input *proto.GroupID) (*proto.GroupResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, input)
	ret0, _ := ret[0].(*proto.GroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGroupRepositoryMockRecorder) Get(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGroupRepository)(nil).Get), ctx, input)
}

// List mocks base method.
func (m *MockGroupRepository) List(tx context.Context, input *proto.GroupListRequest) (*proto.GroupListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", tx, input)
	ret0, _ := ret[0].(*proto.GroupListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockGroupRepositoryMockRecorder) List(tx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGroupRepository)(nil).List), tx, input)
}

// Upsert mocks base method.
func (m *MockGroupRepository) Upsert(ctx context.Context, input *proto.GroupRequest) (*proto.GroupResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, input)
	ret0, _ := ret[0].(*proto.GroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockGroupRepositoryMockRecorder) Upsert(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockGroupRepository)(nil).Upsert), ctx, input)
}
