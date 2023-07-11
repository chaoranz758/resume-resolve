package model

import (
	"database/sql"
	"time"
)

type PostCategory struct {
	PostCategoryId       int64        `db:"post_category_id"`
	PostCategoryParentId int64        `db:"post_category_parent_id"`
	PostCategoryLevel    int8         `db:"post_category_level"`
	PostCategoryName     string       `db:"post_category_name"`
	CreatedAt            time.Time    `db:"created_at"`
	UpdatedAt            time.Time    `db:"updated_at"`
	DeletedAt            sql.NullTime `db:"deleted_at"`
}

func (postCategory PostCategory) TableName() string {
	return "post_category"
}
