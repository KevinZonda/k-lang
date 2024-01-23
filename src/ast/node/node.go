package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type Node interface {
	TokenValue() string
	GetToken() token.Token
}

type Expr interface {
	Node
	expr() // flag to show is expr
}

type Stmt interface {
	Node
	stmt() // flag to show is stmt
}

type Block interface {
	Node
	block() // flag to show is block
}
