package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token string `mapstructure:"token"`
	Env   string `mapstructure:"env"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
