package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"math"
	rand "math/rand"
	"time"
)

type StdMathLib struct {
	m    map[string]*obj.Object
	seed *int64
}

func NewStdMathLib() *StdMathLib {
	o := map[string]*obj.Object{
		"pi": obj.NewImmutableObj(obj.Value, math.Pi),
		"e":  obj.NewImmutableObj(obj.Value, math.E),
	}
	return &StdMathLib{o, nil}
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
	case "randInt":
		if c.seed == nil {
			return resultVal(rand.Int())
		}
		r := rand.New(rand.NewSource(*c.seed))
		return resultVal(r.Int())
	case "randIntN":
		ensureArgsLen(args, 1)
		if c.seed == nil {
			return resultVal(rand.Intn(i64(args[0])))
		}
		r := rand.New(rand.NewSource(*c.seed))
		return resultVal(r.Intn(i64(args[0])))
	case "randNum":
		if c.seed == nil {
			return resultVal(rand.Float64())
		}
		r := rand.New(rand.NewSource(*c.seed))
		return resultVal(r.Float64())
	case "randAlphabet":
		if c.seed == nil {
			return resultVal(string(rune(rand.Intn(26) + 65)))
		}
	case "shuffle":
		ensureArgsLen(args, 1)
		argT := args[0].([]any)
		if c.seed == nil {
			rand.Shuffle(len(argT), func(i, j int) {
				argT[i], argT[j] = argT[j], argT[i]
			})
			return resultNoVal()
		}
		r := rand.New(rand.NewSource(*c.seed))
		r.Shuffle(len(argT), func(i, j int) {
			argT[i], argT[j] = argT[j], argT[i]
		})
		return resultNoVal()

	case "setRandSeed", "randSeed":
		ensureArgsLen(args, 1)
		seed := int64(i64(args[0]))
		c.seed = &seed
		return resultNoVal()
	case "getRandSeed":
		if c.seed == nil {
			return resultVal(nil)
		}
		return resultVal(*c.seed)
	case "unsetRandSeed":
		c.seed = nil
		return resultNoVal()
	case "randSeedTime":
		x := time.Now().UnixNano()
		c.seed = &x
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
