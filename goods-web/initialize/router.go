package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/goods-web/api"
	goodsRouter "mxshop_api/goods-web/router"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 健康检查
	router.GET("/health", api.Health)

	// 业务路由
	group := router.Group("/v1")
	goodsRouter.InitGoodsRouter(group)
	return router
}
