// Code generated by hertz generator. DO NOT EDIT.

package department

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	department "resume-resolving/internal/app/web/handler/department"
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
				_department := _v1.Group("/department", _departmentMw()...)
				_department.POST("/append", append(_appenddepartmentMw(), department.AppendDepartment)...)
				_department.POST("/delete", append(_deletedepartmentMw(), department.DeleteDepartment)...)
				_department.GET("/gets", append(_getsdepartmentMw(), department.GetsDepartment)...)
				_department.POST("/update", append(_updatedepartmentMw(), department.UpdateDepartment)...)
				{
					_city := _department.Group("/city", _cityMw()...)
					_city.POST("/append", append(_appendcityMw(), department.AppendCity)...)
					_city.POST("/delete", append(_deletecityMw(), department.DeleteCity)...)
					_city.GET("/gets", append(_getscityMw(), department.GetsCity)...)
					_city.GET("/gets-department", append(_getscitybydepartmentMw(), department.GetsCityByDepartment)...)
				}
			}
		}
	}
}
