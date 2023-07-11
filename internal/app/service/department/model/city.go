package model

import (
	"database/sql"
	"time"
)

type City struct {
	CityId    int64        `db:"city_id"`
	CityName  string       `db:"city_name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (c City) TableName() string {
	return "city"
}
