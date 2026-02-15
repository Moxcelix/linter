package linters

import (
	"fmt"
	"linters/pkg/logcheck"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	fmt.Println("Register")
	register.Plugin("logcheck", New)
}

type LogcheckPlugin struct {
	config logcheck.Config
}

func New(settings any) (register.LinterPlugin, error) {
	fmt.Println("New")
	cfg, err := register.DecodeSettings[logcheck.Config](settings)
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg)
	return &LogcheckPlugin{config: cfg}, nil
}

func (plugin *LogcheckPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	fmt.Println("Build")
	linter := logcheck.NewLogcheckLinter(&plugin.config)
	analyzer := linter.BuildAnalizer()

	return []*analysis.Analyzer{analyzer}, nil
}

func (plugin *LogcheckPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
