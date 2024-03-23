package httpExternalServer

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/common"
	"github.com/gin-gonic/gin"
	"reflect"
)

type TreeNode struct {
	Type NodeType
	Node map[string]*TreeNode
	argT common.FuncArgs
	Func any
}

type NodeType string

const (
	FuncNodeType    NodeType = "func"
	VarNodeType     NodeType = "var"
	PackageNodeType NodeType = "node"
)

func (t *TreeNode) Engine(packName string) *gin.Engine {
	g := gin.New()
	t.Register(g.Group("/" + packName))
	return g
}

func (t *TreeNode) EngineWithNoPackName() *gin.Engine {
	g := gin.New()
	t.Register(g)
	return g
}

func (t *TreeNode) Register(g gin.IRouter) {
	for k, v := range t.Node {
		if v.Type == FuncNodeType {
			if v.argT == nil {
				v.argT = readFuncArg(v.Func)
			}
			hdlr := createFuncHandler(reflect.ValueOf(v.Func), v.argT)
			if hdlr == nil {
				continue
			}
			g.POST("/"+k, hdlr)
			g.GET("/"+k, func(ctx *gin.Context) {
				ctx.JSON(200, common.PackageInfoElement{
					Type: common.Function,
					Args: v.argT,
				})
			})
			continue
		}
		if v.Type == PackageNodeType {
			v.Register(g.Group("/" + k))
			continue
		}
	}
}

func NewPackage() *TreeNode {
	return &TreeNode{
		Type: PackageNodeType,
		Node: map[string]*TreeNode{},
	}
}

func (t *TreeNode) AppendFx(fx any) *TreeNode {
	return t.AppendFxWithName(FuncName(fx), fx)
}

func (t *TreeNode) AppendFxWithName(name string, fx any) *TreeNode {
	t.Node[name] = &TreeNode{
		Type: FuncNodeType,
		Func: fx,
		argT: readFuncArg(fx),
	}
	return t
}

func (t *TreeNode) AppendPackage(name string, n *TreeNode) *TreeNode {
	t.Node[name] = n
	return t
}
