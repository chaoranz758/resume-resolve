package model

import (
	"database/sql"
	"time"
)

type ProjectExperience struct {
	Id                 int64        `db:"id"`
	ResumeId           int64        `db:"resume_id"`
	ProjectName        string       `db:"project_name"`
	ProjectRole        string       `db:"project_role"`
	ProjectDescription string       `db:"project_description"`
	ProjectUrl         string       `db:"project_url"`
	StartTime          time.Time    `db:"start_time"`
	EndTime            time.Time    `db:"end_time"`
	CreatedAt          time.Time    `db:"created_at"`
	UpdatedAt          time.Time    `db:"updated_at"`
	DeletedAt          sql.NullTime `db:"deleted_at"`
}

func (projectExperience ProjectExperience) TableName() string {
	return "project_experience"
}
