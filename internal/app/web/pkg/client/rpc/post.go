package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"resume-resolving/api/idl/service/post/kitex_gen/post/postrpcservice"
	"resume-resolving/internal/app/web/config"
)

type PostClient struct {
	Client postrpcservice.Client
	config *config.Config
}

func (postClient *PostClient) Init(r discovery.Resolver) (err error) {
	postClient.Client, err = postrpcservice.NewClient(
		postClient.config.ConfigInNacos.RequestGRPCServer.PostService.Name,
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		//RPC超时
		client.WithRPCTimeout(rpcTimeout),
		//连接超时
		client.WithConnectTimeout(connectTimeout))
	return
}

func NewPostClient(config *config.Config) *PostClient {
	return &PostClient{
		config: config,
	}
}
