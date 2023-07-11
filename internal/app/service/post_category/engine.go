package post_category

import (
	"resume-resolving/internal/app/service/post_category/config"
	"sync"
	"sync/atomic"
)

//GlobalEngine is the global engine variable in user service
var GlobalEngine = &Engine{}

const (
	_ uint32 = iota
	statusInitialized
	statusClosed
)

const (
	errInit  = "不要重复创建"
	errClose = "还没有初始化或已经关闭过或没有运行，不能关闭"
)

type Engine struct {
	status uint32
	mutex  sync.Mutex
	//配置文件
	Config *config.Config
	//变量信息
	Options *Options
}

func NewEngine(opts ...Option) *Engine {
	c := config.NewConfig()
	GlobalEngine = &Engine{
		status:  0,
		mutex:   sync.Mutex{},
		Config:  c,
		Options: NewOptions(c, opts...),
	}
	return GlobalEngine
}

func (Engine *Engine) Init() {
	Engine.mutex.Lock()
	defer Engine.mutex.Unlock()
	if atomic.LoadUint32(&Engine.status) != 0 {
		panic(errInit)
	}
	if err := Engine.Config.Init(); err != nil {
		panic(err)
	}
	if err := Engine.Options.Init(); err != nil {
		panic(err)
	}
	atomic.StoreUint32(&Engine.status, statusInitialized)
}

func (Engine *Engine) Close() {
	Engine.mutex.Lock()
	defer Engine.mutex.Unlock()
	if atomic.LoadUint32(&Engine.status) != statusInitialized {
		panic(errClose)
	}
	if err := Engine.Options.Close(); err != nil {
		panic(err)
	}
	atomic.StoreUint32(&Engine.status, statusClosed)
}
