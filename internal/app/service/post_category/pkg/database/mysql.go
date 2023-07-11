package database

import (
	"fmt"
	"resume-resolving/internal/app/service/post_category/config"
	"resume-resolving/internal/pkg"
	"resume-resolving/internal/pkg/database"
)

type Mysql struct {
	config *config.Config
}

func (mysql *Mysql) GetName() string {
	return pkg.DbNameMysql
}

func (mysql *Mysql) SetConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysql.config.ConfigInNacos.Mysql.User,
		mysql.config.ConfigInNacos.Mysql.Password,
		mysql.config.ConfigInNacos.Mysql.Host,
		mysql.config.ConfigInNacos.Mysql.Port,
		mysql.config.ConfigInNacos.Mysql.Dbname,
	)
}

func NewMysql(config *config.Config) database.Database {
	return &Mysql{
		config: config,
	}
}
