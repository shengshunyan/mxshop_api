package router

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")

	{
		UserRouter.GET("list", api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
	}
}
