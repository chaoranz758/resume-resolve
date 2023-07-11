package service

var (
	errCreatePostCategory              = "create post category failed"
	errUpdatePostCategory              = "update post category failed"
	errDeletePostCategory              = "delete post category failed"
	errCreatePostCategoryLevel1ToZSet  = "create post category level1 to zSet filed"
	errDeletePostCategoryFromRedis     = "delete post category from redis failed"
	errDeletePostCategoryLevel1ToZSet  = "delete post category level1 from zSet filed"
	errDeletePostCategoryLevel2ToZSet  = "delete post category level2 from zSet filed"
	errGetsPostCategoryLevel1FromRedis = "gets post category from redis failed"
	errGetsPostCategoryLevel1FromMysql = "gets post category from mysql failed"
	errSetPostCategoryNotFound         = "set post category not found in redis to redis failed"
	errJsonUnMarshal                   = "json unmarshal failed"
	errInputLevel                      = "input level wrong"
	errSetsPostCategoryLevel1ToRedis   = "set post category level1 to redis failed"
	errGetsPostCategoryFromMysql       = "gets post category information from mysql failed"
	errRpcService                      = "rpc service error"
	errRpcBizService                   = "rpc service biz error"
)
