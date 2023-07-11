package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	hertz := server.Default()
	hertz.Spin()
}
