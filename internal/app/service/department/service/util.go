package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	department1 "resume-resolving/internal/app/service/department"
	pkg1 "resume-resolving/internal/app/service/pkg"
	"resume-resolving/internal/pkg"
	"strconv"
)

//GetsCacheWithNotFoundList 输入要查询的key-list，输出结果和在缓存中未查询到的key-list
func GetsCacheWithNotFoundList(idList []int64, keysString []string) ([]interface{}, []int64, error) {
	results, err := department1.GlobalEngine.Options.Cache.Gets(context.Background(), pkg.RedisDataStructureString, keysString)

	if err != nil {
		return nil, nil, err
	}

	idNotFoundList := make([]int64, 0, len(idList))
	for i := 0; i < len(results); i++ {
		if results[i] == nil {
			idNotFoundList = append(idNotFoundList, idList[i])
		}
	}

	return results, idNotFoundList, nil
}

func SetCacheNotFoundToRedis(data []interface{}, keys []string) (err error) {
	datas := make([]interface{}, 0, len(data))
	expireTimeList := pkg1.SetRandomExpireTime(pkg.Random, len(data), pkg.ExpireTime)

	for i := 0; i < len(data); i++ {
		byteData, err := json.Marshal(data[i])
		if err != nil {
			klog.Error(pkg1.ErrTypeConvert, err)
			return err
		}
		datas = append(datas, string(byteData))
	}

	if len(data) != 0 {
		isSets, err := department1.GlobalEngine.Options.Cache.Sets(context.Background(), pkg.RedisDataStructureString, keys, datas, true, expireTimeList...)
		if err != nil || isSets == false {
			return err
		}
	}

	return nil
}

func GetCacheZSet(key []string, min, max, offset, count interface{}) (resultList []int64, result string, err error) {
	result1, err := department1.GlobalEngine.Options.Cache.Gets(
		context.Background(),
		pkg.RedisDataStructureZSet, key,
		min,
		max,
		offset,
		count,
	)
	if err != nil {
		return
	}

	result, ok := result1[0].(string)
	if !ok {
		return nil, "0", errors.New(pkg1.ErrTypeConvert)
	}

	for i := 0; i < len(result1); i++ {
		result2, ok1 := result1[i].(string)
		if !ok1 {
			return nil, "0", errors.New(pkg1.ErrTypeConvert)
		}
		result3, err := strconv.ParseInt(result2, 10, 64)
		if err != nil {
			klog.Error(errors.New(pkg1.ErrTypeConvert), err)
			return nil, "0", err
		}
		resultList = append(resultList, result3)
	}

	return
}
