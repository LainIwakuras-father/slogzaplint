package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/LainIwakuras-father/slogzaplint/pkg/config"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}
	analyzer, err := NewAnalyzer(*config.Default())
	if err != nil {
		t.Fatalf("Failed Default config or analyzer engine: %s", err)
	}
	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, analyzer, "slog", "zap")
}
