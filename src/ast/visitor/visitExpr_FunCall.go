package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitFuncCall(ctx *parser.FuncCallContext) *node.FuncCall {
	fc := &node.FuncCall{
		Token:  token.FromAntlrToken(ctx.GetStart()),
		Caller: v.visitIdentifier(ctx.Identifier()),
	}
	if fca := ctx.FuncCallArgs(); fca != nil {
		args := v.visitFuncCallArgs(ctx.FuncCallArgs().(*parser.FuncCallArgsContext))
		if args != nil {
			fc.Args = args
		}
	}
	return fc
}

func (v *AntlrVisitor) visitFuncCallArgs(ctx *parser.FuncCallArgsContext) []node.Expr {
	var args []node.Expr
	for _, expr := range ctx.AllExpr() {
		args = append(args, v.visitExpr(expr.(*parser.ExprContext)))
	}
	return args
}
