package visualizer

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"github.com/xlab/treeprint"
	"reflect"
	"strings"
)

func VisualizeAny(name string, o any) *VNode {
	if o == nil {
		return &VNode{Val: nil}
	}

	switch vT := o.(type) {
	case int, float64, string, bool:
		return &VNode{Val: vT, Name: name}
	case []any:
		children := make([]*VNode, len(vT))
		for i, k := range vT {
			children[i] = VisualizeAny("", k)
		}
		return &VNode{Val: children, Name: name}
	case *obj.Object:
		return Visualize(name, vT)
	case *obj.StructField:
		f := vT.Fields
		children := map[string]*VNode{}
		for kvp := f.Oldest(); kvp != nil; kvp = kvp.Next() {
			// ignore visualize func
			if _, ok := kvp.Value.(*node.FuncBlock); ok {
				continue
			}
			vN := VisualizeAny(kvp.Key, kvp.Value)
			vN.Name = kvp.Key
			children[kvp.Key] = vN
		}
		return &VNode{Val: children, Name: name}
	default:
		return &VNode{Val: vT, Name: name}
	}
}

func TreeAnyW(name string, o any, ignoreFx bool) treeprint.Tree {
	return obj.TreeAnyW(name, o, ignoreFx)
}



func Visualize(name string, o *obj.Object) *VNode {
	if o == nil {
		return &VNode{Val: nil}
	}
	fmt.Println(reflect.TypeOf(o.Value()))

	switch o.Kind {
	case obj.EvalObj:
		return &VNode{Val: "Eval", Name: name}
	case obj.Func:
		return &VNode{Val: "Func", Name: o.ToFunc().Name.Value}
	case obj.Value:
		return VisualizeAny(name, o.Value())
	case obj.Lambda:
		return &VNode{Val: "Lambda"}
	case obj.Library:
		return &VNode{Val: "Library"}
	case obj.StructDef:
		return &VNode{Val: "StructDef", Name: o.ToStructDef().Name}
	case obj.Struct:
		return VisualizeAny(name, o.Value())
	}

	switch vT := o.Value().(type) {
	case int, float64, string, bool, []any:
		return VisualizeAny(name, vT)
	case *obj.Object:
		return Visualize(name, vT)
	case *obj.StructField:

	}
	return nil
}

type VNode struct {
	Name string
	Val  any
}

func (v *VNode) StringIdent(i int) string {
	sb := &strings.Builder{}
	_, _ = fmt.Fprint(sb, ident(i))
	_, _ = fmt.Fprint(sb, "name:", v.Name)
	_, _ = fmt.Fprint(sb, ident(i))
	if v.Val == nil {
		_, _ = fmt.Fprint(sb, "val : nil")
		return sb.String()
	}
	switch vT := v.Val.(type) {
	case int, float64, string, bool:
		_, _ = fmt.Fprint(sb, "val :", vT)
	case []*VNode:
		_, _ = fmt.Fprint(sb, "val :")
		for _, k := range vT {
			_, _ = fmt.Fprint(sb, k.StringIdent(i+1))
		}
	case map[string]*VNode:
		_, _ = fmt.Fprint(sb, "val :")
		for _, val := range vT {
			_, _ = fmt.Fprint(sb, val.StringIdent(i+1))
		}
	case *VNode:
		_, _ = fmt.Fprint(sb, "val :")
		_, _ = fmt.Fprint(sb, vT.StringIdent(i+1))
	}

	return sb.String()
}

func ident(i int) string {
	s := "\n"
	for i > 0 {
		s += "    "
		i--
	}
	return s
}
