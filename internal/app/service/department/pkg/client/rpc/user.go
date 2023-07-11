package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"resume-resolving/api/idl/service/user/kitex_gen/user/userrpcservice"
	"resume-resolving/internal/app/service/department/config"
	"time"
)

const (
	rpcTimeout     = 3 * time.Second
	connectTimeout = 50 * time.Millisecond
)

type UserClient struct {
	Client userrpcservice.Client
	config *config.Config
}

func (UserClient *UserClient) Init(r discovery.Resolver) (err error) {
	UserClient.Client, err = userrpcservice.NewClient(
		UserClient.config.ConfigInNacos.RequestGRPCServer.UserService.Name,
		client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()),
		//RPC超时
		client.WithRPCTimeout(rpcTimeout),
		//连接超时
		client.WithConnectTimeout(connectTimeout),
	)
	return
}

func NewUserClient(config *config.Config) *UserClient {
	return &UserClient{
		config: config,
	}
}
