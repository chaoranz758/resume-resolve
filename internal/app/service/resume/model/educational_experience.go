package model

import (
	"database/sql"
	"time"
)

type EducationalExperience struct {
	Id         int64        `db:"id"`
	ResumeId   int64        `db:"resume_id"`
	School     string       `db:"school"`
	Education  string       `db:"education"`
	Speciality string       `db:"speciality"`
	Ranking    string       `db:"ranking"`
	StartTime  time.Time    `db:"start_time"`
	EndTime    time.Time    `db:"end_time"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
}

func (educationalExperience EducationalExperience) TableName() string {
	return "educational_experience"
}
