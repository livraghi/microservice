package configuration

type configurationConfig struct {
	Path         string
	EnvVariables bool
	Files        []configurationFile
}

type configurationFile struct {
	Type ConfigType
	Name string
}

type ConfigType string

const (
	Env        ConfigType = "env"
	Json       ConfigType = "json"
	Properties ConfigType = "properties"
	Yaml       ConfigType = "yaml"
	Toml       ConfigType = "toml"
	Ini        ConfigType = "ini"
)

type Option func(config *configurationConfig)

func newConfigurationConfig(opts ...Option) *configurationConfig {
	cfg := &configurationConfig{
		Path: "./",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func WithConfigurationsBasePath(basePath string) Option {
	return func(config *configurationConfig) {
		config.Path = basePath
	}
}

func WithEnvVariables() Option {
	return func(config *configurationConfig) {
		config.EnvVariables = true
	}
}

func WithConfigFile(configType ConfigType, name string) Option {
	return func(config *configurationConfig) {
		config.Files = append(config.Files, configurationFile{
			Type: configType,
			Name: name,
		})
	}
}

//
//func WithConfigPath(path string) Option {
//	return func(config *configurationConfig) {
//		config.Path = append(config.Path, path)
//	}
//}
//
//func WithConfigName(name string) Option {
//	return func(config *configurationConfig) {
//		config.Name = name
//	}
//}
//
//func WithConfigType(configType ConfigType) Option {
//	return func(config *configurationConfig) {
//		config.Type = configType
//	}
//}
//
//func WithAppName(name string) Option {
//	return func(config *configurationConfig) {
//		config.AppName = name
//	}
//}
//
//func WithAppVersion(version string) Option {
//	return func(config *configurationConfig) {
//		config.AppVersion = version
//	}
//}
//
//func WithAppRevision(revision int) Option {
//	return func(config *configurationConfig) {
//		config.AppRevision = revision
//	}
//}
//
//func WithAppEnvironment(environment string) Option {
//	return func(config *configurationConfig) {
//		config.AppEnvironment = environment
//	}
//}
//
//const (
//	Json ConfigType = "json"
//	Env  ConfigType = "env"
//
//	DefaultConfigPath     = "./config"
//	DefaultConfigName     = ".env"
//	DefaultConfigType     = Env
//	DefaultAppName        = "simple-app"
//	DefaultAppVersion     = "0.0.0"
//	DefaultAppRevision    = 0
//	DefaultAppEnvironment = "development"
//)
