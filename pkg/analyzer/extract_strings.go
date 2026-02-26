package analyzer

import (
	"go/ast"
	"go/token"
	"strings"
)

func extractString(call *ast.CallExpr) (string, token.Pos, bool) {
	if len(call.Args) < 1 {
		return "", 0, false
	}
	expr := call.Args[0]
	pos := expr.Pos()
	var sb strings.Builder
	var foundLiteral bool

	var walk func(ast.Expr)
	walk = func(e ast.Expr) {
		switch ex := e.(type) {
		case *ast.BasicLit:
			if ex.Kind == token.STRING {
				sb.WriteString(strings.Trim(ex.Value, `"`))
				foundLiteral = true
			}
		case *ast.BinaryExpr:
			if ex.Op == token.ADD {
				walk(ex.X)
				walk(ex.Y)
			}
		}
	}

	walk(expr)
	if !foundLiteral {
		return "", 0, false
	}
	return sb.String(), pos, true
}
