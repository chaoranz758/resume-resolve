package main

import (
	"log"
	post "resume-resolving/api/idl/service/post/kitex_gen/post/postrpcservice"
)

func main() {
	svr := post.NewServer(new(PostRPCServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
