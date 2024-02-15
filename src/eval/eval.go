package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"path/filepath"
)

type Eval struct {
	ast          tree.Ast
	objTable     *TableStack
	basePath     string
	loopLvl      int
	currentToken token.Token
	builtin      builtin.BuiltIn

	FeatStaticType bool
}

func (e *Eval) SetPath(path string) {
	e.basePath = path
}

func (e *Eval) CurrentToken() token.Token {
	return e.currentToken
}

var openedFiles map[string]*Eval

func ResetGlobal() {
	openedFiles = make(map[string]*Eval)
}

func New(ast tree.Ast, inputFile string) *Eval {
	if openedFiles == nil {
		openedFiles = map[string]*Eval{}
	}
	path := filepath.Dir(inputFile)
	return &Eval{
		ast:      ast,
		objTable: NewObjectTable(),
		basePath: path,
		builtin:  builtin.NewBuiltIn(),
	}
}

func (e *Eval) SetAST(ast tree.Ast) {
	e.ast = ast
	e.currentToken = token.Token{}
}

func (e *Eval) runAst(ast tree.Ast, breaks ...string) DetailedRunResult {
	result := DetailedRunResult{}
	for idx, n := range ast {
		for _, key := range breaks {
			if e.objTable.HasKeyAtTop(key) {
				goto end
			}
		}
		switch nT := n.(type) {
		case *node.CodeBlock:
			e.EvalCodeBlock(nT)
		case node.Expr:
			exprR := e.EvalExpr(nT)
			if idx == len(e.ast)-1 {
				switch n.(type) {
				case *node.AssignStmt:
					result.IsLastExpr = false
				default:
					result.IsLastExpr = true
					result.LastExprVal = exprR
				}
			}
		case node.Stmt:
			e.EvalStmt(nT)
		case *node.FuncBlock:
			e.objTable.Set(nT.Name.Value, nT)
		case *node.StructBlock:
			e.objTable.Set(nT.Name, nT)
		case *node.OpenBlock:
			for _, stmt := range nT.Openers {
				e.EvalStmt(stmt)
			}
		default:
			panic("not implemented")
		}
	}
end:
	result.ReturnObj, result.HasReturn = e.objTable.Get(reserved.Return)
	if result.HasReturn && result.ReturnObj != nil {
		result.ReturnValue = result.ReturnObj.Value()
	}

	result.CurrentToken = e.CurrentToken()
	return result
}

func (e *Eval) runWithBreak(breaks ...string) DetailedRunResult {
	return e.runAst(e.ast, breaks...)
}

func (e *Eval) Do() DetailedRunResult {
	r := e.runWithBreak(reserved.Return)
	if r.HasReturn {
		return r
	}

	retV := e.EvalMain()
	r.HasReturn = retV != nil
	if r.HasReturn {
		r.ReturnValue = retV
	}
	return r
}

func (e *Eval) DoSafely() (rst DetailedRunResult) {
	defer func() {
		rst.CurrentToken = e.CurrentToken()
		if r := recover(); r != nil {
			rst.IsPanic = true
			rst.PanicMsg = fmt.Sprint(r)
		}
	}()
	rst.stderr = e.GetStdErr()
	rst = e.Do()
	return
}

func (e *Eval) EvalMain() any {
	fn, ok := e.objTable.Get("main")
	if !ok || !fn.Is(obj.Func) {
		return nil
	}
	fx := fn.ToFunc()
	e.frameStart(true)

	result := e.runAst((tree.Ast)(fx.Body.Nodes), reserved.Return)

	e.frameEnd()

	return result.ReturnValue
}
