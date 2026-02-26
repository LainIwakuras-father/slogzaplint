package analyzer

import (
	"go/ast"
	"slices"

	"github.com/LainIwakuras-father/loglint/pkg/config"
	"github.com/LainIwakuras-father/loglint/pkg/rules"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func run(pass *analysis.Pass, cfg config.Config) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}
		// veryfy call
		if !isLogCall(pass, call) {
			return
		}
		msg, pos, ok := extractString(call)
		if !ok {
			return
		}

		// check rules
		if slices.Contains(cfg.EnabledRules, "lowercase") {
			err := rules.Lowercase(msg)
			if err != nil {
				pass.Reportf(pos, "%s", err.Error())
			}
		}

		if slices.Contains(cfg.EnabledRules, "english") {
			err := rules.English(msg)
			if err != nil {
				pass.Reportf(pos, "%s", err.Error())
			}
		}
		if slices.Contains(cfg.EnabledRules, "no-special") {
			err := rules.NoSpecial(msg)
			if err != nil {
				pass.Reportf(pos, "%s", err.Error())
			}
		}
		if slices.Contains(cfg.EnabledRules, "no-sensitive") {
			err := rules.NoSensitive(msg, cfg.SensitivePatterns)
			if err != nil {
				pass.Reportf(pos, "%s", err.Error())
			}
		}
	})

	return nil, nil
}
