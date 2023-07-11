// Code generated by hertz generator. DO NOT EDIT.

package post_category

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	post_category "resume-resolving/internal/app/web/handler/post_category"
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
				_post_category := _v1.Group("/post-category", _post_categoryMw()...)
				_post_category.POST("/append", append(_appendpostcategoryMw(), post_category.AppendPostCategory)...)
				_post_category.POST("/delete", append(_deletepostcategoryMw(), post_category.DeletePostCategory)...)
				_post_category.GET("/gets", append(_getspostcategoryMw(), post_category.GetsPostCategory)...)
				_post_category.POST("/update", append(_updatepostcategoryMw(), post_category.UpdatePostCategory)...)
			}
		}
	}
}
