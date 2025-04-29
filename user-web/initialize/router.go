package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user-web/api"
	userRouter "mxshop_api/user-web/router"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 健康检查
	router.GET("/health", api.Health)

	// 业务路由
	group := router.Group("/v1")
	userRouter.InitBaseRouter(group)
	userRouter.InitUserRouter(group)
	return router
}
