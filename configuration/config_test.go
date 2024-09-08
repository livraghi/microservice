package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfiguration_Nothing(t *testing.T) {
	setupEnvs(t)

	var cfg configuration
	err := ReadConfiguration(&cfg)

	assert.NoError(t, err, "read configuration should not return an error")
	assert.NotNil(t, cfg, "configuration should not be nil")
	assert.Zero(t, cfg, "configuration should be empty")
}

func TestReadConfiguration_env(t *testing.T) {
	setupEnvs(t)

	var cfg configuration
	err := ReadConfiguration(&cfg, WithEnvVariables())

	assert.NoError(t, err, "read configuration should not return an error")
	assert.NotNil(t, cfg, "configuration should not be nil")
	expected := *newConfigurationSample()
	expected.Map = `{"key1":"value1","key2":"value2"}` // viper does not support map[string]string for env variables
	assert.Equal(t, expected, cfg, "configuration should be equal to the expected configuration")
}

func TestReadConfiguration_envFile(t *testing.T) {
	var cfg configuration
	err := ReadConfiguration(&cfg, WithConfigFile(Env, "unit-test.env"))

	assert.NoError(t, err, "read configuration should not return an error")
	assert.NotNil(t, cfg, "configuration should not be nil")
	expected := *newConfigurationSample()
	expected.Map = `{"key1":"value1","key2":"value2"}` // viper does not support map[string]string for env variables
	assert.Equal(t, expected, cfg, "configuration should be equal to the expected configuration")
}

func TestReadConfiguration_jsonFile(t *testing.T) {
	var cfg configuration
	err := ReadConfiguration(&cfg, WithConfigFile(Json, "unit-test.json"))

	assert.NoError(t, err, "read configuration should not return an error")
	assert.NotNil(t, cfg, "configuration should not be nil")
	assert.Equal(t, *newConfigurationSample(), cfg, "configuration should be equal to the expected configuration")
}

func TestReadConfiguration_propertiesFile(t *testing.T) {
	var cfg configuration
	err := ReadConfiguration(&cfg, WithConfigFile(Properties, "unit-test.properties"))

	assert.NoError(t, err, "read configuration should not return an error")
	assert.NotNil(t, cfg, "configuration should not be nil")
	assert.Equal(t, *newConfigurationSample(), cfg, "configuration should be equal to the expected configuration")
}

func TestReadConfiguration_yamlFile(t *testing.T) {
	var cfg configuration
	err := ReadConfiguration(&cfg, WithConfigFile(Yaml, "unit-test.yaml"))

	assert.NoError(t, err, "read configuration should not return an error")
	assert.NotNil(t, cfg, "configuration should not be nil")
	assert.Equal(t, *newConfigurationSample(), cfg, "configuration should be equal to the expected configuration")
}

func TestReadConfiguration_iniFile(t *testing.T) {
	var cfg configuration
	err := ReadConfiguration(&cfg, WithConfigFile(Ini, "unit-test.ini"))

	assert.NoError(t, err, "read configuration should not return an error")
	assert.NotNil(t, cfg, "configuration should not be nil")
	assert.Equal(t, *newConfigurationSample(), cfg, "configuration should be equal to the expected configuration")
}

func setupEnvs(t *testing.T) {
	err := godotenv.Load("unit-test.env")
	if err != nil {
		t.Fatal(fmt.Sprintf("cannot load test environment variables: %v", err))
	}
}

type configuration struct {
	Name    string              `mapstructure:"name"`
	Version string              `mapstructure:"version"`
	Port    int32               `mapstructure:"port"`
	List    []int32             `mapstructure:"list"`
	Map     any                 `mapstructure:"map"`
	Struct  configurationStruct `mapstructure:"struct"`
}

type configurationStruct struct {
	StringField string `mapstructure:"string_field"`
	IntField    int    `mapstructure:"int_field"`
	BoolField   bool   `mapstructure:"bool_field"`
}

func newConfigurationSample() *configuration {
	return &configuration{
		Name:    "app name",
		Version: "1.0.0",
		Port:    3210,
		List:    []int32{1, 2, 3},
		Map: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
		Struct: struct {
			StringField string `mapstructure:"string_field"`
			IntField    int    `mapstructure:"int_field"`
			BoolField   bool   `mapstructure:"bool_field"`
		}{
			StringField: "string_value",
			IntField:    123,
			BoolField:   true,
		},
	}
}
