package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type TryCatchStmt struct {
	Token token.Token
	Try   *CodeBlock
	Catch []*CatchBranch
}

func (t *TryCatchStmt) stmt() {}

type CatchBranch struct {
	VarType  *Type
	VarName  string
	Content  *CodeBlock
	CatchAll bool
}

func (t *TryCatchStmt) TokenValue() string {
	return t.Token.Value
}
