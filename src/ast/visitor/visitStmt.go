package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
)

func (v *AntlrVisitor) visitStmt(ctx *parser.StmtContext) node.Stmt {
	if ctx.FuncCall() != nil {
		return v.visitFuncCall(ctx.FuncCall().(*parser.FuncCallContext))
	}
	if ctx.AssignStmt() != nil {
		return v.visitAssignStmt(ctx.AssignStmt().(*parser.AssignStmtContext))
	}
	if ctx.CodeBlock() != nil {
		return v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext))
	}
	if ctx.IfStmt() != nil {
		return v.visitIfStmt(ctx.IfStmt().(*parser.IfStmtContext))
	}
	if ctx.JumpStmt() != nil {
		return v.visitJumpStmt(ctx.JumpStmt().(*parser.JumpStmtContext))
	}
	if ctx.LoopStmt() != nil {
		return v.visitLoopStmt(ctx.LoopStmt().(*parser.LoopStmtContext))
	}
	if ctx.MatchStmt() != nil {
		return v.visitMathStmt(ctx.MatchStmt().(*parser.MatchStmtContext))
	}
	fmt.Println(ctx.GetText())
	v.Errs = append(v.Errs, VisitorError{
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
		EndLine:   ctx.GetStop().GetLine(),
		EndColumn: ctx.GetStop().GetColumn() + len([]rune(ctx.GetStop().GetText())),
		Msg:       "unknown stmt: " + ctx.GetText(),
		Text:      ctx.GetText(),
		Raw:       fmt.Errorf("unknown stmt: %s", ctx.GetText()),
	})
	return nil
}

func (v *AntlrVisitor) visitJumpStmt(ctx *parser.JumpStmtContext) node.Stmt {
	if ctx.Return() != nil {
		ret := node.ReturnStmt{
			Token: token.FromAntlrToken(ctx.GetStart()),
		}
		ret.Value = toExpr(v.visitExprWithLambda(toPtr[parser.ExprWithLambdaContext](ctx.ExprWithLambda())))

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

func (v *AntlrVisitor) visitIfStmt(ctx *parser.IfStmtContext) *node.IfStmt {
	cdx := ctx.AllCodeBlock()
	i := node.IfStmt{
		Token:     token.FromAntlrToken(ctx.GetStart()),
		IfTrue:    v.visitCodeBlock(cdx[0].(*parser.CodeBlockContext)),
		Condition: v.visitExpr(ctx.Expr().(*parser.ExprContext)),
	}
	if len(cdx) == 2 {
		i.IfFalse = v.visitCodeBlock(cdx[1].(*parser.CodeBlockContext))
	}

	return &i
}

func (v *AntlrVisitor) visitAssignStmt(ctx *parser.AssignStmtContext) *node.AssignStmt {
	n := node.AssignStmt{
		Token: token.FromAntlrToken(ctx.Assign().GetSymbol()),
		Type:  v.visitType(toPtr[parser.TypeContext](ctx.Type_())),
		Var:   v.visitVar(ctx.Var_().(*parser.VarContext)),
		Value: v.visitExprWithLambda(ctx.ExprWithLambda().(*parser.ExprWithLambdaContext)),
	}
	return &n
}

func (v *AntlrVisitor) visitIdentifier(n antlr.TerminalNode) *node.Identifier {
	return &node.Identifier{
		Token: token.FromAntlrToken(n.GetSymbol()),
		Value: n.GetText(),
	}
}

func (v *AntlrVisitor) visitType(ctx *parser.TypeContext) *node.Type {
	// fmt.Println("VisitType")
	if ctx == nil || ctx.TypeName == nil {
		return nil
	}

	t := node.Type{}
	if ctx.PackageName != nil {
		t.Package = ctx.PackageName.GetText()
	}
	t.Name = ctx.TypeName.GetText()
	t.Token = token.FromAntlrToken(ctx.GetStart())

	return &t
}

func (v *AntlrVisitor) visitVar(ctx *parser.VarContext) *node.Variable {
	var bs []*node.BaseVariable
	bvs := ctx.AllBaseVar()
	for _, bv := range bvs {
		bs = append(bs, v.visitBaseVar(bv.(*parser.BaseVarContext)))
	}
	return &node.Variable{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Value: bs,
	}
}

func (v *AntlrVisitor) visitBaseVar(ctx *parser.BaseVarContext) *node.BaseVariable {
	id := ctx.Identifier()
	return &node.BaseVariable{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Name: &node.Identifier{
			Token: token.FromAntlrToken(id.GetSymbol()),
			Value: id.GetText(),
		},
		Index: v.visitIndexes(
			toPtr[parser.IndexesContext](
				ctx.Indexes(),
			),
		),
	}
}

func (v *AntlrVisitor) visitIndexes(ctx *parser.IndexesContext) []node.Expr {
	if ctx == nil || ctx.GetChildCount() == 0 {
		return nil
	}
	var idx []node.Expr
	for _, kid := range ctx.GetChildren() {
		idx = append(idx, v.visitIndex(kid.(*parser.IndexContext)))
	}
	return idx
}

func (v *AntlrVisitor) visitIndex(ctx *parser.IndexContext) node.Expr {
	// [ expr ] <- RSquare
	// ^
	// LSquare
	return v.visitExpr(ctx.GetChild(1).(*parser.ExprContext))
}
