package main

import (
	"github.com/LainIwakuras-father/slogzaplint/pkg/analyzer"
	"github.com/LainIwakuras-father/slogzaplint/pkg/config"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	analyzer, err := analyzer.NewAnalyzer(*config.Default())
	if err != nil {
		panic(err)
	}
	singlechecker.Main(analyzer)
}
