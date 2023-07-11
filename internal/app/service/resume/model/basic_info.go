package model

import (
	"database/sql"
	"time"
)

type BasicInfo struct {
	ResumeId       int64        `db:"resume_id"`
	UserId         int64        `db:"user_id"`
	Name           string       `db:"name"`
	Phone          string       `db:"phone"`
	ResumeUrl      string       `db:"resume_url"`
	Email          string       `db:"email"`
	SelfEvaluation string       `db:"self_evaluation"`
	Birthday       time.Time    `db:"birthday"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      time.Time    `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
}

func (basicInfo BasicInfo) TableName() string {
	return "basic_info"
}
