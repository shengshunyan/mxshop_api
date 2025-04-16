package router

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("base")

	{
		UserRouter.GET("captcha", api.GetCaptcha)
		UserRouter.POST("send_sms", api.SendSms)
	}
}
