package resolve

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
	"resume-resolving/internal/app/service/post_category/config"
	"resume-resolving/internal/pkg/resolve"
)

type EtcdResolver struct {
	config *config.Config
}

func (e *EtcdResolver) Resolve() (discovery.Resolver, error) {
	return etcd.NewEtcdResolver([]string{fmt.Sprintf("%s:%d",
		e.config.ConfigInNacos.EtcdServer.Ip,
		e.config.ConfigInNacos.EtcdServer.Port)})
}

func NewEtcdResolver(config *config.Config) resolve.Resolver {
	return &EtcdResolver{
		config: config,
	}
}
