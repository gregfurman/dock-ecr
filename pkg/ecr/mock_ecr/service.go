// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ecr/service.go

// Package mock_ecr is a generated GoMock package.
package mock_ecr

import (
	reflect "reflect"

	types "github.com/aws/aws-sdk-go-v2/service/ecr/types"
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

// CreateEcrRepository mocks base method.
func (m *MockService) CreateEcrRepository(repositoryName string, isMutableImageTags bool, repositoryTags map[string]string) (*types.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEcrRepository", repositoryName, isMutableImageTags, repositoryTags)
	ret0, _ := ret[0].(*types.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEcrRepository indicates an expected call of CreateEcrRepository.
func (mr *MockServiceMockRecorder) CreateEcrRepository(repositoryName, isMutableImageTags, repositoryTags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEcrRepository", reflect.TypeOf((*MockService)(nil).CreateEcrRepository), repositoryName, isMutableImageTags, repositoryTags)
}

// GetAllRepositories mocks base method.
func (m *MockService) GetAllRepositories() ([]types.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRepositories")
	ret0, _ := ret[0].([]types.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRepositories indicates an expected call of GetAllRepositories.
func (mr *MockServiceMockRecorder) GetAllRepositories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRepositories", reflect.TypeOf((*MockService)(nil).GetAllRepositories))
}

// GetAuth mocks base method.
func (m *MockService) GetAuth() (*types.AuthorizationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuth")
	ret0, _ := ret[0].(*types.AuthorizationData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuth indicates an expected call of GetAuth.
func (mr *MockServiceMockRecorder) GetAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuth", reflect.TypeOf((*MockService)(nil).GetAuth))
}

// GetImageScanResults mocks base method.
func (m *MockService) GetImageScanResults(repositoryName, imageDigest, imageTag string) ([]types.ImageScanFindings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageScanResults", repositoryName, imageDigest, imageTag)
	ret0, _ := ret[0].([]types.ImageScanFindings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageScanResults indicates an expected call of GetImageScanResults.
func (mr *MockServiceMockRecorder) GetImageScanResults(repositoryName, imageDigest, imageTag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageScanResults", reflect.TypeOf((*MockService)(nil).GetImageScanResults), repositoryName, imageDigest, imageTag)
}

// GetImages mocks base method.
func (m *MockService) GetImages(repositoryName string) ([]types.ImageDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImages", repositoryName)
	ret0, _ := ret[0].([]types.ImageDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImages indicates an expected call of GetImages.
func (mr *MockServiceMockRecorder) GetImages(repositoryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImages", reflect.TypeOf((*MockService)(nil).GetImages), repositoryName)
}

// GetRepositories mocks base method.
func (m *MockService) GetRepositories(filter func(types.Repository) bool) ([]types.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositories", filter)
	ret0, _ := ret[0].([]types.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositories indicates an expected call of GetRepositories.
func (mr *MockServiceMockRecorder) GetRepositories(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositories", reflect.TypeOf((*MockService)(nil).GetRepositories), filter)
}

// GetRepository mocks base method.
func (m *MockService) GetRepository(repositoryName string) (*types.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository", repositoryName)
	ret0, _ := ret[0].(*types.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockServiceMockRecorder) GetRepository(repositoryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockService)(nil).GetRepository), repositoryName)
}

// GetRepositoryNamesByPrefix mocks base method.
func (m *MockService) GetRepositoryNamesByPrefix(prefix string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryNamesByPrefix", prefix)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryNamesByPrefix indicates an expected call of GetRepositoryNamesByPrefix.
func (mr *MockServiceMockRecorder) GetRepositoryNamesByPrefix(prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryNamesByPrefix", reflect.TypeOf((*MockService)(nil).GetRepositoryNamesByPrefix), prefix)
}

// ListImages mocks base method.
func (m *MockService) ListImages(repositoryName string, tagStatus types.TagStatus, filter func(types.ImageDetail) bool) ([]types.ImageDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", repositoryName, tagStatus, filter)
	ret0, _ := ret[0].([]types.ImageDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockServiceMockRecorder) ListImages(repositoryName, tagStatus, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockService)(nil).ListImages), repositoryName, tagStatus, filter)
}

// TagImage mocks base method.
func (m *MockService) TagImage(repositoryName, imageDigest, imageTag string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagImage", repositoryName, imageDigest, imageTag)
	ret0, _ := ret[0].(error)
	return ret0
}

// TagImage indicates an expected call of TagImage.
func (mr *MockServiceMockRecorder) TagImage(repositoryName, imageDigest, imageTag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagImage", reflect.TypeOf((*MockService)(nil).TagImage), repositoryName, imageDigest, imageTag)
}
