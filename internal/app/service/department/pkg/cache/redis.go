package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"resume-resolving/internal/app/service/department/config"
	pkg1 "resume-resolving/internal/app/service/pkg"
	"resume-resolving/internal/pkg"
	"resume-resolving/internal/pkg/cache"
	"time"
)

const (
	cancelTime         = 5 * time.Second
	stringSetLength    = 2
	zSetSetLength      = 2
	lengthThreshold    = 500
	setZSetFileName    = "../../../internal/app/service/pkg/lua/set_zset.lua"
	deleteZSetFileName = "../../../internal/app/service/pkg/lua/delete_zset.lua"
	getsZSetFileName   = "../../../internal/app/service/pkg/lua/gets_zset.lua"
	getZSetLength      = 5
	expireTime         = 3 * time.Hour
	random             = 3600
)

var (
	ctx            = context.Background()
	errParamInput  = "input param wrong"
	errTypeConvert = "type convert failed"
	errNotFound    = "func not found"
)

type Redis struct {
	rdb    *redis.Client
	config *config.Config
}

func (r *Redis) Open() error {
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.config.ConfigInNacos.Redis.Host, r.config.ConfigInNacos.Redis.Port),
		Password: r.config.ConfigInNacos.Redis.Password,
		DB:       r.config.ConfigInNacos.Redis.Db,
		PoolSize: r.config.ConfigInNacos.Redis.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), cancelTime)
	defer cancel()
	_, err := r.rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Close() error {
	err := r.rdb.Close()
	if err != nil {
		return err
	}
	return nil
}

// Set
//  Include two situations: 1. set one object to string; 2. set one object to zSet.
//  If dataStructureName = string, len(value)=2 and value[0]=value(interface{}) value[1]=time.Duration in set function.
//  If dataStructureName = zSet, it means to add the member with score to zSet by command 'ZADD' len(value)=2 and value[0]=member(interface{}) value[1]=score(float64).
//  - - - - - -
//  For example, if dataStructureName = string, set(context.Background(), "key", []{"value", 1*time.Hour})
//  if dataStructureName = zSet, set(context.Background(), "key", []{"member", 10.})
func (r *Redis) Set(c context.Context, dataStructureName, key string, value []interface{}) (result bool, err error) {
	switch dataStructureName {
	case pkg.RedisDataStructureString:
		if len(value) != stringSetLength {
			klog.Error(errParamInput)
			return false, errors.New(errParamInput)
		}
		expiration, ok := value[1].(time.Duration)
		if !ok {
			klog.Error(errTypeConvert)
			return false, errors.New(errTypeConvert)
		}
		status, err := r.rdb.Set(ctx, key, value[0], expiration).Result()
		if status == pkg.StatusOK {
			result = true
		}
		return result, err
	case pkg.RedisDataStructureZSet:
		if len(value) != zSetSetLength {
			klog.Error(errParamInput)
			return false, errors.New(errParamInput)
		}
		score, ok := value[1].(float64)
		if !ok {
			klog.Error(errTypeConvert)
			return false, errors.New(errTypeConvert)
		}

		file, err := ioutil.ReadFile(setZSetFileName)
		if err != nil {
			klog.Error(err)
			return false, err
		}

		expire := pkg1.SetRandomExpireTime(random, 1, expireTime)
		script := redis.NewScript(string(file))
		script.Run(ctx, r.rdb, []string{key}, lengthThreshold, score, value[0], expire[0])
		
		return true, nil
	default:
		klog.Error(errNotFound)
		return false, errors.New(errNotFound)
	}
}

// Sets
//
func (r *Redis) Sets(c context.Context, dataStructureName string, key []string, value []interface{}, isExpire bool, expire ...time.Duration) (result bool, err error) {
	switch dataStructureName {
	case pkg.RedisDataStructureString:
		if len(key) != len(value) {
			klog.Error(errParamInput)
			return false, errors.New(errParamInput)
		}
		valueMSet := make([]interface{}, 0, 2*len(key))
		for i := 0; i < len(key); i++ {
			valueMSet = append(valueMSet, key[i])
			valueMSet = append(valueMSet, value[i])
		}
		pipeline := r.rdb.Pipeline()
		pipeline.MSet(ctx, valueMSet)
		for i := 0; i < len(key); i++ {
			pipeline.Expire(ctx, key[i], expire[i])
		}
		_, err = pipeline.Exec(ctx)
		if err != nil {
			klog.Error(err)
			return false, err
		}
		return true, nil
	case pkg.RedisDataStructureZSet:
		if len(value)%2 != 0 {
			klog.Error(errParamInput)
			return false, errors.New(errParamInput)
		}
		var zSetValue []*redis.Z
		for i := 0; i < len(value)/2; i++ {
			score, ok := value[(2*i)+1].(float64)
			if !ok {
				klog.Error(errTypeConvert)
				return false, errors.New(errTypeConvert)
			}
			zSetValue = append(zSetValue, &redis.Z{
				Member: value[2*i],
				Score:  score,
			})
		}
		err = r.rdb.ZAdd(ctx, key[0], zSetValue...).Err()
		if err != nil {
			klog.Error(err)
			return false, err
		}
		if isExpire == true {
			r.rdb.Expire(ctx, key[0], expire[0])
		}
		return true, nil
	default:
		klog.Error(errNotFound)
		return false, errors.New(errNotFound)
	}
}

//Get
//Read data from string by get command
func (r *Redis) Get(c context.Context, key string) (value string, err error) {
	return r.rdb.Get(ctx, key).Result()
}

func (r *Redis) Gets(c context.Context, dataStructureName string, keys []string, value ...interface{}) (result []interface{}, err error) {
	switch dataStructureName {
	case pkg.RedisDataStructureString:
		result, err = r.rdb.MGet(ctx, keys...).Result()
		if err != nil {
			klog.Error(errTypeConvert)
			return nil, err
		}
	case pkg.RedisDataStructureZSet:
		min, ok := value[0].(string)
		if !ok {
			klog.Error(errTypeConvert)
			return nil, errors.New(errTypeConvert)
		}
		max, ok := value[1].(string)
		if !ok {
			klog.Error(errTypeConvert)
			return nil, errors.New(errTypeConvert)
		}
		offset, ok := value[2].(int)
		if !ok {
			klog.Error(errTypeConvert)
			return nil, errors.New(errTypeConvert)
		}
		count, ok := value[3].(int)
		if !ok {
			klog.Error(errTypeConvert)
			return nil, errors.New(errTypeConvert)
		}

		file, err := ioutil.ReadFile(getsZSetFileName)
		if err != nil {
			klog.Error(err)
			return nil, err
		}

		isLimit := 1
		if offset == 0 && count == 0 {
			isLimit = 0
		}

		script := redis.NewScript(string(file))
		result, err = script.Run(ctx, r.rdb, []string{keys[0]}, min, max, offset, count, isLimit).Slice()
		if err != nil {
			klog.Error(err)
			return nil, err
		}
	default:
		klog.Error(errNotFound)
		return nil, errors.New(errNotFound)
	}
	return
}

//Delete
//In terms of deleting data structure,  include two situations:
//1. delete all data structure, like string and zSet - delete command;
//2. delete one member by ZRem command.
//In terms of how to handle server deletion, include two situations: 1. sync - delete command 2. async - unlink command.
func (r *Redis) Delete(c context.Context, key []string, isSync, isDeleteAll bool, value ...interface{}) (result int64, err error) {
	if isDeleteAll == true {
		if isSync == true {
			result, err = r.rdb.Del(ctx, key...).Result()
		} else {
			result, err = r.rdb.Unlink(ctx, key...).Result()
		}
	} else {
		file, err := ioutil.ReadFile(deleteZSetFileName)
		if err != nil {
			klog.Error(err)
			return 0, err
		}
		script := redis.NewScript(string(file))
		script.Run(ctx, r.rdb, []string{key[0]}, value[0])
		result = 1
	}
	return
}

func NewRedis(config2 *config.Config) cache.Cache {
	return &Redis{
		config: config2,
	}
}
