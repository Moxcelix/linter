package rules

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func CheckEnglishRule(pass *analysis.Pass, lit *ast.BasicLit, msg string) {
	for _, r := range msg {
		if r > 127 {
			pass.Reportf(lit.Pos(), "log message should have english characters only: %s", msg)
			return
		}
	}
}
