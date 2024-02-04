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

func (e *Eval) LoadContext(o *Eval) {
	if o == nil {
		e.objTable = NewObjectTable()
		return
	}
	e.objTable = o.objTable
	if e.objTable.Empty() {
		e.objTable = NewObjectTable()
	}
	e.builtin = o.builtin
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
		result.ReturnValue = result.ReturnObj.Val
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
		if r := recover(); r != nil {
			rst.IsPanic = true
			rst.PanicMsg = fmt.Sprint(r)
			rst.CurrentToken = e.CurrentToken()
		}
	}()

	rst = e.Do()
	rst.stderr = e.GetStdErr()
	rst.CurrentToken = e.CurrentToken()
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
