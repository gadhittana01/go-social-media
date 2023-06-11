package post

import (
	"context"
	"database/sql"
	"errors"
	reflect "reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	gomock "github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := NewMockDBTX(ctrl)

	type args struct {
		db DBTX
	}
	tests := []struct {
		name    string
		args    args
		want    *Queries
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				db: dbMock,
			},
			want: &Queries{
				db: dbMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_WithTx(t *testing.T) {
	txMock := sql.Tx{}

	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name     string
		args     args
		initMock func() *Queries
		want     *Queries
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				tx: &txMock,
			},
			initMock: func() *Queries {
				return &Queries{
					db: &txMock,
				}
			},
			want: &Queries{
				db: &txMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			if got := p.WithTx(tt.args.tx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CreatePost(t *testing.T) {
	type args struct {
		ctx context.Context
		arg CreatePostParams
	}

	q := `-- name: CreatePost :one
	INSERT INTO posts (
	  userid, title, description
	) VALUES (
	  $1,$2,$3
	)
	RETURNING id, userid, title, description
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     CreatePostRow
		wantErr  bool
	}{
		{
			name: "success create post",
			args: args{
				ctx: context.Background(),
				arg: CreatePostParams{
					Userid:      1,
					Title:       "Holiday in maldives",
					Description: "Yeah yeah yeah",
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "description"}).AddRow(1, 1, "Holiday in maldives", "Yeah yeah yeah")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, "Holiday in maldives", "Yeah yeah yeah").WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: CreatePostRow{
				ID:          1,
				Userid:      1,
				Title:       "Holiday in maldives",
				Description: "Yeah yeah yeah",
			},
			wantErr: false,
		},
		{
			name: "error create post",
			args: args{
				ctx: context.Background(),
				arg: CreatePostParams{
					Userid:      1,
					Title:       "Holiday in maldives",
					Description: "Yeah yeah yeah",
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, "Holiday in maldives", "Yeah yeah yeah").WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    CreatePostRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
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

func Test_DeletePost(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int32
	}

	q := `-- name: DeletePost :exec
	DELETE FROM posts
	WHERE id = $1
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		wantErr  bool
	}{
		{
			name: "success delete post",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

				return &Queries{
					db: dbMock,
				}
			},
			wantErr: false,
		},
		{
			name: "error delete post",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(1).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			err := p.DeletePost(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_GetPost(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int32
	}

	q := `-- name: GetPost :one
		SELECT id, userid, title, description FROM posts
		WHERE id = $1 LIMIT 1
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     GetPostRow
		wantErr  bool
	}{
		{
			name: "success get post",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "description"}).AddRow(1, 1, "holiday yay", "yeah yeah yeah")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: GetPostRow{
				ID:          1,
				Userid:      1,
				Title:       "holiday yay",
				Description: "yeah yeah yeah",
			},
			wantErr: false,
		},
		{
			name: "error get post",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    GetPostRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.GetPost(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetPosts(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	q := `-- name: GetPosts :many
	SELECT id, userid, title, description FROM posts
	ORDER BY created_at DESC
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     []GetPostsRow
		wantErr  bool
	}{
		{
			name: "success get posts",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "description"}).AddRow(1, 1, "holiday yay", "yeah yeah yeah")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: []GetPostsRow{
				{
					ID:          1,
					Userid:      1,
					Title:       "holiday yay",
					Description: "yeah yeah yeah",
				},
			},
			wantErr: false,
		},
		{
			name: "error scan get tags",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "description"}).AddRow("error", 1, "holiday yay", "yeah yeah yeah")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error get tags",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
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
	type args struct {
		ctx context.Context
		arg UpdatePostParams
	}

	q := `-- name: UpdatePost :one
		UPDATE posts
		set title = $2,
		description = $3
		WHERE id = $1
		RETURNING title, description
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     UpdatePostRow
		wantErr  bool
	}{
		{
			name: "success update tag",
			args: args{
				ctx: context.Background(),
				arg: UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay",
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"title", "description"}).AddRow("holiday yay", "yay yay")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, "holiday yay", "yay yay").WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: UpdatePostRow{
				Title:       "holiday yay",
				Description: "yay yay",
			},
			wantErr: false,
		},
		{
			name: "error update post",
			args: args{
				ctx: context.Background(),
				arg: UpdatePostParams{
					ID:          1,
					Title:       "holiday yay",
					Description: "yay yay",
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, "holiday").WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    UpdatePostRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
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
