package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redsync/redsync/v4"
	"resume-resolving/api/idl/service/department/kitex_gen/department"
	"resume-resolving/api/idl/service/post/kitex_gen/post"
	"resume-resolving/api/idl/service/user/kitex_gen/user"
	department1 "resume-resolving/internal/app/service/department"
	"resume-resolving/internal/app/service/department/model"
	"resume-resolving/internal/app/service/department/pkg/code"
	pkg1 "resume-resolving/internal/app/service/pkg"
	"resume-resolving/internal/pkg"
	"strconv"
	"time"
)

const (
	d             = "department"
	c             = "city"
	departCityMap = "department_city_map:"
	max           = 100
	nilString     = ""
	maxLength     = 50
	expireTime    = 3 * time.Hour
	random        = 3600
)

func AppendDepartment(request *department.AppendDepartmentRPCRequest) (int, string, *department.DepartmentInformation, error) {
	createTime := time.Now()
	departmentId := department1.GlobalEngine.Options.Id.GenId()

	var depart = model.Department{
		DepartmentId:          departmentId,
		DepartmentName:        request.DepartmentName,
		DepartmentDescription: request.DepartmentDescription,
		CreatedAt:             createTime,
	}

	departmentCityMaps := make([]*model.DepartmentCityMap, 0, len(request.CityList))
	for i := 0; i < len(request.CityList); i++ {
		var departmentCityMap = model.DepartmentCityMap{
			Id:           department1.GlobalEngine.Options.Id.GenId(),
			CityId:       request.CityList[i],
			DepartmentId: departmentId,
			CreatedAt:    time.Now(),
		}
		departmentCityMaps = append(departmentCityMaps, &departmentCityMap)
	}

	var value = [][]interface{}{{&depart}, {departmentCityMaps}}

	if err := department1.GlobalEngine.Options.Orm.Transaction([]string{pkg.DbFunctionCreate, pkg.DbFunctionCreate}, value); err != nil {
		klog.Error(errCreateDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	key := GetKey(KeyZSet + d)
	createTime1 := createTime.UnixMilli()
	mutex := department1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
	defer department1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

	mutex = department1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(KeyZSet + d)

	_ = mutex.Lock()

	_, err := department1.GlobalEngine.Options.Cache.Set(context.Background(), pkg.RedisDataStructureZSet, key, []interface{}{departmentId, float64(createTime1)})
	if err != nil {
		_, _ = mutex.Unlock()
		klog.Error(err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}
	_, _ = mutex.Unlock()

	var resultData = &department.DepartmentInformation{
		DepartmentId:          departmentId,
		DepartmentName:        request.DepartmentName,
		DepartmentDescription: request.DepartmentDescription,
	}
	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultData, nil
}

func UpdateDepartment(request *department.UpdateDepartmentRPCRequest) (int, string, *department.DepartmentInformation, error) {
	var updateMap map[string]interface{}
	if request.DepartmentName != nilString {
		if request.DepartmentDescription != nilString {
			updateMap = map[string]interface{}{
				"department_name":        request.DepartmentName,
				"department_description": request.DepartmentDescription,
			}
		} else {
			updateMap = map[string]interface{}{
				"department_name": request.DepartmentName,
			}
		}
	} else {
		updateMap = map[string]interface{}{
			"department_description": request.DepartmentDescription,
		}
	}

	isExist, err := department1.GlobalEngine.Options.Orm.Update(&model.Department{}, updateMap, "department_id = ?", request.DepartmentId)
	if err != nil || isExist == false {
		klog.Error(errUpdateDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	key := GetKey(KeyString + strconv.FormatInt(request.DepartmentId, 10))
	_, err = department1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, true)

	if err != nil {
		klog.Error(errDeleteDepartmentFromRedis, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	var resultData = &department.DepartmentInformation{
		DepartmentId:          request.DepartmentId,
		DepartmentName:        request.DepartmentName,
		DepartmentDescription: request.DepartmentDescription,
	}
	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultData, nil
}

func DeleteDepartment(request *department.DeleteDepartmentRPCRequest) (int, string, error) {
	var deleteValue = [][]interface{}{{&model.Department{}, "department_id = ?", request.DepartmentId}, {&model.DepartmentCityMap{}, "department_id = ?", request.DepartmentId}}
	if err := department1.GlobalEngine.Options.Orm.Transaction([]string{pkg.DbFunctionDelete, pkg.DbFunctionDelete}, deleteValue); err != nil {
		klog.Error(errDeleteDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	mutexDepartment := department1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
	defer department1.GlobalEngine.Options.RedSync.LockPool.Put(mutexDepartment)

	mutexDepartment = department1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(KeyZSet + d)
	key := GetKey(KeyZSet + d)
	_ = mutexDepartment.Lock()

	_, err := department1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, false, request.DepartmentId)
	if err != nil {
		_, _ = mutexDepartment.Unlock()
		klog.Error(errDeleteDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}
	_, _ = mutexDepartment.Unlock()

	mutexMap := department1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
	defer department1.GlobalEngine.Options.RedSync.LockPool.Put(mutexMap)

	mutexMap = department1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(KeyZSet + departCityMap + strconv.FormatInt(request.DepartmentId, 10))
	key = GetKey(KeyZSet + departCityMap + strconv.FormatInt(request.DepartmentId, 10))
	_ = mutexMap.Lock()

	_, err = department1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, true)
	if err != nil {
		_, _ = mutexMap.Unlock()
		klog.Error(errDeleteDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}
	_, _ = mutexMap.Unlock()

	key = GetKey(KeyString + strconv.FormatInt(request.DepartmentId, 10))
	_, err = department1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, true)
	if err != nil {
		klog.Error(errDeleteDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	resp, err := department1.GlobalEngine.Options.UserClient.Client.DeleteHRByDepartmentId(context.Background(), &user.DeleteHRByDepartmentIdRPCRequest{
		DepartmentId: request.DepartmentId,
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

func GetsDepartment(request *department.GetsDepartmentRPCRequest) (int, string, []*department.DepartmentInformation, error) {
	key := GetKey(KeyZSet + d)

	resultDatas := make([]*department.DepartmentInformation, 0, maxLength)

	departmentIdList, result, err := GetCacheZSet([]string{key}, strconv.FormatInt(pkg.MinInt64, 10), strconv.FormatInt(pkg.MaxInt64, 10), 0, 0)
	if err != nil {
		klog.Error(errGetsDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	if result == "0" {
		mutexDepartment := department1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
		defer department1.GlobalEngine.Options.RedSync.LockPool.Put(mutexDepartment)

		mutexDepartment = department1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(KeyZSet + d)
		key = GetKey(KeyZSet + d)
		_ = mutexDepartment.Lock()

		departmentList := make([]model.Department, 0, maxLength)

		_, err = department1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&departmentList,
			"created_at asc",
			[]string{"department_id", "department_name", "department_description", "created_at"},
			pkg.NotUseWhere)
		if err != nil {
			_, _ = mutexDepartment.Unlock()
			klog.Error(errGetsDepartment, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		value := make([]interface{}, 0, len(departmentList))
		for i := 0; i < len(departmentList); i++ {
			value = append(value, departmentList[i].DepartmentId)
			value = append(value, float64(departmentList[i].CreatedAt.UnixMilli()))
		}

		if len(value) != 0 {
			expire := pkg1.SetRandomExpireTime(random, 1, expireTime)
			_, err = department1.GlobalEngine.Options.Cache.Sets(context.Background(), pkg.RedisDataStructureZSet, []string{key}, value, true, expire[0])
			if err != nil {
				_, _ = mutexDepartment.Unlock()
				klog.Error(errGetsDepartment, err)
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
		}

		_, _ = mutexDepartment.Unlock()

		for i := 0; i < len(departmentList); i++ {
			var data = department.DepartmentInformation{
				DepartmentId:          departmentList[i].DepartmentId,
				DepartmentName:        departmentList[i].DepartmentName,
				DepartmentDescription: departmentList[i].DepartmentDescription,
			}
			resultDatas = append(resultDatas, &data)
		}

		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
	} else {
		keysString := pkg1.Int64ToStringListWithPrefix(GetKey(KeyString), departmentIdList)

		results, departmentIdNotFoundList, err := GetsCacheWithNotFoundList(departmentIdList, keysString)
		if err != nil {
			klog.Error(errGetsDepartment, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		departmentNotFoundList := make([]model.Department, 0, len(departmentIdNotFoundList))

		if len(departmentIdNotFoundList) != 0 {
			isExist, err := department1.GlobalEngine.Options.Orm.Query(
				-1,
				-1,
				&departmentNotFoundList,
				"created_at asc",
				[]string{"department_id", "department_name", "department_description"},
				"department_id in ?",
				departmentIdNotFoundList)
			if err != nil || isExist == false {
				klog.Error(errGetsDepartment, err)
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
		}

		if len(departmentIdNotFoundList) != len(departmentNotFoundList) {
			klog.Error(errConcurrent)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, errors.New(errConcurrent)
		}

		keys := make([]string, 0, len(departmentNotFoundList))
		data := make([]interface{}, 0, len(departmentNotFoundList))
		for i := 0; i < len(departmentNotFoundList); i++ {
			keys = append(keys, GetKey(KeyString)+strconv.FormatInt(departmentNotFoundList[i].DepartmentId, 10))
			data = append(data, departmentNotFoundList[i])
		}

		if err = SetCacheNotFoundToRedis(data, keys); err != nil {
			klog.Error(errGetsDepartment, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		j := 0
		for i := 0; i < len(results); i++ {
			if results[i] == nil {
				var r = department.DepartmentInformation{
					DepartmentId:          departmentNotFoundList[j].DepartmentId,
					DepartmentName:        departmentNotFoundList[j].DepartmentName,
					DepartmentDescription: departmentNotFoundList[j].DepartmentDescription,
				}
				resultDatas = append(resultDatas, &r)
				j++
			} else {
				dataDepartment, err := unmarshalDepartmentFromRedis(results[i])
				if err != nil {
					return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
				}
				var r = department.DepartmentInformation{
					DepartmentId:          dataDepartment.DepartmentId,
					DepartmentName:        dataDepartment.DepartmentName,
					DepartmentDescription: dataDepartment.DepartmentDescription,
				}
				resultDatas = append(resultDatas, &r)
			}
		}
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
	}
}

func AppendCity(request *department.AppendCityRPCRequest) (int, string, *department.CityInformation, error) {
	createTime := time.Now()
	cityId := department1.GlobalEngine.Options.Id.GenId()

	var city = model.City{
		CityId:    cityId,
		CityName:  request.CityName,
		CreatedAt: createTime,
	}

	isExist, err := department1.GlobalEngine.Options.Orm.Create(&city)
	if err != nil || isExist == false {
		klog.Error(errCreateCity, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	key := GetKey(KeyZSet + c)
	createTime1 := createTime.UnixMilli()
	mutex := department1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
	defer department1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

	mutex = department1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(KeyZSet + c)

	_ = mutex.Lock()

	_, err = department1.GlobalEngine.Options.Cache.Set(context.Background(), pkg.RedisDataStructureZSet, key, []interface{}{cityId, float64(createTime1)})
	if err != nil {
		_, _ = mutex.Unlock()
		klog.Error(err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}
	_, _ = mutex.Unlock()

	var resultData = &department.CityInformation{
		CityId:   cityId,
		CityName: request.CityName,
	}
	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultData, nil
}

func DeleteCity(request *department.DeleteCityRPCRequest) (int, string, error) {
	var deleteValue = [][]interface{}{
		{&model.City{}, "city_id = ?", request.CityId},
		{&model.DepartmentCityMap{}, "city_id = ?", request.CityId},
	}
	err := department1.GlobalEngine.Options.Orm.Transaction([]string{pkg.DbFunctionDelete, pkg.DbFunctionDelete}, deleteValue)
	if err != nil {
		klog.Error(errDeleteCity, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	mutex := department1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
	defer department1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

	mutex = department1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(KeyZSet + c)
	key := GetKey(KeyZSet + c)
	_ = mutex.Lock()

	_, err = department1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, false, request.CityId)
	if err != nil {
		_, _ = mutex.Unlock()
		klog.Error(errDeleteCity, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}
	_, _ = mutex.Unlock()

	key = GetKey(KeyString + strconv.FormatInt(request.CityId, 10))
	_, err = department1.GlobalEngine.Options.Cache.Delete(context.Background(), []string{key}, true, true)
	if err != nil {
		klog.Error(errDeleteCityFromRedis, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	resp, err := department1.GlobalEngine.Options.PostClient.Client.DeletePostCityByCityId(context.Background(), &post.DeletePostCityByCityIdRPCRequest{
		CityId: request.CityId,
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

func GetsCity(request *department.GetsCityRPCRequest) (int, string, []*department.CityInformation, error) {
	key := GetKey(KeyZSet + c)

	resultDatas := make([]*department.CityInformation, 0, maxLength)

	cityIdList, result, err := GetCacheZSet([]string{key}, strconv.FormatInt(pkg.MinInt64, 10), strconv.FormatInt(pkg.MaxInt64, 10), 0, 0)
	if err != nil {
		klog.Error(errGetsCity, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	if result == "0" {
		mutex := department1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
		defer department1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

		mutex = department1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(KeyZSet + c)
		key = GetKey(KeyZSet + c)
		_ = mutex.Lock()

		cityList := make([]model.City, 0, maxLength)

		_, err = department1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&cityList,
			"created_at asc",
			[]string{"city_id", "city_name", "created_at"},
			pkg.NotUseWhere)
		if err != nil {
			_, _ = mutex.Unlock()
			klog.Error(errGetsCity, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		value := make([]interface{}, 0, len(cityList))
		for i := 0; i < len(cityList); i++ {
			value = append(value, cityList[i].CityId)
			value = append(value, float64(cityList[i].CreatedAt.UnixMilli()))
		}

		if len(value) != 0 {
			expire := pkg1.SetRandomExpireTime(random, 1, expireTime)
			_, err = department1.GlobalEngine.Options.Cache.Sets(context.Background(), pkg.RedisDataStructureZSet, []string{key}, value, true, expire[0])
			if err != nil {
				_, _ = mutex.Unlock()
				klog.Error(errGetsCity, err)
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
		}

		_, _ = mutex.Unlock()

		for i := 0; i < len(cityList); i++ {
			var data = department.CityInformation{
				CityId:   cityList[i].CityId,
				CityName: cityList[i].CityName,
			}
			resultDatas = append(resultDatas, &data)
		}

		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
	} else {
		keysString := pkg1.Int64ToStringListWithPrefix(GetKey(KeyString), cityIdList)

		results, cityIdNotFoundList, err := GetsCacheWithNotFoundList(cityIdList, keysString)
		if err != nil {
			klog.Error(errGetsCity, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		cityNotFoundList := make([]model.City, 0, len(cityIdNotFoundList))
		if len(cityIdNotFoundList) != 0 {
			isExist, err := department1.GlobalEngine.Options.Orm.Query(
				-1,
				-1,
				&cityNotFoundList,
				"created_at asc",
				[]string{"city_id", "city_name"},
				"city_id in ?",
				cityIdNotFoundList)
			if err != nil || isExist == false {
				klog.Error(errGetsCity, err)
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
		}

		if len(cityIdNotFoundList) != len(cityNotFoundList) {
			klog.Error(errConcurrent)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, errors.New(errConcurrent)
		}

		keys := make([]string, 0, len(cityNotFoundList))
		data := make([]interface{}, 0, len(cityNotFoundList))
		for i := 0; i < len(cityNotFoundList); i++ {
			keys = append(keys, GetKey(KeyString)+strconv.FormatInt(cityNotFoundList[i].CityId, 10))
			data = append(data, cityNotFoundList[i])
		}

		if err = SetCacheNotFoundToRedis(data, keys); err != nil {
			klog.Error(errGetsCity, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		j := 0
		for i := 0; i < len(results); i++ {
			if results[i] == nil {
				var r = department.CityInformation{
					CityId:   cityNotFoundList[j].CityId,
					CityName: cityNotFoundList[j].CityName,
				}
				resultDatas = append(resultDatas, &r)
				j++
			} else {
				city, err := unmarshalCityFromRedis(results[i])
				if err != nil {
					return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
				}
				var r = department.CityInformation{
					CityId:   city.CityId,
					CityName: city.CityName,
				}
				resultDatas = append(resultDatas, &r)
			}
		}
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
	}
}

func GetsCityByDepartment(request *department.GetsCityByDepartmentRPCRequest) (int, string, []*department.CityInformation, error) {
	key := GetKey(KeyZSet + departCityMap + strconv.FormatInt(request.DepartmentId, 10))
	resultDatas := make([]*department.CityInformation, 0, maxLength)

	cityIdList, result, err := GetCacheZSet([]string{key}, strconv.FormatInt(pkg.MinInt64, 10), strconv.FormatInt(pkg.MaxInt64, 10), 0, 0)
	if err != nil {
		klog.Error(errGetsCityByDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	if result == "0" {
		mutex := department1.GlobalEngine.Options.RedSync.LockPool.Get().(*redsync.Mutex)
		defer department1.GlobalEngine.Options.RedSync.LockPool.Put(mutex)

		mutex = department1.GlobalEngine.Options.RedSync.RedisPool.NewMutex(KeyZSet + departCityMap + strconv.FormatInt(request.DepartmentId, 10))
		_ = mutex.Lock()

		cityList := make([]model.DepartmentCityMap, 0, maxLength)
		isExist, err := department1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&cityList,
			"created_at asc",
			[]string{"city_id", "created_at"},
			"department_id = ?",
			request.DepartmentId)
		if err != nil || isExist == false {
			_, _ = mutex.Unlock()
			klog.Error(errGetsCityByDepartment, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		value := make([]interface{}, 0, len(cityList))
		for i := 0; i < len(cityList); i++ {
			value = append(value, cityList[i].CityId)
			value = append(value, float64(cityList[i].CreatedAt.UnixMilli()))
		}

		expire := pkg1.SetRandomExpireTime(random, 1, expireTime)
		_, err = department1.GlobalEngine.Options.Cache.Sets(context.Background(), pkg.RedisDataStructureZSet, []string{key}, value, true, expire[0])
		if err != nil {
			_, _ = mutex.Unlock()
			klog.Error(errGetsCityByDepartment, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}

		_, _ = mutex.Unlock()

		if len(cityIdList) == 0 {
			for i := 0; i < len(cityList); i++ {
				cityIdList = append(cityIdList, cityList[i].CityId)
			}
		}
	}
	keysString := pkg1.Int64ToStringListWithPrefix(GetKey(KeyString), cityIdList)

	results, cityIdNotFoundList, err := GetsCacheWithNotFoundList(cityIdList, keysString)
	if err != nil {
		klog.Error(errGetsCityByDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	cityNotFoundList := make([]model.City, 0, len(cityIdNotFoundList))
	if len(cityIdNotFoundList) != 0 {
		isExist, err := department1.GlobalEngine.Options.Orm.Query(
			-1,
			-1,
			&cityNotFoundList,
			"created_at asc",
			[]string{"city_id", "city_name"},
			"city_id in ?",
			cityIdNotFoundList)
		if err != nil || isExist == false {
			klog.Error(errGetsCityByDepartment, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}
	}

	if len(cityIdNotFoundList) != len(cityNotFoundList) {
		klog.Error(errConcurrent)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, errors.New(errConcurrent)
	}

	keys := make([]string, 0, len(cityNotFoundList))
	data := make([]interface{}, 0, len(cityNotFoundList))
	for i := 0; i < len(cityNotFoundList); i++ {
		keys = append(keys, GetKey(KeyString)+strconv.FormatInt(cityNotFoundList[i].CityId, 10))
		data = append(data, cityNotFoundList[i])
	}

	if err = SetCacheNotFoundToRedis(data, keys); err != nil {
		klog.Error(errGetsCityByDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	j := 0
	for i := 0; i < len(results); i++ {
		if results[i] == nil {
			var r = department.CityInformation{
				CityId:   cityNotFoundList[j].CityId,
				CityName: cityNotFoundList[j].CityName,
			}
			resultDatas = append(resultDatas, &r)
			j++
		} else {
			city, err := unmarshalCityFromRedis(results[i])
			if err != nil {
				return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
			}
			var r = department.CityInformation{
				CityId:   city.CityId,
				CityName: city.CityName,
			}
			resultDatas = append(resultDatas, &r)
		}
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
}

func GetDepartmentInfosById(request *department.GetDepartmentInfosByIdRPCRequest) (int, string, []*department.DepartmentInformation, error) {
	data, err := getDepartmentInfosById(request.DepartmentIdList)
	if err != nil {
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), data, nil
}

func GetCityInfoById(request *department.GetCityInfoByIdRPCRequest) (int, string, []*department.CityInformation, error) {
	data, err := getCityInfoById(request.CityId)
	if err != nil {
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), data, nil
}

func GetDepartmentCityInfoById(request *department.GetDepartmentCityInfoByIdRPCRequest) (int, string, []*department.DepartmentInformation, []*department.CityInformation, error) {
	dataDepartment, err := getDepartmentInfosById(request.DepartmentIdList)
	if err != nil {
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil, err
	}

	dataCity, err := getCityInfoById(request.CityIdList)
	if err != nil {
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil, err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), dataDepartment, dataCity, nil
}

func getDepartmentInfosById(departmentIdList []int64) ([]*department.DepartmentInformation, error) {
	departments := make([]model.Department, 0, max)
	_, err := department1.GlobalEngine.Options.Orm.Query(-1, -1, &departments, "", []string{"department_id", "department_name", "department_description"},
		"department_id in ?", departmentIdList)
	if err != nil {
		klog.Error(errGetDepartmentInfosById, err)
		return nil, err
	}

	resultDatas := make([]*department.DepartmentInformation, 0, len(departments))
	for i := 0; i < len(departments); i++ {
		var resultData = department.DepartmentInformation{
			DepartmentId:          departments[i].DepartmentId,
			DepartmentName:        departments[i].DepartmentName,
			DepartmentDescription: departments[i].DepartmentDescription,
		}
		resultDatas = append(resultDatas, &resultData)
	}

	return resultDatas, nil
}

func getCityInfoById(cityIdList []int64) ([]*department.CityInformation, error) {
	cities := make([]model.City, 0, max)
	_, err := department1.GlobalEngine.Options.Orm.Query(-1, -1, &cities, "", []string{"city_id", "city_name"},
		"city_id in ?", cityIdList)
	if err != nil {
		klog.Error(errGetCityInfoById, err)
		return nil, err
	}

	resultDatas := make([]*department.CityInformation, 0, len(cities))
	for i := 0; i < len(cities); i++ {
		var resultData = department.CityInformation{
			CityId:   cities[i].CityId,
			CityName: cities[i].CityName,
		}
		resultDatas = append(resultDatas, &resultData)
	}

	return resultDatas, nil
}

func unmarshalDepartmentFromRedis(result interface{}) (dataDepartment model.Department, err error) {
	result1, ok := result.(string)
	if !ok {
		klog.Error(pkg1.ErrTypeConvert)
		return model.Department{}, errors.New(pkg1.ErrTypeConvert)
	}

	if err = json.Unmarshal([]byte(result1), &dataDepartment); err != nil {
		klog.Error(errJsonUnMarshal, err)
		return model.Department{}, err
	}

	return
}

func unmarshalCityFromRedis(result interface{}) (city model.City, err error) {
	result1, ok := result.(string)
	if !ok {
		klog.Error(pkg1.ErrTypeConvert)
		return model.City{}, errors.New(pkg1.ErrTypeConvert)
	}

	if err = json.Unmarshal([]byte(result1), &city); err != nil {
		klog.Error(errJsonUnMarshal, err)
		return model.City{}, err
	}

	return
}
