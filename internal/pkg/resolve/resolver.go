package resolve

import (
	"github.com/cloudwego/kitex/pkg/discovery"
)

type Resolver interface {
	Resolve() (discovery.Resolver, error)
}
