package pruner

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
)

func Prune(ast tree.Ast) tree.Ast {
	pAst := make(tree.Ast, 0)
	for _, node := range ast {
		pNode := prune(node)
		if pNode != nil {
			pAst = append(pAst, pNode)
		}
	}
	return pAst
}

func prune(n node.Node) node.Node {
	if n == nil {
		return nil
	}
	switch n.(type) {
	case node.Expr:
		n = pruneExpr(n.(node.Expr)).(node.Expr)
	case node.Stmt:
		n = pruneStmt(n.(node.Stmt)).(node.Stmt)
	}
	return n
}

func pruneUnaryExpr(n *node.UnaryOperExpr) any {
	right := pruneExprAndVal(n.Expr)
	switch right.(type) {
	case bool:
		if n.Token.Kind == token.Not {
			return !right.(bool)
		}
	}
	return n
}

func pruneBinaryExpr(n *node.BinaryOperExpr) any {
	left := reconstruct(pruneExprAndVal(n.Left)).(node.Expr)
	right := reconstruct(pruneExprAndVal(n.Right)).(node.Expr)

	switch left.(type) {
	case node.Node:
		n.Left = left.(node.Expr)
		n.Right = reconstruct(right).(node.Expr)
		return n
	default:
		switch right.(type) {
		case node.Node:
			n.Left = left.(node.Expr)
			n.Right = reconstruct(right).(node.Expr)
			return n
		default:
			return binaryOperEval.BinaryOper(n.Token.Kind, left, right)
		}
	}
	return n
}

func pruneExprAndVal(n node.Expr) any {
	switch n.(type) {
	case *node.BinaryOperExpr:
		return pruneBinaryExpr(n.(*node.BinaryOperExpr))
	case *node.UnaryOperExpr:
		return pruneUnaryExpr(n.(*node.UnaryOperExpr))
	case node.ILiteralValue:
		return n.(node.ILiteralValue).ConstVal()
	}
	return n
}

func pruneExpr(n node.Expr) any {
	var r any
	switch n.(type) {
	case *node.BinaryOperExpr:
		r = pruneBinaryExpr(n.(*node.BinaryOperExpr))
	case *node.UnaryOperExpr:
		r = pruneUnaryExpr(n.(*node.UnaryOperExpr))
	case node.ILiteralValue:
		return n
	}
	if r != nil {
		// reconstruct the node
		return reconstruct(r)
	}
	return n
}

func reconstruct(r any) node.Node {
	if r == nil {
		return nil
	}
	switch r.(type) {
	case node.Node:
		return r.(node.Node)
	case int:
		return &node.IntLiteral{Value: r.(int), Token: token.Token{Kind: token.IntegerLiteral, Value: fmt.Sprint(r.(int))}}
	case float64:
		return &node.FloatLiteral{Value: r.(float64), Token: token.Token{Kind: token.NumberLiteral, Value: fmt.Sprint(r.(float64))}}
	case string:
		return &node.StringLiteral{Value: r.(string), Token: token.Token{Kind: token.StringLiteral, Value: fmt.Sprint(r.(string))}}
	case bool:
		kind := token.False
		if r.(bool) {
			kind = token.True
		}
		return &node.BoolLiteral{Value: r.(bool), Token: token.Token{Kind: kind, Value: fmt.Sprint(r.(bool))}}
	}
	panic("not supported type")
}

func pruneStmt(n node.Stmt) any {
	return n
}
