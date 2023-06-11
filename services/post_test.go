package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/gadhittana01/socialmedia/pkg/post"
	"github.com/gadhittana01/socialmedia/pkg/post_tags"
	"github.com/gadhittana01/socialmedia/pkg/tag"
	"github.com/golang/mock/gomock"
)

func TestNewPostService(t *testing.T) {
	ctrl := gomock.NewController(t)

	postMock := NewMockPostResource(ctrl)
	postTagMock := NewMockPostTagResource(ctrl)
	tagMock := NewMockTagResource(ctrl)

	type args struct {
		PR  PostResource
		TR  TagResource
		PTR PostTagResource
	}
	tests := []struct {
		name    string
		args    args
		want    PostService
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				PR:  postMock,
				TR:  tagMock,
				PTR: postTagMock,
			},
			want: &postService{
				pr:  postMock,
				tr:  tagMock,
				ptr: postTagMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPostService(tt.args.PR, tt.args.TR, tt.args.PTR)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
		arg CreatePostParams
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *postService
		want    CreatePostRow
		wantErr bool
	}{
		{
			name: "success create post",
			args: args{
				ctx: ctx,
				arg: CreatePostParams{
					Userid:      1,
					Title:       "holiday yay",
					Description: "Yes Holiday",
					TagID:       []int32{1, 2, 3},
				},
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().CreatePost(gomock.Any(), post.CreatePostParams{
					Userid:      1,
					Title:       "holiday yay",
					Description: "Yes Holiday",
				}).Return(post.CreatePostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "Yes Holiday",
				}, nil)

				postTagMock.EXPECT().CreatePostTag(gomock.Any(), post_tags.CreatePostTagParams{
					Postid: 1,
					Tagid:  1,
				}).Return(post_tags.CreatePostTagRow{
					ID:     1,
					Postid: 1,
					Tagid:  1,
				}, nil)

				postTagMock.EXPECT().CreatePostTag(gomock.Any(), post_tags.CreatePostTagParams{
					Postid: 1,
					Tagid:  2,
				}).Return(post_tags.CreatePostTagRow{
					ID:     2,
					Postid: 1,
					Tagid:  2,
				}, nil)

				postTagMock.EXPECT().CreatePostTag(gomock.Any(), post_tags.CreatePostTagParams{
					Postid: 1,
					Tagid:  3,
				}).Return(post_tags.CreatePostTagRow{
					ID:     3,
					Postid: 1,
					Tagid:  3,
				}, nil)

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want: CreatePostRow{
				ID:          1,
				Userid:      1,
				Title:       "holiday yay",
				Description: "Yes Holiday",
				TagID:       []int32{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "error create post",
			args: args{
				ctx: ctx,
				arg: CreatePostParams{
					Userid:      1,
					Title:       "holiday yay",
					Description: "Yes Holiday",
					TagID:       []int32{1, 2, 3},
				},
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().CreatePost(gomock.Any(), post.CreatePostParams{
					Userid:      1,
					Title:       "holiday yay",
					Description: "Yes Holiday",
				}).Return(post.CreatePostRow{}, errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want:    CreatePostRow{},
			wantErr: true,
		},
		{
			name: "success create post tag",
			args: args{
				ctx: ctx,
				arg: CreatePostParams{
					Userid:      1,
					Title:       "holiday yay",
					Description: "Yes Holiday",
					TagID:       []int32{1, 2, 3},
				},
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().CreatePost(gomock.Any(), post.CreatePostParams{
					Userid:      1,
					Title:       "holiday yay",
					Description: "Yes Holiday",
				}).Return(post.CreatePostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "Yes Holiday",
				}, nil)

				postTagMock.EXPECT().CreatePostTag(gomock.Any(), post_tags.CreatePostTagParams{
					Postid: 1,
					Tagid:  1,
				}).Return(post_tags.CreatePostTagRow{}, errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want:    CreatePostRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.CreatePost(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *postService
		want    []GetPostsRow
		wantErr bool
	}{
		{
			name: "success get posts",
			args: args{
				ctx: ctx,
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPosts(gomock.Any()).Return([]post.GetPostsRow{
					{
						ID:          1,
						Userid:      1,
						Title:       "Book A",
						Description: "This is book A",
					},
					{
						ID:          2,
						Userid:      1,
						Title:       "Book B",
						Description: "This is book B",
					},
				}, nil)

				tagMock.EXPECT().GetTagByPostID(gomock.Any(), int32(1)).Return([]tag.GetTagByPostIDRow{
					{
						ID:      1,
						Tagname: "holiday",
					},
					{
						ID:      2,
						Tagname: "reading",
					},
				}, nil)

				tagMock.EXPECT().GetTagByPostID(gomock.Any(), int32(2)).Return([]tag.GetTagByPostIDRow{
					{
						ID:      2,
						Tagname: "reading",
					},
					{
						ID:      3,
						Tagname: "shopping",
					},
				}, nil)

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want: []GetPostsRow{
				{
					ID:          1,
					Userid:      1,
					Title:       "Book A",
					Description: "This is book A",
					Tags: []GetTagByPostIDRow{
						{
							ID:      1,
							Tagname: "holiday",
						},
						{
							ID:      2,
							Tagname: "reading",
						},
					},
				},
				{
					ID:          2,
					Userid:      1,
					Title:       "Book B",
					Description: "This is book B",
					Tags: []GetTagByPostIDRow{
						{
							ID:      2,
							Tagname: "reading",
						},
						{
							ID:      3,
							Tagname: "shopping",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error get posts",
			args: args{
				ctx: ctx,
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPosts(gomock.Any()).Return([]post.GetPostsRow{}, errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want:    []GetPostsRow{},
			wantErr: true,
		},
		{
			name: "error get tag by post id",
			args: args{
				ctx: ctx,
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPosts(gomock.Any()).Return([]post.GetPostsRow{
					{
						ID:          1,
						Userid:      1,
						Title:       "Book A",
						Description: "This is book A",
					},
					{
						ID:          2,
						Userid:      1,
						Title:       "Book B",
						Description: "This is book B",
					},
				}, nil)

				tagMock.EXPECT().GetTagByPostID(gomock.Any(), int32(1)).Return([]tag.GetTagByPostIDRow{}, errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want:    []GetPostsRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.GetPosts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
		arg UpdatePostParams
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *postService
		want    UpdatePostRow
		wantErr bool
	}{
		{
			name: "success update post",
			args: args{
				ctx: ctx,
				arg: UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
					TagID:       []int32{1, 2, 3},
				},
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postMock.EXPECT().UpdatePost(gomock.Any(), post.UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}).Return(post.UpdatePostRow{
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postTagMock.EXPECT().DeletePostTag(gomock.Any(), int32(1))

				postTagMock.EXPECT().CreatePostTag(gomock.Any(), post_tags.CreatePostTagParams{
					Postid: 1,
					Tagid:  1,
				}).Return(post_tags.CreatePostTagRow{
					ID:     1,
					Postid: 1,
					Tagid:  1,
				}, nil)

				postTagMock.EXPECT().CreatePostTag(gomock.Any(), post_tags.CreatePostTagParams{
					Postid: 1,
					Tagid:  2,
				}).Return(post_tags.CreatePostTagRow{
					ID:     2,
					Postid: 1,
					Tagid:  2,
				}, nil)

				postTagMock.EXPECT().CreatePostTag(gomock.Any(), post_tags.CreatePostTagParams{
					Postid: 1,
					Tagid:  3,
				}).Return(post_tags.CreatePostTagRow{
					ID:     3,
					Postid: 1,
					Tagid:  3,
				}, nil)

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want: UpdatePostRow{
				Title:       "holiday yay",
				Description: "yay yay yay",
				TagID:       []int32{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "error get post",
			args: args{
				ctx: ctx,
				arg: UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
					TagID:       []int32{1, 2, 3},
				},
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{}, errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want:    UpdatePostRow{},
			wantErr: true,
		},
		{
			name: "error update post",
			args: args{
				ctx: ctx,
				arg: UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
					TagID:       []int32{1, 2, 3},
				},
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postMock.EXPECT().UpdatePost(gomock.Any(), post.UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}).Return(post.UpdatePostRow{}, errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want:    UpdatePostRow{},
			wantErr: true,
		},
		{
			name: "error delete post tag",
			args: args{
				ctx: ctx,
				arg: UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
					TagID:       []int32{1, 2, 3},
				},
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postMock.EXPECT().UpdatePost(gomock.Any(), post.UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}).Return(post.UpdatePostRow{
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postTagMock.EXPECT().DeletePostTag(gomock.Any(), int32(1)).Return(errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want:    UpdatePostRow{},
			wantErr: true,
		},
		{
			name: "error create post tag",
			args: args{
				ctx: ctx,
				arg: UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
					TagID:       []int32{1, 2, 3},
				},
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postMock.EXPECT().UpdatePost(gomock.Any(), post.UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}).Return(post.UpdatePostRow{
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postTagMock.EXPECT().DeletePostTag(gomock.Any(), int32(1))

				postTagMock.EXPECT().CreatePostTag(gomock.Any(), post_tags.CreatePostTagParams{
					Postid: 1,
					Tagid:  1,
				}).Return(post_tags.CreatePostTagRow{}, errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			want:    UpdatePostRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			got, err := p.UpdatePost(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
		id  int32
	}
	tests := []struct {
		name    string
		args    args
		mock    func() *postService
		wantErr bool
	}{
		{
			name: "success delete post",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postTagMock.EXPECT().DeletePostTag(gomock.Any(), int32(1))

				postMock.EXPECT().DeletePost(gomock.Any(), int32(1)).Return(nil)

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			wantErr: false,
		},
		{
			name: "error get post",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{}, errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			wantErr: true,
		},
		{
			name: "error delete post tag",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postTagMock.EXPECT().DeletePostTag(gomock.Any(), int32(1)).Return(errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			wantErr: true,
		},
		{
			name: "error delete post",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mock: func() *postService {
				postMock := NewMockPostResource(ctrl)
				postTagMock := NewMockPostTagResource(ctrl)
				tagMock := NewMockTagResource(ctrl)

				postMock.EXPECT().GetPost(gomock.Any(), int32(1)).Return(post.GetPostRow{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "yay yay yay",
				}, nil)

				postTagMock.EXPECT().DeletePostTag(gomock.Any(), int32(1))

				postMock.EXPECT().DeletePost(gomock.Any(), int32(1)).Return(errors.New("error"))

				return &postService{
					pr:  postMock,
					tr:  tagMock,
					ptr: postTagMock,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.mock()
			err := p.DeletePost(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
