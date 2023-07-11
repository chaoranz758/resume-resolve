package model

import (
	"database/sql"
	"time"
)

type Language struct {
	Id               int64        `db:"id"`
	ResumeId         int64        `db:"resume_id"`
	LanguageName     string       `db:"language_name"`
	ProficiencyLevel string       `db:"proficiency_level"`
	CreatedAt        time.Time    `db:"created_at"`
	UpdatedAt        time.Time    `db:"updated_at"`
	DeletedAt        sql.NullTime `db:"deleted_at"`
}

func (l Language) TableName() string {
	return "language"
}
