package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redsync/redsync/v4"
	"resume-resolving/api/idl/service/post/kitex_gen/post"
	"resume-resolving/api/idl/service/post_category/kitex_gen/post_category"
	pkg1 "resume-resolving/internal/app/service/pkg"
	post_category1 "resume-resolving/internal/app/service/post_category"
	"resume-resolving/internal/app/service/post_category/model"
	"resume-resolving/internal/app/service/post_category/pkg/code"
	"resume-resolving/internal/pkg"
	"strconv"
	"time"
)

const (
	level1 = iota
	level2
)

const (
	postCategoryLevel1 = "post_category_level1"
	postCategoryLevel2 = "post_category_level2"
	expireTime         = 3 * time.Hour
	maxLength          = 20
	random             = 3600
)

func AppendPostCategory(request *post_category.AppendPostCategoryRPCRequest) (int, string, *post_category.PostCategoryInformation, error) {
	createTime := time.Now()

	postCategoryId := post_category1.GlobalEngine.Options.Id.GenId()
	var postCategoryData = model.PostCategory{
		PostCategoryId:       postCategoryId,
		PostCategoryName:     request.PostCategoryName,
		PostCategoryParentId: request.PostCategoryParentId,
		PostCategoryLevel:    request.PostCategoryLevel,
		CreatedAt:            createTime,
	}
	result, err := post_category1.GlobalEngine.Options.Orm.Create(&postCategoryData)
	if err != nil || result == false {
		klog.Error(errCreatePostCategory, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	switch request.PostCategoryLevel {
	case level1:
		key := GetKey(KeyLevel1PostCategoryZSet)
		if err = appendPostCategory(createTime, postCategoryLevel1, key, postCategoryId); err != nil {
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}
	case level2:
		mutexKey := postCategoryLevel2 + strconv.FormatInt(request.PostCategoryParentId, 10)
		key := GetKey(KeyLevel2PostCategoryZSet + strconv.FormatInt(request.PostCategoryParentId, 10))
		if err = appendPostCategory(createTime, mutexKey, key, postCategoryId); err != nil {
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}
	}

	var data = post_category.PostCategoryInformation{
		PostCategoryId:   postCategoryId,
		PostCategoryName: request.PostCategoryName,
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), &data, nil
}

func appendPostCategory(createTime time.Time, mutexKey, key string, postCategoryId int64) (err error) {
	createTime1 := createTime.UnixMilli()
	mutex := post_category1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
	defer post_category1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

	mutex = post_category1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(mutexKey)

	_ = mutex.Lock()

	_, err = post_category1.GlobalEngine.Options.Cache.Set(context.Background(), pkg.RedisDataStructureZSet, key, []interface{}{postCategoryId, float64(createTime1)})
	if err != nil {
		_, _ = mutex.Unlock()
		klog.Error(errCreatePostCategoryLevel1ToZSet, err)
		_, _ = post_category1.GlobalEngine.Options.Orm.Delete(&model.PostCategory{}, "post_category_id = ?", postCategoryId)
		return err
	}
	_, _ = mutex.Unlock()
	return nil
}

func UpdatePostCategory(request *post_category.UpdatePostCategoryRPCRequest) (int, string, *post_category.PostCategoryInformation, error) {
	updateMap := map[string]interface{}{
		"post_category_name": request.PostCategoryName,
	}

	_, err := post_category1.GlobalEngine.Options.Orm.Update(&model.PostCategory{}, updateMap, "post_category_id = ?", request.PostCategoryId)
	if err != nil {
		klog.Error(errUpdatePostCategory, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	key := GetKey(strconv.FormatInt(request.PostCategoryId, 10))
	_, err = post_category1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, true)

	if err != nil {
		klog.Error(errDeletePostCategoryFromRedis, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	var data = post_category.PostCategoryInformation{
		PostCategoryId:   request.PostCategoryId,
		PostCategoryName: request.PostCategoryName,
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), &data, nil
}

func DeletePostCategory(request *post_category.DeletePostCategoryRPCRequest) (int, string, error) {
	postCategoryIdList := make([]int64, 0, maxLength)
	switch request.Level {
	case level1:
		postCategories := make([]model.PostCategory, 0, maxLength)
		_, err := post_category1.GlobalEngine.Options.Orm.Query(-1, -1, &postCategories, "", []string{"post_category_id"},
			"post_category_parent_id = ?", request.PostCategoryId)
		if err != nil {
			klog.Error(errDeletePostCategory, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}
		for i := 0; i < len(postCategories); i++ {
			postCategoryIdList = append(postCategoryIdList, postCategories[i].PostCategoryId)
		}

		_, err = post_category1.GlobalEngine.Options.Orm.Delete(&model.PostCategory{}, "post_category_id = ? or "+
			"post_category_parent_id = ?", request.PostCategoryId, request.PostCategoryId)
		if err != nil {
			klog.Error(errDeletePostCategory, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}

		mutex := post_category1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
		defer post_category1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

		mutex = post_category1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(postCategoryLevel1)

		_ = mutex.Lock()
		key := GetKey(KeyLevel1PostCategoryZSet)
		_, err = post_category1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, false, request.PostCategoryId)
		if err != nil {
			_, _ = mutex.Unlock()
			klog.Error(errDeletePostCategoryLevel1ToZSet, err)
			updateMap := map[string]interface{}{
				"deleted_at": nil,
			}
			_, _ = post_category1.GlobalEngine.Options.Orm.Update(&model.PostCategory{}, updateMap, "post_category_id = ? or post_category_parent_id = ?", request.PostCategoryId, request.PostCategoryId)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}

		key1 := GetKey(KeyLevel2PostCategoryZSet + strconv.FormatInt(request.PostCategoryParentId, 10))
		_, err = post_category1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key1}, true, true)
		if err != nil {
			_, _ = mutex.Unlock()
			klog.Error(errDeletePostCategoryLevel2ToZSet, err)
			updateMap := map[string]interface{}{
				"deleted_at": nil,
			}
			_, _ = post_category1.GlobalEngine.Options.Orm.Update(&model.PostCategory{}, updateMap, "post_category_id = ? or post_category_parent_id = ?", request.PostCategoryId, request.PostCategoryId)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}

		_, _ = mutex.Unlock()
	case level2:
		postCategoryIdList = append(postCategoryIdList, request.PostCategoryId)
		_, err := post_category1.GlobalEngine.Options.Orm.Delete(&model.PostCategory{}, "post_category_id = ?", request.PostCategoryId)

		mutex := post_category1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
		defer post_category1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

		mutex = post_category1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(postCategoryLevel2 + strconv.FormatInt(request.PostCategoryParentId, 10))

		_ = mutex.Lock()
		key := GetKey(KeyLevel2PostCategoryZSet + strconv.FormatInt(request.PostCategoryParentId, 10))
		_, err = post_category1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, false, request.PostCategoryId)
		if err != nil {
			_, _ = mutex.Unlock()
			klog.Error(errDeletePostCategoryLevel2ToZSet, err)
			updateMap := map[string]interface{}{
				"deleted_at": nil,
			}
			_, _ = post_category1.GlobalEngine.Options.Orm.Update(&model.PostCategory{}, updateMap, "post_category_id = ?", request.PostCategoryId)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}
		_, _ = mutex.Unlock()
	}

	resp, err := post_category1.GlobalEngine.Options.PostClient.Client.DeleteResumeRelativeInfoByPostCategoryIdList(context.Background(), &post.DeleteResumeRelativeInfoByPostCategoryIdListRPCRequest{
		PostCategoryIdList: postCategoryIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil
	}
	if resp.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return int(resp.Code), resp.Message, nil
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func GetsPostCategory(request *post_category.GetsPostCategoryRPCRequest) (int, string, []*post_category.PostCategoryAllInformation, error) {
	var key string

	switch request.Level {
	case level1:
		key = GetKey(KeyLevel1PostCategoryZSet)
	case level2:
		key = GetKey(KeyLevel2PostCategoryZSet + strconv.FormatInt(request.PostCategoryId, 10))
	default:
		klog.Error(errInputLevel)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, errors.New(errInputLevel)
	}

	resultDatas := make([]*post_category.PostCategoryAllInformation, 0, maxLength)

	resultList, err := post_category1.GlobalEngine.Options.Cache.Gets(
		context.Background(),
		pkg.RedisDataStructureZSet, []string{key},
		strconv.FormatInt(pkg.MinInt64, 10),
		strconv.FormatInt(pkg.MaxInt64, 10),
		0,
		0,
	)
	if err != nil {
		klog.Error(errGetsPostCategoryLevel1FromRedis, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	postCategoryIdList := make([]int64, 0, len(resultList))

	result, ok := resultList[0].(string)
	if !ok {
		klog.Error(pkg1.ErrTypeConvert)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, errors.New(pkg1.ErrTypeConvert)
	}

	if result == "0" {
		if request.Level == 0 {
			resultDatas, err = getsPostCategoryWithNotZSet(postCategoryLevel1, key, request.Level, request.PostCategoryId)
			if err != nil {
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
			return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
		} else {
			resultDatas, err = getsPostCategoryWithNotZSet(
				postCategoryLevel2+strconv.FormatInt(request.PostCategoryId, 10),
				key,
				request.Level,
				request.PostCategoryId)
			if err != nil {
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
			return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
		}
	} else {
		postCategoryIdList, err = pkg1.InterfaceToStringToInt64(resultList)
		if err != nil {
			klog.Error(pkg1.ErrTypeConvert, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, errors.New(pkg1.ErrTypeConvert)
		}

		results, postCategoryNotFoundList, err := getPostCategoryByIdWithSet(postCategoryIdList)
		if err != nil {
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		j := 0
		for i := 0; i < len(results); i++ {
			if results[i] == nil {
				var r = post_category.PostCategoryAllInformation{
					PostCategoryId:   postCategoryNotFoundList[j].PostCategoryId,
					PostCategoryName: postCategoryNotFoundList[j].PostCategoryName,
				}
				resultDatas = append(resultDatas, &r)
				j++
			} else {
				dataPostCategory, err := unmarshalPostCategoryFromRedis(results[i])
				if err != nil {
					return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
				}
				var r = post_category.PostCategoryAllInformation{
					PostCategoryId:   dataPostCategory.PostCategoryId,
					PostCategoryName: dataPostCategory.PostCategoryName,
				}
				resultDatas = append(resultDatas, &r)
			}
		}
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
	}
}

func getsPostCategoryWithNotZSet(mutexKey, key string, level int8, postCategoryId int64) (resultDatas []*post_category.PostCategoryAllInformation, err error) {
	mutex := post_category1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
	defer post_category1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

	mutex = post_category1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(mutexKey)

	_ = mutex.Lock()

	postCategoryList := make([]model.PostCategory, 0, maxLength)

	switch level {
	case level1:
		_, err = post_category1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&postCategoryList,
			"created_at asc",
			[]string{"post_category_id", "post_category_name", "created_at"},
			"post_category_level = ?",
			level)
	case level2:
		_, err = post_category1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&postCategoryList,
			"created_at asc",
			[]string{"post_category_id", "post_category_name", "created_at"},
			"post_category_parent_id = ? and post_category_level = ?",
			postCategoryId, level)
	default:
		klog.Error(errInputLevel)
		return nil, errors.New(errInputLevel)
	}

	if err != nil {
		_, _ = mutex.Unlock()
		klog.Error(errGetsPostCategoryLevel1FromMysql, err)
		return
	}

	value := make([]interface{}, 0, len(postCategoryList))
	for i := 0; i < len(postCategoryList); i++ {
		value = append(value, postCategoryList[i].PostCategoryId)
		value = append(value, float64(postCategoryList[i].CreatedAt.UnixMilli()))
	}

	if len(value) != 0 {
		expire := pkg1.SetRandomExpireTime(random, 1, expireTime)
		_, err = post_category1.GlobalEngine.Options.Cache.Sets(context.Background(), pkg.RedisDataStructureZSet, []string{key}, value, true, expire[0])
		if err != nil {
			_, _ = mutex.Unlock()
			klog.Error(errSetsPostCategoryLevel1ToRedis, err)
			return
		}
	}

	for i := 0; i < len(postCategoryList); i++ {
		var data = post_category.PostCategoryAllInformation{
			PostCategoryId:   postCategoryList[i].PostCategoryId,
			PostCategoryName: postCategoryList[i].PostCategoryName,
		}
		resultDatas = append(resultDatas, &data)
	}

	_, _ = mutex.Unlock()

	return
}

func removeDuplicateInt(list []int64) []int64 {
	result := []int64{}
	for _, v := range list {
		exists := false
		for _, rv := range result {
			if v == rv {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, v)
		}
	}
	return result
}

func GetPostCategoryById(request *post_category.GetPostCategoryByIdRPCRequest) (int, string, []*post_category.PostCategoryAllInformation, error) {
	postCateList := removeDuplicateInt(request.PostCategoryId)

	resultDatas := make([]*post_category.PostCategoryAllInformation, 0, maxLength)

	results, postCategoryNotFoundList, err := getPostCategoryByIdWithSet(postCateList)
	if err != nil {
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	parentIdList := make([]int64, 0, len(results))
	j := 0
	for i := 0; i < len(results); i++ {
		if results[i] == nil {
			parentIdList = append(parentIdList, postCategoryNotFoundList[j].PostCategoryParentId)
		} else {
			dataPostCategory, err := unmarshalPostCategoryFromRedis(results[i])
			if err != nil {
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
			parentIdList = append(parentIdList, dataPostCategory.PostCategoryParentId)
		}
	}

	resultsParent, postCategoryNotFoundListParent, err := getPostCategoryByIdWithSet(parentIdList)
	if err != nil {
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	j = 0
	k := 0
	for i := 0; i < len(results); i++ {
		if results[i] == nil {
			var r = post_category.PostCategoryAllInformation{
				PostCategoryId:   postCategoryNotFoundList[j].PostCategoryId,
				PostCategoryName: postCategoryNotFoundList[j].PostCategoryName,
			}
			if resultsParent[i] == nil {
				r.PostCategoryParentId = postCategoryNotFoundListParent[k].PostCategoryId
				r.PostCategoryParentName = postCategoryNotFoundListParent[k].PostCategoryName
				k++
			} else {
				dataPostCategory, err := unmarshalPostCategoryFromRedis(resultsParent[i])
				if err != nil {
					return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
				}
				r.PostCategoryParentId = dataPostCategory.PostCategoryId
				r.PostCategoryParentName = dataPostCategory.PostCategoryName
			}
			resultDatas = append(resultDatas, &r)
			j++
		} else {
			dataPostCategory, err := unmarshalPostCategoryFromRedis(results[i])
			if err != nil {
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
			var r = post_category.PostCategoryAllInformation{
				PostCategoryId:   dataPostCategory.PostCategoryId,
				PostCategoryName: dataPostCategory.PostCategoryName,
			}
			if resultsParent[i] == nil {
				r.PostCategoryParentId = postCategoryNotFoundListParent[k].PostCategoryId
				r.PostCategoryParentName = postCategoryNotFoundListParent[k].PostCategoryName
				k++
			} else {
				dataPostCategory1, err := unmarshalPostCategoryFromRedis(resultsParent[i])
				if err != nil {
					return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
				}
				r.PostCategoryParentId = dataPostCategory1.PostCategoryId
				r.PostCategoryParentName = dataPostCategory1.PostCategoryName
			}
			resultDatas = append(resultDatas, &r)
		}
	}
	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
}

func getPostCategoryByIdWithSet(idList []int64) (results []interface{}, postCategoryNotFoundList []model.PostCategory, err error) {
	keysString := pkg1.Int64ToStringListWithPrefix(GetKey(KeyString), idList)

	results, postCategoryIdNotFoundList, err := GetsCacheWithNotFoundList(idList, keysString)
	if err != nil {
		klog.Error(err)
		return
	}

	postCategoryNotFoundList = make([]model.PostCategory, 0, len(postCategoryIdNotFoundList))

	var isExist bool
	if len(postCategoryIdNotFoundList) != 0 {
		isExist, err = post_category1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&postCategoryNotFoundList,
			"",
			[]string{"post_category_id", "post_category_name", "post_category_parent_id"},
			"post_category_id in ?",
			postCategoryIdNotFoundList)
		if err != nil || isExist == false {
			klog.Error(errGetsPostCategoryFromMysql, err)
			return
		}

		err = setPostCategoryNotFoundToRedis(postCategoryNotFoundList)
		if err != nil {
			return
		}
	}

	return
}

func setPostCategoryNotFoundToRedis(postCategoryNotFoundList []model.PostCategory) error {
	keys := make([]string, 0, len(postCategoryNotFoundList))
	datas := make([]interface{}, 0, len(postCategoryNotFoundList))
	expireTimeList := pkg1.SetRandomExpireTime(random, len(postCategoryNotFoundList), expireTime)

	for i := 0; i < len(postCategoryNotFoundList); i++ {
		keys = append(keys, GetKey(KeyString)+strconv.FormatInt(postCategoryNotFoundList[i].PostCategoryId, 10))
		byteData, err := json.Marshal(postCategoryNotFoundList[i])
		if err != nil {
			klog.Error(pkg1.ErrTypeConvert, err)
			return err
		}
		datas = append(datas, string(byteData))
	}

	isSets, err := post_category1.GlobalEngine.Options.Cache.Sets(context.Background(), pkg.RedisDataStructureString, keys, datas, true, expireTimeList...)
	if err != nil || isSets == false {
		klog.Error(errSetPostCategoryNotFound, err)
		return err
	}

	return nil
}

func unmarshalPostCategoryFromRedis(result interface{}) (dataPostCategory model.PostCategory, err error) {
	result1, ok := result.(string)
	if !ok {
		klog.Error(pkg1.ErrTypeConvert)
		return model.PostCategory{}, errors.New(pkg1.ErrTypeConvert)
	}

	if err = json.Unmarshal([]byte(result1), &dataPostCategory); err != nil {
		klog.Error(errJsonUnMarshal, err)
		return model.PostCategory{}, err
	}

	return
}
