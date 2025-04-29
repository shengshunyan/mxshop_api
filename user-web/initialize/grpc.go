package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	userProto "github.com/shengshunyan/mxshop-proto/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials/insecure"
	"mxshop_api/user-web/global"
)

var conn *grpc.ClientConn

func InitGrpc() {
	// 从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consulConfig := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	//client, err := api.NewClient(cfg)
	//if err != nil {
	//	panic(err)
	//}

	userServer := global.ServerConfig.UserServer
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s" `, userServer.Name))
	//if err != nil {
	//	panic(err)
	//}
	//userServerHost := ""
	//userServerPort := 0
	//for _, service := range data {
	//	userServerHost = service.Address
	//	userServerPort = service.Port
	//}

	// 拨号连接grpc服务
	//conn, err = grpc.NewClient(fmt.Sprintf("%s:%d", userServerHost, userServerPort),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	panic("new client failed" + err.Error())
	//}

	// grpc服务负载均衡
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			consulConfig.Host, consulConfig.Port, userServer.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
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
