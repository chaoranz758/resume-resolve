package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"resume-resolving/api/idl/service/department/kitex_gen/department"
	department2 "resume-resolving/api/idl/service/post/kitex_gen/department"
	"resume-resolving/api/idl/service/post/kitex_gen/post"
	category "resume-resolving/api/idl/service/post/kitex_gen/post_category"
	"resume-resolving/api/idl/service/post_category/kitex_gen/post_category"
	post1 "resume-resolving/internal/app/service/post"
	"resume-resolving/internal/app/service/post/model"
	"resume-resolving/internal/app/service/post/pkg/code"
	"resume-resolving/internal/pkg"
)

const (
	max              = 100
	maxLength        = 10
	maxDeliveryCount = 1
)

var resumeStatusMap = map[int8]string{
	0: "未处理",
	1: "通过",
	2: "未通过",
}

func AppendPost(request *post.AppendPostRPCRequest) (int32, string, error) {
	postId := post1.GlobalEngine.Options.Id.GenId()

	var p = model.Post{
		PostId:              postId,
		HRId:                request.HrId,
		PostCategoryId:      request.PostCategoryId,
		DepartmentId:        request.DepartmentId,
		IsSchoolRecruitment: request.IsSchoolRecruitment,
		IsInternship:        request.IsInternship,
		PostBrief:           request.PostBrief,
		PostDescription:     request.PostDescription,
		PostRequire:         request.PostRequire,
	}

	postCityMaps := make([]*model.PostCityMap, 0, len(request.CityList))
	for i := 0; i < len(request.CityList); i++ {
		var postCityMap = model.PostCityMap{
			Id:     post1.GlobalEngine.Options.Id.GenId(),
			CityId: request.CityList[i],
			PostId: postId,
		}
		postCityMaps = append(postCityMaps, &postCityMap)
	}

	var createValue = [][]interface{}{
		{&p},
		{postCityMaps}}

	err := post1.GlobalEngine.Options.Orm.Transaction([]string{pkg.DbFunctionCreate, pkg.DbFunctionCreate}, createValue)
	if err != nil {
		klog.Error(errCreatePost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func UpdatePost(request *post.UpdatePostRPCRequest) (int32, string, error) {
	var updateMap map[string]interface{}

	if request.PostBrief != pkg.NilString {
		if request.PostDescription != pkg.NilString {
			if request.PostRequire != pkg.NilString {
				updateMap = map[string]interface{}{
					"post_brief":       request.PostBrief,
					"post_description": request.PostDescription,
					"post_require":     request.PostRequire,
				}
			} else {
				updateMap = map[string]interface{}{
					"post_brief":       request.PostBrief,
					"post_description": request.PostDescription,
				}
			}
		} else {
			if request.PostRequire != pkg.NilString {
				updateMap = map[string]interface{}{
					"post_brief":   request.PostBrief,
					"post_require": request.PostRequire,
				}
			} else {
				updateMap = map[string]interface{}{
					"post_brief": request.PostBrief,
				}
			}
		}
	} else {
		if request.PostDescription != pkg.NilString {
			if request.PostRequire != pkg.NilString {
				updateMap = map[string]interface{}{
					"post_description": request.PostDescription,
					"post_require":     request.PostRequire,
				}
			} else {
				updateMap = map[string]interface{}{
					"post_description": request.PostDescription,
				}
			}
		} else {
			updateMap = map[string]interface{}{
				"post_require": request.PostRequire,
			}
		}
	}

	fmt.Printf("%v\n", updateMap)

	isUpdate, err := post1.GlobalEngine.Options.Orm.Update(&model.Post{}, updateMap, "post_id = ?", request.PostId)
	if err != nil || isUpdate == false {
		klog.Error(errUpdatePost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func DeletePost(request *post.DeletePostRPCRequest) (int32, string, error) {
	var deleteValue = [][]interface{}{
		{&model.Post{}, "post_id = ?", request.PostId},
		{&model.CollectPost{}, "post_id = ?", request.PostId},
		{&model.DeliveryPost{}, "post_id = ?", request.PostId},
		{&model.PostCityMap{}, "post_id = ?", request.PostId}}

	if err := post1.GlobalEngine.Options.Orm.Transaction([]string{pkg.DbFunctionDelete, pkg.DbFunctionDelete, pkg.DbFunctionDelete, pkg.DbFunctionDelete}, deleteValue); err != nil {
		klog.Error(errDeletePost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func GetsPostInUser(request *post.GetsPostInUserRPCRequest) (int32, string, []*post.PostInfo, error) {
	resultDatas := make([]*post.PostInfo, 0, request.Limit)

	posts := make([]model.Post, 0, request.Limit)
	if request.IsNew == 1 {
		isExist, err := post1.GlobalEngine.Options.Orm.Query(
			int(request.Limit),
			-1,
			&posts,
			"created_at desc",
			[]string{"post_id", "post_brief", "post_description", "post_require", "is_school_recruitment", "is_internship", "post_category_id", "department_id"},
			pkg.NotUseWhere)
		if err != nil || isExist == false {
			klog.Error(errGetsPostInUser, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}
	} else {
		var query string
		var args []interface{}
		if request.IsInternship != 0 {
			if len(request.DepartmentIdList) != 0 {
				if len(request.PostCategoryIdList) != 0 {
					query = "department_id in ? and post_category_id in ? and is_school_recruitment = ? and is_internship = ?"
					args = append(args, request.DepartmentIdList, request.PostCategoryIdList, request.IsSchoolRecruitment, request.IsInternship)
				} else {
					query = "department_id in ? and is_school_recruitment = ? and is_internship = ?"
					args = append(args, request.DepartmentIdList, request.IsSchoolRecruitment, request.IsInternship)
				}
			} else {
				if len(request.PostCategoryIdList) != 0 {
					query = "post_category_id in ? and is_school_recruitment = ? and is_internship = ?"
					args = append(args, request.PostCategoryIdList, request.IsSchoolRecruitment, request.IsInternship)
				} else {
					query = "is_school_recruitment = ? and is_internship = ?"
					args = append(args, request.IsSchoolRecruitment, request.IsInternship)
				}
			}
		} else {
			if len(request.DepartmentIdList) != 0 {
				if len(request.PostCategoryIdList) != 0 {
					query = "department_id in ? and post_category_id in ?"
					args = append(args, request.DepartmentIdList, request.PostCategoryIdList)
				} else {
					query = "department_id in ?"
					args = append(args, request.DepartmentIdList)
				}
			} else {
				if len(request.PostCategoryIdList) != 0 {
					query = "post_category_id in ?"
					args = append(args, request.PostCategoryIdList)
				} else {
					query = pkg.NotUseWhere
				}
			}
		}
		isExist, err := post1.GlobalEngine.Options.Orm.Query(
			int(request.Limit),
			int(request.Offset),
			&posts,
			"created_at desc",
			[]string{"post_id", "post_brief", "post_description", "post_require", "is_school_recruitment", "is_internship", "post_category_id", "department_id"},
			query,
			args...)
		if err != nil || isExist == false {
			klog.Error(errGetsPostInUser, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}
	}

	if len(posts) == 0 {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil, nil
	}

	postIdList := make([]int64, 0, len(posts))
	postCategoryIdList := make([]int64, 0, len(posts))
	departmentIdList := make([]int64, 0, len(posts))
	for i := 0; i < len(posts); i++ {
		postIdList = append(postIdList, posts[i].PostId)
		postCategoryIdList = append(postCategoryIdList, posts[i].PostCategoryId)
		departmentIdList = append(departmentIdList, posts[i].DepartmentId)
	}

	postCityMap := make([]model.PostCityMap, 0, max)
	isExist, err := post1.GlobalEngine.Options.Orm.Query(-1, -1, &postCityMap, "", []string{"post_id", "city_id"}, "post_id in ?", postIdList)
	if err != nil || isExist == false {
		klog.Error(errGetsPostInUser, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}
	cityIdList := make([]int64, 0, len(postCityMap))
	for i := 0; i < len(postCityMap); i++ {
		cityIdList = append(cityIdList, postCityMap[i].CityId)
	}

	//查岗位类别信息
	dataPostCategory, err := post1.GlobalEngine.Options.PostCategoryClient.Client.GetPostCategoryById(context.Background(), &post_category.GetPostCategoryByIdRPCRequest{
		PostCategoryId: postCategoryIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if dataPostCategory.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return dataPostCategory.Code, dataPostCategory.Message, nil, nil
	}

	//查部门和城市信息
	dataDepartmentCity, err := post1.GlobalEngine.Options.DepartmentClient.Client.GetDepartmentCityInfoById(context.Background(), &department.GetDepartmentCityInfoByIdRPCRequest{
		DepartmentIdList: departmentIdList,
		CityIdList:       cityIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if dataDepartmentCity.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return dataDepartmentCity.Code, dataDepartmentCity.Message, nil, nil
	}

	for i := 0; i < len(posts); i++ {
		var postCategoryInfo category.PostCategoryAllInformation
		for j := 0; j < len(dataPostCategory.PostCategoryInfoList); j++ {
			if dataPostCategory.PostCategoryInfoList[j].PostCategoryId == posts[i].PostCategoryId {
				postCategoryInfo = category.PostCategoryAllInformation{
					PostCategoryId:         dataPostCategory.PostCategoryInfoList[j].PostCategoryId,
					PostCategoryName:       dataPostCategory.PostCategoryInfoList[j].PostCategoryName,
					PostCategoryParentId:   dataPostCategory.PostCategoryInfoList[j].PostCategoryParentId,
					PostCategoryParentName: dataPostCategory.PostCategoryInfoList[j].PostCategoryParentName,
				}
				break
			}
		}

		var departmentInfo department2.DepartmentInformation
		for j := 0; j < len(dataDepartmentCity.DepartmentInfoList); j++ {
			if dataDepartmentCity.DepartmentInfoList[j].DepartmentId == posts[i].DepartmentId {
				departmentInfo = department2.DepartmentInformation{
					DepartmentId:          dataDepartmentCity.DepartmentInfoList[j].DepartmentId,
					DepartmentName:        dataDepartmentCity.DepartmentInfoList[j].DepartmentName,
					DepartmentDescription: dataDepartmentCity.DepartmentInfoList[j].DepartmentDescription,
				}
				break
			}
		}

		cityInfos := make([]*department2.CityInformation, 0, max)

		for j := 0; j < len(postCityMap); j++ {
			if postCityMap[j].PostId == posts[i].PostId {
				var cityInfo department2.CityInformation
				for k := 0; k < len(dataDepartmentCity.CityInfoList); k++ {
					if dataDepartmentCity.CityInfoList[k].CityId == postCityMap[j].CityId {
						cityInfo = department2.CityInformation{
							CityId:   dataDepartmentCity.CityInfoList[k].CityId,
							CityName: dataDepartmentCity.CityInfoList[k].CityName,
						}
						break
					}
				}
				cityInfos = append(cityInfos, &cityInfo)
			}
		}

		var resultData = post.PostInfo{
			PostId:                  posts[i].PostId,
			PostBrief:               posts[i].PostBrief,
			PostDescription:         posts[i].PostDescription,
			PostRequire:             posts[i].PostRequire,
			IsSchoolRecruitment:     posts[i].IsSchoolRecruitment,
			IsInternship:            posts[i].IsInternship,
			PostCategoryInformation: &postCategoryInfo,
			DepartmentInformation:   &departmentInfo,
			CityInformation:         cityInfos,
		}
		resultDatas = append(resultDatas, &resultData)
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
}

func GetsPostInHR(request *post.GetsPostInHRRPCRequest) (int32, string, []*post.PostInfo, error) {
	resultDatas := make([]*post.PostInfo, 0, request.Limit)

	var query string
	var args []interface{}
	if request.IsSchoolRecruitment != 0 {
		if request.HrId == 0 {
			query = "is_school_recruitment = ? and is_internship = ?"
			args = append(args, request.IsSchoolRecruitment, request.IsInternship)
		} else {
			query = "hr_id = ? and is_school_recruitment = ? and is_internship = ?"
			args = append(args, request.HrId, request.IsSchoolRecruitment, request.IsInternship)
		}
	} else {
		if request.HrId == 0 {
			query = pkg.NotUseWhere
		} else {
			query = "hr_id = ?"
			args = append(args, request.HrId)
		}
	}

	posts := make([]model.Post, 0, request.Limit)
	isExist, err := post1.GlobalEngine.Options.Orm.Query(
		int(request.Limit),
		int(request.Offset),
		&posts,
		"created_at desc",
		[]string{"post_id", "post_brief", "post_description", "post_require", "is_school_recruitment", "is_internship", "post_category_id", "department_id"},
		query,
		args...)
	if err != nil || isExist == false {
		klog.Error(errGetsPostInHR, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	if len(posts) == 0 {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil, nil
	}

	postIdList := make([]int64, 0, len(posts))
	postCategoryIdList := make([]int64, 0, len(posts))
	departmentIdList := make([]int64, 0, len(posts))
	for i := 0; i < len(posts); i++ {
		postIdList = append(postIdList, posts[i].PostId)
		postCategoryIdList = append(postCategoryIdList, posts[i].PostCategoryId)
		departmentIdList = append(departmentIdList, posts[i].DepartmentId)
	}

	postCityMap := make([]model.PostCityMap, 0, max)
	isExist, err = post1.GlobalEngine.Options.Orm.Query(-1, -1, &postCityMap, "", []string{"city_id"}, "post_id in ?", postIdList)
	if err != nil || isExist == false {
		klog.Error(errGetsPostInUser, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}
	cityIdList := make([]int64, 0, len(postCityMap))
	for i := 0; i < len(postCityMap); i++ {
		cityIdList = append(cityIdList, postCityMap[i].CityId)
	}

	//查岗位类别信息
	dataPostCategory, err := post1.GlobalEngine.Options.PostCategoryClient.Client.GetPostCategoryById(context.Background(), &post_category.GetPostCategoryByIdRPCRequest{
		PostCategoryId: postCategoryIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if dataPostCategory.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return dataPostCategory.Code, dataPostCategory.Message, nil, nil
	}

	//查部门和城市信息
	dataDepartmentCity, err := post1.GlobalEngine.Options.DepartmentClient.Client.GetDepartmentCityInfoById(context.Background(), &department.GetDepartmentCityInfoByIdRPCRequest{
		DepartmentIdList: departmentIdList,
		CityIdList:       cityIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if dataDepartmentCity.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return dataDepartmentCity.Code, dataDepartmentCity.Message, nil, nil
	}

	for i := 0; i < len(posts); i++ {
		var postCategoryInfo category.PostCategoryAllInformation
		for j := 0; j < len(dataPostCategory.PostCategoryInfoList); j++ {
			if dataPostCategory.PostCategoryInfoList[j].PostCategoryId == posts[i].PostCategoryId {
				postCategoryInfo = category.PostCategoryAllInformation{
					PostCategoryId:         dataPostCategory.PostCategoryInfoList[j].PostCategoryId,
					PostCategoryName:       dataPostCategory.PostCategoryInfoList[j].PostCategoryName,
					PostCategoryParentId:   dataPostCategory.PostCategoryInfoList[j].PostCategoryParentId,
					PostCategoryParentName: dataPostCategory.PostCategoryInfoList[j].PostCategoryParentName,
				}
				break
			}
		}

		var departmentInfo department2.DepartmentInformation
		for j := 0; j < len(dataDepartmentCity.DepartmentInfoList); j++ {
			if dataDepartmentCity.DepartmentInfoList[j].DepartmentId == posts[i].DepartmentId {
				departmentInfo = department2.DepartmentInformation{
					DepartmentId:          dataDepartmentCity.DepartmentInfoList[j].DepartmentId,
					DepartmentName:        dataDepartmentCity.DepartmentInfoList[j].DepartmentName,
					DepartmentDescription: dataDepartmentCity.DepartmentInfoList[j].DepartmentDescription,
				}
				break
			}
		}

		cityInfos := make([]*department2.CityInformation, 0, max)

		for j := 0; j < len(postCityMap); j++ {
			if postCityMap[j].PostId == posts[i].PostId {
				var cityInfo department2.CityInformation
				for k := 0; k < len(dataDepartmentCity.CityInfoList); k++ {
					if dataDepartmentCity.CityInfoList[k].CityId == postCityMap[j].CityId {
						cityInfo = department2.CityInformation{
							CityId:   dataDepartmentCity.CityInfoList[k].CityId,
							CityName: dataDepartmentCity.CityInfoList[k].CityName,
						}
						break
					}
				}
				cityInfos = append(cityInfos, &cityInfo)
			}
		}

		var resultData = post.PostInfo{
			PostId:                  posts[i].PostId,
			PostBrief:               posts[i].PostBrief,
			PostDescription:         posts[i].PostDescription,
			PostRequire:             posts[i].PostRequire,
			IsSchoolRecruitment:     posts[i].IsSchoolRecruitment,
			IsInternship:            posts[i].IsInternship,
			PostCategoryInformation: &postCategoryInfo,
			DepartmentInformation:   &departmentInfo,
			CityInformation:         cityInfos,
		}
		resultDatas = append(resultDatas, &resultData)
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
}

func DeliveryPost(request *post.DeliveryPostRPCRequest) (int32, string, error) {
	if request.IsDelivery == 1 {
		deliveryPostList := make([]model.DeliveryPost, 0, maxDeliveryCount)
		_, err := post1.GlobalEngine.Options.Orm.Query(maxDeliveryCount, -1, &deliveryPostList, "created_at asc",
			[]string{"id"}, "user_id = ? and is_talent_pool = ?", request.UserId, 0)
		if err != nil {
			klog.Error(errDeliveryPost, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}

		if len(deliveryPostList) == maxDeliveryCount {
			klog.Error(errDeliveryOverflow)
			return code.CodeDeliveryOverflow, code.GetMsg(code.CodeDeliveryOverflow), errors.New(errDeliveryOverflow)
		}

		deliveryPostList1 := make([]model.DeliveryPost, 0, 1)
		_, err = post1.GlobalEngine.Options.Orm.Query(1, -1, &deliveryPostList1, "created_at asc",
			[]string{"id"}, "user_id = ? and is_talent_pool = ?", request.UserId, 1)
		if err != nil {
			klog.Error(errDeliveryPost, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}

		if len(deliveryPostList1) == 1 {
			klog.Error(errPostProcessing)
			return code.CodePostProcessing, code.GetMsg(code.CodePostProcessing), errors.New(errPostProcessing)
		}

		var postDelivery = model.DeliveryPost{
			Id:     post1.GlobalEngine.Options.Id.GenId(),
			UserId: request.UserId,
			PostId: request.PostId,
		}

		isCreate, err := post1.GlobalEngine.Options.Orm.Create(&postDelivery)
		if err != nil || isCreate == false {
			klog.Error(errDeliveryPost, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}

		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
	} else {
		isDelete, err := post1.GlobalEngine.Options.Orm.Delete(&model.DeliveryPost{}, "user_id = ? and post_id = ? and resume_status = ?",
			request.UserId, request.PostId, 0)
		if err != nil {
			klog.Error(errDeliveryPost, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}
		if isDelete == false {
			klog.Error(errDeliveryPostProcessing)
			return code.CodeDeliveryPostProcessing, code.GetMsg(code.CodeDeliveryPostProcessing), errors.New(errDeliveryPostProcessing)
		}

		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
	}
}

func CollectPost(request *post.CollectPostRPCRequest) (int32, string, error) {
	if request.IsCollect == 1 {
		var collectPost = model.CollectPost{
			Id:     post1.GlobalEngine.Options.Id.GenId(),
			UserId: request.UserId,
			PostId: request.PostId,
		}

		isCreate, err := post1.GlobalEngine.Options.Orm.Create(&collectPost)
		if err != nil || isCreate == false {
			klog.Error(errCollectPost, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}

		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
	} else {
		isDelete, err := post1.GlobalEngine.Options.Orm.Delete(&model.CollectPost{}, "user_id = ? and post_id = ?",
			request.UserId, request.PostId)
		if err != nil || isDelete == false {
			klog.Error(errCollectPost, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), err
		}

		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
	}
}

func GetDeliveryPost(request *post.GetDeliveryPostRPCRequest) (int32, string, []*post.GetDeliveryPostRPCData, error) {
	resultDatas := make([]*post.GetDeliveryPostRPCData, 0, maxDeliveryCount)

	deliveryPost := make([]model.DeliveryPost, 0, maxDeliveryCount)
	isExist, err := post1.GlobalEngine.Options.Orm.Query(-1, -1, &deliveryPost, "created_at asc", []string{"post_" +
		"id", "resume_status"}, "user_id = ? and is_talent_pool = ?", request.UserId, 0)
	if err != nil {
		klog.Error(errGetDeliveryPost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	if isExist == false {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil, err
	}

	posts := make([]model.Post, 0, len(deliveryPost))
	postIdList := make([]int64, 0, len(deliveryPost))
	for i := 0; i < len(deliveryPost); i++ {
		postIdList = append(postIdList, deliveryPost[i].PostId)
	}
	_, err = post1.GlobalEngine.Options.Orm.Query(-1, -1, &posts, "created_at asc",
		[]string{"post_id", "post_brief", "post_description", "post_require", "is_school_recruitment", "is_internship", "post_category_id", "department_id"},
		"post_id in ?", postIdList)
	if err != nil {
		klog.Error(errGetDeliveryPost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	postsCity := make([]model.PostCityMap, 0, max)
	_, err = post1.GlobalEngine.Options.Orm.Query(-1, -1, &postsCity, "", []string{"post_id", "city_id"}, "post_id in ?", postIdList)
	if err != nil {
		klog.Error(errGetDeliveryPost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	cityIdList := make([]int64, 0, max)
	departmentIdList := make([]int64, 0, len(posts))
	postCategoryIdList := make([]int64, 0, len(posts))
	for i := 0; i < len(posts); i++ {
		departmentIdList = append(departmentIdList, posts[i].DepartmentId)
		postCategoryIdList = append(postCategoryIdList, posts[i].PostCategoryId)
	}
	for i := 0; i < len(postsCity); i++ {
		cityIdList = append(cityIdList, postsCity[i].CityId)
	}

	//查岗位类别信息
	dataPostCategory, err := post1.GlobalEngine.Options.PostCategoryClient.Client.GetPostCategoryById(context.Background(), &post_category.GetPostCategoryByIdRPCRequest{
		PostCategoryId: postCategoryIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if dataPostCategory.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return dataPostCategory.Code, dataPostCategory.Message, nil, nil
	}

	//查部门和城市信息
	dataDepartmentCity, err := post1.GlobalEngine.Options.DepartmentClient.Client.GetDepartmentCityInfoById(context.Background(), &department.GetDepartmentCityInfoByIdRPCRequest{
		DepartmentIdList: departmentIdList,
		CityIdList:       cityIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if dataDepartmentCity.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return dataDepartmentCity.Code, dataDepartmentCity.Message, nil, nil
	}

	for i := 0; i < len(posts); i++ {
		var postCategoryInfo category.PostCategoryAllInformation
		for j := 0; j < len(dataPostCategory.PostCategoryInfoList); j++ {
			if dataPostCategory.PostCategoryInfoList[j].PostCategoryId == posts[i].PostCategoryId {
				postCategoryInfo = category.PostCategoryAllInformation{
					PostCategoryId:         dataPostCategory.PostCategoryInfoList[j].PostCategoryId,
					PostCategoryName:       dataPostCategory.PostCategoryInfoList[j].PostCategoryName,
					PostCategoryParentId:   dataPostCategory.PostCategoryInfoList[j].PostCategoryParentId,
					PostCategoryParentName: dataPostCategory.PostCategoryInfoList[j].PostCategoryParentName,
				}
				break
			}
		}

		var departmentInfo department2.DepartmentInformation
		for j := 0; j < len(dataDepartmentCity.DepartmentInfoList); j++ {
			if dataDepartmentCity.DepartmentInfoList[j].DepartmentId == posts[i].DepartmentId {
				departmentInfo = department2.DepartmentInformation{
					DepartmentId:          dataDepartmentCity.DepartmentInfoList[j].DepartmentId,
					DepartmentName:        dataDepartmentCity.DepartmentInfoList[j].DepartmentName,
					DepartmentDescription: dataDepartmentCity.DepartmentInfoList[j].DepartmentDescription,
				}
				break
			}
		}

		cityInfos := make([]*department2.CityInformation, 0, max)
		for j := 0; j < len(postsCity); j++ {
			if postsCity[j].PostId == posts[i].PostId {
				var cityInfo department2.CityInformation
				for k := 0; k < len(dataDepartmentCity.CityInfoList); k++ {
					if dataDepartmentCity.CityInfoList[k].CityId == postsCity[j].CityId {
						cityInfo = department2.CityInformation{
							CityId:   dataDepartmentCity.CityInfoList[k].CityId,
							CityName: dataDepartmentCity.CityInfoList[k].CityName,
						}
						break
					}
				}
				cityInfos = append(cityInfos, &cityInfo)
			}
		}

		var postInfo = post.PostInfo{
			PostId:                  posts[i].PostId,
			PostBrief:               posts[i].PostBrief,
			PostDescription:         posts[i].PostDescription,
			PostRequire:             posts[i].PostRequire,
			IsSchoolRecruitment:     posts[i].IsSchoolRecruitment,
			IsInternship:            posts[i].IsInternship,
			PostCategoryInformation: &postCategoryInfo,
			DepartmentInformation:   &departmentInfo,
			CityInformation:         cityInfos,
		}
		var resultData = post.GetDeliveryPostRPCData{
			ResumeStatus:    resumeStatusMap[deliveryPost[i].ResumeStatus],
			PostInformation: &postInfo,
		}
		resultDatas = append(resultDatas, &resultData)
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
}

func GetCollectPost(request *post.GetCollectPostRPCRequest) (int32, string, []*post.PostInfo, error) {
	resultDatas := make([]*post.PostInfo, 0, maxLength)

	collectPost := make([]model.CollectPost, 0, max)
	isExist, err := post1.GlobalEngine.Options.Orm.Query(-1, -1, &collectPost, "created_at desc", []string{
		"post_id"}, "user_id = ?", request.UserId)
	if err != nil {
		klog.Error(errGetCollectPost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	if isExist == false {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil, err
	}

	posts := make([]model.Post, 0, len(collectPost))
	postIdList := make([]int64, 0, len(collectPost))
	for i := 0; i < len(collectPost); i++ {
		postIdList = append(postIdList, collectPost[i].PostId)
	}
	_, err = post1.GlobalEngine.Options.Orm.Query(-1, -1, &posts, "created_at asc",
		[]string{"post_id", "post_brief", "post_description", "post_require", "is_school_recruitment", "is_internship", "post_category_id", "department_id"},
		"post_id in ?", postIdList)
	if err != nil {
		klog.Error(errGetCollectPost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	postsCity := make([]model.PostCityMap, 0)
	_, err = post1.GlobalEngine.Options.Orm.Query(-1, -1, &postsCity, "", []string{"post_id", "city_id"}, "post_id in ?", postIdList)
	if err != nil {
		klog.Error(errGetDeliveryPost, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	cityIdList := make([]int64, 0)
	departmentIdList := make([]int64, 0, len(posts))
	postCategoryIdList := make([]int64, 0, len(posts))
	for i := 0; i < len(posts); i++ {
		departmentIdList = append(departmentIdList, posts[i].DepartmentId)
		postCategoryIdList = append(postCategoryIdList, posts[i].PostCategoryId)
	}
	for i := 0; i < len(postsCity); i++ {
		cityIdList = append(cityIdList, postsCity[i].CityId)
	}

	//查岗位类别信息
	dataPostCategory, err := post1.GlobalEngine.Options.PostCategoryClient.Client.GetPostCategoryById(context.Background(), &post_category.GetPostCategoryByIdRPCRequest{
		PostCategoryId: postCategoryIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if dataPostCategory.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return dataPostCategory.Code, dataPostCategory.Message, nil, nil
	}

	//查部门和城市信息
	dataDepartmentCity, err := post1.GlobalEngine.Options.DepartmentClient.Client.GetDepartmentCityInfoById(context.Background(), &department.GetDepartmentCityInfoByIdRPCRequest{
		DepartmentIdList: departmentIdList,
		CityIdList:       cityIdList,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if dataDepartmentCity.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return dataDepartmentCity.Code, dataDepartmentCity.Message, nil, nil
	}

	for i := 0; i < len(posts); i++ {
		var postCategoryInfo category.PostCategoryAllInformation
		for j := 0; j < len(dataPostCategory.PostCategoryInfoList); j++ {
			if dataPostCategory.PostCategoryInfoList[j].PostCategoryId == posts[i].PostCategoryId {
				postCategoryInfo = category.PostCategoryAllInformation{
					PostCategoryId:         dataPostCategory.PostCategoryInfoList[j].PostCategoryId,
					PostCategoryName:       dataPostCategory.PostCategoryInfoList[j].PostCategoryName,
					PostCategoryParentId:   dataPostCategory.PostCategoryInfoList[j].PostCategoryParentId,
					PostCategoryParentName: dataPostCategory.PostCategoryInfoList[j].PostCategoryParentName,
				}
				break
			}
		}

		var departmentInfo department2.DepartmentInformation
		for j := 0; j < len(dataDepartmentCity.DepartmentInfoList); j++ {
			if dataDepartmentCity.DepartmentInfoList[j].DepartmentId == posts[i].DepartmentId {
				departmentInfo = department2.DepartmentInformation{
					DepartmentId:          dataDepartmentCity.DepartmentInfoList[j].DepartmentId,
					DepartmentName:        dataDepartmentCity.DepartmentInfoList[j].DepartmentName,
					DepartmentDescription: dataDepartmentCity.DepartmentInfoList[j].DepartmentDescription,
				}
				break
			}
		}

		cityInfos := make([]*department2.CityInformation, 0, max)
		for j := 0; j < len(postsCity); j++ {
			if postsCity[j].PostId == posts[i].PostId {
				var cityInfo department2.CityInformation
				for k := 0; k < len(dataDepartmentCity.CityInfoList); k++ {
					if dataDepartmentCity.CityInfoList[k].CityId == postsCity[j].CityId {
						cityInfo = department2.CityInformation{
							CityId:   dataDepartmentCity.CityInfoList[k].CityId,
							CityName: dataDepartmentCity.CityInfoList[k].CityName,
						}
						break
					}
				}
				cityInfos = append(cityInfos, &cityInfo)
			}
		}

		var resultData = post.PostInfo{
			PostId:                  posts[i].PostId,
			PostBrief:               posts[i].PostBrief,
			PostDescription:         posts[i].PostDescription,
			PostRequire:             posts[i].PostRequire,
			IsSchoolRecruitment:     posts[i].IsSchoolRecruitment,
			IsInternship:            posts[i].IsInternship,
			PostCategoryInformation: &postCategoryInfo,
			DepartmentInformation:   &departmentInfo,
			CityInformation:         cityInfos,
		}
		resultDatas = append(resultDatas, &resultData)
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
}

func UpdateResumeStatus(request *post.UpdateResumeStatusRPCRequest) (int32, string, error) {
	var updateMap map[string]interface{}
	if request.ResumeOperate == 2 {
		updateMap = map[string]interface{}{
			"user_id":        request.UserId,
			"post_id":        request.PostId,
			"resume_status":  2,
			"is_talent_pool": 1,
		}
	} else if request.ResumeOperate == 1 {
		updateMap = map[string]interface{}{
			"user_id":       request.UserId,
			"post_id":       request.PostId,
			"resume_status": 1,
		}
	} else {
		updateMap = map[string]interface{}{
			"user_id":       request.UserId,
			"post_id":       request.PostId,
			"resume_status": 3,
		}
	}

	isUpdate, err := post1.GlobalEngine.Options.Orm.Update(&model.DeliveryPost{}, updateMap, "user_id = ? and post_id = ?", request.UserId, request.PostId)
	if err != nil || isUpdate == false {
		klog.Error(errUpdateResumeStatus, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func GetUserIdByPostId(request *post.GetUserIdByPostIdRPCRequest) (int32, string, []int64, error) {
	posts := make([]model.DeliveryPost, 0, maxLength)
	if request.IsTalentPool == 1 {
		_, err := post1.GlobalEngine.Options.Orm.Query(int(request.Limit), int(request.Offset), &posts, "created_at asc", []string{"user_id"}, "is_talent_pool = ?", 1)
		if err != nil {
			klog.Error(errGetUserIdByPostId, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}
	} else {
		_, err := post1.GlobalEngine.Options.Orm.Query(int(request.Limit), int(request.Offset), &posts, "created_at asc", []string{"user_id"}, "post_id = ? and is_talent_pool = ?", request.PostId, 0)
		if err != nil {
			klog.Error(errGetUserIdByPostId, err)
			return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
		}
	}

	userIdList := make([]int64, 0, len(posts))
	for i := 0; i < len(posts); i++ {
		userIdList = append(userIdList, posts[i].UserId)
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), userIdList, nil
}

func DeleteResumeRelativeInfoByHRId(request *post.DeleteResumeRelativeInfoByHRIdRPCRequest) (int32, string, error) {
	posts := make([]model.Post, 0, max)
	_, err := post1.GlobalEngine.Options.Orm.Query(-1, -1, &posts, "created_at asc", []string{"post_id"}, "hr_id = ?", request.HrId)
	if err != nil {
		klog.Error(errDeleteResumeRelativeInfoByHRId, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	if len(posts) == 0 {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
	}

	postIdList := make([]int64, 0, len(posts))
	for i := 0; i < len(posts); i++ {
		postIdList = append(postIdList, posts[i].PostId)
	}

	var deleteValue = [][]interface{}{
		{&model.Post{}, "post_id in ?", postIdList},
		{&model.CollectPost{}, "post_id in ?", postIdList},
		{&model.DeliveryPost{}, "post_id in ?", postIdList},
		{&model.PostCityMap{}, "post_id in ?", postIdList}}

	if err = post1.GlobalEngine.Options.Orm.Transaction([]string{pkg.DbFunctionDelete, pkg.DbFunctionDelete, pkg.DbFunctionDelete, pkg.DbFunctionDelete}, deleteValue); err != nil {
		klog.Error(errDeleteResumeRelativeInfoByHRId, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func DeleteResumeRelativeInfoByDepartmentId(request *post.DeleteResumeRelativeInfoByDepartmentIdRPCRequest) (int32, string, error) {
	posts := make([]model.Post, 0, max)
	_, err := post1.GlobalEngine.Options.Orm.Query(-1, -1, &posts, "created_at asc", []string{"post_id"}, "department_id = ?", request.DepartmentId)
	if err != nil {
		klog.Error(errDeleteResumeRelativeInfoByHRId, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	if len(posts) == 0 {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
	}

	postIdList := make([]int64, 0, len(posts))
	for i := 0; i < len(posts); i++ {
		postIdList = append(postIdList, posts[i].PostId)
	}

	var deleteValue = [][]interface{}{
		{&model.Post{}, "post_id in ?", postIdList},
		{&model.CollectPost{}, "post_id in ?", postIdList},
		{&model.DeliveryPost{}, "post_id in ?", postIdList},
		{&model.PostCityMap{}, "post_id in ?", postIdList}}

	if err = post1.GlobalEngine.Options.Orm.Transaction([]string{pkg.DbFunctionDelete, pkg.DbFunctionDelete, pkg.DbFunctionDelete, pkg.DbFunctionDelete}, deleteValue); err != nil {
		klog.Error(errDeleteResumeRelativeInfoByHRId, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func DeleteResumeRelativeInfoByPostCategoryIdList(request *post.DeleteResumeRelativeInfoByPostCategoryIdListRPCRequest) (int32, string, error) {
	posts := make([]model.Post, 0, max)
	_, err := post1.GlobalEngine.Options.Orm.Query(-1, -1, &posts, "created_at asc", []string{"post_id"}, "post_category_id in ?", request.PostCategoryIdList)
	if err != nil {
		klog.Error(errDeleteResumeRelativeInfoByHRId, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	if len(posts) == 0 {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
	}

	postIdList := make([]int64, 0, len(posts))
	for i := 0; i < len(posts); i++ {
		postIdList = append(postIdList, posts[i].PostId)
	}

	var deleteValue = [][]interface{}{
		{&model.Post{}, "post_id in ?", postIdList},
		{&model.CollectPost{}, "post_id in ?", postIdList},
		{&model.DeliveryPost{}, "post_id in ?", postIdList},
		{&model.PostCityMap{}, "post_id in ?", postIdList}}

	if err = post1.GlobalEngine.Options.Orm.Transaction([]string{pkg.DbFunctionDelete, pkg.DbFunctionDelete, pkg.DbFunctionDelete, pkg.DbFunctionDelete}, deleteValue); err != nil {
		klog.Error(errDeleteResumeRelativeInfoByHRId, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

func DeletePostCityByCityId(request *post.DeletePostCityByCityIdRPCRequest) (int32, string, error) {
	isDelete, err := post1.GlobalEngine.Options.Orm.Delete(&model.PostCityMap{}, "city_id = ?", request.CityId)
	if err != nil || isDelete == false {
		klog.Error(errDeletePostCityByCityId, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}
