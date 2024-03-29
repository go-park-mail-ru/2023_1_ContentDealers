// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package mock_search is a generated GoMock package.
package mock_search

import (
	context "context"
	reflect "reflect"

	domain "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockExtender is a mock of Extender interface.
type MockExtender struct {
	ctrl     *gomock.Controller
	recorder *MockExtenderMockRecorder
}

// MockExtenderMockRecorder is the mock recorder for MockExtender.
type MockExtenderMockRecorder struct {
	mock *MockExtender
}

// NewMockExtender creates a new mock instance.
func NewMockExtender(ctrl *gomock.Controller) *MockExtender {
	mock := &MockExtender{ctrl: ctrl}
	mock.recorder = &MockExtenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtender) EXPECT() *MockExtenderMockRecorder {
	return m.recorder
}

// Extend mocks base method.
func (m *MockExtender) Extend(ctx context.Context, query string) (func(*domain.SearchResult), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Extend", ctx, query)
	ret0, _ := ret[0].(func(*domain.SearchResult))
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Extend indicates an expected call of Extend.
func (mr *MockExtenderMockRecorder) Extend(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Extend", reflect.TypeOf((*MockExtender)(nil).Extend), ctx, query)
}
