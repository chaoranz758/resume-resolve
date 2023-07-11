package database

import (
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"resume-resolving/internal/app/service/post/config"
	"resume-resolving/internal/pkg"
	"resume-resolving/internal/pkg/database"
)

const (
	createLength = 1
	queryLength  = 6
	updateLength = 4
	deleteLength = 3
)

var (
	errNoDb        = "can not find db, for example, mysql..."
	errParamInput  = "input param wrong"
	errTypeConvert = "type convert failed"
	errNotFound    = "func not found"
)

type Gorm struct {
	db       *gorm.DB
	database database.Database
	config   *config.Config
}

func (g *Gorm) Open() error {
	dsn := g.database.SetConfig()
	switch g.database.GetName() {
	case pkg.DbNameMysql:
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		g.db = db
		if err != nil {
			return err
		}
	default:
		panic(errNoDb)
	}
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	//空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(g.config.ConfigInNacos.Mysql.MaxIdleConns)
	//打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(g.config.ConfigInNacos.Mysql.MaxOpenConns)
	return nil
}

func (g *Gorm) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	if err = sqlDB.Close(); err != nil {
		return err
	}
	return nil
}

func (g *Gorm) Create(v interface{}) (bool, error) {
	result := g.db.Create(v)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (g *Gorm) Query(limit, offset int, dest interface{}, orderValue interface{}, querySelect []string, queryWhere interface{}, args ...interface{}) (bool, error) {
	q, ok := queryWhere.(string)
	if !ok {
		return false, errors.New(errTypeConvert)
	}
	var result *gorm.DB
	if q == pkg.NotUseWhere {
		result = g.db.Select(querySelect).Order(orderValue).Limit(limit).Offset(offset).Find(dest)
	} else {
		result = g.db.Select(querySelect).Where(queryWhere, args...).Order(orderValue).Limit(limit).Offset(offset).Find(dest)
	}
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (g *Gorm) Update(valueModel interface{}, valueUpdates interface{}, query interface{}, args ...interface{}) (bool, error) {
	result := g.db.Model(valueModel).Where(query, args...).Updates(valueUpdates)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (g *Gorm) Delete(v interface{}, query interface{}, args ...interface{}) (bool, error) {
	result := g.db.Where(query, args...).Delete(v)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (g *Gorm) Transaction(fs []string, value [][]interface{}) (err error) {
	return g.db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(fs); i++ {
			switch fs[i] {
			case pkg.DbFunctionCreate:
				if len(value[i]) != createLength {
					klog.Error(errParamInput)
					return errors.New(errParamInput)
				}
				if err = tx.Create(value[i][0]).Error; err != nil {
					klog.Error(err)
					return err
				}
			case pkg.DbFunctionQuery:
				if len(value[i]) != queryLength {
					klog.Error(errParamInput)
					return errors.New(errParamInput)
				}
				querySelect, ok := value[i][4].([]string)
				if !ok {
					klog.Error(errTypeConvert)
					return errors.New(errTypeConvert)
				}
				q, ok := value[i][5].(string)
				if !ok {
					klog.Error(errTypeConvert)
					return errors.New(errTypeConvert)
				}
				limit, ok := value[i][0].(int)
				if !ok {
					klog.Error(errTypeConvert)
					return errors.New(errTypeConvert)
				}
				offset, ok := value[i][1].(int)
				if !ok {
					klog.Error(errTypeConvert)
					return errors.New(errTypeConvert)
				}
				if q == pkg.NotUseWhere {
					if err = tx.Select(querySelect).Order(value[i][3]).Limit(limit).Offset(offset).Find(value[i][2]).Error; err != nil {
						klog.Error(err)
						return err
					}
				} else {
					if err = tx.Select(querySelect).Where(value[i][5], value[i][6]).Order(value[i][3]).Limit(limit).Offset(offset).Find(value[i][2]).Error; err != nil {
						klog.Error(err)
						return err
					}
				}
			case pkg.DbFunctionUpdate:
				if len(value[i]) != updateLength {
					klog.Error(errParamInput)
					return errors.New(errParamInput)
				}
				if err = tx.Model(value[i][0]).Where(value[i][2], value[i][3]).Updates(value[i][1]).Error; err != nil {
					klog.Error(err)
					return err
				}
			case pkg.DbFunctionDelete:
				if len(value[i]) != deleteLength {
					klog.Error(errParamInput)
					return errors.New(errParamInput)
				}
				if err = tx.Where(value[i][1], value[i][2]).Delete(value[i][0]).Error; err != nil {
					klog.Error(err)
					return err
				}
			default:
				klog.Error(errNotFound)
				return errors.New(errNotFound)
			}
		}
		return nil
	})
}

func NewGorm(config *config.Config, database database.Database) database.Orm {
	return &Gorm{
		database: database,
		config:   config,
	}
}
