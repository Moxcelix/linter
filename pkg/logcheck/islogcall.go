package logcheck

import (
	"go/ast"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func IsLogCall(call *ast.CallExpr, pass *analysis.Pass, config *Config) bool {
	selExpr, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	receiver, ok := selExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.Uses[receiver]
	if obj == nil {
		return false
	}

	receiverType := obj.Type()
	if receiverType == nil {
		return false
	}

	typeStr := strings.Trim(receiverType.String(), "*")

	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	pkgName := pkg.Name()

	for _, loggerConfig := range config.Loggers {
		for _, objConfig := range loggerConfig.LoggerObj {
			fullTypeName := loggerConfig.PkgName + "." + objConfig.Name
			if typeStr == fullTypeName {
				return slices.Contains(objConfig.Methods, selExpr.Sel.Name)
			}
		}

		if pkgName == loggerConfig.PkgName {
			return slices.Contains(loggerConfig.Funcs, selExpr.Sel.Name)
		}
	}

	return false
}
