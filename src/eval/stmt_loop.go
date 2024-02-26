package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func (e *Eval) evalIterArray(n *node.IterStyleFor, iters []any) {
	e.currentToken = n.GetToken()
	e.loopFrame()
	defer e.frameEnd()

	for _, iter := range iters {
		isContinue, isBreak, isReturn := e.loopCentral(n.Body, func() {
			e.memory.Top().SetValue(n.Variable.Value, iter)
		})
		if isContinue {
			continue
		}
		if isReturn || isBreak {
			return
		}
	}
	return
}

func (e *Eval) loopCentral(body *node.CodeBlock, onNewFrame func()) (isContinue, isBreak, isReturn bool) {
	e.EvalLoopCodeBlockWithHook(body, onNewFrame)
	top := e.memory.Top()
	if top.Has(reserved.Continue) {
		top.Remove(reserved.Continue)
		isContinue = true
		return
	}
	if top.Has(reserved.Return) {
		isReturn = true
		return
	}
	if top.Has(reserved.Break) {
		top.Remove(reserved.Break)
		isBreak = true
		return
	}
	return
}

func (e *Eval) evalIterMap(n *node.IterStyleFor, iters map[any]any) {
	e.currentToken = n.GetToken()
	e.loopFrame()
	defer e.frameEnd()

	for key, val := range iters {
		isContinue, isBreak, isReturn := e.loopCentral(n.Body, func() {
			field := orderedmap.New[string, any]()
			field.Set("key", key)
			field.Set("val", val)
			e.memory.Top().SetValue(n.Variable.Value, &obj.StructField{
				Fields: field,
			})
		})
		if isContinue {
			continue
		}
		if isReturn || isBreak {
			return
		}
	}
	return
}

func (e *Eval) EvalIterStyleForStmt(styleFor *node.IterStyleFor) {
	e.currentToken = styleFor.GetToken()

	iters := e.EvalExpr(styleFor.Iterator).EnsureValue()
	switch itersT := iters.(type) {
	case string:
		// str to []rune to []string
		rs := []rune(itersT)
		as := make([]any, len(rs))
		for i, r := range rs {
			as[i] = string(rune(r))
		}
		e.evalIterArray(styleFor, as)
		return
	case []any:
		e.evalIterArray(styleFor, itersT)
		return
	case map[any]any:
		e.evalIterMap(styleFor, itersT)
		return
	}
	panic("Not Supported Iteration Type")
}

func (e *Eval) EvalWhileForStmt(n *node.WhileStyleFor) {
	e.currentToken = n.GetToken()

	e.loopFrame()
	defer e.frameEnd()

	for {
		if n.ConditionExpr != nil {
			rst := e.EvalExpr(n.ConditionExpr).EnsureValue().(bool)
			if !rst {
				return
			}
		}
		isContinue, isBreak, isReturn := e.loopCentral(n.Body, nil)
		if isContinue {
			continue
		}
		if isReturn || isBreak {
			return
		}
	}
}

func (e *Eval) EvalCStyleFrStmt(n *node.CStyleFor) {
	e.currentToken = n.GetToken()

	e.loopFrame()
	defer e.frameEnd()

	if n.InitialExpr != nil {
		switch initExpr := n.InitialExpr.(type) {
		case node.Stmt:
			e.EvalStmt(initExpr)
		case node.Expr:
			e.EvalExpr(initExpr)
		}
	}
	for {
		if n.ConditionExpr != nil {
			rst := e.EvalExpr(n.ConditionExpr).EnsureValue().(bool)
			if !rst {
				return
			}
		}
		isContinue, isBreak, isReturn := e.loopCentral(n.Body, nil)
		if isContinue {
			e.EvalExpr(n.AfterIterExpr)
		}
		if isContinue {
			continue
		}
		if isReturn || isBreak {
			return
		}
		e.EvalExpr(n.AfterIterExpr)
	}
}

func (e *Eval) EvalLoopCodeBlockWithHook(fc *node.CodeBlock, onNewFrame func()) {
	e.currentToken = fc.GetToken()
	e.frameStart(false)
	defer e.frameEndWithAll()

	if onNewFrame != nil {
		onNewFrame()
	}

	_ = e.runAst(fc.Nodes, reserved.Return, reserved.Break, reserved.Continue)
}
