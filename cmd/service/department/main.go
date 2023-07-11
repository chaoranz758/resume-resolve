package main

import (
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"
	department1 "resume-resolving/api/idl/service/department/kitex_gen/department/departmentrpcservice"
	"resume-resolving/internal/app/service/department"
	"resume-resolving/internal/app/service/department/handler"
)

func main() {
	engine := department.NewEngine()
	defer engine.Close()
	engine.Init()
	register, err := engine.Options.Registry.Register()
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":8801")
	serve := department1.NewServer(new(handler.DepartmentRPCServiceImpl),
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
