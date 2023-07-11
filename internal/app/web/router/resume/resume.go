// Code generated by hertz generator. DO NOT EDIT.

package resume

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	resume "resume-resolving/internal/app/web/handler/resume"
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
				_resume := _v1.Group("/resume", _resumeMw()...)
				_resume.GET("/get-id", append(_getresumebyidMw(), resume.GetResumeById)...)
				_resume.GET("/gets-post", append(_getresumebypostMw(), resume.GetResumeByPost)...)
				_resume.POST("/upload", append(_uploadstructresumeMw(), resume.UploadStructResume)...)
				_resume.POST("/upload-file", append(_uploadresumefileMw(), resume.UploadResumeFile)...)
			}
		}
	}
}
