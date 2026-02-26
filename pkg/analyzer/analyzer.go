package analyzer

import (
	"github.com/LainIwakuras-father/loglint/pkg/config"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func NewAnalyzer(cfg config.Config) (*analysis.Analyzer, error) {
	return &analysis.Analyzer{
		Name: "loglint",
		Doc:  "reports integer additions",
		URL:  "https://gitlab.com/LainIwakuras-father/slogzaplint",
		Run: func(pass *analysis.Pass) (any, error) {
			return run(pass, cfg)
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}, nil
}
