package resthttp

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gadhittana01/socialmedia/services"
	"github.com/golang/mock/gomock"
)

func Test_NewUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	userMock := NewMockUserService(ctrl)

	type args struct {
		userService UserService
	}
	tests := []struct {
		name string
		args args
		want *UserHandler
	}{
		{
			args: args{
				userService: userMock,
			},
			want: &UserHandler{
				userService: userMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("GET", "http://localhost:8000/users", strings.NewReader(``))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("GET", "http://localhost:8000/users", strings.NewReader(``))
	internalServerErrResp := httptest.NewRecorder()

	type fields struct {
		userService UserService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() UserHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)
				userMock.EXPECT().GetUsers(gomock.Any()).Return([]services.GetUsersRow{
					{
						ID:       1,
						Fullname: "Giri Putra Adhittana",
					},
					{
						ID:       2,
						Fullname: "Giri Adhittana",
					},
				}, nil)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)
				userMock.EXPECT().GetUsers(gomock.Any()).Return([]services.GetUsersRow{}, errors.New("error"))

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   internalServerErrResp,
				req: internalServerErrReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			field.GetUsers(tt.args.w, tt.args.req)
		})
	}
}

func Test_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("POST", "http://localhost:8000/user", strings.NewReader(`{
		"fullname" : "Giri Putra Adhittana"
	}`))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("POST", "http://localhost:8000/user", strings.NewReader(`{
		"fullname" : "Giri Putra Adhittana"
	}`))
	internalServerErrResp := httptest.NewRecorder()

	emptyFullnameReq := httptest.NewRequest("POST", "http://localhost:8000/user", strings.NewReader(`{
		"fullname" : ""
	}`))
	emptyFullnameResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("POST", "http://localhost:8000/user", strings.NewReader(""))
	badResp := httptest.NewRecorder()

	type fields struct {
		userService UserService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() UserHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)
				userMock.EXPECT().CreateUser(gomock.Any(), "Giri Putra Adhittana").Return(services.CreateUserRow{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}, nil)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test bad request",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test empty fullname",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   emptyFullnameResp,
				req: emptyFullnameReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)
				userMock.EXPECT().CreateUser(gomock.Any(), "Giri Putra Adhittana").Return(services.CreateUserRow{}, errors.New("error"))

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   internalServerErrResp,
				req: internalServerErrReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			field.CreateUser(tt.args.w, tt.args.req)
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("PUT", "http://localhost:8000/user?id=1", strings.NewReader(`{
		"fullname" : "Giri Putra Adhittana"
	}`))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("PUT", "http://localhost:8000/user?id=1", strings.NewReader(`{
		"fullname" : "Giri Putra Adhittana"
	}`))
	internalServerErrResp := httptest.NewRecorder()

	idErrParseReq := httptest.NewRequest("POST", "http://localhost:8000/user?id='error'", strings.NewReader(`{
		"fullname" : "Giri Putra Adhittana"
	}`))
	idErrParseResp := httptest.NewRecorder()

	idErrReq := httptest.NewRequest("PUT", "http://localhost:8000/user", strings.NewReader(`{
		"fullname" : "Giri Putra Adhittana"
	}`))
	idErrResp := httptest.NewRecorder()

	emptyFullnameReq := httptest.NewRequest("PUT", "http://localhost:8000/user?id=1", strings.NewReader(`{
		"fullname" : ""
	}`))
	emptyFullnameResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("PUT", "http://localhost:8000/user?id=1", strings.NewReader(""))
	badResp := httptest.NewRecorder()

	type fields struct {
		userService UserService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() UserHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				userMock.EXPECT().UpdateUser(gomock.Any(), services.UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}).Return(services.UpdateUserRow{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}, nil)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test bad request",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test empty fullname",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   emptyFullnameResp,
				req: emptyFullnameReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				userMock.EXPECT().UpdateUser(gomock.Any(), services.UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}).Return(services.UpdateUserRow{}, errors.New("error"))

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   internalServerErrResp,
				req: internalServerErrReq,
			},
		},
		{
			name: "test id not provided",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   idErrResp,
				req: idErrReq,
			},
		},
		{
			name: "test error parsed id",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   idErrParseResp,
				req: idErrParseReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			field.UpdateUser(tt.args.w, tt.args.req)
		})
	}
}

func Test_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("DELETE", "http://localhost:8000/user?id=1", strings.NewReader(``))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("DELETE", "http://localhost:8000/user?id=1", strings.NewReader(``))
	internalServerErrResp := httptest.NewRecorder()

	idErrParseReq := httptest.NewRequest("DELETE", "http://localhost:8000/user?id='error'", strings.NewReader(``))
	idErrParseResp := httptest.NewRecorder()

	idErrReq := httptest.NewRequest("DELETE", "http://localhost:8000/user", strings.NewReader(``))
	idErrResp := httptest.NewRecorder()

	type fields struct {
		userService UserService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() UserHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				userMock.EXPECT().DeleteUser(gomock.Any(), int32(1)).Return(nil)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				userMock.EXPECT().DeleteUser(gomock.Any(), int32(1)).Return(errors.New("error"))

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   internalServerErrResp,
				req: internalServerErrReq,
			},
		},
		{
			name: "test id not provided",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   idErrResp,
				req: idErrReq,
			},
		},
		{
			name: "test error parsed id",
			fields: func() UserHandler {
				userMock := NewMockUserService(ctrl)

				return UserHandler{
					userService: userMock,
				}
			},
			args: args{
				w:   idErrParseResp,
				req: idErrParseReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := tt.fields()
			field.DeleteUser(tt.args.w, tt.args.req)
		})
	}
}
