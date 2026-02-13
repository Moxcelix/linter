package logcheck

import (
	"main/analysis/logcheck/app"
	"main/analysis/logcheck/data"
	"main/analysis/logcheck/rules"
	"main/analysis/logcheck/service"

	"go.uber.org/fx"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var CommonModules = fx.Options(
	rules.Module,
	service.Module,
	data.Module,
	app.Module,
)

func SetupApp(controller *app.Controller) {
	analyzer := &analysis.Analyzer{
		Name:     "logcheck",
		Doc:      "checks log messages for proper format",
		Run:      controller.Run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	singlechecker.Main(analyzer)
}

func StartApp() {
	opts := fx.Options(
		fx.Invoke(SetupApp),
	)

	app := fx.New(
		CommonModules,
		opts,
	)

	app.Run()
}
