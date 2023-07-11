package post_category

import (
	config2 "resume-resolving/internal/app/service/post_category/config"
	cache1 "resume-resolving/internal/app/service/post_category/pkg/cache"
	"resume-resolving/internal/app/service/post_category/pkg/client/rpc"
	database1 "resume-resolving/internal/app/service/post_category/pkg/database"
	"resume-resolving/internal/app/service/post_category/pkg/id"
	"resume-resolving/internal/app/service/post_category/pkg/redsync"
	registry1 "resume-resolving/internal/app/service/post_category/pkg/registry"
	resolve1 "resume-resolving/internal/app/service/post_category/pkg/resolve"
	"resume-resolving/internal/pkg/cache"
	"resume-resolving/internal/pkg/database"
	id2 "resume-resolving/internal/pkg/id"
	"resume-resolving/internal/pkg/logger"
	"resume-resolving/internal/pkg/registry"
	"resume-resolving/internal/pkg/resolve"
)

type Option func(options *Options)

type Options struct {
	PostClient *rpc.PostClient

	Logger   logger.Logger
	Database database.Database
	Orm      database.Orm
	Cache    cache.Cache
	Registry registry.Registry
	Resolver resolve.Resolver
	Id       id2.DistributedId

	RedSync *redsync.RedSync

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

	if err = options.PostClient.Init(resolver); err != nil {
		return err
	}

	if err = options.RedSync.Init(); err != nil {
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
		//Logger:     logger1.NewKitexLogrus(config),
		Database:   database1.NewMysql(config),
		Cache:      cache1.NewRedis(config),
		Registry:   registry1.NewEtcdRegistry(config),
		Resolver:   resolve1.NewEtcdResolver(config),
		Id:         id.NewSnowFlake(config),
		RedSync:    redsync.New(config),
		config:     config,
		PostClient: rpc.NewPostClient(config),
	}
	options.Orm = database1.NewGorm(config, options.Database)
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
