package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	goodsProto "github.com/shengshunyan/mxshop-proto/goods/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop_api/goods-web/global"
)

var conn *grpc.ClientConn

func InitGrpc() {
	// 从注册中心获取到服务的信息
	consulConfig := global.ServerConfig.ConsulInfo
	goodsServer := global.ServerConfig.GoodsServer

	// 拨号连接grpc服务
	var err error
	conn, err = grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			consulConfig.Host, consulConfig.Port, goodsServer.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 负载均衡
	)
	if err != nil {
		panic("new client failed" + err.Error())
	}

	global.Stub.GoodsClient = goodsProto.NewGoodsClient(conn)

	zap.S().Infow("init grpc success")
}

func CloseGrpc() {
	_ = conn.Close()
}
