package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func CheckLowercaseRule(pass *analysis.Pass, lit *ast.BasicLit, msg string) {
	if len(msg) == 0 {
		return
	}

	firstChar := []rune(msg)[0]
	if unicode.IsUpper(firstChar) {
		pass.Reportf(lit.Pos(), "log message should start with a lowercase letter: %s", msg)
	}
}
