package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"resume-resolving/api/idl/service/resume/kitex_gen/resume/resumerpcservice"
	"resume-resolving/internal/app/web/config"
)

type ResumeClient struct {
	Client resumerpcservice.Client
	config *config.Config
}

func (resumeClient *ResumeClient) Init(r discovery.Resolver) (err error) {
	resumeClient.Client, err = resumerpcservice.NewClient(
		resumeClient.config.ConfigInNacos.RequestGRPCServer.ResumeService.Name,
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		//RPC超时
		client.WithRPCTimeout(rpcTimeout),
		//连接超时
		client.WithConnectTimeout(connectTimeout))
	return
}

func NewResumeClient(config *config.Config) *ResumeClient {
	return &ResumeClient{
		config: config,
	}
}
