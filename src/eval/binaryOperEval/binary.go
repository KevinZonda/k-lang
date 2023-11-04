package binaryOperEval

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

func BinaryOper(oper token.Kind, left any, right any) any {

	switch oper {
	case token.Add:
		return Add(left, right)
	case token.Sub:
		return Sub(left, right)
	case token.Mul:
		return Mul(left, right)
	case token.Pow:
		return Pow(left, right)
	case token.Div:
		return Div(left, right)
	case token.Mod:
		return Mod(left, right)
	case token.And:
		return And(left, right)
	case token.Equals:
		return Eq(left, right)
	case token.NotEq:
		return !Eq(left, right)
	case token.Or:
		return Or(left, right)
	case token.Less:
		return Less(left, right)
	case token.LessEq:
		return LessEq(left, right)
	case token.Greater:
		return Greater(left, right)
	case token.GreaterEq:
		return GreaterEq(left, right)
	}
	panic("not implemented")
}
