package config

type ServerConfig struct {
	Name        string            `mapstructure:"name" json:"name"`
	Port        int               `mapstructure:"port" json:"port"`
	GoodsServer GoodsServerConfig `mapstructure:"goods_srv"`
	JWTInfo     JWTConfig         `mapstructure:"jwt"`
	ConsulInfo  ConsulConfig      `mapstructure:"consul"`
}

type GoodsServerConfig struct {
	Name string `mapstructure:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
