package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop_api/user-web/forms"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/global/response"
	"mxshop_api/user-web/middlewares"
	"mxshop_api/user-web/models"
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
	client := global.Stub.UserClient
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

	// 验证图形验证码
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	client := global.Stub.UserClient
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

	// 生成token
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(rsp.Id),
		NickName:    rsp.Nickname,
		AuthorityId: uint(rsp.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                          // 签名生效时间
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // 30天过期
			Issuer:    "imooc",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         rsp.Id,
		"nick_name":  rsp.Nickname,
		"token":      token,
		"expired_at": time.Now().Add(time.Hour*24*30).Unix() * 1000,
	})
}

// 用户注册
func Register(ctx *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBindJSON(&registerForm); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	value, err := global.Rdb.Get(context.Background(), registerForm.Mobile).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "验证码获取出错",
		})
		return
	}

	if value != registerForm.Code {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	rsp, err := global.Stub.UserClient.CreateUser(context.Background(), &userProto.CreateUserInfo{
		Mobile:   registerForm.Mobile,
		Password: registerForm.Password,
		Nickname: registerForm.Mobile,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "注册失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":        rsp.Id,
		"nick_name": rsp.Nickname,
	})
}
