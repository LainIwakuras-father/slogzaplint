# slogzaplint 
Custom plugin module for 'golangci-lint' standardization of messages in logs

## Assembly and Launch Instructions

### Locally 
```bash
git clone https://gitlab.com/LainIwakuras-father/slogzaplint.git
cd slogzaplint
go run ./cmd/slogzaplint/main.go  $PATH_YOUR_PROJECT
```
### As plugin golangci-lint



1. Install [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/)

2. add file .golangci.yml your project and configure 
```yaml
#Example
version: "2"

linters:
  default: none
  enable:
    - slogzaplint
  settings:
    custom:
      slogzaplint:
        type: "module"
        description: This is an logging usage of a plugin linter.
        original-url: "https://github.com/LainIwakuras-father/slogzaplint"
        settings:
          enabled-rules:
            - lowercase
            - english
            - no-special
            - no-sensitive
        sensitive-patterns:
          - "api_key"

```
. add file .custom-gcl your project and configure
```yaml
version: v2.10.1
plugins:
  - module: 'github.com/LainIwakuras-father/slogzaplint'
    import: 'github.com/LainIwakuras-father/slogzaplint/pkg/analyzer'
    path: path/to/plugin_lint

```

4. build binary + plugin slogzaplint
```bash
cd path/to/plugin_lint
golangci-lint custom -v
cp custom-gcl path/to/your_project
```
5. run linter (optional, move binary file ./custom-gcl in directory with your project)
```bash
./custom-gcl run main.go
``` 


## 📁 Project structure 
```
slogzaplint/
├── cmd/                       # Standalone executable entry point
│   └── slogzaplint/
│       └── main.go            # CLI entry for running the linter directly
├── pkg/                       # Core packages (reusable across the project)
│   ├── analyzer/               # Main analyzer logic
│   │   ├── analyzer.go         # Analyzer definition and run function
│   │   ├── analyzer_test.go         # Analyzer definition and run function
│   │   ├── checher.go         # Analyzer definition and run function
│   │   ├── golangci.go         # Analyzer definition and run function
│   │   ├── extractString.go          # Helpers to extract string messages from AST
│   │   └── isLogCall.go      # Detection of log function calls
│   ├── rules/                   # Individual lint rules
│   │   ├── msg.go      
│   │   ├── lowercase.go        # Rule: message starts with lowercase
│   │   ├── english.go          # Rule: only English letters
│   │   ├── nospecial.go        # Rule: no special chars/emojis
│   │   └── sensitive.go        # Rule: no sensitive data
│   ├── config/                  # Configuration handling
│      └── config.go           # Load and validate YAML config
├── testdata/                      # Test fixtures for analysistest
│   ├── src/                       # Source files under test
│       ├── slog/        
│       │   ├── lowercase.go     
│       │   ├── english.go     
│       │   ├── nospecial.go  
│       │   └── sensitive.go
│       └── zap/             
│           ├── lowercase.go 
│           ├── english.go  
│           ├── nospecial.go
│           └── sensitive.go
├── .gitlab-ci.yml                  # GitLab CI configuration
├── .golangci.yml                    # Example configuration for golangci-lint
├── .custom-gcl.yml                    # Example configuration for golangci-lint
├── go.mod
├── go.sum
└── README.md                        # This file
```

## Author 

Developed as a test task for Backend Developer (Golang).

