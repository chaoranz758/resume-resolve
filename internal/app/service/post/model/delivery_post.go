package model

import (
	"database/sql"
	"time"
)

type DeliveryPost struct {
	Id           int64        `db:"id"`
	UserId       int64        `db:"user_id"`
	PostId       int64        `db:"post_id"`
	ResumeStatus int8         `db:"resume_status"`
	IsTalentPool int8         `db:"is_talent_pool"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func (d DeliveryPost) TableName() string {
	return "delivery_post"
}
