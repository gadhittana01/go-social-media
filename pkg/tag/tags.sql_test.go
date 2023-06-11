package tag

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

func Test_CreateTag(t *testing.T) {
	type args struct {
		ctx     context.Context
		tagname string
	}

	q := `-- name: CreateTag :one
		INSERT INTO tags (
		tagname
		) VALUES (
		$1
		)
		RETURNING id, tagname
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     CreateTagRow
		wantErr  bool
	}{
		{
			name: "success create tag",
			args: args{
				ctx:     context.Background(),
				tagname: "holiday",
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "tagname"}).AddRow(1, "holiday")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("holiday").WillReturnRows(rows)

				return &Queries{
					db: dbMock,
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
				ctx:     context.Background(),
				tagname: "holiday",
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("holiday").WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    CreateTagRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
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

func Test_DeleteTag(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int32
	}

	q := `-- name: DeleteTag :exec
		DELETE FROM tags
		WHERE id = $1
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		wantErr  bool
	}{
		{
			name: "success delete tag",
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
			name: "error delete tag",
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
			err := p.DeleteTag(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_GetTag(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int32
	}

	q := `-- name: GetTag :one
		SELECT id, tagname FROM tags
		WHERE id = $1 LIMIT 1
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     GetTagRow
		wantErr  bool
	}{
		{
			name: "success get tag",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "tagname"}).AddRow(1, "holiday")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: GetTagRow{
				ID:      1,
				Tagname: "holiday",
			},
			wantErr: false,
		},
		{
			name: "error get tag",
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
			want:    GetTagRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.GetTag(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetTagByPostID(t *testing.T) {
	type args struct {
		ctx    context.Context
		postid int32
	}

	q := `-- name: GetTagByPostID :many
		SELECT 
			b.id,
			b.tagname
		FROM post_tags a JOIN tags b
		ON a.tagID = b.id
		WHERE a.postid = $1
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     []GetTagByPostIDRow
		wantErr  bool
	}{
		{
			name: "success get tag by post id",
			args: args{
				ctx:    context.Background(),
				postid: 1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "tagname"}).AddRow(1, "holiday").AddRow(2, "reading")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: []GetTagByPostIDRow{
				{
					ID:      1,
					Tagname: "holiday",
				},
				{
					ID:      2,
					Tagname: "reading",
				},
			},
			wantErr: false,
		},
		{
			name: "error scan get tag by post id",
			args: args{
				ctx:    context.Background(),
				postid: 1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "tagname"}).AddRow("holiday", 1)
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error get tag by post id",
			args: args{
				ctx:    context.Background(),
				postid: 1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnError(errors.New("error"))

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
			got, err := p.GetTagByPostID(tt.args.ctx, tt.args.postid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagByPostID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTagByPostID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetTags(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	q := `-- name: GetTags :many
		SELECT id, tagname FROM tags
	`

	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     []GetTagsRow
		wantErr  bool
	}{
		{
			name: "success get tags",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "tagname"}).AddRow(1, "holiday").AddRow(2, "reading")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: []GetTagsRow{
				{
					ID:      1,
					Tagname: "holiday",
				},
				{
					ID:      2,
					Tagname: "reading",
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
				rows := sqlmock.NewRows([]string{"id", "tagname"}).AddRow("holiday", 1)
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

func Test_UpdateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		arg UpdateTagParams
	}

	q := `-- name: UpdateTag :one
		UPDATE tags
		set tagname = $2
		WHERE id = $1
		RETURNING id, tagname
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     UpdateTagRow
		wantErr  bool
	}{
		{
			name: "success update tag",
			args: args{
				ctx: context.Background(),
				arg: UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "fullname"}).AddRow(1, "holiday")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, "holiday").WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: UpdateTagRow{
				ID:      1,
				Tagname: "holiday",
			},
			wantErr: false,
		},
		{
			name: "error update tag",
			args: args{
				ctx: context.Background(),
				arg: UpdateTagParams{
					ID:      1,
					Tagname: "holiday",
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, "holiday").WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    UpdateTagRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.UpdateTag(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
