package compressor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"github.com/brimdata/zed/zson"
)

func NewUnmarshal() *zson.UnmarshalContext {
	um := zson.NewUnmarshaler()
	um.Bind(
		node.DotExpr{},
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

		node.LambdaExpr{},

		node.ArrayLiteral{},
		node.BoolLiteral{},
		node.FloatLiteral{},
		node.IntLiteral{},
		node.MapLiteral{},
		node.MapPairLiteral{},
		node.StringLiteral{},

		node.MatchStmt{},
		node.MatchCase{},

		node.OpenStmt{},

		node.ReturnStmt{},

		node.StructBlock{},
		node.UnaryOperExpr{},

		node.BaseVariable{},
		node.Variable{},
	)
	return um
}
