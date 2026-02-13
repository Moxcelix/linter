package logcheck

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"main/internal/rules"
)

func BuildAnalizer(config *Config) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "logcheck",
		Doc:      "checks log messages for proper format",
		Run:      makeRunWithConfig(config),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

type Rule interface {
	Check(msg string) error
}

func makeRunWithConfig(cfg *Config) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

		var rulesList []Rule

		if cfg.Rules.English.Enabled {
			rulesList = append(rulesList, rules.NewEnglishRule())
		}

		if cfg.Rules.Lowercase.Enabled {
			rulesList = append(rulesList, rules.NewLowercaseRule())
		}

		if cfg.Rules.Secret.Enabled {
			rulesList = append(rulesList, rules.NewSecretRule(SecretProvider{cfg.Rules.Secret.Words}))
		}

		if cfg.Rules.Special.Enabled {
			rulesList = append(rulesList, rules.NewSpecialRule(SpecialProvider{cfg.Rules.Special.Chars}))
		}

		nodeFilter := []ast.Node{
			(*ast.CallExpr)(nil),
		}

		inspect.Preorder(nodeFilter, func(n ast.Node) {
			ce := n.(*ast.CallExpr)

			if !IsLogCall(ce, pass, cfg) {
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

			for _, rule := range rulesList {
				if err := rule.Check(msg); err != nil {
					pass.Reportf(lit.Pos(), err.Error()+": %s", render(pass.Fset, ce))
				}
			}
		})

		return nil, nil
	}
}

func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
