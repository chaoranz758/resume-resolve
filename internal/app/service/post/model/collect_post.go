package model

import (
	"database/sql"
	"time"
)

type CollectPost struct {
	Id        int64        `db:"id"`
	UserId    int64        `db:"user_id"`
	PostId    int64        `db:"post_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (c CollectPost) TableName() string {
	return "collect_post"
}
