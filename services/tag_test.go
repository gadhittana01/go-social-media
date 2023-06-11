package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/gadhittana01/socialmedia/pkg/tag"
	"github.com/golang/mock/gomock"
)

func TestNewTagService(t *testing.T) {
	ctrl := gomock.NewController(t)

	tagMock := NewMockTagResource(ctrl)

	type args struct {
		TR TagResource
	}
	tests := []struct {
		name    string
		args    args
		want    TagService
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				TR: tagMock,
			},
			want: &tagService{
				tr: tagMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTagService(tt.args.TR)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTagService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *tagService
		want    []GetTagsRow
		wantErr bool
	}{
		{
			name: "success get tags",
			args: args{
				ctx: ctx,
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().GetTags(gomock.Any()).Return([]tag.GetTagsRow{
					{
						ID:      1,
						Tagname: "holiday",
					},
				}, nil)

				return &tagService{
					tr: tagMock,
				}
			},
			want: []GetTagsRow{
				{
					ID:      1,
					Tagname: "holiday",
				},
			},
			wantErr: false,
		},
		{
			name: "error get tags",
			args: args{
				ctx: ctx,
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().GetTags(gomock.Any()).Return([]tag.GetTagsRow{}, errors.New("error"))

				return &tagService{
					tr: tagMock,
				}
			},
			want:    []GetTagsRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.GetTags(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreateTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		tagname string
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *tagService
		want    CreateTagRow
		wantErr bool
	}{
		{
			name: "success create tag",
			args: args{
				ctx:     ctx,
				tagname: "holiday",
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().CreateTag(gomock.Any(), "holiday").Return(tag.CreateTagRow{
					ID:      1,
					Tagname: "holiday",
				}, nil)

				return &tagService{
					tr: tagMock,
				}
			},
			want: CreateTagRow{
				ID:      1,
				Tagname: "holiday",
			},
			wantErr: false,
		},
		{
			name: "error create tag",
			args: args{
				ctx:     ctx,
				tagname: "holiday",
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().CreateTag(gomock.Any(), "holiday").Return(tag.CreateTagRow{}, errors.New("error"))

				return &tagService{
					tr: tagMock,
				}
			},
			want:    CreateTagRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.CreateTag(tt.args.ctx, tt.args.tagname)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UpdateTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
		arg UpdateTagParams
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *tagService
		want    UpdateTagRow
		wantErr bool
	}{
		{
			name: "success update tag",
			args: args{
				ctx: ctx,
				arg: UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				},
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().GetTag(gomock.Any(), int32(1)).Return(tag.GetTagRow{
					ID:      1,
					Tagname: "holiday",
				}, nil)

				tagMock.EXPECT().UpdateTag(gomock.Any(), tag.UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				}).Return(tag.UpdateTagRow{
					ID:      1,
					Tagname: "holiday",
				}, nil)

				return &tagService{
					tr: tagMock,
				}
			},
			want: UpdateTagRow{
				ID:      1,
				Tagname: "holiday",
			},
			wantErr: false,
		},
		{
			name: "error get tag",
			args: args{
				ctx: ctx,
				arg: UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				},
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().GetTag(gomock.Any(), int32(1)).Return(tag.GetTagRow{}, errors.New("error"))

				return &tagService{
					tr: tagMock,
				}
			},
			want:    UpdateTagRow{},
			wantErr: true,
		},
		{
			name: "error update tag",
			args: args{
				ctx: ctx,
				arg: UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				},
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().GetTag(gomock.Any(), int32(1)).Return(tag.GetTagRow{
					ID:      1,
					Tagname: "holiday",
				}, nil)

				tagMock.EXPECT().UpdateTag(gomock.Any(), tag.UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				}).Return(tag.UpdateTagRow{}, errors.New("error"))

				return &tagService{
					tr: tagMock,
				}
			},
			want:    UpdateTagRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.UpdateTag(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DeleteTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
		id  int32
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *tagService
		wantErr bool
	}{
		{
			name: "success delete tag",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().GetTag(gomock.Any(), int32(1)).Return(tag.GetTagRow{
					ID:      1,
					Tagname: "holiday",
				}, nil)

				tagMock.EXPECT().DeleteTag(gomock.Any(), int32(1)).Return(nil)

				return &tagService{
					tr: tagMock,
				}
			},
			wantErr: false,
		},
		{
			name: "error get tag",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().GetTag(gomock.Any(), int32(1)).Return(tag.GetTagRow{}, errors.New("error"))

				return &tagService{
					tr: tagMock,
				}
			},
			wantErr: true,
		},
		{
			name: "error delete tag",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *tagService {
				tagMock := NewMockTagResource(ctrl)

				tagMock.EXPECT().GetTag(gomock.Any(), int32(1)).Return(tag.GetTagRow{
					ID:      1,
					Tagname: "holiday",
				}, nil)

				tagMock.EXPECT().DeleteTag(gomock.Any(), int32(1)).Return(errors.New("error"))

				return &tagService{
					tr: tagMock,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			err := p.DeleteTag(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
