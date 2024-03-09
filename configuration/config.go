package configuration

import (
	"github.com/spf13/viper"
	"runtime/debug"
	"time"
)

const (
	BuildVersion = "0.0.0"
)

func LoadConfigurations(opts ...Option) (*MicroserviceConfiguration, error) {
	cfg := newConfigurationConfig(opts...)

	for _, path := range cfg.Path {
		viper.AddConfigPath(path)
	}
	viper.SetConfigName(cfg.Name)
	viper.SetConfigType(string(cfg.Type))

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	microserviceCfg := &MicroserviceConfiguration{
		Port:                    8080,
		GracefulShutdownTimeout: 2 * time.Second,
	}
	err = viper.Unmarshal(microserviceCfg)
	if err != nil {
		return nil, err
	}

	info, ok := debug.ReadBuildInfo()
	if ok {
		microserviceCfg.Version = info.Main.Version
	}

	return microserviceCfg, nil
}

type MicroserviceConfiguration struct {
	Version                 string
	Port                    int32         `mapstructure:"port"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}
