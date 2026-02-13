package logcheck

import (
	"go/ast"
	"slices"

	"golang.org/x/tools/go/analysis"
)

var loggerMethods = map[string][]string{
	"slog.Logger": {
		"Debug", "Info", "Warn", "Error",
		"DebugContext", "InfoContext", "WarnContext", "ErrorContext",
		"Log", "LogContext",
		"Enabled", "Handler", "With", "WithGroup",
	},
	"zap.Logger": {
		"Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal",
		"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
		"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw",
		"With", "Named", "WithOptions", "Core", "Check", "Sugar",
	},
	"zap.SugaredLogger": {
		"Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal",
		"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
		"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw",
		"With", "Named", "Desugar",
	},
}

var loggerFuncs = map[string][]string{
	"slog": {
		"Debug", "Info", "Warn", "Error",
		"DebugContext", "InfoContext", "WarnContext", "ErrorContext",
		"Log", "LogContext",
		"Default", "SetDefault", "New", "NewJSONHandler", "NewTextHandler",
		"With", "NewRecord", "NewLogLogger",
	},
}

func IsLogCall(call *ast.CallExpr, pass *analysis.Pass) bool {
	if expr, ok := call.Fun.(*ast.SelectorExpr); ok {
		if selector, ok := expr.X.(*ast.Ident); ok {
			if obj := pass.TypesInfo.Uses[selector]; obj != nil {
				selectorName := obj.Type().String()
				if selectorName[0] == '*' {
					selectorName = selectorName[1:]
				}
				if methods, exists := loggerMethods[selectorName]; exists {
					return slices.Contains(methods, expr.Sel.Name)
				}

				selectorPkg := obj.Pkg().Name()
				if funcs, exists := loggerFuncs[selectorPkg]; exists {
					return slices.Contains(funcs, expr.Sel.Name)
				}
			}
		}
	}

	return false
}
