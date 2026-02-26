package analyzer

import (
	"fmt"

	"github.com/LainIwakuras-father/slogzaplint/pkg/config"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("slogzaplint", New)
}

type plugin struct {
	cfg *config.Config
}

var _ register.LinterPlugin = (*plugin)(nil)

func New(settings any) (register.LinterPlugin, error) {
	p := &plugin{
		cfg: &config.Config{
			EnabledRules:      []string{},
			SensitivePatterns: []string{},
		},
	}

	if settings != nil {
		m, ok := settings.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("settings must be a map, got %T", settings)
		}

		if rules, ok := m["enabled-rules"]; ok {
			if rulesSlice, ok := rules.([]any); ok {
				for _, r := range rulesSlice {
					if ruleStr, ok := r.(string); ok {
						p.cfg.EnabledRules = append(p.cfg.EnabledRules, ruleStr)
					}
				}
			}
		}

		if patterns, ok := m["sensitive-patterns"]; ok {
			if patternsSlice, ok := patterns.([]any); ok {
				for _, pat := range patternsSlice {
					if patStr, ok := pat.(string); ok {
						p.cfg.SensitivePatterns = append(p.cfg.SensitivePatterns, patStr)
					}
				}
			}
		}

	} else {
		p.cfg = config.Default()
	}

	return p, nil
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	analyzer, err := NewAnalyzer(*p.cfg)
	if err != nil {
		return nil, err
	}
	return []*analysis.Analyzer{
		analyzer,
	}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
