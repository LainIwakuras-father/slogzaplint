package analyzer

import (
	"fmt"
	"os"

	"github.com/LainIwakuras-father/loglint/pkg/config"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", New)
}

type plugin struct {
	cfg *config.Config
}

var _ register.LinterPlugin = (*plugin)(nil)

func New(settings any) (register.LinterPlugin, error) {
	p := &plugin{
		// Начинаем с пустой конфигурации, чтобы потом полностью перезаписать из конфига
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

		// Отладка
		fmt.Fprintf(os.Stderr, "DEBUG: raw settings: %#v\n", m)

		// Парсим enabled-rules (прямо из m, без вложенного settings)
		if rulesVal, ok := m["enabled-rules"]; ok {
			if rulesSlice, ok := rulesVal.([]any); ok {
				for _, r := range rulesSlice {
					if ruleStr, ok := r.(string); ok {
						p.cfg.EnabledRules = append(p.cfg.EnabledRules, ruleStr)
					}
				}
			}
		}

		// Парсим sensitive-patterns
		if patternsVal, ok := m["sensitive-patterns"]; ok {
			if patternsSlice, ok := patternsVal.([]any); ok {
				for _, pat := range patternsSlice {
					if patStr, ok := pat.(string); ok {
						p.cfg.SensitivePatterns = append(p.cfg.SensitivePatterns, patStr)
					}
				}
			}
		}

		fmt.Fprintf(os.Stderr, "DEBUG: final config: %+v\n", p.cfg)
	} else {
		// Если конфиг не передан, можно использовать дефолтный
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
