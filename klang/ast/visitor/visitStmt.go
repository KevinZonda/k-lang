package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/parser"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/utils/jout"
	"reflect"
)

func (v *AntlrVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	if ctx.AssignStmt() != nil {
		return v.VisitAssignStmt(ctx.AssignStmt().(*parser.AssignStmtContext))
	}
	panic("implement me")

}

func (v *AntlrVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	for i, n := range ctx.GetChildren() {
		fmt.Println(i, n, reflect.TypeOf(n))
	}
	n := node.AssignStmt{
		Token: token.FromAntlrToken(ctx.Assign().GetSymbol()),
		Type:  v.VisitType(ctx.Type_().(*parser.TypeContext)).(*node.Identifier),
		Var:   v.VisitVar(ctx.Var_().(*parser.VarContext)).(*node.Variable),
		Value: v.VisitExprWithLambda(ctx.ExprWithLambda().(*parser.ExprWithLambdaContext)).(node.Expr),
	}
	jout.Println(n)
	return &n
}

func (v *AntlrVisitor) VisitType(ctx *parser.TypeContext) any {
	// fmt.Println("VisitType")
	if ctx == nil || ctx.Identifier() == nil {
		return nil
	}
	n := ctx.Identifier()
	id := node.Identifier{
		Token: token.FromAntlrToken(n.GetSymbol()),
		Value: n.GetText(),
	}
	// fmt.Println(id)
	return &id
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
