package config

import "mxshop_api/common/config"

type ServerConfig struct {
	Name        string              `mapstructure:"name" json:"name"`
	Host        string              `mapstructure:"host" json:"host"`
	Port        int                 `mapstructure:"port" json:"port"`
	GoodsServer GoodsServerConfig   `mapstructure:"goods_srv"`
	JWTInfo     JWTConfig           `mapstructure:"jwt"`
	ConsulInfo  config.ConsulConfig `mapstructure:"consul"`
}

type GoodsServerConfig struct {
	Name string `mapstructure:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}
