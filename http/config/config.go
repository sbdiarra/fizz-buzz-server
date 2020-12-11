package config

import (
	"github.com/spf13/viper"
	"strings"
)

type appConfig struct {
	GIN_MODE string
}

var config *appConfig

func Config() *appConfig {
	return config
}

const (
	DOT        = "."
	UNDERSCORE = "_"
)

func loadFromEnv() *appConfig {
	viper.SetEnvKeyReplacer(strings.NewReplacer(DOT, UNDERSCORE))

	return &appConfig{
		GIN_MODE: viper.GetString("gin.mode"),
	}
}

func init() {
	config = loadFromEnv()
}
