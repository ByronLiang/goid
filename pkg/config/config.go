package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const DefaultEnv = "local"

func NewConfig() error {
	viper.AutomaticEnv()
	env := viper.GetString("CONFIG_ENV")
	if env == "" {
		env = DefaultEnv
	}
	filename := fmt.Sprintf("%s.config", env)
	viper.SetConfigName(filename)
	if viper.GetString("CONFIG_PATH") == "" {
		viper.AddConfigPath(".")
	} else {
		viper.AddConfigPath(viper.GetString("CONFIG_PATH"))
	}
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
