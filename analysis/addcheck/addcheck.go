package addcheck

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"

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
		(*ast.BinaryExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		be := n.(*ast.BinaryExpr)
		if be.Op != token.ADD {
			return
		}

		if _, ok := be.X.(*ast.BasicLit); !ok {
			return
		}

		if _, ok := be.Y.(*ast.BasicLit); !ok {
			return
		}

		isInteger := func(expr ast.Expr) bool {
			t := pass.TypesInfo.TypeOf(expr)
			if t == nil {
				return false
			}

			bt, ok := t.Underlying().(*types.Basic)
			if !ok {
				return false
			}

			if (bt.Info() & types.IsInteger) == 0 {
				return false
			}

			return true
		}

		if !isInteger(be.X) || !isInteger(be.Y) {
			return
		}

		pass.Reportf(be.Pos(), "integer addition found %q",
			render(pass.Fset, be))
	})

	return nil, nil
}
