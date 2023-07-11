package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"
	post1 "resume-resolving/api/idl/service/post/kitex_gen/post/postrpcservice"
	"resume-resolving/internal/app/service/post"
	"resume-resolving/internal/app/service/post/handler"
)

func main() {
	engine := post.NewEngine()
	defer engine.Close()
	engine.Init()
	register, err := engine.Options.Registry.Register()
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":8802")
	serve := post1.NewServer(new(handler.PostRPCServiceImpl),
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
