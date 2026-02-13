package data

import "main/analysis/logcheck/service"

type LoggerProvider struct {
}

func NewLoggerProvider() service.LoggerProvider {
	return &LoggerProvider{}
}

func (p *LoggerProvider) ProvideLoggerFuncs() map[string][]string {
	return map[string][]string{
		"log/slog": {
			"Debug", "Info", "Warn", "Error",
			"DebugContext", "InfoContext", "WarnContext", "ErrorContext",
			"Log", "LogContext",
			"Default", "SetDefault", "New", "NewJSONHandler", "NewTextHandler",
			"With", "NewRecord", "NewLogLogger",
		},
	}
}

func (p *LoggerProvider) ProvideLoggerMethods() map[string][]string {
	return map[string][]string{
		"log/slog.Logger": {
			"Debug", "Info", "Warn", "Error",
			"DebugContext", "InfoContext", "WarnContext", "ErrorContext",
			"Log", "LogContext",
			"Enabled", "Handler", "With", "WithGroup",
		},
		"go.uber.org/zap.Logger": {
			"Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal",
			"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
			"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw",
			"With", "Named", "WithOptions", "Core", "Check", "Sugar",
		},
		"go.uber.org/zap.SugaredLogger": {
			"Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal",
			"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
			"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw",
			"With", "Named", "Desugar",
		},
	}
}
