package pkg

import "time"

const (
	MinInt64                 = 0
	MaxInt64                 = 9223372036854775806
	DbNameMysql              = "mysql"
	NotUseWhere              = "not use where"
	DbFunctionCreate         = "create"
	DbFunctionQuery          = "query"
	DbFunctionUpdate         = "update"
	DbFunctionDelete         = "delete"
	RedisDataStructureString = "string"
	RedisDataStructureZSet   = "zSet"
	StatusOK                 = "OK"
	Random                   = 3600
	ExpireTime               = 3 * time.Hour
	CacheZSetGet             = "get_zSet"
	CacheZSetSet             = "set_zSet"
	NilString                = ""
)
