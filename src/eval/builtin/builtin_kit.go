package builtin

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"

type Func struct {
	f func(b obj.StdIO, args []any) any
}

type Kit struct {
	m map[string]Func
}

func NewKit() *Kit {
	return &Kit{
		m: make(map[string]Func),
	}
}

func (k *Kit) AddFunc(name string, f Func) *Kit {
	k.m[name] = f
	return k
}

func NewFuncSingle[T any](f func(any) T) Func {
	return NewFunc(func(b obj.StdIO, args []any) any {
		return f(args[0])
	})
}

func NewFunc(f func(b obj.StdIO, args []any) any) Func {
	return Func{f: f}
}

func (k *Kit) Call(b obj.StdIO, name string, args []any) any {
	f, ok := k.m[name]
	if !ok {
		panic("func not found: " + name)
	}
	return f.f(b, args)
}

var Set *Kit = NewKit().
	AddFunc("print", NewFunc(Print)).
	AddFunc("println", NewFunc(Println)).
	AddFunc("typeOf", NewFuncSingle[string](TypeOf)).
	AddFunc("len", NewFuncSingle[int](Len)).
	AddFunc("range", NewFuncSingle[[]any](Range)).
	AddFunc("panic", NewFunc(Panic))
