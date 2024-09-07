package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadConfiguration_Ptr(t *testing.T) {
	setupEnvs(t)

	os.Getenv("NAME")
	var cfg configuration
	err := ReadConfiguration(&cfg)
	assert.NoError(t, err, "read configuration should not return an error")
	assert.NotNil(t, cfg, "configuration should not be nil")
}

func setupEnvs(t *testing.T) {
	err := godotenv.Load("unit-test.env")
	if err != nil {
		t.Fatal(fmt.Sprintf("cannot load test environment variables: %v", err))
	}
}

type configuration struct {
	Name    string         `mapstructure:"name"`
	Version string         `mapstructure:"version"`
	Port    int32          `mapstructure:"port"`
	List    []any          `mapstructure:"list"`
	Map     map[string]any `mapstructure:"map"`
	Struct  struct {
		StringField string `mapstructure:"string_field"`
		IntField    int    `mapstructure:"int_field"`
		BoolField   bool   `mapstructure:"bool_field"`
	} `mapstructure:"struct"`
}
