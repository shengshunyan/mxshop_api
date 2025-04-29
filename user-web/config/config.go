package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Namespace string `mapstructure:"namespace"`
	Group     string `mapstructure:"group"`
	DataId    string `mapstructure:"dataid"`
}

type ServerConfig struct {
	Name       string           `mapstructure:"name" json:"name"`
	Port       int              `mapstructure:"port" json:"port"`
	UserServer UserServerConfig `mapstructure:"user_srv"`
	JWTInfo    JWTConfig        `mapstructure:"jwt"`
	Redis      RedisConfig      `mapstructure:"redis"`
	ConsulInfo ConsulConfig     `mapstructure:"consul"`
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

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
