package service

import (
	"go.uber.org/fx"
)

type Service struct {
	rulesService  *RulesService
	loggerService *LoggerService
}

func NewService(
	rulesService *RulesService,
	loggerService *LoggerService,
) *Service {
	return &Service{
		rulesService:  rulesService,
		loggerService: loggerService,
	}
}

func (s *Service) CheckLoggerExpression(pkg string, selector string, method string, msg string, callback Callback) {
	if !s.loggerService.IsLoggerCall(pkg, selector, method) {
		return
	}

	s.rulesService.CheckRules(msg, callback)
}

var Module = fx.Options(
	fx.Provide(NewLoggerService),
	fx.Provide(NewRulesService),
	fx.Provide(NewService),
)
