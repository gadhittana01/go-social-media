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

func Test_NewPostHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	postMock := NewMockPostService(ctrl)

	type args struct {
		postService PostService
	}
	tests := []struct {
		name string
		args args
		want *PostHandler
	}{
		{
			args: args{
				postService: postMock,
			},
			want: &PostHandler{
				postService: postMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostHandler(tt.args.postService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("GET", "http://localhost:8000/posts", strings.NewReader(``))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("GET", "http://localhost:8000/posts", strings.NewReader(``))
	internalServerErrResp := httptest.NewRecorder()

	type fields struct {
		postService PostService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() PostHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)
				postMock.EXPECT().GetPosts(gomock.Any()).Return([]services.GetPostsRow{
					{
						ID:          1,
						Userid:      1,
						Title:       "title A",
						Description: "description A",
						Tags: []services.GetTagByPostIDRow{
							{
								ID:      1,
								Tagname: "holiday",
							},
						},
					},
				}, nil)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)
				postMock.EXPECT().GetPosts(gomock.Any()).Return([]services.GetPostsRow{}, errors.New("error"))

				return PostHandler{
					postService: postMock,
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
			field.GetPosts(tt.args.w, tt.args.req)
		})
	}
}

func Test_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("POST", "http://localhost:8000/post", strings.NewReader(`{
		"user_id" : 1,
		"title" : "Upa",
		"description" : "Dayo",
		"tag_ids" : [3, 4]
	}`))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("POST", "http://localhost:8000/post", strings.NewReader(`{
		"user_id" : 1,
		"title" : "Upa",
		"description" : "Dayo",
		"tag_ids" : [3, 4]
	}`))
	internalServerErrResp := httptest.NewRecorder()

	emptyUserIDReq := httptest.NewRequest("POST", "http://localhost:8000/post", strings.NewReader(`{
		"title" : "Upa",
		"description" : "Dayo",
		"tag_ids" : [3, 4]
	}`))
	emptyUserIDResp := httptest.NewRecorder()

	emptyTitleReq := httptest.NewRequest("POST", "http://localhost:8000/post", strings.NewReader(`{
		"user_id" : 1,
		"description" : "Dayo",
		"tag_ids" : [3, 4]
	}`))
	emptyTitleResp := httptest.NewRecorder()

	emptyDescriptionReq := httptest.NewRequest("POST", "http://localhost:8000/post", strings.NewReader(`{
		"user_id" : 1,
		"title" : "Upa",
		"tag_ids" : [3, 4]
	}`))
	emptyDescriptionResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("POST", "http://localhost:8000/post", strings.NewReader(""))
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
		fields func() PostHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				postMock.EXPECT().CreatePost(gomock.Any(), services.CreatePostParams{
					Userid:      1,
					Title:       "Upa",
					Description: "Dayo",
					TagID:       []int32{3, 4},
				}).Return(services.CreatePostRow{
					ID:          1,
					Userid:      1,
					Title:       "Upa",
					Description: "Dayo",
					TagID:       []int32{3, 4},
				}, nil)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test bad request",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test empty userid",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   emptyUserIDResp,
				req: emptyUserIDReq,
			},
		},
		{
			name: "test empty title",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   emptyTitleResp,
				req: emptyTitleReq,
			},
		},
		{
			name: "test empty description",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   emptyDescriptionResp,
				req: emptyDescriptionReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				postMock.EXPECT().CreatePost(gomock.Any(), services.CreatePostParams{
					Userid:      1,
					Title:       "Upa",
					Description: "Dayo",
					TagID:       []int32{3, 4},
				}).Return(services.CreatePostRow{}, errors.New("error"))

				return PostHandler{
					postService: postMock,
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
			field.CreatePost(tt.args.w, tt.args.req)
		})
	}
}

func Test_UpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("PUT", "http://localhost:8000/post?id=1", strings.NewReader(`{
		"title" : "uaya",
		"description" : "Yoyooy",
		"tag_ids" : [1, 2]
	}`))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("PUT", "http://localhost:8000/post?id=1", strings.NewReader(`{
		"title" : "uaya",
		"description" : "Yoyooy",
		"tag_ids" : [1, 2]
	}`))
	internalServerErrResp := httptest.NewRecorder()

	idErrParseReq := httptest.NewRequest("PUT", "http://localhost:8000/post?id='error'", strings.NewReader(`{
		"title" : "uaya",
		"description" : "Yoyooy",
		"tag_ids" : [1, 2]
	}`))
	idErrParseResp := httptest.NewRecorder()

	idErrReq := httptest.NewRequest("PUT", "http://localhost:8000/post", strings.NewReader(`{
		"title" : "uaya",
		"description" : "Yoyooy",
		"tag_ids" : [1, 2]
	}`))
	idErrResp := httptest.NewRecorder()

	emptyTitleReq := httptest.NewRequest("POST", "http://localhost:8000/post?id=1", strings.NewReader(`{
		"description" : "Yoyooy",
		"tag_ids" : [1, 2]
	}`))
	emptyTitleResp := httptest.NewRecorder()

	emptyDescriptionReq := httptest.NewRequest("PUT", "http://localhost:8000/post?id=1", strings.NewReader(`{
		"title" : "uaya",
		"tag_ids" : [1, 2]
	}`))
	emptyDescriptionResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("PUT", "http://localhost:8000/post?id=1", strings.NewReader(""))
	badResp := httptest.NewRecorder()

	type fields struct {
		postService PostService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() PostHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				postMock.EXPECT().UpdatePost(gomock.Any(), services.UpdatePostParams{
					ID:          1,
					Title:       "uaya",
					Description: "Yoyooy",
					TagID:       []int32{1, 2},
				}).Return(services.UpdatePostRow{
					Title:       "uaya",
					Description: "Yoyooy",
					TagID:       []int32{1, 2},
				}, nil)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test bad request",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test empty title",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   emptyTitleResp,
				req: emptyTitleReq,
			},
		},
		{
			name: "test empty description",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   emptyDescriptionResp,
				req: emptyDescriptionReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				postMock.EXPECT().UpdatePost(gomock.Any(), services.UpdatePostParams{
					ID:          1,
					Title:       "uaya",
					Description: "Yoyooy",
					TagID:       []int32{1, 2},
				}).Return(services.UpdatePostRow{}, errors.New("error"))

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   internalServerErrResp,
				req: internalServerErrReq,
			},
		},
		{
			name: "test id not provided",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   idErrResp,
				req: idErrReq,
			},
		},
		{
			name: "test error parsed id",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
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
			field.UpdatePost(tt.args.w, tt.args.req)
		})
	}
}

func Test_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("DELETE", "http://localhost:8000/post?id=1", strings.NewReader(``))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("DELETE", "http://localhost:8000/post?id=1", strings.NewReader(``))
	internalServerErrResp := httptest.NewRecorder()

	idErrParseReq := httptest.NewRequest("DELETE", "http://localhost:8000/post?id='error'", strings.NewReader(``))
	idErrParseResp := httptest.NewRecorder()

	idErrReq := httptest.NewRequest("DELETE", "http://localhost:8000/post", strings.NewReader(``))
	idErrResp := httptest.NewRecorder()

	type fields struct {
		postService PostService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() PostHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				postMock.EXPECT().DeletePost(gomock.Any(), int32(1)).Return(nil)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				postMock.EXPECT().DeletePost(gomock.Any(), int32(1)).Return(errors.New("error"))

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   internalServerErrResp,
				req: internalServerErrReq,
			},
		},
		{
			name: "test id not provided",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
				}
			},
			args: args{
				w:   idErrResp,
				req: idErrReq,
			},
		},
		{
			name: "test error parsed id",
			fields: func() PostHandler {
				postMock := NewMockPostService(ctrl)

				return PostHandler{
					postService: postMock,
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
			field.DeletePost(tt.args.w, tt.args.req)
		})
	}
}
