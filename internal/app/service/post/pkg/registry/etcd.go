package registry

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/registry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"resume-resolving/internal/app/service/post/config"
	registry1 "resume-resolving/internal/pkg/registry"
)

type EtcdRegistry struct {
	config *config.Config
}

func (e *EtcdRegistry) Register() (registry.Registry, error) {
	return etcd.NewEtcdRegistry([]string{fmt.Sprintf("%s:%d",
		e.config.ConfigInNacos.EtcdServer.Ip,
		e.config.ConfigInNacos.EtcdServer.Port)})
}

func NewEtcdRegistry(config *config.Config) registry1.Registry {
	return &EtcdRegistry{
		config: config,
	}
}
