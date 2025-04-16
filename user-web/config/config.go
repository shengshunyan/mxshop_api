package config

type ServerConfig struct {
	Name       string           `mapstructure:"name"`
	Port       int              `mapstructure:"port"`
	UserServer UserServerConfig `mapstructure:"user_srv"`
	JWTInfo    JWTConfig        `mapstructure:"jwt"`
	Redis      RedisConfig      `mapstructure:"redis"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
