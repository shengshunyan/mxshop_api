package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"mxshop_api/user-web/forms"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/global/response"
	"net/http"
	"strconv"
	"time"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err == nil {
		return
	}

	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"msg": e.Message(),
			})
		case codes.Internal:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "内部错误",
			})
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数错误",
			})
		case codes.Unavailable:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "用户服务不可用",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "其他错误",
			})
		}
	}
}

// 获取用户列表
func GetUserList(ctx *gin.Context) {
	// 拨号连接grpc服务
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d",
		global.ServerConfig.UserServer.Host,
		global.ServerConfig.UserServer.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("new client failed" + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("close client failed" + err.Error())
		}
	}(conn)

	client := userProto.NewUserClient(conn)
	pn := ctx.DefaultQuery("pn", "1")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := client.GetUserList(context.Background(), &userProto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("failed to get user list")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]response.UserResponse, 0, len(rsp.Data))
	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			Password: value.Password,
			Mobile:   value.Mobile,
			Nickname: value.Nickname,
			Birthday: time.Unix(int64(value.Birthday), 0),
			Gender:   value.Gender,
			Role:     value.Role,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}

// 表单验证
func PasswordLogin(ctx *gin.Context) {
	passwordLoginForm := forms.PasswordLoginForm{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 拨号连接grpc服务
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d",
		global.ServerConfig.UserServer.Host,
		global.ServerConfig.UserServer.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("new client failed" + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("close client failed" + err.Error())
		}
	}(conn)

	client := userProto.NewUserClient(conn)

	rsp, err := client.GetUserByMobile(context.Background(), &userProto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败",
		})
		return
	}

	checkRsp, checkErr := client.CheckPassword(context.Background(), &userProto.CheckInfo{
		Password:          passwordLoginForm.Password,
		EncryptedPassword: rsp.Password,
	})
	if checkErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登录失败",
		})
		return
	}
	if !checkRsp.Success {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "密码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
