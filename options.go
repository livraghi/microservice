package microservice

import "time"

const (
	defaultPort                    = 8080
	defaultGracefulShutdownTimeout = 0
)

type serverConfig struct {
	port                    int32
	gracefulShutdownTimeout time.Duration
}

func newServerConfig(opts ...ServerOption) *serverConfig {
	cfg := &serverConfig{
		port:                    defaultPort,
		gracefulShutdownTimeout: defaultGracefulShutdownTimeout,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

type ServerOption func(cfg *serverConfig)

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
