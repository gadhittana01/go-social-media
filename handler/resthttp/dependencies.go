package resthttp

import (
	"context"

	"github.com/gadhittana01/socialmedia/services"
)

type (
	UserService interface {
		CreateUser(ctx context.Context, fullname string) (services.CreateUserRow, error)
		GetUsers(ctx context.Context) ([]services.GetUsersRow, error)
		UpdateUser(ctx context.Context, arg services.UpdateUserParams) (services.UpdateUserRow, error)
		DeleteUser(ctx context.Context, id int32) error
	}

	TagService interface {
		GetTags(ctx context.Context) ([]services.GetTagsRow, error)
		CreateTag(ctx context.Context, tagname string) (services.CreateTagRow, error)
		UpdateTag(ctx context.Context, arg services.UpdateTagParams) (services.UpdateTagRow, error)
		DeleteTag(ctx context.Context, id int32) error
	}

	PostService interface {
		CreatePost(ctx context.Context, arg services.CreatePostParams) (services.CreatePostRow, error)
		GetPosts(ctx context.Context) ([]services.GetPostsRow, error)
		UpdatePost(ctx context.Context, arg services.UpdatePostParams) (services.UpdatePostRow, error)
		DeletePost(ctx context.Context, id int32) error
	}
)
