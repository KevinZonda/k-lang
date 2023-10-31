package visitor_test

import (
	"testing"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
)

func TestSingleExpressionParse(t *testing.T) {
	ast := parserHelper.Ast("1 + 2.2")
	if len(ast) != 1 {
		t.Error("Expected 1 node, got", len(ast))
	}
	expr, ok := ast[0].(node.Expr)
	if !ok {
		t.Error("Expected Expr node, got", ast[0])
	}
	if expr.TokenValue() != "+" {
		t.Error("Expected +, got", expr.TokenValue())
	}
	e, ok := expr.(*node.BinaryOperExpr)
	if !ok {
		t.Error("Expected BinaryOperExpr, got", expr)
	}
	if e.Left.TokenValue() != "1" {
		t.Error("Expected 1, got", e.Left.TokenValue())
	}
	if e.Right.TokenValue() != "2.2" {
		t.Error("Expected 2.2, got", e.Right.TokenValue())
	}
	if e.Oper != "+" {
		t.Error("Expected +, got", e.Oper)
	}
}
