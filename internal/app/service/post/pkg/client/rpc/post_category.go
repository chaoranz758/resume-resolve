package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"resume-resolving/api/idl/service/post_category/kitex_gen/post_category/postcategoryrpcservice"
	"resume-resolving/internal/app/service/post/config"
)

type PostCategoryClient struct {
	Client postcategoryrpcservice.Client
	config *config.Config
}

func (postCategoryClient *PostCategoryClient) Init(r discovery.Resolver) (err error) {
	postCategoryClient.Client, err = postcategoryrpcservice.NewClient(
		postCategoryClient.config.ConfigInNacos.RequestRpcServer.PostCategoryService.Name,
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		//RPC超时
		client.WithRPCTimeout(rpcTimeout),
		//连接超时
		client.WithConnectTimeout(connectTimeout),
	)
	return
}

func NewPostCategoryClient(config *config.Config) *PostCategoryClient {
	return &PostCategoryClient{
		config: config,
	}
}
