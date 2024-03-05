// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/docker/service.go

// Package mock_docker is a generated GoMock package.
package mock_docker

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockService) Build(dockerfile string, tags ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{dockerfile}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Build", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockServiceMockRecorder) Build(dockerfile interface{}, tags ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dockerfile}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockService)(nil).Build), varargs...)
}

// Pull mocks base method.
func (m *MockService) Pull(imageRefURL, registryAuth string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", imageRefURL, registryAuth)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockServiceMockRecorder) Pull(imageRefURL, registryAuth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockService)(nil).Pull), imageRefURL, registryAuth)
}

// Push mocks base method.
func (m *MockService) Push(imageRefURL, registryAuth string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", imageRefURL, registryAuth)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockServiceMockRecorder) Push(imageRefURL, registryAuth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockService)(nil).Push), imageRefURL, registryAuth)
}

// Tag mocks base method.
func (m *MockService) Tag(src, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag", src, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockServiceMockRecorder) Tag(src, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockService)(nil).Tag), src, dest)
}