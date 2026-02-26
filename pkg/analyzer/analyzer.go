package analyzer

import (
	"github.com/LainIwakuras-father/slogzaplint/pkg/config"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func NewAnalyzer(cfg config.Config) (*analysis.Analyzer, error) {
	return &analysis.Analyzer{
		Name: "slogzaplint",
		Doc:  "reports non-standart message logs",
		URL:  "https://github.com/LainIwakuras-father/slogzaplint",
		Run: func(pass *analysis.Pass) (any, error) {
			return run(pass, cfg)
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}, nil
}
