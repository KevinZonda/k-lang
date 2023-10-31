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
