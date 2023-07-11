package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"
	resume1 "resume-resolving/api/idl/service/resume/kitex_gen/resume/resumerpcservice"
	"resume-resolving/internal/app/service/resume"
	"resume-resolving/internal/app/service/resume/handler"
)

func main() {
	engine := resume.NewEngine()
	defer engine.Close()
	engine.Init()
	register, err := engine.Options.Registry.Register()
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":8804")
	serve := resume1.NewServer(new(handler.ResumeRPCServiceImpl),
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
