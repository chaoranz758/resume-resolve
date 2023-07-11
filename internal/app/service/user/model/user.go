package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserId       int64        `db:"user_id"`
	DepartmentId int64        `db:"department_id"`
	Role         int8         `db:"role"`
	Username     string       `db:"username"`
	Password     string       `db:"password"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func (u User) TableName() string {
	return "user"
}
