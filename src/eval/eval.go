package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/memory"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"path/filepath"
)

type Eval struct {
	memory       *memory.Memory
	basePath     string
	loopLvl      int
	currentToken token.Token
	builtin      builtin.BuiltIn

	FeatStaticType bool
	FeatVerbose    bool
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

func New(basePath string) *Eval {
	if openedFiles == nil {
		openedFiles = map[string]*Eval{}
	}
	path := filepath.Dir(basePath)
	return &Eval{
		memory:   memory.NewMemory(),
		basePath: path,
		builtin:  builtin.NewBuiltIn(),
	}
}

func (e *Eval) runAst(ast tree.Ast, breaks ...string) DetailedRunResult {
	result := DetailedRunResult{}
	for idx, n := range ast {
		if e.memory.Top().HasAny(breaks...) {
			goto end
		}

		switch nT := n.(type) {
		case *node.CodeBlock:
			e.EvalCodeBlock(nT)
		case node.Expr:
			exprR := e.EvalExpr(nT)
			if idx != len(ast)-1 {
				continue
			}
			result.IsLastExpr = exprR.HasValue
			result.LastExprVal = exprR.Value
		case node.Stmt:
			e.EvalStmt(nT)
		case *node.FuncBlock:
			e.memory.Set(nT.Name.Value, nT)
		case *node.StructBlock:
			e.memory.Set(nT.Name, nT)
		case *node.OpenBlock:
			for _, stmt := range nT.Openers {
				e.EvalOpenStmt(stmt)
			}
		default:
			panic("not implemented")
		}
	}
end:
	result.ReturnObj, result.HasReturn = e.memory.Get(reserved.Return)
	if result.HasReturn && result.ReturnObj != nil {
		result.ReturnValue = result.ReturnObj.Value()
	}

	result.CurrentToken = e.CurrentToken()
	return result
}

func (e *Eval) Do(ast tree.Ast) DetailedRunResult {
	r := e.runAst(ast, reserved.Return)
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

func (e *Eval) DoSafely(ast tree.Ast) (rst DetailedRunResult) {
	defer func() {
		rst.CurrentToken = e.CurrentToken()
		if r := recover(); r != nil {
			rst.IsPanic = true
			rst.PanicMsg = fmt.Sprint(r)
			if e.FeatVerbose {
				rst.PanicMsg += "\n"
				rst.PanicMsg += e.MemStr()
			}
		}
	}()
	rst.stderr = e.GetStdErr()
	rst = e.Do(ast)
	return
}

func (e *Eval) EvalMain() any {
	fn, ok := e.memory.Get("main")
	if !ok || !fn.Is(obj.Func) {
		return nil
	}
	fx := fn.ToFunc()
	e.frameStart(true)

	result := e.runAst((tree.Ast)(fx.Body.Nodes), reserved.Return)

	e.frameEnd()

	return result.ReturnValue
}
