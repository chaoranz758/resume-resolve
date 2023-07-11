package model

import (
	"database/sql"
	"time"
)

type InternshipWorkExperience struct {
	Id           int64        `db:"id"`
	ResumeId     int64        `db:"resume_id"`
	Company      string       `db:"company"`
	Position     string       `db:"position"`
	Description  string       `db:"description"`
	IsInternship int8         `db:"is_internship"`
	StartTime    time.Time    `db:"start_time"`
	EndTime      time.Time    `db:"end_time"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func (internshipWorkExperience InternshipWorkExperience) TableName() string {
	return "internship_work_experience"
}
