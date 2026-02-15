package linter

import (
	"go.uber.org/zap"
)

func b() {
	logger, _ := zap.NewProduction()

	logger.Info("PepeваФпа")
	logger.Info("ApiKey: dsfsdfsdfsdf ✅")
	logger.Info("/")
	logger.Info("!")
	logger.Info("!")
}
