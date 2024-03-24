package observability

import (
	"time"
)

type tracingConfig struct {
	exporterFactory NewTraceExporter
	exporterOptions []ExporterOption
}

type TracingOption func(config *tracingConfig)

func newTracingConfig(opts ...TracingOption) *tracingConfig {
	cfg := &tracingConfig{
		exporterFactory: newGRpcOtelTraceExporter,
		exporterOptions: nil,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func WithOtelGrpcTraceExporter(opts ...ExporterOption) TracingOption {
	return func(config *tracingConfig) {
		config.exporterFactory = newGRpcOtelTraceExporter
		config.exporterOptions = opts
	}
}

func WithOtelHttpTraceExporter(opts ...ExporterOption) TracingOption {
	return func(config *tracingConfig) {
		config.exporterFactory = newHttpOtelTraceExporter
		config.exporterOptions = opts
	}
}

type exporterConfig struct {
	host     string
	grpcPort int32
	httpPort int32
	timeout  time.Duration
}

type ExporterOption func(config *exporterConfig)

func newExporterConfig(opts ...ExporterOption) *exporterConfig {
	cfg := &exporterConfig{
		host:     defaultExporterHost,
		grpcPort: defaultExporterGrpcPort,
		httpPort: defaultExporterHttpPort,
		timeout:  defaultExporterTimeout,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func WithExporterHost(host string) ExporterOption {
	return func(config *exporterConfig) {
		config.host = host
	}
}

func WithExporterGrpcPort(port int32) ExporterOption {
	return func(config *exporterConfig) {
		config.grpcPort = port
	}
}

func WithExporterHttpPort(port int32) ExporterOption {
	return func(config *exporterConfig) {
		config.httpPort = port
	}
}

func WithExporterTimeout(timeout time.Duration) ExporterOption {
	return func(config *exporterConfig) {
		config.timeout = timeout
	}
}

const (
	defaultExporterHost     = "localhost"
	defaultExporterGrpcPort = 4317
	defaultExporterHttpPort = 4318
	defaultExporterTimeout  = 2 * time.Second
)
