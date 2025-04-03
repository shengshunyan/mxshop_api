package config

type ServerConfig struct {
	Name       string           `mapstructure:"name"`
	Port       int              `mapstructure:"port"`
	UserServer UserServerConfig `mapstructure:"user_srv"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
