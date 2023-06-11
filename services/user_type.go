package services

import "database/sql"

type CreateUserRow struct {
	ID       int32  `json:"id"`
	Fullname string `json:"fullname"`
}

type User struct {
	ID        int32        `json:"id"`
	Fullname  string       `json:"fullname"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type UpdateUserParams struct {
	ID       int32
	Fullname string
}

type UpdateUserRow struct {
	ID       int32  `json:"id"`
	Fullname string `json:"fullname"`
}

type GetUsersRow struct {
	ID       int32  `json:"id"`
	Fullname string `json:"fullname"`
}

type GetUserRow struct {
	ID       int32
	Fullname string
}
