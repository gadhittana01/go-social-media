package services

import (
	"context"

	"github.com/gadhittana01/socialmedia/pkg/tag"
)

type TagService interface {
	GetTags(ctx context.Context) ([]GetTagsRow, error)
	CreateTag(ctx context.Context, tagname string) (CreateTagRow, error)
	UpdateTag(ctx context.Context, arg UpdateTagParams) (UpdateTagRow, error)
	DeleteTag(ctx context.Context, id int32) error
}

type tagService struct {
	tr TagResource
}

func NewTagService(TR TagResource) (TagService, error) {
	return &tagService{
		tr: TR,
	}, nil
}

func (ts *tagService) GetTags(ctx context.Context) ([]GetTagsRow, error) {
	var result []GetTagsRow = []GetTagsRow{}
	res, err := ts.tr.GetTags(ctx)
	if err != nil {
		return result, err
	}

	for _, tag := range res {
		result = append(result, GetTagsRow(tag))
	}

	return result, nil
}

func (ts *tagService) CreateTag(ctx context.Context, tagname string) (CreateTagRow, error) {
	var result CreateTagRow = CreateTagRow{}
	res, err := ts.tr.CreateTag(context.Background(), tagname)
	if err != nil {
		return result, err
	}
	result = CreateTagRow{
		ID:      res.ID,
		Tagname: res.Tagname,
	}
	return result, nil
}

func (ts *tagService) UpdateTag(ctx context.Context, arg UpdateTagParams) (UpdateTagRow, error) {
	var result UpdateTagRow = UpdateTagRow{}
	_, err := ts.tr.GetTag(context.Background(), int32(arg.ID))
	if err != nil {
		return result, err
	}

	res, err := ts.tr.UpdateTag(context.Background(), tag.UpdateTagParams{
		ID:      arg.ID,
		Tagname: arg.Tagname,
	})
	if err != nil {
		return result, err
	}
	result = UpdateTagRow{
		ID:      res.ID,
		Tagname: res.Tagname,
	}
	return result, nil
}

func (ts *tagService) DeleteTag(ctx context.Context, id int32) error {
	_, err := ts.tr.GetTag(context.Background(), int32(id))
	if err != nil {
		return err
	}

	err = ts.tr.DeleteTag(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}
