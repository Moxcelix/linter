package app

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"main/analysis/logcheck/service"
	"strings"

	"go.uber.org/fx"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (l *Controller) Run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ce := n.(*ast.CallExpr)

		if len(ce.Args) == 0 {
			return
		}

		lit, ok := ce.Args[0].(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return
		}

		if expr, ok := ce.Fun.(*ast.SelectorExpr); ok {
			if selector, ok := expr.X.(*ast.Ident); ok {
				if obj := pass.TypesInfo.Uses[selector]; obj != nil {
					l.service.CheckLoggerExpression(
						obj.Pkg().Name(),
						obj.Type().String(),
						expr.Sel.Name,
						strings.Trim(lit.Value, "\""),
						func(err error) {
							pass.Reportf(lit.Pos(), err.Error()+": %s", l.render(pass.Fset, ce))
						})
				}
			}
		}
	})

	return nil, nil
}

func (l *Controller) render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}

var Module = fx.Options(
	fx.Provide(NewController),
)
