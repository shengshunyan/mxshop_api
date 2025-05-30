package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop_api/user-web/global"
)

var conn *grpc.ClientConn

func InitGrpc() {
	// 从注册中心获取到用户服务的信息
	consulConfig := global.ServerConfig.ConsulInfo
	userServer := global.ServerConfig.UserServer
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	// 拨号连接grpc服务
	var err error
	conn, err = grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			consulConfig.Host, consulConfig.Port, userServer.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 负载均衡
	)
	if err != nil {
		panic("new client failed" + err.Error())
	}

	global.Stub.UserClient = userProto.NewUserClient(conn)

	zap.S().Infow("init grpc success")
}

func CloseGrpc() {
	_ = conn.Close()
}
