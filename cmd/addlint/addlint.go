package main

import (
	"main/analysis/addcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(addcheck.Analyzer)
}
