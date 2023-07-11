package web

import (
	"resume-resolving/internal/app/web/config"
	rpc2 "resume-resolving/internal/app/web/pkg/client/rpc"
	"resume-resolving/internal/app/web/pkg/jwt"
	"resume-resolving/internal/app/web/pkg/oss"
	"resume-resolving/internal/app/web/pkg/resolve"
	logger2 "resume-resolving/internal/pkg/logger"
	oss2 "resume-resolving/internal/pkg/oss"
	resolve2 "resume-resolving/internal/pkg/resolve"
)

type Option func(options *Options)

type Options struct {
	Jwt                *jwt.JWT
	UserClient         *rpc2.UserClient
	DepartmentClient   *rpc2.DepartmentClient
	PostClient         *rpc2.PostClient
	PostCategoryClient *rpc2.PostCategoryClient
	ResumeClient       *rpc2.ResumeClient

	Oss oss2.OSS

	Logger   logger2.Logger
	Resolver resolve2.Resolver
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

	resolver, err := options.Resolver.Resolve()
	if err != nil {
		return err
	}

	if err = options.Oss.Init(); err != nil {
		return
	}

	if err = options.UserClient.Init(resolver); err != nil {
		return err
	}

	if err = options.DepartmentClient.Init(resolver); err != nil {
		return err
	}

	if err = options.PostClient.Init(resolver); err != nil {
		return err
	}

	if err = options.PostCategoryClient.Init(resolver); err != nil {
		return err
	}

	if err = options.ResumeClient.Init(resolver); err != nil {
		return err
	}

	if err = options.Jwt.Init(options.UserClient); err != nil {
		return err
	}

	return nil
}

func (options *Options) Close() (err error) {
	return nil
}

func NewOptions(config *config.Config, opts ...Option) *Options {
	options := &Options{
		//Logger:             logger.NewHertzLogrus(config),
		Resolver:           resolve.NewEtcdResolver(config),
		Oss:                oss.NewAliYunOSS(config),
		UserClient:         rpc2.NewUserClient(config),
		DepartmentClient:   rpc2.NewDepartmentClient(config),
		ResumeClient:       rpc2.NewResumeClient(config),
		PostCategoryClient: rpc2.NewPostCategoryClient(config),
		PostClient:         rpc2.NewPostClient(config),
		Jwt:                jwt.NewJWT(config),
	}
	options.apply(opts...)
	return options
}

func (options *Options) WithLogger(logger logger2.Logger) Option {
	return func(options *Options) {
		options.Logger = logger
	}
}

func (options *Options) WithResolver(resolver resolve2.Resolver) Option {
	return func(options *Options) {
		options.Resolver = resolver
	}
}
