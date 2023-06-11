package services

import (
	"context"

	"github.com/gadhittana01/socialmedia/pkg/post"
	"github.com/gadhittana01/socialmedia/pkg/post_tags"
	"github.com/gadhittana01/socialmedia/pkg/tag"
	"github.com/gadhittana01/socialmedia/pkg/user"
)

type (
	UserResource interface {
		CreateUser(ctx context.Context, fullname string) (user.CreateUserRow, error)
		GetUsers(ctx context.Context) ([]user.GetUsersRow, error)
		UpdateUser(ctx context.Context, arg user.UpdateUserParams) (user.UpdateUserRow, error)
		DeleteUser(ctx context.Context, id int32) error
		GetUser(ctx context.Context, id int32) (user.GetUserRow, error)
	}

	PostResource interface {
		CreatePost(ctx context.Context, arg post.CreatePostParams) (post.CreatePostRow, error)
		GetPosts(ctx context.Context) ([]post.GetPostsRow, error)
		UpdatePost(ctx context.Context, arg post.UpdatePostParams) (post.UpdatePostRow, error)
		DeletePost(ctx context.Context, id int32) error
		GetPost(ctx context.Context, id int32) (post.GetPostRow, error)
	}

	TagResource interface {
		CreateTag(ctx context.Context, tagname string) (tag.CreateTagRow, error)
		GetTagByPostID(ctx context.Context, postid int32) ([]tag.GetTagByPostIDRow, error)
		GetTags(ctx context.Context) ([]tag.GetTagsRow, error)
		UpdateTag(ctx context.Context, arg tag.UpdateTagParams) (tag.UpdateTagRow, error)
		DeleteTag(ctx context.Context, id int32) error
		GetTag(ctx context.Context, id int32) (tag.GetTagRow, error)
	}

	PostTagResource interface {
		CreatePostTag(ctx context.Context, arg post_tags.CreatePostTagParams) (post_tags.CreatePostTagRow, error)
		DeletePostTag(ctx context.Context, postid int32) error
	}
)
