package model

import (
	"database/sql"
	"time"
)

type PostCityMap struct {
	Id        int64        `db:"id"`
	CityId    int64        `db:"city_id"`
	PostId    int64        `db:"post_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (p PostCityMap) TableName() string {
	return "post_city_map"
}
