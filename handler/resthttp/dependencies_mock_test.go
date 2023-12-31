// Code generated by MockGen. DO NOT EDIT.
// Source: ./handler/resthttp/dependencies.go

// Package mock_resthttp is a generated GoMock package.
package resthttp

import (
	context "context"
	reflect "reflect"

	services "github.com/gadhittana01/socialmedia/services"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(ctx context.Context, fullname string) (services.CreateUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, fullname)
	ret0, _ := ret[0].(services.CreateUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(ctx, fullname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), ctx, fullname)
}

// DeleteUser mocks base method.
func (m *MockUserService) DeleteUser(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserService)(nil).DeleteUser), ctx, id)
}

// GetUsers mocks base method.
func (m *MockUserService) GetUsers(ctx context.Context) ([]services.GetUsersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx)
	ret0, _ := ret[0].([]services.GetUsersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServiceMockRecorder) GetUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserService)(nil).GetUsers), ctx)
}

// UpdateUser mocks base method.
func (m *MockUserService) UpdateUser(ctx context.Context, arg services.UpdateUserParams) (services.UpdateUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, arg)
	ret0, _ := ret[0].(services.UpdateUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServiceMockRecorder) UpdateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserService)(nil).UpdateUser), ctx, arg)
}

// MockTagService is a mock of TagService interface.
type MockTagService struct {
	ctrl     *gomock.Controller
	recorder *MockTagServiceMockRecorder
}

// MockTagServiceMockRecorder is the mock recorder for MockTagService.
type MockTagServiceMockRecorder struct {
	mock *MockTagService
}

// NewMockTagService creates a new mock instance.
func NewMockTagService(ctrl *gomock.Controller) *MockTagService {
	mock := &MockTagService{ctrl: ctrl}
	mock.recorder = &MockTagServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagService) EXPECT() *MockTagServiceMockRecorder {
	return m.recorder
}

// CreateTag mocks base method.
func (m *MockTagService) CreateTag(ctx context.Context, tagname string) (services.CreateTagRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", ctx, tagname)
	ret0, _ := ret[0].(services.CreateTagRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockTagServiceMockRecorder) CreateTag(ctx, tagname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockTagService)(nil).CreateTag), ctx, tagname)
}

// DeleteTag mocks base method.
func (m *MockTagService) DeleteTag(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTag", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag.
func (mr *MockTagServiceMockRecorder) DeleteTag(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockTagService)(nil).DeleteTag), ctx, id)
}

// GetTags mocks base method.
func (m *MockTagService) GetTags(ctx context.Context) ([]services.GetTagsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags", ctx)
	ret0, _ := ret[0].([]services.GetTagsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags.
func (mr *MockTagServiceMockRecorder) GetTags(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockTagService)(nil).GetTags), ctx)
}

// UpdateTag mocks base method.
func (m *MockTagService) UpdateTag(ctx context.Context, arg services.UpdateTagParams) (services.UpdateTagRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTag", ctx, arg)
	ret0, _ := ret[0].(services.UpdateTagRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTag indicates an expected call of UpdateTag.
func (mr *MockTagServiceMockRecorder) UpdateTag(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTag", reflect.TypeOf((*MockTagService)(nil).UpdateTag), ctx, arg)
}

// MockPostService is a mock of PostService interface.
type MockPostService struct {
	ctrl     *gomock.Controller
	recorder *MockPostServiceMockRecorder
}

// MockPostServiceMockRecorder is the mock recorder for MockPostService.
type MockPostServiceMockRecorder struct {
	mock *MockPostService
}

// NewMockPostService creates a new mock instance.
func NewMockPostService(ctrl *gomock.Controller) *MockPostService {
	mock := &MockPostService{ctrl: ctrl}
	mock.recorder = &MockPostServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostService) EXPECT() *MockPostServiceMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockPostService) CreatePost(ctx context.Context, arg services.CreatePostParams) (services.CreatePostRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, arg)
	ret0, _ := ret[0].(services.CreatePostRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockPostServiceMockRecorder) CreatePost(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostService)(nil).CreatePost), ctx, arg)
}

// DeletePost mocks base method.
func (m *MockPostService) DeletePost(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostServiceMockRecorder) DeletePost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostService)(nil).DeletePost), ctx, id)
}

// GetPosts mocks base method.
func (m *MockPostService) GetPosts(ctx context.Context) ([]services.GetPostsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", ctx)
	ret0, _ := ret[0].([]services.GetPostsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockPostServiceMockRecorder) GetPosts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockPostService)(nil).GetPosts), ctx)
}

// UpdatePost mocks base method.
func (m *MockPostService) UpdatePost(ctx context.Context, arg services.UpdatePostParams) (services.UpdatePostRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", ctx, arg)
	ret0, _ := ret[0].(services.UpdatePostRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostServiceMockRecorder) UpdatePost(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostService)(nil).UpdatePost), ctx, arg)
}
