package main

import (
	"log"
	user "resume-resolving/api/idl/service/user/kitex_gen/user/userrpcservice"
)

func main() {
	svr := user.NewServer(new(UserRPCServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
