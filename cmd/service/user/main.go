package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"
	user1 "resume-resolving/api/idl/service/user/kitex_gen/user/userrpcservice"
	"resume-resolving/internal/app/service/user"
	"resume-resolving/internal/app/service/user/handler"
)

func main() {
	engine := user.NewEngine()
	defer engine.Close()
	engine.Init()
	register, err := engine.Options.Registry.Register()
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":8805")
	serve := user1.NewServer(new(handler.UserRPCServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: engine.Config.ConfigInNacos.Server.Name}),
		server.WithRegistry(register),
		server.WithLimit(&limit.Option{
			MaxConnections: engine.Config.ConfigInNacos.Server.MaxConnection,
			MaxQPS:         engine.Config.ConfigInNacos.Server.MaxQps}))
	if err = serve.Run(); err != nil {
		panic(err)
	}
}
