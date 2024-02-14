package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
)

func (v *AntlrVisitor) visitStmt(ctx parser.IStmtContext) node.Stmt {
	if ctx.FuncCall() != nil {
		return v.visitFuncCall(ctx.FuncCall())
	}
	if ctx.AssignStmt() != nil {
		return v.visitAssignStmt(ctx.AssignStmt())
	}
	if ctx.CodeBlock() != nil {
		return v.visitCodeBlock(ctx.CodeBlock())
	}
	if ctx.IfStmt() != nil {
		return v.visitIfStmt(ctx.IfStmt())
	}
	if ctx.JumpStmt() != nil {
		return v.visitJumpStmt(ctx.JumpStmt())
	}
	if ctx.LoopStmt() != nil {
		return v.visitLoopStmt(ctx.LoopStmt())
	}
	if ctx.TryCatchSmt() != nil {
		return v.visitTryCatch(ctx.TryCatchSmt())
	}
	if ctx.MatchStmt() != nil {
		return v.visitMathStmt(ctx.MatchStmt())
	}
	if ctx.DeclareStmt() != nil {
		return v.visitDeclareStmt(ctx.DeclareStmt())
	}
	v.appendErr(ctx, "unknown stmt: "+ctx.GetText(), nil)
	return nil
}

func (v *AntlrVisitor) visitJumpStmt(ctx parser.IJumpStmtContext) node.Stmt {
	if ctx.Return() != nil {
		var retV []node.Expr
		if len(ctx.AllExprWithLambda()) > 0 {
			for _, kid := range ctx.AllExprWithLambda() {
				retV = append(retV, v.visitExprWithLambda(kid))
			}
		}
		ret := node.ReturnStmt{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			Value: retV,
		}
		return &ret
	}
	if ctx.Break() != nil {
		return &node.BreakStmt{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		}
	}

	if ctx.Continue() != nil {
		return &node.ContinueStmt{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		}
	}
	v.appendErr(ctx, "unknown jump statement", nil)
	return nil
}

func (v *AntlrVisitor) visitIfStmt(ctx parser.IIfStmtContext) *node.IfStmt {
	if ctx == nil {
		return nil
	}
	return &node.IfStmt{
		Token:            token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		IfTrue:           v.visitCodeBlock(ctx.CodeBlock(0)),
		IfFalseCodeBlock: v.visitCodeBlock(ctx.CodeBlock(1)),
		IfFalseIfStmt:    v.visitIfStmt(ctx.IfStmt()),
		Condition:        v.visitExpr(ctx.Expr()),
	}
}

func (v *AntlrVisitor) visitAssignStmt(ctx parser.IAssignStmtContext) *node.AssignStmt {
	var assignee []*node.Assignee
	var val node.Expr
	if ctx.Vars() != nil {
		assignee = v.visitAssignStmtAssignee(ctx.Vars())
		val = v.visitCommaExpr(ctx.CommaExpr())
	} else {
		assignee = append(assignee, &node.Assignee{
			Type: v.visitType(ctx.Type_()),
			Var:  v.visitVar(ctx.Var_()),
			Ref:  ctx.Ref() != nil,
		})
		val = v.visitExprWithLambda(ctx.ExprWithLambda())
	}
	n := node.AssignStmt{
		Token:    token.FromAntlrToken(ctx.Assign().GetSymbol()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Assignee: assignee,
		Value:    val,
	}
	return &n
}

func (v *AntlrVisitor) visitAssignStmtAssignee(ctx parser.IVarsContext) []*node.Assignee {
	var assignee []*node.Assignee
	for _, kid := range ctx.AllVar_() {
		assignee = append(assignee, &node.Assignee{
			Var: v.visitVar(kid),
		})
	}
	return assignee
}

func (v *AntlrVisitor) visitIdentifier(n antlr.TerminalNode) *node.Identifier {
	return &node.Identifier{
		Token: token.FromAntlrToken(n.GetSymbol()), // Identifier
		Value: n.GetText(),
	}
}

func (v *AntlrVisitor) visitType(ctx parser.ITypeContext) *node.Type {
	// fmt.Println("VisitType")
	if ctx == nil {
		return nil
	}
	if ctx.Map() != nil {
		return &node.Type{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			Map:   true,
		}
	}

	if ctx.Function() != nil {
		return &node.Type{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			Func:  true,
		}
	}

	if ctx.GetTypeName() == nil {
		return nil
	}

	t := node.Type{
		Token:    token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Name:     ctx.GetTypeName().GetText(),
		Nullable: ctx.Question() != nil,
		Array:    ctx.LSquare() != nil,
	}
	if ctx.GetPackageName() != nil {
		t.Package = ctx.GetPackageName().GetText()
	}

	return &t
}

func (v *AntlrVisitor) visitVar(ctx parser.IVarContext) *node.Variable {
	var bs []*node.BaseVariable
	bvs := ctx.AllBaseVar()
	for _, bv := range bvs {
		bs = append(bs, v.visitBaseVar(bv))
	}
	return &node.Variable{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Value: bs,
	}
}

func (v *AntlrVisitor) visitBaseVar(ctx parser.IBaseVarContext) *node.BaseVariable {
	id := ctx.Identifier()
	return &node.BaseVariable{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Name: &node.Identifier{
			Token: token.FromAntlrToken(id.GetSymbol()).WithEnd(ctx.GetStop()),
			Value: id.GetText(),
		},
		Index: v.visitIndexes(ctx.Indexes()),
	}
}

func (v *AntlrVisitor) visitIndexes(ctx parser.IIndexesContext) []node.Expr {
	if ctx == nil || ctx.GetChildCount() == 0 {
		return nil
	}
	var idx []node.Expr
	for _, kid := range ctx.AllIndex() {
		idx = append(idx, v.visitIndex(kid))
	}
	return idx
}

func (v *AntlrVisitor) visitIndex(ctx parser.IIndexContext) node.Expr {
	// [ expr ] <- RSquare
	// ^
	// LSquare
	return v.visitExpr(ctx.GetChild(1).(*parser.ExprContext))
}
