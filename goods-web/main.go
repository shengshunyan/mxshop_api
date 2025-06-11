package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"mxshop_api/goods-web/global"
	"mxshop_api/goods-web/initialize"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化logger
	initialize.InitLogger()
	defer initialize.CloseLogger()

	// 初始化config
	initialize.InitConfig()

	// 初始化grpc
	initialize.InitGrpc()
	defer initialize.CloseGrpc()

	// 初始化router
	router := initialize.InitRouter()

	// 服务注册
	initialize.InitRegister()
	defer initialize.CloseRegister()

	// 创建 HTTP Server 实例
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.ServerConfig.Port),
		Handler: router,
	}

	go func() {
		zap.S().Infof("start server on port %d", global.ServerConfig.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Panic("failed to start server", zap.Error(err))
		}
	}()

	// 捕获中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("shutting down server...")

	// 创建 context 用于优雅关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 优雅关闭 HTTP 服务
	if err := server.Shutdown(ctx); err != nil {
		zap.S().Error("server shutdown failed", zap.Error(err))
	}
}
