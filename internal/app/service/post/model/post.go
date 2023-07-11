package model

import (
	"database/sql"
	"time"
)

type Post struct {
	PostId              int64        `db:"post_id"`
	HRId                int64        `db:"hr_id"`
	PostCategoryId      int64        `db:"post_category_id"`
	DepartmentId        int64        `db:"department_id"`
	IsSchoolRecruitment int8         `db:"is_school_recruitment"`
	IsInternship        int8         `db:"is_internship"`
	PostBrief           string       `db:"post_brief"`
	PostDescription     string       `db:"post_description"`
	PostRequire         string       `db:"post_require"`
	CreatedAt           time.Time    `db:"created_at"`
	UpdatedAt           time.Time    `db:"updated_at"`
	DeletedAt           sql.NullTime `db:"deleted_at"`
}

func (p Post) TableName() string {
	return "post"
}
