package initialize

import (
	"mxshop_api/common/config"
	"mxshop_api/common/initialize"
	"mxshop_api/user-web/global"
)

// 服务注册
func InitRegister() {
	serverConfig := global.ServerConfig
	initialize.InitRegister(&config.ServerConfig{
		Name:       serverConfig.Name,
		Host:       serverConfig.Host,
		Port:       serverConfig.Port,
		ConsulInfo: serverConfig.ConsulInfo,
	})
}

func CloseRegister() {
	initialize.CloseRegister()
}
