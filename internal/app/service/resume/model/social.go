package model

import (
	"database/sql"
	"time"
)

type Social struct {
	Id             int64        `db:"id"`
	ResumeId       int64        `db:"resume_id"`
	SocialPlatform string       `db:"social_platform"`
	PlatformUrl    string       `db:"platform_url"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      time.Time    `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
}

func (s Social) TableName() string {
	return "social"
}
