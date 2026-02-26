# slogzaplint🔍 

**slogzaplint** - custom plugin module for 'golangci-lint' standardization of messages in logs. It supports most popular LogLib - "log/slog", "go.uber.org/zap"

## 📦 Installation and usage

### locally <D-F2>(without golangci-lint) 
```bash
git clone https://gitlab.com/lainiwakuras-father/slogzaplint.git
cd slogzaplint
go run ./cmd/slogzaplint/main.go  $path_your_project
```
### As a golangci-lint plugin module (recommended)

1. Install [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/)

2. Add a plugin configuration file .custom-gcl in root of the project
```yaml
#Example
version: v2.10.1
plugins:
  - module: 'github.com/lainiwakuras-father/slogzaplint'
    import: 'github.com/lainiwakuras-father/slogzaplint/pkg/analyzer'
    version: main
```

3. Add a Golangci-lint configuration file .golangci.yml in root of the project 
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
        description: This is an standart message logs usage of a plugin linter.
        original-url: "https://github.com/lainiwakuras-father/slogzaplint"
        settings:
          enabled-rules:
            - lowercase
            - english
            - no-special
            - no-sensitive
        sensitive-patterns:
          - "api_key"
          - "password"
          - "token"

```

4. build the custom linter binary
```bash
cd /path/to/your_project
golangci-lint custom --name your_name --destination /your/path/
```
5. 
```bash
./custom-gcl run main.go
```
6. (Optional), Integrate with github action

Add the following job to your .github/workflows/ci.yml to run the linter automatically

```
 lint:
    name:  Lint with slogzaplint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v5
      - uses: actions/setup-go@v6
        with:
          go-version: stable
      - name: run golangci-lint (with custom plugin)
        uses: golangci/golangci-lint-action@v9
        with:
          version: v2.10.1

```


## 📁 project structure 
```
slogzaplint/
├── cmd/                               # Standalone entry point
│   └── slogzaplint/
│       └── main.go
├── pkg/                               # Core packages
│   ├── analyzer/                       # Main analyzer logic
│   │   ├── analyzer.go
│   │   ├── analyzer_test.go
│   │   ├── checker.go                   # AST traversal & rule application
│   │   ├── golangci.go                   # Plugin glue code
│   │   ├── extractstring.go               # Extracts log message from AST
│   │   └── islogcall.go                    # Detects logging function calls
│   ├── rules/                           # Individual lint rules
│   │   ├── msg.go
│   │   ├── lowercase.go
│   │   ├── english.go
│   │   ├── nospecial.go
│   │   └── sensitive.go
│   └── config/                          # Configuration loader
│       └── config.go
├── testdata/                            # Test fixtures for analysistest
│   ├── src/
│   │   ├── slog/
│   │   │   ├── lowercase.go
│   │   │   ├── english.go
│   │   │   ├── nospecial.go
│   │   │   └── sensitive.go
│   │   └── zap/
│   │       ├── lowercase.go
│   │       ├── english.go
│   │       ├── nospecial.go
│   │       └── sensitive.go
├── .golangci-example.yml                # Example golangci-lint config
├── .custom-gcl-example.yml              # Example plugin config
├── go.mod
├── go.sum
└── README.md                            # This file
```

## HOW ITS WORK
My project - [Valentine-VK-Bot]()

CI workflows
![Example](docs/1.jpg)

## Author 

Developed as a test task for Backend Developer (Golang).

