// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/docker/client.go

// Package mock_docker is a generated GoMock package.
package mock_docker

import (
	reflect "reflect"

	types "github.com/docker/docker/api/types"
	registry "github.com/docker/docker/api/types/registry"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockClient) Build(options types.ImageBuildOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockClientMockRecorder) Build(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockClient)(nil).Build), options)
}

// Login mocks base method.
func (m *MockClient) Login(auth registry.AuthConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", auth)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockClientMockRecorder) Login(auth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockClient)(nil).Login), auth)
}

// Pull mocks base method.
func (m *MockClient) Pull(refStr string, options types.ImagePullOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", refStr, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockClientMockRecorder) Pull(refStr, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockClient)(nil).Pull), refStr, options)
}

// Push mocks base method.
func (m *MockClient) Push(image string, options types.ImagePushOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", image, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockClientMockRecorder) Push(image, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockClient)(nil).Push), image, options)
}

// Tag mocks base method.
func (m *MockClient) Tag(source, target string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag", source, target)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockClientMockRecorder) Tag(source, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockClient)(nil).Tag), source, target)
}
