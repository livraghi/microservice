package microservice

import (
	"github.com/livraghi/microservice/configuration"
	"time"
)

type serverConfig struct {
	name                    string
	version                 string
	revision                int
	environment             string
	port                    int32
	gracefulShutdownTimeout time.Duration
}

func newServerConfig(opts ...ServerOption) *serverConfig {
	cfg := &serverConfig{
		name:                    configuration.DefaultAppName,
		version:                 configuration.DefaultAppVersion,
		revision:                configuration.DefaultAppRevision,
		environment:             configuration.DefaultAppEnvironment,
		port:                    configuration.DefaultPort,
		gracefulShutdownTimeout: configuration.DefaultGracefulTimeout,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

type ServerOption func(cfg *serverConfig)

func WithConfiguration(microserviceConfiguration *configuration.MicroserviceConfiguration) ServerOption {
	return func(cfg *serverConfig) {
		if cfg != nil {
			WithName(microserviceConfiguration.Name)(cfg)
			WithVersion(microserviceConfiguration.Version)(cfg)
			WithRevision(microserviceConfiguration.Revision)(cfg)
			WithEnvironment(microserviceConfiguration.Environment)(cfg)
			WithPort(microserviceConfiguration.Port)(cfg)
			WithGracefulShutdownTimeout(microserviceConfiguration.GracefulShutdownTimeout)(cfg)
		}
	}
}

func WithName(name string) ServerOption {
	return func(cfg *serverConfig) {
		if name != "" {
			cfg.name = name
		}
	}
}

func WithVersion(version string) ServerOption {
	return func(cfg *serverConfig) {
		if version != "" {
			cfg.version = version
		}
	}
}

func WithRevision(revision int) ServerOption {
	return func(cfg *serverConfig) {
		if revision > 0 {
			cfg.revision = revision
		}
	}
}

func WithEnvironment(environment string) ServerOption {
	return func(cfg *serverConfig) {
		if environment != "" {
			cfg.environment = environment
		}
	}
}

func WithPort(port int32) ServerOption {
	return func(cfg *serverConfig) {
		if port > 0 {
			cfg.port = port
		}
	}
}

func WithGracefulShutdownTimeout(timeout time.Duration) ServerOption {
	return func(cfg *serverConfig) {
		if timeout >= 0 {
			cfg.gracefulShutdownTimeout = timeout
		}
	}
}
