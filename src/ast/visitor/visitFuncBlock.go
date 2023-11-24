package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitFuncBlock(ctx parser.IFuncBlockContext) *node.FuncBlock {
	fx := v.visitFuncSig(ctx.FuncSig().(*parser.FuncSigContext))
	fx.Token = token.FromAntlrToken(ctx.GetStart())
	fx.Body = v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext))
	return fx
}

func (v *AntlrVisitor) visitLambda(ctx parser.ILambdaContext) *node.LambdaExpr {
	fb := &node.LambdaExpr{
		Token:   token.FromAntlrToken(ctx.GetStart()),
		Body:    v.visitCodeBlock(ctx.CodeBlock()),
		Args:    v.visitFuncSignArgs(ctx.FuncSignArgs()),
		RetType: v.visitType(ctx.Type_()),
	}
	return fb
}

func (v *AntlrVisitor) visitFuncSig(ctx parser.IFuncSigContext) *node.FuncBlock {
	fb := &node.FuncBlock{
		// Token: token.FromAntlrToken(ctx.GetStart()),
		Name:    v.visitIdentifier(ctx.Identifier()),
		Args:    v.visitFuncSignArgs(ctx.FuncSignArgs()),
		RetType: v.visitType(ctx.Type_()),
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
	return rst
}

func (v *AntlrVisitor) visitFuncSignArgItem(ctx parser.IFuncSignArgItemContext) *node.FuncArg {
	a := node.FuncArg{
		Name: v.visitIdentifier(ctx.Identifier()),
		Type: v.visitType(ctx.Type_()),
	}
	return &a
}

func (v *AntlrVisitor) visitCodeBlock(ctx parser.ICodeBlockContext) *node.CodeBlock {
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
		Token: token.FromAntlrToken(ctx.GetStart()),
	}
}
