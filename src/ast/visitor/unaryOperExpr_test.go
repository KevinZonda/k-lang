package visitor_test

import (
	"testing"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
)

func TestUnaryOperExpr(t *testing.T) {
	a := parserHelper.Ast("-1")
	if len(a) != 1 {
		t.Error("Expected 1 node, got", len(a))
	}
	expr, ok := a[0].(node.Expr)
	if !ok {
		t.Error("Expected Expr node, got", a[0])
	}
	if expr.TokenValue() != "-" {
		t.Error("Expected -, got", expr.TokenValue())
	}
	e, ok := expr.(*node.UnaryOperExpr)
	if !ok {
		t.Error("Expected UnaryOperExpr, got", expr)
	}
	if e.Expr.TokenValue() != "1" {
		t.Error("Expected 1, got", e.Expr.TokenValue())
	}
	if e.Oper != "-" {
		t.Error("Expected -, got", e.Oper)
	}
}
