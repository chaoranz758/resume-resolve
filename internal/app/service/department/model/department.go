package model

import (
	"database/sql"
	"time"
)

type Department struct {
	DepartmentId          int64        `db:"department_id"`
	DepartmentName        string       `db:"department_name"`
	DepartmentDescription string       `db:"department_description"`
	CreatedAt             time.Time    `db:"created_at"`
	UpdatedAt             time.Time    `db:"updated_at"`
	DeletedAt             sql.NullTime `db:"deleted_at"`
}

func (d Department) TableName() string {
	return "department"
}
