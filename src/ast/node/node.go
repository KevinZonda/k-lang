package node

type Node interface {
	TokenValue() string
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

type baseStmt struct{}

func (b baseStmt) stmt() {}

type baseExpr struct{}

func (e baseExpr) expr() {}
