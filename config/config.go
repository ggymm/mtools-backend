package config

import (
	"github.com/BurntSushi/toml"
)

type GlobalConfig struct {
	App   appConfig
	Coder CoderConfig
}

type appConfig struct {
	Addr             string
	AllowCrossDomain bool `toml:"allow_cross_domain"`
}

type CoderConfig struct {
	ConfigFolder string `toml:"config_folder"`
}

func InitConfig() (*GlobalConfig, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func NewConfig() (*GlobalConfig, error) {
	config := new(GlobalConfig)
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return nil, err
	}
	return config, nil
}
