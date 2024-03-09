package configuration

type configurationConfig struct {
	Path []string
	Name string
	Type ConfigType
}

type ConfigType string

type Option func(config *configurationConfig)

func newConfigurationConfig(opts ...Option) *configurationConfig {
	cfg := &configurationConfig{
		Path: nil,
		Name: DefaultConfigName,
		Type: DefaultConfigType,
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

const (
	JSON ConfigType = "json"
	ENV  ConfigType = "env"

	DefaultConfigPath = "./config"
	DefaultConfigName = ".env"
	DefaultConfigType = ENV
)
