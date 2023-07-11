package database

type Orm interface {
	Open() error
	Close() error
	Create(v interface{}) (bool, error)
	Query(limit, offset int, dest interface{}, orderValue interface{}, querySelect []string, query interface{}, args ...interface{}) (bool, error)
	Update(valueModel interface{}, valueUpdates interface{}, query interface{}, args ...interface{}) (bool, error)
	Delete(v interface{}, query interface{}, args ...interface{}) (bool, error)
	Transaction(fs []string, value [][]interface{}) (err error)
}
