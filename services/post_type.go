package services

import "database/sql"

type Post struct {
	ID          int32               `json:"id"`
	Userid      int32               `json:"user_id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	CreatedAt   sql.NullTime        `json:"created_at"`
	UpdatedAt   sql.NullTime        `json:"updated_at"`
	DeletedAt   sql.NullTime        `json:"deleted_at"`
	Tags        []GetTagByPostIDRow `json:"tags"`
}

type CreatePostParams struct {
	Userid      int32
	Title       string
	Description string
	TagID       []int32
}

type CreatePostRow struct {
	ID          int32   `json:"id"`
	Userid      int32   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	TagID       []int32 `json:"tag_ids"`
}

type UpdatePostParams struct {
	ID          int32
	Title       string
	Description string
	TagID       []int32
}

type UpdatePostRow struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	TagID       []int32 `json:"tag_ids"`
}

type GetPostRow struct {
	ID          int32  `json:"id"`
	Userid      int32  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetPostsRow struct {
	ID          int32               `json:"id"`
	Userid      int32               `json:"user_id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Tags        []GetTagByPostIDRow `json:"tags"`
}
