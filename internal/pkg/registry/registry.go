package registry

import (
	"github.com/cloudwego/kitex/pkg/registry"
)

type Registry interface {
	Register() (registry.Registry, error)
}
