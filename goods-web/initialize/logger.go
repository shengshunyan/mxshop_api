package initialize

import "go.uber.org/zap"

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewDevelopment()

	zap.ReplaceGlobals(logger)
}

func CloseLogger() {
	_ = logger.Sync()
}
