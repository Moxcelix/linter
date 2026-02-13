package main

import (
	"main/analysis/logcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(logcheck.Analyzer)
}
