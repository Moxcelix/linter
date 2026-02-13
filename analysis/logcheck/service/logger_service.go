package service

import (
	"slices"
	"strings"
)

type LoggerProvider interface {
	ProvideLoggerFuncs() map[string][]string
	ProvideLoggerMethods() map[string][]string
}

type LoggerService struct {
	loggerFuncs   map[string][]string
	loggerMethods map[string][]string
}

func NewLoggerService(provider LoggerProvider) *LoggerService {
	return &LoggerService{
		loggerFuncs:   provider.ProvideLoggerFuncs(),
		loggerMethods: provider.ProvideLoggerMethods(),
	}
}

func (s *LoggerService) IsLoggerCall(pkg string, selector string, method string) bool {
	selector = strings.Trim(selector, "*")
	if methods, exists := s.loggerMethods[selector]; exists {
		return slices.Contains(methods, method)
	}

	if funcs, exists := s.loggerFuncs[pkg]; exists {
		return slices.Contains(funcs, method)
	}

	return false
}
