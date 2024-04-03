package obj

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"github.com/xlab/treeprint"
)

var VizAny = false

func TreeAnyW(name string, o any, ignoreFx bool) treeprint.Tree {
	t := treeprint.New()
	return treeAny(&context{addrs: map[string]bool{}}, t, name, o, ignoreFx)
}

func TreeAnyT(name string, o any, ignoreFx bool) any {
	o = UnboxToEnd(o)
	switch t := o.(type) {
	case int, float64, string, bool:
		return t
	case IVisualize:
		return t.Visualize()
	default:
		return TreeAnyW(name, o, ignoreFx).String()
	}
}

type IVisualize interface {
	Visualize() string
}

func treeAny(ctx *context, t treeprint.Tree, name string, o any, ignoreFx bool) treeprint.Tree {
	if o == nil {
		t.SetValue(name + ": <nil>")
		return t
	}
	switch vT := o.(type) {
	case int, float64, string, bool:
		t.SetValue(fmt.Sprintf("%s: %v", name, vT))
		return t
	case []any:
		t.SetValue(name + ": []")
		for idx, k := range vT {
			treeAny(ctx, t.AddBranch(""), fmt.Sprint(idx), k, ignoreFx)
		}
		return t
	case map[any]any:
		t.SetValue(name + ": map")
		for k, v := range vT {
			treeAny(ctx, t.AddBranch(""), fmt.Sprint(k), v, ignoreFx)
		}
		return t
	case *node.FuncBlock:
		t.SetValue(name + ": " + vT.String())
		return t
	case *Object:
		return treeObj(ctx, t, name, vT, ignoreFx)
	case *StructField:
		f := vT.Fields
		for kvp := f.Oldest(); kvp != nil; kvp = kvp.Next() {
			if _, ok := kvp.Value.(*node.FuncBlock); ok && ignoreFx {
				continue
			}
			treeAny(ctx, t.AddBranch(kvp.Key), kvp.Key, kvp.Value, ignoreFx)
		}
		return t
	default:
		return nil
	}
}

func (c *context) addAddr(addr any) {
	if addr == nil {
		return
	}
	c.addrs[fmt.Sprintf("%p", addr)] = true
}

func (c *context) hasAddr(addr any) bool {
	if addr == nil {
		return false
	}
	_, ok := c.addrs[fmt.Sprintf("%p", addr)]
	return ok
}

func treeObj(ctx *context, t treeprint.Tree, name string, o *Object, ignoreFx bool) treeprint.Tree {
	if t == nil {
		t = treeprint.New()
	}
	if ctx.hasAddr(o) || ctx.hasAddr(o.Ref) {
		t.SetValue(fmt.Sprintf("%s: %s", name, "Possible Loop Object"))
		return t
	}
	ctx.addAddr(o)
	ctx.addAddr(o.Ref)

	if o == nil {
		t.SetValue("<nil>")
		return t
	}

	switch o.Kind {
	case EvalObj:
		t.SetValue(name + ": Eval")
	case Func:
		t.SetValue(name + ": Func")
	case Lambda:
		t.SetValue(name + ":Lambda")
	case Library:
		t.SetValue(name + ":Library")
	case StructDef:
		t.SetValue(name + ": StructDef")
	case Struct:
		treeAny(ctx, t, name, o.Value(), ignoreFx)
	case Value:
		treeAny(ctx, t, name, o.Value(), ignoreFx)
	}

	return t
}

type context struct {
	addrs map[string]bool
}
