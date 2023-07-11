package service

import (
	"context"
	"resume-resolving/internal/app/service/post_category"
	"resume-resolving/internal/pkg"
)

//GetsCacheWithNotFoundList 输入要查询的key-list，输出结果和在缓存中未查询到的key-list
func GetsCacheWithNotFoundList(idList []int64, keysString []string) ([]interface{}, []int64, error) {
	results, err := post_category.GlobalEngine.Options.Cache.Gets(context.Background(), pkg.RedisDataStructureString, keysString)

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
