package main

import (
	"encoding/json"
	"fmt"
	"linters/pkg/logcheck"

	_ "github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	cfg := logcheck.DefaultConfig()
	fmt.Println(conf)
	if conf != nil {
		bytes, err := json.Marshal(conf)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(bytes))
		var userCfg logcheck.Config
		if err := json.Unmarshal(bytes, &userCfg); err != nil {
			return nil, err
		}
		fmt.Println(userCfg)
		cfg = &userCfg
	}
	fmt.Println(cfg)
	plugin := logcheck.NewLogcheckLinter(cfg)
	analyzer := plugin.BuildAnalizer()
	return []*analysis.Analyzer{analyzer}, nil
}
