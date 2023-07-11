package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"resume-resolving/api/idl/service/department/kitex_gen/department/departmentrpcservice"
	"resume-resolving/internal/app/web/config"
	"time"
)

const (
	rpcTimeout     = 3 * time.Second
	connectTimeout = 50 * time.Millisecond
)

type DepartmentClient struct {
	Client departmentrpcservice.Client
	config *config.Config
}

func (departmentClient *DepartmentClient) Init(r discovery.Resolver) (err error) {
	departmentClient.Client, err = departmentrpcservice.NewClient(
		departmentClient.config.ConfigInNacos.RequestGRPCServer.DepartmentService.Name,
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		//RPC超时
		client.WithRPCTimeout(rpcTimeout),
		//连接超时
		client.WithConnectTimeout(connectTimeout),
	)
	return
}

func NewDepartmentClient(config *config.Config) *DepartmentClient {
	return &DepartmentClient{
		config: config,
	}
}
