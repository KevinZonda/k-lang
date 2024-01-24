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

func (e *Eval) evalAssignStmt(n *node.AssignStmt, v any) any {
	e.currentToken = n.GetToken()

	var from any = e

	for i, bvar := range n.Var.Value[:len(n.Var.Value)-1] {
		switch from.(type) {
		case *Eval:
			if i == 0 {
				ok := false
				from, ok = from.(*Eval).objTable.Get(bvar.Name.Value)
				if !ok {
					panic(fmt.Sprintf("object %s not found", bvar.Name.Value))
				}
			} else {
				ok := false
				from, ok = from.(*Eval).objTable.Bottom().Get(bvar.Name.Value)
				if !ok {
					panic(fmt.Sprintf("object %s not found", bvar.Name.Value))
				}
			}
		default:
			panic("not supported type: " + reflect.TypeOf(from).String())
		}
		// TODO: INDEX!
	}
	fmt.Println("FROM", reflect.TypeOf(from))
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
	switch from.(type) {
	case *obj.StructField:
		ok := false
		_, ok = from.(*obj.StructField).Fields.Get(lastVar.Name.Value)
		if !ok {
			panic(fmt.Sprintf("field %s not found", lastVar.Name.Value))
		}
		from.(*obj.StructField).Fields.Set(lastVar.Name.Value, v)
		// TODO: More case !
		// TODO: INDEX!
	}

	fmt.Println("SET VAL", reflect.TypeOf(from), from, "->", v)

	return nil
}

func (e *Eval) EvalAssignStmt(n *node.AssignStmt) any {
	v := e.EvalExpr(n.Value)
	v = clone(v)

	e.currentToken = n.GetToken()
	// TODO: arr
	if len(n.Var.Value) > 1 {
		return e.evalAssignStmt(n, v)
	}
	baseV := n.Var.Value[0]
	obj, ok := e.objTable.Get(baseV.Name.Value)
	if ok && len(baseV.Index) != 0 {
		switch obj.Val.(type) {
		case []any:
			bIdx := baseV.Index[:len(baseV.Index)-1]
			if tgt := e.evalValWithIdx(bIdx, &obj.Val); tgt != nil {
				switch (*tgt).(type) {
				case []any:
					((*tgt).([]any))[e.EvalExpr(baseV.Index[len(baseV.Index)-1]).(int)] = v
				}
			}
			// Consider we use pointer to modify the value, we are not okay to
			// set back with new value by using objTable.Set(..., v) method.
			e.objTable.Set(baseV.Name.Value, obj)
			return v
		case map[any]any:
			m := obj.Val.(map[any]any)
			m[e.EvalExpr(baseV.Index[len(baseV.Index)-1])] = v
			e.objTable.Set(baseV.Name.Value, obj)
			return v
		}
	}
	e.objTable.Set(baseV.Name.Value, v)
	return v
}
