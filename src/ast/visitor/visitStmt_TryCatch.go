package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitTryCatch(ctx parser.ITryCatchSmtContext) *node.TryCatchStmt {
	if ctx == nil {
		return nil
	}
	tcs := node.TryCatchStmt{
		Token: token.FromAntlrToken(ctx.Try().GetSymbol()).WithBegin(ctx.GetStart()),
		Try:   v.visitCodeBlock(ctx.CodeBlock()),
		Catch: v.visitCatch(ctx.CatchStmt()),
	}
	//cs := ctx.AllCatchStmt()
	//for _, c := range cs {
	//	if cb := v.visitCatch(c); cb != nil {
	//		tcs.Catch = append(tcs.Catch, cb)
	//	}
	//
	//}
	return &tcs
}

func (v *AntlrVisitor) visitCatch(ctx parser.ICatchStmtContext) *node.CatchBranch {
	if ctx == nil {
		return nil
	}
	cr := node.CatchBranch{
		Content: v.visitCodeBlock(ctx.CodeBlock()),
	}
	if ctx.LParen() == nil {
		//cr.CatchAll = true
		return &cr
	}
	cr.VarName = v.visitIdentifier(ctx.Identifier()).Value
	cr.VarType = v.visitType(ctx.Type_())
	return &cr
}
