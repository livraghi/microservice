package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

func LoadConfigurations(opts ...Option) (*MicroserviceConfiguration, error) {
	cfg := newConfigurationConfig(opts...)

	viper.SetDefault("name", cfg.AppName)
	viper.SetDefault("version", cfg.AppVersion)
	viper.SetDefault("revision", cfg.AppRevision)
	viper.SetDefault("environment", cfg.AppEnvironment)
	viper.SetDefault("port", DefaultPort)
	viper.SetDefault("graceful_shutdown_timeout", DefaultGracefulTimeout)
	viper.SetDefault("k_service", "")
	viper.SetDefault("k_revision", "")
	viper.SetDefault("k_configuration", "")

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

	microserviceCfg := &MicroserviceConfiguration{}
	err = viper.Unmarshal(microserviceCfg)
	if err != nil {
		return nil, err
	}

	kNativeCfg := &kNativeConfiguration{}
	err = viper.Unmarshal(kNativeCfg)
	if err != nil {
		return nil, err
	}
	microserviceCfg.applyKNative(kNativeCfg)

	//info, ok := debug.ReadBuildInfo()
	//if ok {
	//	microserviceCfg.Version = info.Main.Version
	//}

	return microserviceCfg, nil
}

type MicroserviceConfiguration struct {
	Name                    string        `mapstructure:"name"`
	Version                 string        `mapstructure:"version"`
	Revision                int           `mapstructure:"revision"`
	Environment             string        `mapstructure:"environment"`
	Port                    int32         `mapstructure:"port"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type kNativeConfiguration struct {
	Port          int32  `mapstructure:"port"`
	Service       string `mapstructure:"k_service"`
	Revision      string `mapstructure:"k_revision"`
	Configuration string `mapstructure:"k_configuration"`
}

func (cfg *MicroserviceConfiguration) applyKNative(kNativeCfg *kNativeConfiguration) {
	if kNativeCfg.Port > 0 {
		cfg.Port = kNativeCfg.Port
	}
	if kNativeCfg.Service != "" {
		cfg.Name = kNativeCfg.Service
	}
	if kNativeCfg.Revision != "" {
		rev, err := func(value string) (int, error) {
			revision, _ := strings.CutPrefix(value, fmt.Sprintf("%s-", kNativeCfg.Service))
			return strconv.Atoi(revision)
		}(kNativeCfg.Revision)
		if err == nil {
			cfg.Revision = rev
		}
	}
	if kNativeCfg.Configuration != "" {
	}
}

const (
	DefaultPort            = 8080
	DefaultGracefulTimeout = 2 * time.Second
)
