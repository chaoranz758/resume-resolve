package id

import (
	sf "github.com/bwmarrin/snowflake"
	"resume-resolving/internal/app/service/post_category/config"
	"resume-resolving/internal/pkg/id"
	"time"
)

const (
	standardTimeLayout = "2006-01-02"
)

type SnowFlake struct {
	node   *sf.Node
	config *config.Config
}

func (s *SnowFlake) Init() (err error) {
	var st time.Time
	st, err = time.Parse(standardTimeLayout, s.config.ConfigInNacos.SnowFlake.StartTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	s.node, err = sf.NewNode(s.config.ConfigInNacos.SnowFlake.MachineID)
	return
}

func (s *SnowFlake) GenId() int64 {
	return s.node.Generate().Int64()
}

func NewSnowFlake(config *config.Config) id.DistributedId {
	return &SnowFlake{
		config: config,
	}
}
