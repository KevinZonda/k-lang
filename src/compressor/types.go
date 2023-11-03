package compressor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"github.com/brimdata/zed/zson"
)

func NewUnmarshal() *zson.UnmarshalContext {
	um := zson.NewUnmarshaler()
	um.Bind(
		node.AssignStmt{},

		node.BinaryOperExpr{},

		node.DeclareStmt{},

		node.CStyleFor{},
		node.IterStyleFor{},
		node.WhileStyleFor{},

		node.CodeBlock{},
		node.FuncArg{},
		node.FuncBlock{},

		node.FuncCall{},

		node.Identifier{},

		node.IfStmt{},

		node.ArrayLiteral{},
		node.BoolLiteral{},
		node.FloatLiteral{},
		node.IntLiteral{},
		node.MapLiteral{},
		node.MapPairLiteral{},
		node.StringLiteral{},

		node.ReturnStmt{},
		
		node.Struct{},
		node.UnaryOperExpr{},
	)
	return um
}
