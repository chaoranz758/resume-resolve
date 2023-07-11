// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
)

// register registers all routers.
func Register(r *server.Hertz) {
	r.Use(cors.Default())

	GeneratedRegister(r)

	customizedRegister(r)
}
