package model

import (
	"database/sql"
	"time"
)

type DepartmentCityMap struct {
	Id           int64        `db:"id"`
	CityId       int64        `db:"city_id"`
	DepartmentId int64        `db:"department_id"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func (d DepartmentCityMap) TableName() string {
	return "department_city_map"
}
