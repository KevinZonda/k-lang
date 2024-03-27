package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitFuncBlock(ctx parser.IFuncBlockContext) *node.FuncBlock {
	fx := v.visitFuncSig(ctx.FuncSig().(*parser.FuncSigContext))
	fx.Token = token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop())
	fx.Body = v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext))
	return fx
}

func (v *AntlrVisitor) visitLambda(ctx parser.ILambdaContext) *node.LambdaExpr {
	fb := &node.LambdaExpr{
		Token:   token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Body:    v.visitCodeBlock(ctx.CodeBlock()),
		Args:    v.visitFuncSignArgs(ctx.FuncSignArgs()),
		RetType: v.visitFuncReturnType(ctx.FuncReturnType()),
	}
	if ctx.ExprWithLambda() != nil {
		expr := v.visitExprWithLambda(ctx.ExprWithLambda())
		if expr != nil {
			fb.Body = &node.CodeBlock{
				Token: expr.GetToken(),
				Nodes: []node.Node{&node.ReturnStmt{
					Token: expr.GetToken(),
					Value: []node.Expr{expr},
				}},
			}
		}
	}
	return fb
}

func (v *AntlrVisitor) visitFuncReturnType(ctx parser.IFuncReturnTypeContext) []*node.Type {
	if ctx == nil {
		return nil
	}
	var ret []*node.Type
	ts := ctx.AllType_()
	l := len(ts)
	for _, t := range ts {
		_t := v.visitType(t)
		if _t.IsPlainType(node.TypeVoid) && l > 1 {
			v.appendErr(ctx, "void function only allow void as return type", nil)
			return nil
		}
		ret = append(ret, v.visitType(t))

	}
	return ret
}

func (v *AntlrVisitor) visitFuncSig(ctx parser.IFuncSigContext) *node.FuncBlock {
	fb := &node.FuncBlock{
		// Token: token.FromAntlrToken(ctx.GetStart()),
		Name:    v.visitIdentifier(ctx.Identifier()),
		Args:    v.visitFuncSignArgs(ctx.FuncSignArgs()),
		RetType: v.visitFuncReturnType(ctx.FuncReturnType()),
	}
	return fb
}

func (v *AntlrVisitor) visitFuncSignArgs(ctx parser.IFuncSignArgsContext) []*node.FuncArg {
	// []*node.FuncArg
	if ctx == nil {
		return nil
	}
	ts := ctx.AllFuncSignArgItem()
	var rst []*node.FuncArg
	for _, t := range ts {
		rst = append(rst, v.visitFuncSignArgItem(t))
	}
	// check has same name
	names := make(map[string]bool)
	for _, a := range rst {
		_, ok := names[a.Name.Value]
		if ok {
			v.appendErr(ctx, "duplicate arg name: "+a.Name.Value, nil)
		} else {
			names[a.Name.Value] = true
		}
	}
	return rst
}

func (v *AntlrVisitor) visitFuncSignArgItem(ctx parser.IFuncSignArgItemContext) *node.FuncArg {
	a := node.FuncArg{
		Name: v.visitIdentifier(ctx.Identifier()),
		Type: v.visitType(ctx.Type_()),
		Ref:  ctx.Ref() != nil,
	}
	return &a
}

func (v *AntlrVisitor) visitCodeBlock(ctx parser.ICodeBlockContext) *node.CodeBlock {
	if ctx == nil {
		return nil
	}
	var body []node.Node
	for _, m := range ctx.GetChildren() {
		switch m.(type) {
		case *parser.StmtContext:
			body = append(body, v.visitStmt(m.(*parser.StmtContext)))
		case *parser.ExprContext:
			body = append(body, v.visitExpr(m.(*parser.ExprContext)))
		}
	}
	return &node.CodeBlock{
		Nodes: body,
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
	}
}
