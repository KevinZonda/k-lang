package eval

import (
	"fmt"
	"reflect"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func clone(v any) any {
	if v == nil {
		return nil
	}
	switch vT := v.(type) {
	case *obj.StructField:
		fields := orderedmap.New[string, any]()
		for pair := vT.Fields.Oldest(); pair != nil; pair = pair.Next() {
			fields.Set(pair.Key, clone(pair.Value))
		}
		return &obj.StructField{
			TypeAs: vT.TypeAs,
			Fields: fields,
		}
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
	default:
		fmt.Println("[DEBUG] CLONE DEFAULT VAR ->", reflect.TypeOf(v))
		return v
	}
}

func (e *Eval) EvalAssignStmt(n *node.AssignStmt) {
	if len(n.Assignee) == 1 {
		e.EvalAssignStmtX(n, n.Assignee[0], n.Value)
		return
	}

}

func (e *Eval) EvalAssignStmtX(n *node.AssignStmt, assignee *node.Assignee, value any) {
	var v any = nil
	if assignee.Ref {
		v = e.evalExpr(n.Value, true)
	} else {
		v = clone(e.EvalExpr(n.Value))
	}
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
		switch fromObj.Val.(type) {
		case *obj.StructField:
			from = fromObj.Val.(*obj.StructField)
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
				o, ok := e.objTable.Get(lastVar.Name.Value)
				if ok && n.Token.Value != ":=" {
					o.Val = v
				} else {
					e.objTable.SetAtTop(lastVar.Name.Value, cons(v))
				}
			} else {
				o, ok := fromT.objTable.Bottom().Get(lastVar.Name.Value)
				if ok {
					o.Val = v
				} else {
					// f.objTable.SetAtTop(lastVar.Name.Value, cons(v))
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
	switch from.(type) {
	case *Eval:
		if canFromLocalVar {
			ok := false
			from, ok = from.(*Eval).objTable.Get(field)
			if !ok {
				panic(fmt.Sprintf("object %s not found", field))
			}
		} else {
			ok := false
			from, ok = from.(*Eval).objTable.Bottom().Get(field)
			if !ok {
				panic(fmt.Sprintf("object %s not found", field))
			}
		}
		return from
	// TODO Support struct
	case *obj.StructField:
		ok := false
		from, ok = from.(*obj.StructField).Fields.Get(field)
		if !ok {
			panic(fmt.Sprintf("field %s not found", field))
		}
		return from

	default:
		panic("not supported type: " + reflect.TypeOf(from).String())
	}
	return nil
}

func (e *Eval) unboxObj(from any) (any, bool) {
	if from == nil {
		return nil, false
	}
	switch from.(type) {
	case *obj.Object:
		return from.(*obj.Object).Val, true
	}
	return from, false
}

func (e *Eval) evalObjByIndex(from any, indexes []node.Expr) any {
	if len(indexes) == 0 {
		return from
	}

	for _, idxExpr := range indexes {
		from, _ = e.unboxObj(from)
		idx := e.EvalExpr(idxExpr)
		switch fromT := from.(type) {
		case []any:
			from = fromT[idx.(int)]
		case map[any]any:
			from = fromT[idx]
		default:
			panic(fmt.Sprint("not supported type to access by index: "+reflect.TypeOf(from).String(), "INDEX:", idx))
		}
	}
	return from
}

func (e *Eval) assignObjIndexValue(root any, index node.Expr, v any) {
	lastIndex := e.EvalExpr(index)
	switch rT := root.(type) {
	case []any:
		if rT != nil {
			rT[lastIndex.(int)] = v
		}

		// Consider we use pointer to modify the value, we are not okay to
		// set back with new value by using objTable.Set(..., v) method.
		// e.objTable.Set(baseV.Name.Value, obj)
		return
	case map[any]any:
		rT[lastIndex] = v
		// e.objTable.Set(baseV.Name.Value, obj)
		return
	default:
		panic(fmt.Sprint("not supported type to access by index: "+reflect.TypeOf(root).String(), "INDEX:", index))
	}
}
