package eval

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func (e *Eval) EvalTryCatchStmt(t *node.TryCatchStmt) {
	e.currentToken = t.GetToken()
	defer func() {
		r := recover()
		if r == nil || t.Catch == nil {
			return
		}
		e.evalCatch(t.Catch, r)
	}()
	e.EvalCodeBlock(t.Try)
}

//func (e *Eval) EvalCatchs(t *node.TryCatchStmt, exc any) {
//	for _, c := range t.Catch {
//		if !c.CatchAll {
//			// TODO: Catch Pattern Match
//		}
//		e.evalCatch(c, exc)
//		return
//	}
//	// TODO: Miss Catch
//	panic(exc)
//}

func (e *Eval) evalCatch(c *node.CatchBranch, exc any) {
	e.objTable.PushEmpty(false)
	if c.VarName != "" {
		e.objTable.SetAtTop(c.VarName, exc)
	}
	e.EvalCodeBlock(c.Content)
	e.objTable.Pop()
}
