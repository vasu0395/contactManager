// Code generated by MockGen. DO NOT EDIT.
// Source: ContactManager\service\service.go
// Package mock_service is a generated GoMock package.


package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockContactStore is a mock of ContactStore interface.
type MockContactStore struct {
	ctrl     *gomock.Controller
	recorder *MockContactStoreMockRecorder
}

// MockContactStoreMockRecorder is the mock recorder for MockContactStore.
type MockContactStoreMockRecorder struct {
	mock *MockContactStore
}

// NewMockContactStore creates a new mock instance.
func NewMockContactStore(ctrl *gomock.Controller) *MockContactStore {
	mock := &MockContactStore{ctrl: ctrl}
	mock.recorder = &MockContactStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContactStore) EXPECT() *MockContactStoreMockRecorder {
	return m.recorder
}

// DeleteContactNumber mocks base method.
func (m *MockContactStore) DeleteContactNumber(ctx context.Context, contactName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContactNumber", ctx, contactName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContactNumber indicates an expected call of DeleteContactNumber.
func (mr *MockContactStoreMockRecorder) DeleteContactNumber(ctx, contactName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContactNumber", reflect.TypeOf((*MockContactStore)(nil).DeleteContactNumber), ctx, contactName)
}

// GetContactNumber mocks base method.
func (m *MockContactStore) GetContactNumber(ctx context.Context, contactName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactNumber", ctx, contactName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactNumber indicates an expected call of GetContactNumber.
func (mr *MockContactStoreMockRecorder) GetContactNumber(ctx, contactName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactNumber", reflect.TypeOf((*MockContactStore)(nil).GetContactNumber), ctx, contactName)
}

// SetContactMapping mocks base method.
func (m *MockContactStore) SetContactMapping(ctx context.Context, contactName, phoneNumber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContactMapping", ctx, contactName, phoneNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetContactMapping indicates an expected call of SetContactMapping.
func (mr *MockContactStoreMockRecorder) SetContactMapping(ctx, contactName, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContactMapping", reflect.TypeOf((*MockContactStore)(nil).SetContactMapping), ctx, contactName, phoneNumber)
}

// UpdateContactNumber mocks base method.
func (m *MockContactStore) UpdateContactNumber(ctx context.Context, contactName, updatedNumber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContactNumber", ctx, contactName, updatedNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContactNumber indicates an expected call of UpdateContactNumber.
func (mr *MockContactStoreMockRecorder) UpdateContactNumber(ctx, contactName, updatedNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContactNumber", reflect.TypeOf((*MockContactStore)(nil).UpdateContactNumber), ctx, contactName, updatedNumber)
}