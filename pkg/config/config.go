package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const DefaultName = "local"

func NewConfig() error {
	viper.AutomaticEnv()
	name := viper.GetString("CONFIG_NAME")
	if name == "" {
		name = DefaultName
	}
	filename := fmt.Sprintf("%s.config", name)
	viper.SetConfigName(filename)
	if viper.GetString("CONFIG_PATH") == "" {
		viper.AddConfigPath(".")
	} else {
		viper.AddConfigPath(viper.GetString("CONFIG_PATH"))
	}
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
