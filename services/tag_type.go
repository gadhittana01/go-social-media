package services

import "database/sql"

type CreateTagRow struct {
	ID      int32  `json:"id"`
	Tagname string `json:"tagname"`
}

type GetTagByPostIDRow struct {
	ID      int32  `json:"id"`
	Tagname string `json:"tagname"`
}

type Tag struct {
	ID        int32  `json:"id"`
	Tagname   string `json:"tagname"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type UpdateTagParams struct {
	ID      int32
	Tagname string
}

type UpdateTagRow struct {
	ID      int32  `json:"id"`
	Tagname string `json:"tagname"`
}

type GetTagsRow struct {
	ID      int32  `json:"id"`
	Tagname string `json:"tagname"`
}
