package user

import (
	config2 "resume-resolving/internal/app/service/user/config"
	cache2 "resume-resolving/internal/app/service/user/pkg/cache"
	"resume-resolving/internal/app/service/user/pkg/client/rpc"
	database2 "resume-resolving/internal/app/service/user/pkg/database"
	"resume-resolving/internal/app/service/user/pkg/id"
	registry2 "resume-resolving/internal/app/service/user/pkg/registry"
	resolve2 "resume-resolving/internal/app/service/user/pkg/resolve"
	"resume-resolving/internal/pkg/cache"
	"resume-resolving/internal/pkg/database"
	id2 "resume-resolving/internal/pkg/id"
	"resume-resolving/internal/pkg/logger"
	"resume-resolving/internal/pkg/registry"
	"resume-resolving/internal/pkg/resolve"
)

type Option func(options *Options)

type Options struct {
	DepartmentClient *rpc.DepartmentClient
	PostClient       *rpc.PostClient

	Logger   logger.Logger
	Database database.Database
	Orm      database.Orm
	Cache    cache.Cache
	Registry registry.Registry
	Resolver resolve.Resolver
	Id       id2.DistributedId

	config *config2.Config
}

func (options *Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(options)
	}
}

func (options *Options) Init() (err error) {

	//if err = options.Logger.Init(); err != nil {
	//	return err
	//}

	if err = options.Orm.Open(); err != nil {
		return err
	}

	if err = options.Cache.Open(); err != nil {
		return err
	}

	if err = options.Id.Init(); err != nil {
		return
	}

	resolver, err := options.Resolver.Resolve()
	if err != nil {
		return err
	}

	if err = options.DepartmentClient.Init(resolver); err != nil {
		return err
	}

	if err = options.PostClient.Init(resolver); err != nil {
		return err
	}

	return nil
}

func (options *Options) Close() (err error) {
	if err = options.Orm.Close(); err != nil {
		return err
	}
	if err = options.Cache.Close(); err != nil {
		return err
	}
	return nil
}

func NewOptions(config *config2.Config, opts ...Option) *Options {
	options := &Options{
		//Logger:           logger2.NewKitexLogrus(config),
		Database:         database2.NewMysql(config),
		Cache:            cache2.NewRedis(config),
		Registry:         registry2.NewEtcdRegistry(config),
		Resolver:         resolve2.NewEtcdResolver(config),
		Id:               id.NewSnowFlake(config),
		config:           config,
		DepartmentClient: rpc.NewDepartmentClient(config),
		PostClient:       rpc.NewPostClient(config),
	}
	options.Orm = database2.NewGorm(config, options.Database)
	options.apply(opts...)
	return options
}

func (options *Options) WithLogger(logger logger.Logger) Option {
	return func(options *Options) {
		options.Logger = logger
	}
}

func (options *Options) WithDatabase(database database.Database) Option {
	return func(options *Options) {
		options.Database = database
	}
}

func (options *Options) WithOrm(orm database.Orm) Option {
	return func(options *Options) {
		options.Orm = orm
	}
}

func (options *Options) WithCache(cache cache.Cache) Option {
	return func(options *Options) {
		options.Cache = cache
	}
}

func (options *Options) WithRegistry(registry registry.Registry) Option {
	return func(options *Options) {
		options.Registry = registry
	}
}

func (options *Options) WithResolver(resolver resolve.Resolver) Option {
	return func(options *Options) {
		options.Resolver = resolver
	}
}
