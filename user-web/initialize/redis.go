package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"mxshop_api/user-web/global"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			global.ServerConfig.Redis.Host,
			global.ServerConfig.Redis.Port), // Redis 服务器地址和端口
		Password: "", // 密码（可选）
		DB:       0,  // 数据库编号（默认为0）
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	zap.S().Infow("connect redis success")
	global.Rdb = rdb
}

func CloseRedis() {
	_ = global.Rdb.Close()
}
