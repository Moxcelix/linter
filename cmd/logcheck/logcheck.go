package main

import (
	"linters/pkg/logcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	plugin := logcheck.NewLogcheckLinter(logcheck.DefaultConfig())
	singlechecker.Main(plugin.BuildAnalizer())
}
