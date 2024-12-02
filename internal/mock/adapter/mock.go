// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/adapter/adapter.go

// Package mock_adapter is a generated GoMock package.
package mock_adapter

import (
	context "context"
	model "post-storage-service/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPostProvider is a mock of PostProvider interface.
type MockPostProvider struct {
	ctrl     *gomock.Controller
	recorder *MockPostProviderMockRecorder
}

// MockPostProviderMockRecorder is the mock recorder for MockPostProvider.
type MockPostProviderMockRecorder struct {
	mock *MockPostProvider
}

// NewMockPostProvider creates a new mock instance.
func NewMockPostProvider(ctrl *gomock.Controller) *MockPostProvider {
	mock := &MockPostProvider{ctrl: ctrl}
	mock.recorder = &MockPostProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostProvider) EXPECT() *MockPostProviderMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockPostProvider) Fetch(ctx context.Context, limit, offset int) ([]model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, limit, offset)
	ret0, _ := ret[0].([]model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockPostProviderMockRecorder) Fetch(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockPostProvider)(nil).Fetch), ctx, limit, offset)
}

// GetTotalPosts mocks base method.
func (m *MockPostProvider) GetTotalPosts(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalPosts", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalPosts indicates an expected call of GetTotalPosts.
func (mr *MockPostProviderMockRecorder) GetTotalPosts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalPosts", reflect.TypeOf((*MockPostProvider)(nil).GetTotalPosts), ctx)
}
