package post_tags

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

func Test_CreatePostTag(t *testing.T) {
	type args struct {
		ctx context.Context
		arg CreatePostTagParams
	}

	q := `-- name: CreatePostTag :one
	INSERT INTO post_tags (
	  postid, tagid
	) VALUES (
	  $1,$2
	)
	RETURNING id, postid, tagid
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     CreatePostTagRow
		wantErr  bool
	}{
		{
			name: "success create post tag",
			args: args{
				ctx: context.Background(),
				arg: CreatePostTagParams{
					Postid: 1,
					Tagid:  1,
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "postid", "tagid"}).AddRow(1, 1, 1)
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, 1).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: CreatePostTagRow{
				ID:     1,
				Postid: 1,
				Tagid:  1,
			},
			wantErr: false,
		},
		{
			name: "error create post tag",
			args: args{
				ctx: context.Background(),
				arg: CreatePostTagParams{
					Postid: 1,
					Tagid:  1,
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, 1).WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    CreatePostTagRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.CreatePostTag(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePostTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePostTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DeleteTag(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int32
	}

	q := `-- name: DeletePostTag :exec
		DELETE FROM post_tags
		WHERE postid = $1
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		wantErr  bool
	}{
		{
			name: "success delete post tag",
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
			name: "error delete post tag",
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
			err := p.DeletePostTag(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePostTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
