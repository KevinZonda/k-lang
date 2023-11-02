package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
)

func (v *AntlrVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if ctx.FuncCall() != nil {
		return v.VisitFuncCall(ctx.FuncCall().(*parser.FuncCallContext))
	}
	if ctx.AssignStmt() != nil {
		return v.VisitAssignStmt(ctx.AssignStmt().(*parser.AssignStmtContext))
	}
	if ctx.CodeBlock() != nil {
		return v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext))
	}
	if ctx.IfStmt() != nil {
		return v.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext))
	}
	if ctx.JumpStmt() != nil {
		return v.VisitJumpStmt(ctx.JumpStmt().(*parser.JumpStmtContext))
	}
	if ctx.LoopStmt() != nil {
		return v.VisitLoopStmt(ctx.LoopStmt().(*parser.LoopStmtContext))
	}
	panic("implement me")

}

func (v *AntlrVisitor) VisitLoopStmt(ctx *parser.LoopStmtContext) interface{} {
	if ctx.CStyleFor() != nil {
		return v.VisitCStyleFor(ctx.CStyleFor().(*parser.CStyleForContext))
	}
	if ctx.WhileStyleFor() != nil {
		return v.VisitWhileStyleFor(ctx.WhileStyleFor().(*parser.WhileStyleForContext))
	}
	if ctx.IterFor() != nil {
		return v.VisitIterStyleFor(ctx.IterFor().(*parser.IterForContext))
	}
	panic("implement me")
}

func (v *AntlrVisitor) VisitCStyleFor(ctx *parser.CStyleForContext) any {
	n := node.CStyleFor{
		Token: token.FromAntlrToken(ctx.For().GetSymbol()),
		Body:  v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock),
	}
	if ctx.ExprOrAssign() != nil {
		n.InitialExpr = v.VisitExprOrAssign(ctx.ExprOrAssign().(*parser.ExprOrAssignContext)).(node.Expr)
	}

	if ctx.GetOnCondition() != nil {
		n.ConditionExpr = v.VisitExpr(ctx.GetOnCondition().(*parser.ExprContext)).(node.Expr)
	}

	if ctx.GetOnEnd() != nil {
		n.AfterIterExpr = v.VisitExpr(ctx.GetOnEnd().(*parser.ExprContext)).(node.Expr)
	}

	panic("not supported stmt")
}

func (v *AntlrVisitor) VisitWhileStyleFor(ctx *parser.WhileStyleForContext) any {
	n := node.WhileStyleFor{
		Token: token.FromAntlrToken(ctx.For().GetSymbol()),
		Body:  v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock),
	}
	if ctx.Expr() != nil {
		n.ConditionExpr = v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr)
	}
	return &n
}

func (v *AntlrVisitor) VisitIterStyleFor(ctx *parser.IterForContext) any {
	n := node.IterStyleFor{
		Token:    token.FromAntlrToken(ctx.For().GetSymbol()),
		Variable: v.visitIdentifier(ctx.Identifier()),
		Iterator: v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr),
		Body:     v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock),
	}
	if ctx.Type_() != nil {
		t := v.VisitType(typeCastToPtr[parser.TypeContext](ctx.Type_()))
		if t != nil {
			n.Type = t.(*node.Identifier)
		}
	}
	return &n
}

func (v *AntlrVisitor) VisitJumpStmt(ctx *parser.JumpStmtContext) interface{} {
	if ctx.Return() != nil {
		ret := node.ReturnStmt{
			Token: token.FromAntlrToken(ctx.GetStart()),
		}
		if ctx.Expr() != nil {
			ret.Value = v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr)
		}
		return &ret
	}
	if ctx.Break() != nil {
		return &node.BreakStmt{
			Token: token.FromAntlrToken(ctx.GetStart()),
		}
	}

	if ctx.Continue() != nil {
		return &node.ContinueStmt{
			Token: token.FromAntlrToken(ctx.GetStart()),
		}
	}
	panic("Unknown jumpStmt")
}

func (v *AntlrVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cdx := ctx.AllCodeBlock()
	i := node.IfStmt{
		Token:     token.FromAntlrToken(ctx.GetStart()),
		IfTrue:    v.VisitCodeBlock(cdx[0].(*parser.CodeBlockContext)).(*node.CodeBlock),
		Condition: v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr),
	}
	if len(cdx) == 2 {
		i.IfFalse = v.VisitCodeBlock(cdx[1].(*parser.CodeBlockContext)).(*node.CodeBlock)
	}

	return &i
}

func (v *AntlrVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	n := node.AssignStmt{
		Token: token.FromAntlrToken(ctx.Assign().GetSymbol()),
		Type:  typeCastToPtr[node.Identifier](v.VisitType(typeCastToPtr[parser.TypeContext](ctx.Type_()))),
		Var:   v.VisitVar(ctx.Var_().(*parser.VarContext)).(*node.Variable),
		Value: v.VisitExprWithLambda(ctx.ExprWithLambda().(*parser.ExprWithLambdaContext)).(node.Expr),
	}
	return &n
}

func (v *AntlrVisitor) visitIdentifier(n antlr.TerminalNode) *node.Identifier {
	return &node.Identifier{
		Token: token.FromAntlrToken(n.GetSymbol()),
		Value: n.GetText(),
	}
}

func (v *AntlrVisitor) VisitType(ctx *parser.TypeContext) any {
	// fmt.Println("VisitType")
	if ctx == nil || ctx.Identifier() == nil {
		return nil
	}
	return v.visitIdentifier(ctx.Identifier())
}

func (v *AntlrVisitor) VisitVar(ctx *parser.VarContext) any {
	var bs []*node.BaseVariable
	bvs := ctx.AllBaseVar()
	for _, bv := range bvs {
		bs = append(bs, v.VisitBaseVar(bv.(*parser.BaseVarContext)).(*node.BaseVariable))
	}
	return &node.Variable{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Value: bs,
	}
}

func (v *AntlrVisitor) VisitBaseVar(ctx *parser.BaseVarContext) any {
	id := ctx.Identifier()
	return &node.BaseVariable{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Name: &node.Identifier{
			Token: token.FromAntlrToken(id.GetSymbol()),
			Value: id.GetText(),
		},
		Index: typeCastToArr[node.Expr](
			v.VisitIndexes(
				typeCastToPtr[parser.IndexesContext](
					ctx.Indexes(),
				),
			)),
	}
}

func typeCastToPtr[T any](i any) *T {
	if i == nil {
		return nil
	}
	return i.(*T)
}

func typeCastToArr[T any](v any) []T {
	if v == nil {
		return nil
	}
	return v.([]T)
}

func (v *AntlrVisitor) VisitIndexes(ctx *parser.IndexesContext) any {
	if ctx == nil || ctx.GetChildCount() == 0 {
		return nil
	}
	var idx []node.Expr
	for _, kid := range ctx.GetChildren() {
		idx = append(idx, v.VisitIndex(kid.(*parser.IndexContext)).(node.Expr))
	}
	return idx
}

func (v *AntlrVisitor) VisitIndex(ctx *parser.IndexContext) any {
	// [ expr ] <- RSquare
	// ^
	// LSquare
	return v.VisitExpr(ctx.GetChild(1).(*parser.ExprContext))
}
