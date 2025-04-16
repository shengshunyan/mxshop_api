package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"mxshop_api/user-web/forms"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/utils"
	"net/http"
	"time"
)

func SendSms(ctx *gin.Context) {
	code := utils.GenerateNumericCode(6)
	sendSmsForm := &forms.SendSmsForm{}
	if err := ctx.ShouldBindJSON(sendSmsForm); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	global.Rdb.Set(context.Background(), sendSmsForm.Mobile, code, 300*time.Second)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
