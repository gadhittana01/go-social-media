package user

import (
	"context"
	"database/sql"
	"errors"
	reflect "reflect"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
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

func Test_CreateUser(t *testing.T) {
	type args struct {
		ctx      context.Context
		fullname string
	}

	q := `-- name: CreateUser :one
		INSERT INTO users (
		fullname
		) VALUES (
		$1
		)
		RETURNING id, fullname
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     CreateUserRow
		wantErr  bool
	}{
		{
			name: "success create user",
			args: args{
				ctx:      context.Background(),
				fullname: "Giri Putra Adhittana",
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "fullname"}).AddRow(1, "Giri Putra Adhittana")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("Giri Putra Adhittana").WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: CreateUserRow{
				ID:       1,
				Fullname: "Giri Putra Adhittana",
			},
			wantErr: false,
		},
		{
			name: "error create user",
			args: args{
				ctx:      context.Background(),
				fullname: "Giri Putra Adhittana",
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs("Giri Putra Adhittana").WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    CreateUserRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.CreateUser(tt.args.ctx, tt.args.fullname)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DeleteUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int32
	}

	q := `-- name: DeleteUser :exec
		DELETE FROM users
		WHERE id = $1
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		wantErr  bool
	}{
		{
			name: "success delete user",
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
			name: "error delete user",
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
			err := p.DeleteUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_GetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int32
	}

	q := `-- name: GetUser :one
		SELECT id, fullname FROM users
		WHERE id = $1 LIMIT 1
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     GetUserRow
		wantErr  bool
	}{
		{
			name: "success get user",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "fullname"}).AddRow(1, "Giri Putra Adhittana")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: GetUserRow{
				ID:       1,
				Fullname: "Giri Putra Adhittana",
			},
			wantErr: false,
		},
		{
			name: "error get user",
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
			want:    GetUserRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.GetUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	q := `-- name: GetUsers :many
		SELECT id, fullname FROM users
		ORDER BY created_at DESC
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     []GetUsersRow
		wantErr  bool
	}{
		{
			name: "success get users",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "fullname"}).AddRow(1, "Giri Putra Adhittana").AddRow(2, "Giri Adhittana")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: []GetUsersRow{
				{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				},
				{
					ID:       2,
					Fullname: "Giri Adhittana",
				},
			},
			wantErr: false,
		},
		{
			name: "error get users",
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
		{
			name: "error scan get users",
			args: args{
				ctx: context.Background(),
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "fullname"}).AddRow("Giri Putra Adhittana", 1)
				mock.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)

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
			got, err := p.GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		arg UpdateUserParams
	}

	q := `-- name: UpdateUser :one
		UPDATE users
		set fullname = $2
		WHERE id = $1
		RETURNING id, fullname
	`
	tests := []struct {
		name     string
		initMock func() *Queries
		args     args
		want     UpdateUserRow
		wantErr  bool
	}{
		{
			name: "success update user",
			args: args{
				ctx: context.Background(),
				arg: UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				rows := sqlmock.NewRows([]string{"id", "fullname"}).AddRow(1, "Giri Putra Adhittana")
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, "Giri Putra Adhittana").WillReturnRows(rows)

				return &Queries{
					db: dbMock,
				}
			},
			want: UpdateUserRow{
				ID:       1,
				Fullname: "Giri Putra Adhittana",
			},
			wantErr: false,
		},
		{
			name: "error update user",
			args: args{
				ctx: context.Background(),
				arg: UpdateUserParams{
					ID:       1,
					Fullname: "Giri Putra Adhittana",
				},
			},
			initMock: func() *Queries {
				dbMock, mock, _ := sqlmock.New()
				mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(1, "Giri Putra Adhittana").WillReturnError(errors.New("error"))

				return &Queries{
					db: dbMock,
				}
			},
			want:    UpdateUserRow{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.initMock()
			got, err := p.UpdateUser(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
