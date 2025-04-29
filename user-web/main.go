package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/initialize"
	myValidator "mxshop_api/user-web/validator"
)

func main() {
	// 初始化logger
	initialize.InitLogger()
	defer initialize.CloseLogger()

	// 初始化config
	initialize.InitConfig()

	// 初始化redis
	initialize.InitRedis()
	defer initialize.CloseRedis()

	// 初始化grpc
	initialize.InitGrpc()
	defer initialize.CloseGrpc()

	// 初始化router
	router := initialize.InitRouter()

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myValidator.ValidateMobile)
	}

	zap.S().Infof("start server on port %d", global.ServerConfig.Port)
	err := router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
	if err != nil {
		zap.S().Panic("failed to start server", zap.Error(err))
	}
}
