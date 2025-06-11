package config

import "mxshop_api/common/config"

type ServerConfig struct {
	Name       string              `mapstructure:"name" json:"name"`
	Host       string              `mapstructure:"host" json:"host"`
	Port       int                 `mapstructure:"port" json:"port"`
	UserServer UserServerConfig    `mapstructure:"user_srv"`
	JWTInfo    JWTConfig           `mapstructure:"jwt"`
	Redis      RedisConfig         `mapstructure:"redis"`
	ConsulInfo config.ConsulConfig `mapstructure:"consul"`
}

type UserServerConfig struct {
	Name string `mapstructure:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
