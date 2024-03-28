package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"math"
	rand "math/rand"
	"time"
)

type StdMathLib struct {
	FBLibrary
	seed *int64
}

func NewStdMathLib() *StdMathLib {
	c := &StdMathLib{}
	c.FBLibrary = FBLibrary{
		P: map[string]*obj.Object{
			"pi": obj.NewImmutableObj(obj.Value, math.Pi),
			"e":  obj.NewImmutableObj(obj.Value, math.E),
		},
		F: map[string]*node.FuncBlock{
			"cos":   FxToFuncBlock(math.Cos),
			"sin":   FxToFuncBlock(math.Sin),
			"tan":   FxToFuncBlock(math.Tan),
			"acos":  FxToFuncBlock(math.Acos),
			"asin":  FxToFuncBlock(math.Asin),
			"atan":  FxToFuncBlock(math.Atan),
			"atan2": FxToFuncBlock(math.Atan2),
			"cosh":  FxToFuncBlock(math.Cosh),
			"sinh":  FxToFuncBlock(math.Sinh),
			"tanh":  FxToFuncBlock(math.Tanh),
			"acosh": FxToFuncBlock(math.Acosh),
			"asinh": FxToFuncBlock(math.Asinh),
			"atanh": FxToFuncBlock(math.Atanh),
			"exp":   FxToFuncBlock(math.Exp),
			"exp2":  FxToFuncBlock(math.Exp2),
			"expm1": FxToFuncBlock(math.Expm1),
			"log":   FxToFuncBlock(math.Log),
			"log2":  FxToFuncBlock(math.Log2),
			"log10": FxToFuncBlock(math.Log10),
			"log1p": FxToFuncBlock(math.Log1p),
			"pow":   FxToFuncBlock(math.Pow),
			"sqrt":  FxToFuncBlock(math.Sqrt),
			"cbrt":  FxToFuncBlock(math.Cbrt),
			"hypot": FxToFuncBlock(math.Hypot),
			"erf":   FxToFuncBlock(math.Erf),
			"erfc":  FxToFuncBlock(math.Erfc),
			"gamma": FxToFuncBlock(math.Gamma),
			"randInt": FxToFuncBlock(func() int {
				if c.seed == nil {
					return rand.Int()
				}
				r := rand.New(rand.NewSource(*c.seed))
				return r.Int()
			}),
			"randIntN": FxToFuncBlock(func(n int) int {
				if c.seed == nil {
					return rand.Intn(n)
				}
				r := rand.New(rand.NewSource(*c.seed))
				return r.Intn(n)
			}),
			"randNum": FxToFuncBlock(func() float64 {
				if c.seed == nil {
					return rand.Float64()
				}
				r := rand.New(rand.NewSource(*c.seed))
				return r.Float64()
			}),
			"randAlphabet": FxToFuncBlock(func() string {
				if c.seed == nil {
					return string(rune(rand.Intn(26) + 65))
				}
				r := rand.New(rand.NewSource(*c.seed))
				return string(rune(r.Intn(26) + 65))
			}),
			"shuffle": FxToFuncBlock(func(a []any) {
				if c.seed == nil {
					rand.Shuffle(len(a), func(i, j int) {
						a[i], a[j] = a[j], a[i]
					})
				}
				r := rand.New(rand.NewSource(*c.seed))
				r.Shuffle(len(a), func(i, j int) {
					a[i], a[j] = a[j], a[i]
				})
			}),
			"setRandSeed": FxToFuncBlock(func(seed int) {
				i := int64(seed)
				c.seed = &i
			}),
			"randSeed": FxToFuncBlock(func(seed int) {
				i := int64(seed)
				c.seed = &i
			}),
			"getRandSeed": FxToFuncBlock(func() int {
				if c.seed == nil {
					return 0
				}
				return int(*c.seed)
			}),
			"unsetRandSeed": FxToFuncBlock(func() {
				c.seed = nil
			}),
			"randSeedTime": FxToFuncBlock(func() {
				x := time.Now().UnixNano()
				c.seed = &x
			}),
		},
	}
	return c
}

var _ obj.ILibrary = (*StdMathLib)(nil)
