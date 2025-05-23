package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取商品列表
func GetGoodsList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
	})
}
