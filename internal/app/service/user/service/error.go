package service

import "errors"

var (
	errGetUserByName          = errors.New("get user information by username failed")
	errUserNotExist           = errors.New("username input failed, user not exist")
	errUserExist              = errors.New("username register failed, user exist")
	errPasswordEncrypt        = errors.New("password encrypt failed")
	errPasswordInput          = errors.New("password failed")
	errUpdateUserPassword     = errors.New("update user's password failed")
	errCreateUser             = errors.New("create user failed")
	errDeleteHR               = errors.New("delete hr by hr id failed")
	errGetHRByDepartment      = errors.New("get hr information by department id failed")
	errGetHRByName            = errors.New("get hr by name failed")
	errDeleteHRByDepartmentId = errors.New("delete hr by department id failed")
	errRpcService             = "rpc service error"
	errRpcBizService          = "rpc service biz error"
)
