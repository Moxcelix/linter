package main

import (
	"encoding/json"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"main/pkg/logcheck"
)

type Plugin struct {
	config *logcheck.Config
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	analyzer := logcheck.BuildAnalizer(p.config)
	return []*analysis.Analyzer{analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func NewPlugin(settings any) (register.LinterPlugin, error) {
	cfg := logcheck.DefaultConfig()

	if settings != nil {
		bytes, err := json.Marshal(settings)
		if err != nil {
			return nil, err
		}

		var userCfg logcheck.Config
		if err := json.Unmarshal(bytes, &userCfg); err != nil {
			return nil, err
		}

		cfg = &userCfg
	}

	return &Plugin{config: cfg}, nil
}

func main() {
	register.Plugin("logcheck", NewPlugin)
}
