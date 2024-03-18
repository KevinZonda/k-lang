package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitFuncCall(ctx parser.IFuncCallContext) *node.FuncCall {
	fc := &node.FuncCall{
		Token:  token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Caller: v.visitIdentifier(ctx.Identifier()),
		Args:   v.visitFuncCallArgs(ctx.FuncCallArgs()),
	}
	return fc
}

func (v *AntlrVisitor) visitFuncCallArgs(ctx parser.IFuncCallArgsContext) []node.Expr {
	if ctx == nil {
		return nil
	}
	var args []node.Expr
	for _, expr := range ctx.AllExprWithLambda() {
		args = append(args, v.visitExprWithLambda(expr))
	}
	return args
}
