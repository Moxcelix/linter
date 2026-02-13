package logcheck

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"main/analysis/logcheck/rules"
	"strings"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "addlint",
	Doc:      "reports integer additions",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ce := n.(*ast.CallExpr)

		if !IsLogCall(ce, pass) {
			return
		}

		if len(ce.Args) == 0 {
			return
		}

		lit, ok := ce.Args[0].(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return
		}

		msg := strings.Trim(lit.Value, "\"")

		if err := rules.CheckEnglishRule(msg); err != nil {
			pass.Reportf(lit.Pos(), err.Error()+": %s", render(pass.Fset, ce))
		}

		if err := rules.CheckLowercaseRule(msg); err != nil {
			pass.Reportf(lit.Pos(), err.Error()+": %s", render(pass.Fset, ce))
		}

		if err := rules.CheckSecretRule(msg); err != nil {
			pass.Reportf(lit.Pos(), err.Error()+": %s", render(pass.Fset, ce))
		}
	})

	return nil, nil
}
