package linter

import "go.uber.org/zap"

func c() {
	logger, _ := zap.NewProduction()

	logger.Info("adekvatny log")
}
