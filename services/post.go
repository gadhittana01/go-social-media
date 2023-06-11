package services

import (
	"context"

	"github.com/gadhittana01/socialmedia/pkg/post"
	"github.com/gadhittana01/socialmedia/pkg/post_tags"
)

type PostService interface {
	CreatePost(ctx context.Context, arg CreatePostParams) (CreatePostRow, error)
	GetPosts(ctx context.Context) ([]GetPostsRow, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (UpdatePostRow, error)
	DeletePost(ctx context.Context, id int32) error
}

type postService struct {
	pr  PostResource
	tr  TagResource
	ptr PostTagResource
}

func NewPostService(PR PostResource, TR TagResource, PTR PostTagResource) (PostService, error) {
	return &postService{
		pr:  PR,
		tr:  TR,
		ptr: PTR,
	}, nil
}

func (ps *postService) CreatePost(ctx context.Context, arg CreatePostParams) (CreatePostRow, error) {
	var result CreatePostRow = CreatePostRow{}
	var tagIDs []int32
	res, err := ps.pr.CreatePost(ctx, post.CreatePostParams{
		Userid:      arg.Userid,
		Title:       arg.Title,
		Description: arg.Description,
	})
	if err != nil {
		return result, err
	}

	for _, tagID := range arg.TagID {
		resPostTag, err := ps.ptr.CreatePostTag(ctx, post_tags.CreatePostTagParams{
			Postid: res.ID,
			Tagid:  tagID,
		})
		if err != nil {
			return result, err
		}
		tagIDs = append(tagIDs, resPostTag.Tagid)
	}

	result = CreatePostRow{
		ID:          res.ID,
		Userid:      res.Userid,
		Title:       res.Title,
		Description: res.Description,
		TagID:       tagIDs,
	}

	return result, nil
}

func (ps *postService) GetPosts(ctx context.Context) ([]GetPostsRow, error) {
	var result []GetPostsRow = []GetPostsRow{}
	res, err := ps.pr.GetPosts(context.Background())
	if err != nil {
		return result, err
	}

	for _, item := range res {
		res, err := ps.tr.GetTagByPostID(context.Background(), item.ID)
		if err != nil {
			return result, err
		}

		var tags = []GetTagByPostIDRow{}
		for _, tag := range res {
			tags = append(tags, GetTagByPostIDRow{
				ID:      tag.ID,
				Tagname: tag.Tagname,
			})
		}

		result = append(result, GetPostsRow{
			ID:          item.ID,
			Userid:      item.Userid,
			Title:       item.Title,
			Description: item.Description,
			Tags:        tags,
		})
	}

	return result, nil
}

func (ps *postService) UpdatePost(ctx context.Context, arg UpdatePostParams) (UpdatePostRow, error) {
	var result UpdatePostRow = UpdatePostRow{}
	_, err := ps.pr.GetPost(ctx, arg.ID)
	if err != nil {
		return result, err
	}

	res, err := ps.pr.UpdatePost(ctx, post.UpdatePostParams{
		ID:          arg.ID,
		Title:       arg.Title,
		Description: arg.Description,
	})
	if err != nil {
		return result, err
	}

	err = ps.ptr.DeletePostTag(ctx, arg.ID)
	if err != nil {
		return result, err
	}

	var tagIDs []int32
	for _, tagID := range arg.TagID {
		resPostTag, err := ps.ptr.CreatePostTag(ctx, post_tags.CreatePostTagParams{
			Postid: arg.ID,
			Tagid:  tagID,
		})
		if err != nil {
			return result, err
		}
		tagIDs = append(tagIDs, resPostTag.Tagid)
	}

	result = UpdatePostRow{
		Title:       res.Title,
		Description: res.Description,
		TagID:       tagIDs,
	}

	return result, nil
}

func (ps *postService) DeletePost(ctx context.Context, id int32) error {
	_, err := ps.pr.GetPost(ctx, id)
	if err != nil {
		return err
	}

	err = ps.ptr.DeletePostTag(ctx, id)
	if err != nil {
		return err
	}

	err = ps.pr.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
