package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	fc := &node.FuncCall{
		Token:  token.FromAntlrToken(ctx.GetStart()),
		Caller: v.VisitVar(ctx.Var_().(*parser.VarContext)).(*node.Variable),
	}
	if fca := ctx.FuncCallArgs(); fca != nil {
		args := v.VisitFuncCallArgs(ctx.FuncCallArgs().(*parser.FuncCallArgsContext))
		if args != nil {
			fc.Args = args.([]node.Expr)
		}
	}
	return fc
}

func (v *AntlrVisitor) VisitFuncCallArgs(ctx *parser.FuncCallArgsContext) interface{} {
	var args []node.Expr
	for _, expr := range ctx.AllExpr() {
		args = append(args, v.VisitExpr(expr.(*parser.ExprContext)).(node.Expr))
	}
	return args
}
