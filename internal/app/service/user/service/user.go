package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"resume-resolving/api/idl/service/department/kitex_gen/department"
	"resume-resolving/api/idl/service/post/kitex_gen/post"
	"resume-resolving/api/idl/service/user/kitex_gen/base"
	"resume-resolving/api/idl/service/user/kitex_gen/user"
	user1 "resume-resolving/internal/app/service/user"
	"resume-resolving/internal/app/service/user/model"
	"resume-resolving/internal/app/service/user/pkg/code"
	"resume-resolving/internal/app/service/user/pkg/encrypt"
	"strings"
)

// UserChangePassword implements the user or hr or admin change password
//   logic
//   1.Query whether the user exists based on the username role and
//   return the encrypted password stored in the database
//   2.Encrypt the old password entered by the user and compare
//   it with the encrypted password in the database
//   3.Encrypt the new password and update it into the database
func UserChangePassword(request *user.UserChangePasswordRPCRequest) (int, string, error) {
	var u model.User

	c, message, err := userLogin(&u, request.Username, request.Password, request.Role)
	if err != nil {
		return c, message, err
	}

	passwordNew, err := encrypt.ScryptPw(request.Password)
	if err != nil {
		klog.Error(errPasswordEncrypt, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	updateMap := map[string]interface{}{
		"password": passwordNew,
	}
	_, err = user1.GlobalEngine.Options.Orm.Update(&model.User{}, updateMap, "username = ? and role = ?", request.Username, request.Role)
	if err != nil {
		klog.Error(errUpdateUserPassword, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}
	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

// UserLogin implements the user or hr or admin login
//   logic
//   1.Query whether the user exists based on the username role and
//   return the encrypted password stored in the database
//   2.Encrypt the old password entered by the user and compare
//   it with the encrypted password in the database
func UserLogin(request *user.UserLoginRPCRequest) (int, string, error) {
	var u model.User

	c, message, err := userLogin(&u, request.Username, request.Password, request.Role)
	if err != nil {
		return c, message, err
	}
	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

// userLogin is the user or hr or admin login's util function
//   logic
//   1.Query whether the user exists based on the username role and
//   return the encrypted password stored in the database
//   2.Encrypt the old password entered by the user and compare
//   it with the encrypted password in the database
func userLogin(u *model.User, username, password string, role int8) (int, string, error) {
	isExist, err := user1.GlobalEngine.Options.Orm.Query(
		1,
		-1,
		u,
		"created_at asc",
		[]string{"password"},
		"username = ? and role = ?",
		username, role)
	if err != nil {
		klog.Error(errGetUserByName, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	if isExist == false {
		klog.Error(errUserNotExist)
		return code.CodeUsernameInputFailed, code.GetMsg(code.CodeUsernameInputFailed), errUserNotExist
	}

	passwordIn, err := encrypt.ScryptPw(password)
	if err != nil {
		klog.Error(errPasswordEncrypt, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	if !strings.EqualFold(passwordIn, u.Password) {
		klog.Error(errPasswordInput, err)
		return code.CodePasswordInputFailed, code.GetMsg(code.CodePasswordInputFailed), err
	}

	return 0, "", nil
}

// UserRegister implements the user register
//   logic
//   1.Query whether a user exists based on their username and role
//   2.If existed, create user to database
func UserRegister(request *user.UserRegisterRPCRequest) (int, string, error) {
	var u model.User

	isExist, err := user1.GlobalEngine.Options.Orm.Query(
		1,
		-1,
		&u,
		"created_at asc",
		//don't need to return the form
		[]string{"username"},
		"username = ? and role = ?",
		request.Username, 0)

	if err != nil {
		klog.Error(errGetUserByName, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	if isExist == true {
		klog.Error(errUserExist)
		return code.CodeUserExist, code.GetMsg(code.CodeUserExist), errUserExist
	}

	userId := user1.GlobalEngine.Options.Id.GenId()
	password, err := encrypt.ScryptPw(request.Password)
	if err != nil {
		klog.Error(errPasswordEncrypt, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	u1 := model.User{
		UserId:   userId,
		Username: request.Username,
		Password: password,
		Role:     0,
	}

	_, err = user1.GlobalEngine.Options.Orm.Create(&u1)
	if err != nil {
		klog.Error(errCreateUser, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

// HRRegister implements the hr register, and return relative information
//   logic
//   1.Query whether a hr exists based on their username and role
//   2.If existed, create hr to database
//   3.Return relative information
func HRRegister(request *user.HRRegisterRPCRequest) (int, string, error) {
	var u model.User

	isExist, err := user1.GlobalEngine.Options.Orm.Query(
		1,
		-1,
		&u,
		"created_at asc",
		[]string{"username"},
		"username = ? and role = ?",
		request.Username, 1)

	if err != nil {
		klog.Error(errGetUserByName, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	if isExist == true {
		klog.Error(errUserExist)
		return code.CodeUserExist, code.GetMsg(code.CodeUserExist), errUserNotExist
	}

	userId := user1.GlobalEngine.Options.Id.GenId()
	password, err := encrypt.ScryptPw(user1.GlobalEngine.Config.ConfigInNacos.Password.Hr)
	if err != nil {
		klog.Error(errPasswordEncrypt, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	u1 := model.User{
		UserId:       userId,
		Username:     request.Username,
		Password:     password,
		Role:         1,
		DepartmentId: request.DepartmentId,
	}

	_, err = user1.GlobalEngine.Options.Orm.Create(&u1)
	if err != nil {
		klog.Error(errCreateUser, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	//部门微服务调用相关信息

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil

}

// HRDelete implements the hr delete by admin
//   logic
//   delete hr by id and role
func HRDelete(request *user.HRDeleteRPCRequest) (int, string, error) {
	isDelete, err := user1.GlobalEngine.Options.Orm.Delete(
		&model.User{},
		"user_id = ? and role = ?",
		request.UserId, 1)

	if err != nil || isDelete == false {
		klog.Error(errDeleteHR, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	response, err := user1.GlobalEngine.Options.PostClient.Client.DeleteResumeRelativeInfoByHRId(context.Background(), &post.DeleteResumeRelativeInfoByHRIdRPCRequest{
		HrId: request.UserId,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil
	}
	if response.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return int(response.Code), response.Message, nil
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}

// GetHRByDepartment implements getting hrs information
//   If department_id = 0, function return the whole hr information in user table by paging technology
//   If department_id != 0, function return the hr information in user table by department_id and paging technology
//   Moreover, if last_time = 0, it means the first query; if last_time != 0, it means is not the
//   first query and the department_id means the last department_id from last time
func GetHRByDepartment(request *user.GetHRByDepartmentRPCRequest) (int, string, []*base.UserInfo, error) {
	limit := int(request.Limit)
	offset := int(request.Offset)

	var err error
	users := make([]model.User, 0, limit)

	if request.DepartmentId == 0 {
		_, err = user1.GlobalEngine.Options.Orm.Query(
			limit,
			offset,
			&users,
			"created_at asc",
			[]string{"user_id", "username", "department_id", "created_at"},
			"role = ?",
			1,
		)
	} else {
		_, err = user1.GlobalEngine.Options.Orm.Query(
			limit,
			offset,
			&users,
			"created_at asc",
			[]string{"user_id", "username", "department_id", "created_at"},
			"department_id = ? and role = ?",
			request.DepartmentId, 1,
		)
	}

	if err != nil {
		klog.Error(errGetHRByDepartment, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	if len(users) == 0 {
		return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil, nil
	}

	departmentId := make([]int64, 0, len(users))
	for i := 0; i < len(users); i++ {
		departmentId = append(departmentId, users[i].DepartmentId)
	}

	//从department_service查出相关信息并整合数据返回
	response, err := user1.GlobalEngine.Options.DepartmentClient.Client.GetDepartmentInfosById(context.Background(), &department.GetDepartmentInfosByIdRPCRequest{
		DepartmentIdList: departmentId,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if response.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return int(response.Code), response.Message, nil, nil
	}

	resultDatas := make([]*base.UserInfo, 0, len(users))
	for i := 0; i < len(users); i++ {
		var resultData = base.UserInfo{
			UserId:         users[i].UserId,
			Username:       users[i].Username,
			DepartmentId:   users[i].DepartmentId,
			DepartmentName: response.DepartmentInfoList[i].DepartmentName,
		}
		resultDatas = append(resultDatas, &resultData)
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), resultDatas, nil
}

func GetUserByName(request *user.GetUserByNameRPCRequest) (int, string, *base.CommonUserInfo, error) {
	var u model.User
	isExist, err := user1.GlobalEngine.Options.Orm.Query(1, -1, &u, "", []string{"user_id", "username"},
		"username = ? and role = ?", request.Username, 0)
	if err != nil || isExist == false {
		klog.Error(errGetUserByName, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	var data = base.CommonUserInfo{
		UserId:   u.UserId,
		Username: u.Username,
	}
	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), &data, nil
}

func GetHRByName(request *user.GetHRByNameRPCRequest) (int, string, *base.UserInfo, error) {
	var u model.User
	isExist, err := user1.GlobalEngine.Options.Orm.Query(1, -1, &u, "", []string{"user_id", "username", "department_id"},
		"username = ? and role = ?", request.Username, 1)
	if err != nil || isExist == false {
		klog.Error(errGetHRByName, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, err
	}

	//从department_service查出相关信息并整合数据返回
	response, err := user1.GlobalEngine.Options.DepartmentClient.Client.GetDepartmentInfosById(context.Background(), &department.GetDepartmentInfosByIdRPCRequest{
		DepartmentIdList: []int64{u.DepartmentId},
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil, nil
	}
	if response.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return int(response.Code), response.Message, nil, nil
	}

	var data = base.UserInfo{
		UserId:         u.UserId,
		Username:       u.Username,
		DepartmentId:   u.DepartmentId,
		DepartmentName: response.DepartmentInfoList[0].DepartmentName,
	}
	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), &data, nil
}

func DeleteHRByDepartmentId(request *user.DeleteHRByDepartmentIdRPCRequest) (int, string, error) {
	_, err := user1.GlobalEngine.Options.Orm.Delete(&model.User{}, "department_id = ? and role = ?", request.DepartmentId, 1)
	if err != nil {
		klog.Error(errDeleteHRByDepartmentId, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), err
	}

	response, err := user1.GlobalEngine.Options.PostClient.Client.DeleteResumeRelativeInfoByDepartmentId(context.Background(), &post.DeleteResumeRelativeInfoByDepartmentIdRPCRequest{
		DepartmentId: request.DepartmentId,
	})
	if err != nil {
		klog.Error(errRpcService, err)
		return code.CodeInternal, code.GetMsg(code.CodeInternal), nil
	}
	if response.Code != code.CodeSuccess {
		klog.Error(errRpcBizService)
		return int(response.Code), response.Message, nil
	}

	return code.CodeSuccess, code.GetMsg(code.CodeSuccess), nil
}
