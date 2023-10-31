package visitor_test

import (
	"testing"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/parserHelper"
)

func TestParenthesesExpr(t *testing.T) {
	ast := parserHelper.Ast("(1+2.2) * (3 -8)")
	if len(ast) != 1 {
		t.Error("Expected 1 node, got", len(ast))
	}
	expr, ok := ast[0].(node.Expr)
	if !ok {
		t.Error("Expected Expr node, got", ast[0])
	}
	if expr.TokenValue() != "*" {
		t.Error("Expected +, got", expr.TokenValue())
	}

	e, ok := expr.(*node.BinaryOperExpr)

	if !ok {
		t.Error("Expected BinaryOperExpr, got", expr)
	}
	if e.Left.TokenValue() != "+" {
		t.Error("Expected +, got", e.Left.TokenValue())
	}
	if e.Right.TokenValue() != "-" {
		t.Error("Expected -, got", e.Right.TokenValue())
	}
	if e.Oper != "*" {
		t.Error("Expected *, got", e.Oper)
	}
	lhs := e.Left.(*node.BinaryOperExpr)
	if lhs.Left.TokenValue() != "1" {
		t.Error("Expected 1, got", lhs.Left.TokenValue())
	}
	if lhs.Right.TokenValue() != "2.2" {
		t.Error("Expected 2.2, got", lhs.Right.TokenValue())
	}
	if lhs.Oper != "+" {
		t.Error("Expected +, got", lhs.Oper)
	}

	rhs := e.Right.(*node.BinaryOperExpr)
	if rhs.Left.TokenValue() != "3" {
		t.Error("Expected 3, got", rhs.Left.TokenValue())
	}
	if rhs.Right.TokenValue() != "8" {
		t.Error("Expected 8, got", rhs.Right.TokenValue())
	}
	if rhs.Oper != "-" {
		t.Error("Expected -, got", rhs.Oper)
	}
}
