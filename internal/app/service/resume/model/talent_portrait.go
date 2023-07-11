package model

import (
	"database/sql"
	"time"
)

type TalentPortrait struct {
	Id               int64        `db:"id"`
	UserId           int64        `db:"user_id"`
	Age              int8         `db:"age"`
	SchoolLevel      int8         `db:"school_level"`
	WorkingSeniority int8         `db:"working_seniority"`
	MaxEducation     string       `db:"max_education"`
	GraduatedSchool  string       `db:"graduated_school"`
	CreatedAt        time.Time    `db:"created_at"`
	UpdatedAt        time.Time    `db:"updated_at"`
	DeletedAt        sql.NullTime `db:"deleted_at"`
}

func (t TalentPortrait) TableName() string {
	return "talent_portrait"
}
