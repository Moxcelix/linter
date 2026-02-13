package main

import (
	"main/pkg/logcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(logcheck.BuildAnalizer(logcheck.DefaultConfig()))
}
