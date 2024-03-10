package configuration

type configurationConfig struct {
	Path           []string
	Name           string
	Type           ConfigType
	AppName        string
	AppVersion     string
	AppRevision    int
	AppEnvironment string
}

type ConfigType string

type Option func(config *configurationConfig)

func newConfigurationConfig(opts ...Option) *configurationConfig {
	cfg := &configurationConfig{
		Path:           nil,
		Name:           DefaultConfigName,
		Type:           DefaultConfigType,
		AppName:        DefaultAppName,
		AppVersion:     DefaultAppVersion,
		AppRevision:    DefaultAppRevision,
		AppEnvironment: DefaultAppEnvironment,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	if len(cfg.Path) == 0 {
		cfg.Path = append(cfg.Path, DefaultConfigPath)
	}
	return cfg
}

func WithConfigPath(path string) Option {
	return func(config *configurationConfig) {
		config.Path = append(config.Path, path)
	}
}

func WithConfigName(name string) Option {
	return func(config *configurationConfig) {
		config.Name = name
	}
}

func WithConfigType(configType ConfigType) Option {
	return func(config *configurationConfig) {
		config.Type = configType
	}
}

func WithAppName(name string) Option {
	return func(config *configurationConfig) {
		config.AppName = name
	}
}

func WithAppVersion(version string) Option {
	return func(config *configurationConfig) {
		config.AppVersion = version
	}
}

func WithAppRevision(revision int) Option {
	return func(config *configurationConfig) {
		config.AppRevision = revision
	}
}

func WithAppEnvironment(environment string) Option {
	return func(config *configurationConfig) {
		config.AppEnvironment = environment
	}
}

const (
	JSON ConfigType = "json"
	ENV  ConfigType = "env"

	DefaultConfigPath     = "./config"
	DefaultConfigName     = ".env"
	DefaultConfigType     = ENV
	DefaultAppName        = "simple-app"
	DefaultAppVersion     = "0.0.0"
	DefaultAppRevision    = 0
	DefaultAppEnvironment = "development"
)
