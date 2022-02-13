package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const DefaultEnv = "local"

func NewConfig(env string) error {
	if env == "" {
		env = DefaultEnv
	}
	filename := fmt.Sprintf("%s.config", env)
	viper.SetConfigName(filename)
	if os.Getenv("CONFIG_PATH") == "" {
		viper.AddConfigPath(".")
	} else {
		viper.AddConfigPath(os.Getenv("CONFIG_PATH"))
	}
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
