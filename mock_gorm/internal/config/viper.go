package config

import (
	"github.com/spf13/viper"
)

type Viper interface {
	GetString(key string) string
}

func NewViper() Viper {
	viperConfig := viper.New()
	viperConfig.SetConfigFile(".env")
	if err := viperConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	return viperConfig
}
