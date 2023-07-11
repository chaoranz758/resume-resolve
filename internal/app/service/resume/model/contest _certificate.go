package model

import (
	"database/sql"
	"time"
)

type ContestCertificate struct {
	Id          int64        `db:"id"`
	ResumeId    int64        `db:"resume_id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	IsContest   int8         `db:"is_contest"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

func (contestCertificate ContestCertificate) TableName() string {
	return "contest_certificate"
}
