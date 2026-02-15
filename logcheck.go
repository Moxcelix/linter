package linters

import (
	"linters/pkg/logcheck"

	"github.com/golangci/plugin-module-register/register"
	_ "go.uber.org/zap"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("logcheck", New)
}

type LogcheckPlugin struct {
	config logcheck.Config
}

func New(settings any) (register.LinterPlugin, error) {
	cfg, err := register.DecodeSettings[logcheck.Config](settings)
	if err != nil {
		return nil, err
	}
	return &LogcheckPlugin{config: cfg}, nil
}

func (plugin *LogcheckPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	linter := logcheck.NewLogcheckLinter(&plugin.config)
	analyzer := linter.BuildAnalizer()

	return []*analysis.Analyzer{analyzer}, nil
}

func (plugin *LogcheckPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
