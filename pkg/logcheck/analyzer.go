package logcheck

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"linters/internal/rules"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type LogcheckLinter struct {
	cfg *Config
}

func NewLogcheckLinter(cfg *Config) *LogcheckLinter {
	return &LogcheckLinter{cfg: cfg}
}

func (linter *LogcheckLinter) BuildAnalizer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "logcheck",
		Doc:  "checks log messages for proper format",
		Run:  linter.run,
	}
}

type Rule interface {
	Check(msg string) error
}

func (linter *LogcheckLinter) run(pass *analysis.Pass) (interface{}, error) {
	var rulesList []Rule
	if linter.cfg.Rules.English.Enabled == "true" {
		rulesList = append(rulesList, rules.NewEnglishRule())
	}

	if linter.cfg.Rules.Lowercase.Enabled == "true" {
		rulesList = append(rulesList, rules.NewLowercaseRule())
	}

	if linter.cfg.Rules.Secret.Enabled == "true" {
		rulesList = append(rulesList, rules.NewSecretRule(SecretProvider{linter.cfg.Rules.Secret.Words}))
	}

	if linter.cfg.Rules.Special.Enabled == "true" {
		rulesList = append(rulesList, rules.NewSpecialRule(SpecialProvider{linter.cfg.Rules.Special.Chars}))
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if ce, ok := n.(*ast.CallExpr); ok {
				if !linter.IsLogCall(ce, pass) {
					return true
				}

				if len(ce.Args) == 0 {
					return true
				}

				lit, ok := ce.Args[0].(*ast.BasicLit)
				if !ok || lit.Kind != token.STRING {
					return true
				}

				msg := strings.Trim(lit.Value, "\"")

				for _, rule := range rulesList {
					if err := rule.Check(msg); err != nil {
						pass.Reportf(lit.Pos(), err.Error()+": %s !", render(pass.Fset, ce))
					}
				}
			}

			return true
		})
	}

	return nil, nil
}

func (linter *LogcheckLinter) IsLogCall(call *ast.CallExpr, pass *analysis.Pass) bool {
	if call == nil || pass == nil || pass.TypesInfo == nil {
		return false
	}

	selExpr, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	receiverType := pass.TypesInfo.TypeOf(selExpr.X)
	if receiverType == nil {
		return false
	}

	switch t := receiverType.(type) {
	case *types.Pointer:
		receiverType = t.Elem()
	}

	named, ok := receiverType.(*types.Named)
	if !ok {
		return false
	}

	pkg := named.Obj().Pkg()
	if pkg == nil {
		return false
	}

	pkgPath := pkg.Path()
	pkgName := pkg.Name()
	typeName := named.Obj().Name()

	for _, loggerConfig := range linter.cfg.Loggers {
		for _, objConfig := range loggerConfig.LoggerObj {
			if pkgPath == loggerConfig.PkgName || pkgName == loggerConfig.PkgName {
				if typeName == objConfig.Name {
					if slices.Contains(objConfig.Methods, selExpr.Sel.Name) {
						return true
					}
				}
			}
		}

		if pkgName == loggerConfig.PkgName || strings.HasSuffix(pkgPath, loggerConfig.PkgName) {
			if slices.Contains(loggerConfig.Funcs, selExpr.Sel.Name) {
				return true
			}
		}
	}

	return false
}

func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
