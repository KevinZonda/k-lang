package eval

import (
	"fmt"
	"reflect"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func clone(v any) any {
	if v == nil {
		return nil
	}
	switch vT := v.(type) {
	case *obj.StructField:
		sf := obj.NewStruct(vT.ParentEval)
		sf.TypeAs = vT.TypeAs
		for pair := vT.Fields.Oldest(); pair != nil; pair = pair.Next() {
			sf.With(pair.Key, clone(pair.Value))
		}
		return sf
	case []any:
		a := make([]any, len(vT))
		copy(a, v.([]any))
		return a
	case map[any]any:
		m := make(map[any]any)
		for k, val := range vT {
			m[k] = val
		}
		return m
	case int, float64, string, bool:
		return v
	case *node.LambdaExpr, *node.FuncBlock:
		return v
	case *Eval:
		return v
	default:
		fmt.Println("[DEBUG] CLONE DEFAULT VAR ->", reflect.TypeOf(v))
		return v
	}
}

func (e *Eval) evalExprOrRef(n node.Expr, ref bool) any {
	var v any
	if e.FeatRefAll {
		return e.evalExpr(n, true).EnsureValue()
	}
	if ref {
		v = e.evalExpr(n, true).EnsureValue()
		if vT, ok := v.(*obj.Object); ok {
			v = vT.CreateRef()
		}
		return v
	}

	return clone(e.EvalExpr(n).EnsureValue())
}

func (e *Eval) EvalAssignStmt(n *node.AssignStmt) {
	if len(n.Assignee) == 1 {
		v := e.evalExprOrRef(n.Value, n.Assignee[0].Ref)
		e.EvalAssignStmtX(n, n.Assignee[0], v)
		return
	}
	vals, ok := e.EvalExpr(n.Value).EnsureValue().([]any)
	if !ok {
		panic("eval expr not supported type: " + reflect.TypeOf(n.Value).String())
	}
	if len(vals) != len(n.Assignee) {
		panic("assign cannot happened if have different numbers of receivers and senders have: " + fmt.Sprint(len(n.Assignee)) + " receivers, but have: " + fmt.Sprint(len(vals)) + " senders")
	}

	if !e.FeatRefAll {
		vals = clone(vals).([]any)
	}
	for i, assignee := range n.Assignee {
		e.EvalAssignStmtX(n, assignee, vals[i])
	}

}

func (e *Eval) EvalAssignStmtX(n *node.AssignStmt, assignee *node.Assignee, value any) {
	var v any = value
	e.currentToken = n.GetToken()

	nVar := assignee.Var

	var from any = e

	for i, bvar := range nVar.Value[:len(nVar.Value)-1] {
		from, _ = e.unboxObj(from)
		from = e.evalObjByField(from, i == 0, bvar.Name.Value)
		from, _ = e.unboxObj(from)
		from = e.evalObjByIndex(from, bvar.Index)
	}
	// fmt.Println("FROM", reflect.TypeOf(from))
	switch from.(type) {
	case *obj.Object:
		fromObj := from.(*obj.Object)
		switch fromObj.Value().(type) {
		case *obj.StructField:
			from = fromObj.ToStruct()
		}
		//fmt.Println("SET VAL", reflect.TypeOf(from.(*obj.Object).Val), from.(*obj.Object).Val, "->", v)
		//from.(*obj.Object).Val = v
	}

	lastVar := nVar.Value[len(nVar.Value)-1]
	if len(lastVar.Index) == 0 {
		switch fromT := from.(type) {
		case *obj.StructField:
			ok := false
			_, ok = fromT.Fields.Get(lastVar.Name.Value)
			if !ok {
				panic(fmt.Sprintf("field %s not found", lastVar.Name.Value))
			}
			fromT.Fields.Set(lastVar.Name.Value, v)

		case *Eval:
			if e == fromT {
				o, ok := e.memory.Get(lastVar.Name.Value)
				if o.Is(obj.EvalObj, obj.Func, obj.StructDef) {
					panic("cannot assign to " + lastVar.Name.Value)
				}

				shadow := n.Token.Value == ":=" || assignee.Type != nil

				if ok && !shadow {
					// following is not possible because ref syntax
					// foo(&x) will let set val not possible
					// e.memory.Set(lastVar.Name.Value, Construct(v))
					if e.FeatStaticType {
						v = e.TypeCast(o.Type, v)
						o.SetValue(v)
					} else {
						o.WithType(e.AutoType(v)).SetValue(v)
					}

					//if o.Ref != nil {
					//	o = o.NoRef()
					//}
				} else {
					typeV := assignee.Type
					if typeV != nil {
						v = e.TypeCast(typeV, v)
					} else {
						typeV = e.AutoType(v)
					}
					e.memory.Top().SetValue(lastVar.Name.Value, obj.Construct(v).WithType(typeV))
				}
			} else {
				o, ok := fromT.memory.Bottom().Get(lastVar.Name.Value)
				if o.Is(obj.EvalObj, obj.Func, obj.StructDef) {
					panic("cannot assign to " + lastVar.Name.Value)
				}

				if ok {
					if e.FeatStaticType {
						e.TypeCheckOrPanic(o.Type, v)
					}
					o.SetValue(v)
				} else {
					panic("not support create new variable in other file")
				}
			}

		}
	} else {
		// TODO: TEST
		// TODO: INDEX!
		from = e.evalObjByField(from, len(nVar.Value) == 0, lastVar.Name.Value)
		from, _ = e.unboxObj(from)
		from = e.evalObjByIndex(from, lastVar.Index[0:len(lastVar.Index)-1])
		from, _ = e.unboxObj(from)
		e.assignObjIndexValue(from, lastVar.Index[len(lastVar.Index)-1], v)
	}

	// fmt.Println("SET VAL", reflect.TypeOf(from), from, "->", v)

	return
}

func (e *Eval) evalObjByField(from any, canFromLocalVar bool, field string) any {
	ok := false
	switch fromT := from.(type) {
	case *Eval:
		if canFromLocalVar {
			if from, ok = fromT.memory.Get(field); ok {
				return from
			}
			panic(fmt.Sprintf("object %s not found", field))
		}
		from, ok = fromT.memory.Bottom().Get(field)
		if !ok {
			panic(fmt.Sprintf("object %s not found", field))
		}
		return from
	// TODO Support struct
	case *obj.StructField:
		from, ok = fromT.Fields.Get(field)
		if !ok {
			panic(fmt.Sprintf("field %s not found", field))
		}
		return from

	default:
		panic("assign to field not supported type: " + reflect.TypeOf(from).String())
	}
	return nil
}

func (e *Eval) unboxObj(from any) (any, bool) {
	if from == nil {
		return nil, false
	}
	switch fromT := from.(type) {
	case *obj.Object:
		return fromT.Value(), true
	}
	return from, false
}

func (e *Eval) evalObjByIndex(from any, indexes []node.Expr) any {
	if len(indexes) == 0 {
		return from
	}

	for _, idxExpr := range indexes {
		from, _ = e.unboxObj(from)
		idx := e.EvalExpr(idxExpr).EnsureValue()
		switch fromT := from.(type) {
		case []any:
			from = fromT[asType[int](idx)]
		case map[any]any:
			from = fromT[idx]
		default:
			panic(fmt.Sprint("not supported type to access by index: "+reflect.TypeOf(from).String(), "INDEX:", idx))
		}
	}
	return from
}

func (e *Eval) assignObjIndexValue(root any, index node.Expr, v any) {
	lastIndex := e.EvalExpr(index).EnsureValue()
	switch rT := root.(type) {
	case []any:
		if rT != nil {
			rT[asType[int](lastIndex)] = v
		}

		// Consider we use pointer to modify the value, we are not okay to
		// set back with new value by using memory.Set(..., v) method.
		// e.memory.Set(baseV.Name.Value, obj)
		return
	case map[any]any:
		rT[lastIndex] = v
		// e.memory.Set(baseV.Name.Value, obj)
		return
	default:
		panic(fmt.Sprint("not supported type to access by index: "+reflect.TypeOf(root).String(), "INDEX:", index))
	}
}
