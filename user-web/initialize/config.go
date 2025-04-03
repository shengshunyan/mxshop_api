package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop_api/user-web/global"
)

const ENV_KEY = "MXSHOP_ENV"

func getEnv() string {
	err := viper.BindEnv(ENV_KEY)
	if err != nil {
		panic(err)
	}
	env := viper.GetString(ENV_KEY)

	return env
}

func InitConfig() {
	env := getEnv()

	v := viper.New()
	if env == "dev" {
		v.SetConfigFile("user-web/config/config-dev.yaml")
	} else {
		v.SetConfigFile("user-web/config/config-prod.yaml")
	}
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infow("[config] init get config", "serverConfig", &global.ServerConfig)

	// 动态监控功能
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := v.Unmarshal(&global.ServerConfig); err != nil {
			panic(err)
		}
		zap.S().Infow("[config] watch config change", "serverConfig", &global.ServerConfig)
	})
}
