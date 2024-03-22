package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"math"
)

type StdMathLib struct {
	m map[string]*obj.Object
}

func constObj(v any) *obj.Object {
	o := obj.NewObj(obj.Value, v)
	o.Immutable = true
	return o
}

func NewStdMathLib() *StdMathLib {
	o := map[string]*obj.Object{
		"pi": constObj(math.Pi),
		"e":  constObj(math.E),
	}
	return &StdMathLib{o}
}

func (c *StdMathLib) GetObjList() map[string]*obj.Object {
	return c.m
}

func (c *StdMathLib) FuncCall(b obj.StdIO, caller string, args []any) obj.ILibraryCall {
	switch caller {
	case "cos":
		ensureArgsLen(args, 1)
		return resultVal(math.Cos(f64(args[0])))
	case "sin":
		ensureArgsLen(args, 1)
		return resultVal(math.Sin(f64(args[0])))
	case "tan":
		ensureArgsLen(args, 1)
		return resultVal(math.Tan(f64(args[0])))
	case "acos":
		ensureArgsLen(args, 1)
		return resultVal(math.Acos(f64(args[0])))
	case "asin":
		ensureArgsLen(args, 1)
		return resultVal(math.Asin(f64(args[0])))
	case "atan":
		ensureArgsLen(args, 1)
		return resultVal(math.Atan(f64(args[0])))
	case "atan2":
		ensureArgsLen(args, 2)
		return resultVal(math.Atan2(f64(args[0]), f64(args[1])))
	case "cosh":
		ensureArgsLen(args, 1)
		return resultVal(math.Cosh(f64(args[0])))
	case "sinh":
		ensureArgsLen(args, 1)
		return resultVal(math.Sinh(f64(args[0])))
	case "tanh":
		ensureArgsLen(args, 1)
		return resultVal(math.Tanh(f64(args[0])))
	case "acosh":
		ensureArgsLen(args, 1)
		return resultVal(math.Acosh(f64(args[0])))
	case "asinh":
		ensureArgsLen(args, 1)
		return resultVal(math.Asinh(f64(args[0])))
	case "atanh":
		ensureArgsLen(args, 1)
		return resultVal(math.Atanh(f64(args[0])))
	case "exp":
		ensureArgsLen(args, 1)
		return resultVal(math.Exp(f64(args[0])))
	case "exp2":
		ensureArgsLen(args, 1)
		return resultVal(math.Exp2(f64(args[0])))
	case "expm1":
		ensureArgsLen(args, 1)
		return resultVal(math.Expm1(f64(args[0])))
	case "log":
		ensureArgsLen(args, 1)
		return resultVal(math.Log(f64(args[0])))
	case "log2":
		ensureArgsLen(args, 1)
		return resultVal(math.Log2(f64(args[0])))
	case "log10":
		ensureArgsLen(args, 1)
		return resultVal(math.Log10(f64(args[0])))
	case "log1p":
		ensureArgsLen(args, 1)
		return resultVal(math.Log1p(f64(args[0])))
	case "pow":
		ensureArgsLen(args, 2)
		return resultVal(math.Pow(f64(args[0]), f64(args[1])))
	case "sqrt":
		ensureArgsLen(args, 1)
		return resultVal(math.Sqrt(f64(args[0])))
	case "cbrt":
		ensureArgsLen(args, 1)
		return resultVal(math.Cbrt(f64(args[0])))
	case "hypot":
		ensureArgsLen(args, 2)
		return resultVal(math.Hypot(f64(args[0]), f64(args[1])))
	case "erf":
		ensureArgsLen(args, 1)
		return resultVal(math.Erf(f64(args[0])))
	case "erfc":
		ensureArgsLen(args, 1)
		return resultVal(math.Erfc(f64(args[0])))
	case "gamma":
		ensureArgsLen(args, 1)
		return resultVal(math.Gamma(f64(args[0])))
	}
	panic("Unknown function: " + caller)
}

func f64(arg any) float64 {
	switch argT := arg.(type) {
	case int:
		return float64(argT)
	case float64:
		return argT
	default:
		panic("Invalid argument type, expected int or float64")
	}
}

func i64(arg any) int {
	switch argT := arg.(type) {
	case int:
		return int(argT)
	case float64:
		return int(argT)
	default:
		panic("Invalid argument type, expected int or int64")
	}
}

var _ obj.ILibrary = (*StdMathLib)(nil)
