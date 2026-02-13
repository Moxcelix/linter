package data

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLoggerProvider),
	fx.Provide(NewSecretProvider),
	fx.Provide(NewSpecialProvider),
)
