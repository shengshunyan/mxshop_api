package initialize

import (
	"github.com/gin-gonic/gin"
	userRouter "mxshop_api/user-web/router"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	group := router.Group("/v1")

	userRouter.InitUserRouter(group)
	return router
}
