package visualizer

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type edgeContext struct {
	addrs map[string]string
	edge  map[string]map[string]bool
}

func (c *edgeContext) addEdge(from, to string) {
	if c.edge[from] == nil {
		c.edge[from] = map[string]bool{}
	}
	c.edge[from][to] = true
}

func (c *edgeContext) registerAddr(addr, name string) {
	c.addrs[addr] = name
}

func (c *edgeContext) Mermaid() string {
	var s string = "graph TD\n"
	for from, tos := range c.edge {
		for to := range tos {
			s += from + " --> " + to + "\n"
		}
	}
	return s
}

func EdgeObj(name string, o *obj.Object) string {
	c := &edgeContext{addrs: map[string]string{}, edge: map[string]map[string]bool{}}
	return edgeObj(c, name, o)
}

func edgeStruct(ctx *edgeContext, name string, oT *obj.StructField) {
	for pair := oT.Fields.Oldest(); pair != nil; pair = pair.Next() {
		switch pairT := pair.Value.(type) {
		case *obj.Object:
			edgeObj(ctx, pair.Key, pairT)
		case *obj.StructField:
			ctx.addEdge(name, addr(pairT))
			ctx.registerAddr(addr(pairT), pair.Key)
			edgeStruct(ctx, pair.Key, pairT)
		}
	}
}

func edgeObj(ctx *edgeContext, name string, o *obj.Object) string {
	if o == nil {
		return ""
	}
	ctx.registerAddr(addr(o), name)
	if o.Ref != nil {
		ctx.addEdge(addr(o), addr(o.Ref))
		ctx.registerAddr(addr(o.Ref), name+"Ref")
		edgeObj(ctx, name, o.Ref)
	}
	if o.Kind == obj.Struct {
		oT := o.ToStruct()
		edgeStruct(ctx, name, oT)
	}
	return ctx.Mermaid()
}

func addr(v any) string {
	return fmt.Sprintf("%p", v)
}
