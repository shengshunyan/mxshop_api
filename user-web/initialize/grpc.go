package initialize

import (
	"fmt"
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop_api/user-web/global"
)

var conn *grpc.ClientConn

func InitGrpc() {
	// 拨号连接grpc服务
	var err error
	conn, err = grpc.NewClient(fmt.Sprintf("%s:%d",
		global.ServerConfig.UserServer.Host,
		global.ServerConfig.UserServer.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("new client failed" + err.Error())
	}

	global.Stub.UserClient = userProto.NewUserClient(conn)
	zap.S().Infow("init grpc success")
}

func CloseGrpc() {
	_ = conn.Close()
}
