package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/gadhittana01/socialmedia/pkg/user"
	"github.com/golang/mock/gomock"
)

func TestNewUserService(t *testing.T) {
	ctrl := gomock.NewController(t)

	userMock := NewMockUserResource(ctrl)

	type args struct {
		UR UserResource
	}
	tests := []struct {
		name    string
		args    args
		want    UserService
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				UR: userMock,
			},
			want: &userService{
				ur: userMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserService(tt.args.UR)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx      context.Context
		fullname string
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *userService
		want    CreateUserRow
		wantErr bool
	}{
		{
			name: "success create user",
			args: args{
				ctx:      ctx,
				fullname: "Giri Putra Adhittana",
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().CreateUser(gomock.Any(), "Giri Putra Adhittana").Return(user.CreateUserRow{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}, nil)

				return &userService{
					ur: userMock,
				}
			},
			want: CreateUserRow{
				ID:       1,
				Fullname: "Giri Putra Adhittana",
			},
			wantErr: false,
		},
		{
			name: "error create user",
			args: args{
				ctx:      ctx,
				fullname: "Giri Putra Adhittana",
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().CreateUser(gomock.Any(), "Giri Putra Adhittana").Return(user.CreateUserRow{}, errors.New("error"))

				return &userService{
					ur: userMock,
				}
			},
			want:    CreateUserRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.CreateUser(tt.args.ctx, tt.args.fullname)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *userService
		want    []GetUsersRow
		wantErr bool
	}{
		{
			name: "success get users",
			args: args{
				ctx: ctx,
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().GetUsers(gomock.Any()).Return([]user.GetUsersRow{
					{
						ID:       1,
						Fullname: "Giri Putra Adhittana",
					},
					{
						ID:       2,
						Fullname: "Giri Adhittana",
					},
				}, nil)

				return &userService{
					ur: userMock,
				}
			},
			want: []GetUsersRow{
				{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				},
				{
					ID:       2,
					Fullname: "Giri Adhittana",
				},
			},
			wantErr: false,
		},
		{
			name: "error get users",
			args: args{
				ctx: ctx,
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().GetUsers(gomock.Any()).Return([]user.GetUsersRow{}, errors.New("error"))

				return &userService{
					ur: userMock,
				}
			},
			want:    []GetUsersRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
		arg UpdateUserParams
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *userService
		want    UpdateUserRow
		wantErr bool
	}{
		{
			name: "success update user",
			args: args{
				ctx: ctx,
				arg: UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				},
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().GetUser(gomock.Any(), int32(1)).Return(user.GetUserRow{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}, nil)

				userMock.EXPECT().UpdateUser(gomock.Any(), user.UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}).Return(user.UpdateUserRow{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}, nil)

				return &userService{
					ur: userMock,
				}
			},
			want: UpdateUserRow{
				ID:       1,
				Fullname: "Giri Putra Adhittana",
			},
			wantErr: false,
		},
		{
			name: "error get user",
			args: args{
				ctx: ctx,
				arg: UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				},
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().GetUser(gomock.Any(), int32(1)).Return(user.GetUserRow{}, errors.New("error"))

				return &userService{
					ur: userMock,
				}
			},
			want:    UpdateUserRow{},
			wantErr: true,
		},
		{
			name: "error update user",
			args: args{
				ctx: ctx,
				arg: UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				},
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().GetUser(gomock.Any(), int32(1)).Return(user.GetUserRow{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}, nil)

				userMock.EXPECT().UpdateUser(gomock.Any(), user.UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}).Return(user.UpdateUserRow{}, errors.New("error"))

				return &userService{
					ur: userMock,
				}
			},
			want:    UpdateUserRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.UpdateUser(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
		id  int32
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *userService
		wantErr bool
	}{
		{
			name: "success delete user",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().GetUser(gomock.Any(), int32(1)).Return(user.GetUserRow{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}, nil)

				userMock.EXPECT().DeleteUser(gomock.Any(), int32(1)).Return(nil)

				return &userService{
					ur: userMock,
				}
			},
			wantErr: false,
		},
		{
			name: "error get user",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().GetUser(gomock.Any(), int32(1)).Return(user.GetUserRow{}, errors.New("error"))

				return &userService{
					ur: userMock,
				}
			},
			wantErr: true,
		},
		{
			name: "error delete user",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *userService {
				userMock := NewMockUserResource(ctrl)

				userMock.EXPECT().GetUser(gomock.Any(), int32(1)).Return(user.GetUserRow{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				}, nil)

				userMock.EXPECT().DeleteUser(gomock.Any(), int32(1)).Return(errors.New("error"))

				return &userService{
					ur: userMock,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			err := p.DeleteUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
