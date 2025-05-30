package initialize

import (
	"mxshop_api/common/initialize"
	"mxshop_api/common/utils"
	"mxshop_api/user-web/config"
	"mxshop_api/user-web/global"
)

func InitConfig() {
	env := utils.GetEnv()
	configFilePath := ""
	if env == "dev" {
		configFilePath = "user-web/config/config-dev.yaml"
	} else {
		configFilePath = "user-web/config/config-prod.yaml"
	}

	global.ServerConfig = initialize.GetConfig[config.ServerConfig](configFilePath)
}
