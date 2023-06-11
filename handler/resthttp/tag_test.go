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

func Test_NewTagHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	tagMock := NewMockTagService(ctrl)

	type args struct {
		tagService TagService
	}
	tests := []struct {
		name string
		args args
		want *TagHandler
	}{
		{
			args: args{
				tagService: tagMock,
			},
			want: &TagHandler{
				tagService: tagMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTagHandler(tt.args.tagService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("GET", "http://localhost:8000/tags", strings.NewReader(``))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("GET", "http://localhost:8000/tags", strings.NewReader(``))
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
		fields func() TagHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)
				tagMock.EXPECT().GetTags(gomock.Any()).Return([]services.GetTagsRow{
					{
						ID:      1,
						Tagname: "holiday",
					},
					{
						ID:      2,
						Tagname: "reading",
					},
				}, nil)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)
				tagMock.EXPECT().GetTags(gomock.Any()).Return([]services.GetTagsRow{}, errors.New("error"))

				return TagHandler{
					tagService: tagMock,
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
			field.GetTags(tt.args.w, tt.args.req)
		})
	}
}

func Test_CreateTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("POST", "http://localhost:8000/tag", strings.NewReader(`{
		"tagname" : "holiday"
	}`))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("POST", "http://localhost:8000/tag", strings.NewReader(`{
		"tagname" : "holiday"
	}`))
	internalServerErrResp := httptest.NewRecorder()

	emptyTagnameReq := httptest.NewRequest("POST", "http://localhost:8000/tag", strings.NewReader(`{
		"tagname" : ""
	}`))
	emptyTagnameResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("POST", "http://localhost:8000/tag", strings.NewReader(""))
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
		fields func() TagHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				tagMock.EXPECT().CreateTag(gomock.Any(), "holiday").Return(services.CreateTagRow{
					ID:      1,
					Tagname: "holiday",
				}, nil)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test bad request",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test empty fullname",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   emptyTagnameResp,
				req: emptyTagnameReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				tagMock.EXPECT().CreateTag(gomock.Any(), "holiday").Return(services.CreateTagRow{}, errors.New("error"))

				return TagHandler{
					tagService: tagMock,
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
			field.CreateTag(tt.args.w, tt.args.req)
		})
	}
}

func Test_UpdateTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("PUT", "http://localhost:8000/tag?id=1", strings.NewReader(`{
		"tagname" : "holiday"
	}`))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("PUT", "http://localhost:8000/tag?id=1", strings.NewReader(`{
		"tagname" : "holiday"
	}`))
	internalServerErrResp := httptest.NewRecorder()

	idErrParseReq := httptest.NewRequest("POST", "http://localhost:8000/tag?id='error'", strings.NewReader(`{
		"tagname" : "holiday"
	}`))
	idErrParseResp := httptest.NewRecorder()

	idErrReq := httptest.NewRequest("PUT", "http://localhost:8000/tag", strings.NewReader(`{
		"tagname" : "holiday"
	}`))
	idErrResp := httptest.NewRecorder()

	emptyTagnameReq := httptest.NewRequest("PUT", "http://localhost:8000/tag?id=1", strings.NewReader(`{
		"tagname" : ""
	}`))
	emptyTagnameResp := httptest.NewRecorder()

	badReq := httptest.NewRequest("PUT", "http://localhost:8000/tag?id=1", strings.NewReader(""))
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
		fields func() TagHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				tagMock.EXPECT().UpdateTag(gomock.Any(), services.UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				}).Return(services.UpdateTagRow{
					ID:      1,
					Tagname: "holiday",
				}, nil)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test bad request",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   badResp,
				req: badReq,
			},
		},
		{
			name: "test empty fullname",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   emptyTagnameResp,
				req: emptyTagnameReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				tagMock.EXPECT().UpdateTag(gomock.Any(), services.UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				}).Return(services.UpdateTagRow{}, errors.New("error"))

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   internalServerErrResp,
				req: internalServerErrReq,
			},
		},
		{
			name: "test id not provided",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   idErrResp,
				req: idErrReq,
			},
		},
		{
			name: "test error parsed id",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				return TagHandler{
					tagService: tagMock,
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
			field.UpdateTag(tt.args.w, tt.args.req)
		})
	}
}

func Test_DeleteTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	sampleReq := httptest.NewRequest("DELETE", "http://localhost:8000/tag?id=1", strings.NewReader(``))
	sampleResp := httptest.NewRecorder()

	internalServerErrReq := httptest.NewRequest("DELETE", "http://localhost:8000/tag?id=1", strings.NewReader(``))
	internalServerErrResp := httptest.NewRecorder()

	idErrParseReq := httptest.NewRequest("DELETE", "http://localhost:8000/tag?id='error'", strings.NewReader(``))
	idErrParseResp := httptest.NewRecorder()

	idErrReq := httptest.NewRequest("DELETE", "http://localhost:8000/tag", strings.NewReader(``))
	idErrResp := httptest.NewRecorder()

	type fields struct {
		tagService TagService
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields func() TagHandler
		args   args
	}{
		{
			name: "test normal flow",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				tagMock.EXPECT().DeleteTag(gomock.Any(), int32(1)).Return(nil)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   sampleResp,
				req: sampleReq,
			},
		},
		{
			name: "test internal server error",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				tagMock.EXPECT().DeleteTag(gomock.Any(), int32(1)).Return(errors.New("error"))

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   internalServerErrResp,
				req: internalServerErrReq,
			},
		},
		{
			name: "test id not provided",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				return TagHandler{
					tagService: tagMock,
				}
			},
			args: args{
				w:   idErrResp,
				req: idErrReq,
			},
		},
		{
			name: "test error parsed id",
			fields: func() TagHandler {
				tagMock := NewMockTagService(ctrl)

				return TagHandler{
					tagService: tagMock,
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
			field.DeleteTag(tt.args.w, tt.args.req)
		})
	}
}
