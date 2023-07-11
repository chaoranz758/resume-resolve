package cache

import (
	"context"
	"time"
)

type Cache interface {
	Open() error
	Close() error
	Set(c context.Context, dataStructureName, key string, value []interface{}) (result bool, err error)
	Sets(c context.Context, dataStructureName string, key []string, value []interface{}, isExpire bool, expire ...time.Duration) (result bool, err error)
	Get(c context.Context, key string) (value string, err error)
	Gets(c context.Context, dataStructureName string, keys []string, value ...interface{}) (result []interface{}, err error)
	Delete(c context.Context, key []string, isSync, isDeleteAll bool, value ...interface{}) (result int64, err error)
}
