package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/shengshunyan/mxshop-proto/user/proto"
	"mxshop_api/user-web/config"
)

type GrpcStub struct {
	UserClient proto.UserClient
}

var (
	ServerConfig = &config.ServerConfig{}
	Rdb          *redis.Client
	Stub         = &GrpcStub{}
)
