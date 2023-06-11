package services

import (
	"context"

	"github.com/gadhittana01/socialmedia/pkg/user"
)

type UserService interface {
	CreateUser(ctx context.Context, fullname string) (CreateUserRow, error)
	GetUsers(ctx context.Context) ([]GetUsersRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
	DeleteUser(ctx context.Context, id int32) error
}

type userService struct {
	ur UserResource
}

func NewUserService(UR UserResource) (UserService, error) {
	return &userService{
		ur: UR,
	}, nil
}

func (us *userService) CreateUser(ctx context.Context, fullname string) (CreateUserRow, error) {
	var result CreateUserRow = CreateUserRow{}
	res, err := us.ur.CreateUser(context.Background(), fullname)
	if err != nil {
		return result, err
	}

	result = CreateUserRow{
		ID:       res.ID,
		Fullname: res.Fullname,
	}

	return result, nil
}

func (us *userService) GetUsers(ctx context.Context) ([]GetUsersRow, error) {
	var result []GetUsersRow = []GetUsersRow{}
	res, err := us.ur.GetUsers(context.Background())
	if err != nil {
		return result, err
	}
	for _, item := range res {
		result = append(result, GetUsersRow{
			ID:       item.ID,
			Fullname: item.Fullname,
		})
	}
	return result, nil
}

func (us *userService) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	var result UpdateUserRow = UpdateUserRow{}
	_, err := us.ur.GetUser(ctx, arg.ID)
	if err != nil {
		return result, err
	}

	res, err := us.ur.UpdateUser(context.Background(), user.UpdateUserParams{
		ID:       arg.ID,
		Fullname: arg.Fullname,
	})
	if err != nil {
		return result, err
	}
	result = UpdateUserRow{
		ID:       res.ID,
		Fullname: res.Fullname,
	}
	return result, nil
}

func (us *userService) DeleteUser(ctx context.Context, id int32) error {
	_, err := us.ur.GetUser(ctx, id)
	if err != nil {
		return err
	}

	err = us.ur.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
