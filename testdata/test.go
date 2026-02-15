package testdata

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	logger.Info("Привет это лог!")
	logger.Info("api_key: dsfsdfsdfsdf")
	logger.Info("Start server")
	logger.Info("server started")
	logger.Info("you!")
}
