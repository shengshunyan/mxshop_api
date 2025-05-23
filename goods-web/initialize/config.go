package initialize

import (
	"mxshop_api/common/initialize"
	"mxshop_api/common/utils"
	"mxshop_api/goods-web/config"
	"mxshop_api/goods-web/global"
)

func InitConfig() {
	env := utils.GetEnv()
	configFilePath := ""
	if env == "dev" {
		configFilePath = "goods-web/config/config-dev.yaml"
	} else {
		configFilePath = "goods-web/config/config-prod.yaml"
	}

	global.ServerConfig = initialize.GetConfig[config.ServerConfig](configFilePath)
}
