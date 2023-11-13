package visitor

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func toPtr[T any](i any) *T {
	if i == nil {
		return nil
	}
	return i.(*T)
}

func toArr[T any](v any) []T {
	if v == nil {
		return nil
	}
	return v.([]T)
}

func toExpr(x any) node.Expr {
	if x == nil {
		return nil
	}
	return x.(node.Expr)
}

func toNode(x any) node.Node {
	if x == nil {
		return nil
	}
	return x.(node.Node)
}
