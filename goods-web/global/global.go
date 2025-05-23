package global

import (
	"github.com/shengshunyan/mxshop-proto/goods/proto"
	"mxshop_api/goods-web/config"
)

type GrpcStub struct {
	GoodsClient proto.GoodsClient
}

var (
	ServerConfig = &config.ServerConfig{}
	Stub         = &GrpcStub{}
)
