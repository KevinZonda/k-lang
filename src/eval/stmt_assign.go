package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"reflect"
)

func clone(v any) any {
	if v == nil {
		return nil
	}
	switch v.(type) {
	case *obj.StructField:
		sf := v.(*obj.StructField)
		fields := orderedmap.New[string, any]()
		for pair := sf.Fields.Oldest(); pair != nil; pair = pair.Next() {
			fields.Set(pair.Key, clone(pair.Value))
		}
		return &obj.StructField{
			TypeAs: sf.TypeAs,
			Fields: fields,
		}
	case []any:
		a := make([]any, len(v.([]any)))
		copy(a, v.([]any))
		return a
	case map[any]any:
		m := make(map[any]any)
		for k, v := range v.(map[any]any) {
			m[k] = v
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
	v := e.EvalExpr(n.Value)
	v = clone(v)

	e.currentToken = n.GetToken()

	var from any = e

	for i, bvar := range n.Var.Value[:len(n.Var.Value)-1] {
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

	lastVar := n.Var.Value[len(n.Var.Value)-1]
	if len(lastVar.Index) == 0 {
		switch from.(type) {
		case *obj.StructField:
			ok := false
			_, ok = from.(*obj.StructField).Fields.Get(lastVar.Name.Value)
			if !ok {
				panic(fmt.Sprintf("field %s not found", lastVar.Name.Value))
			}
			from.(*obj.StructField).Fields.Set(lastVar.Name.Value, v)

		case *Eval:
			e.objTable.Set(lastVar.Name.Value, v)
		}
	} else {
		// TODO: TEST
		// TODO: INDEX!
		from = e.evalObjByField(from, len(n.Var.Value) == 0, lastVar.Name.Value)
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
		switch from.(type) {
		case []any:
			from = from.([]any)[idx.(int)]
		case map[any]any:
			from = from.(map[any]any)[idx]
		default:
			panic(fmt.Sprint("not supported type to access by index: "+reflect.TypeOf(from).String(), "INDEX:", idx))
		}
	}
	return from
}

func (e *Eval) assignObjIndexValue(root any, index node.Expr, v any) {
	lastIndex := e.EvalExpr(index)
	switch root.(type) {
	case []any:
		if root != nil {
			switch (root).(type) {
			case []any:
				(root.([]any))[lastIndex.(int)] = v
			}
		}

		// Consider we use pointer to modify the value, we are not okay to
		// set back with new value by using objTable.Set(..., v) method.
		// e.objTable.Set(baseV.Name.Value, obj)
		return
	case map[any]any:
		m := root.(map[any]any)
		m[lastIndex] = v
		// e.objTable.Set(baseV.Name.Value, obj)
		return
	}
}
