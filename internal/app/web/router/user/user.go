// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "resume-resolving/internal/app/web/handler/user"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			{
				_user := _v1.Group("/user", _userMw()...)
				_user.POST("/change-password", append(_userchangepasswordMw(), user.UserChangePassword)...)
				_user.GET("/get", append(_getuserbynameMw(), user.GetUserByName)...)
				_user.GET("/get-hr", append(_gethrbynameMw(), user.GetHRByName)...)
				_user.GET("/gets-department", append(_gethrbydepartmentMw(), user.GetHRByDepartment)...)
				_user.POST("/hr-delete", append(_hrdeleteMw(), user.HRDelete)...)
				_user.POST("/hr-register", append(_hrregisterMw(), user.HRRegister)...)
				_user.POST("/login", append(_userloginMw(), user.UserLogin)...)
				_user.POST("/register", append(_userregisterMw(), user.UserRegister)...)
			}
		}
	}
}
