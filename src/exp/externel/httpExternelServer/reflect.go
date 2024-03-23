package httpExternelServer

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/externel"
	"reflect"
	"runtime"
	"strings"
)

func readFuncArg(fx any) externel.FuncArgs {
	f := reflect.TypeOf(fx)
	var args externel.FuncArgs
	n := f.NumIn()
	for i := 0; i < n; i++ {
		inT := f.In(i)
		args = append(args, externel.ArgType(inT.String()))
	}
	return args
}

func FuncName(v any) string {
	fns := strings.Split(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name(), ".")
	return fns[len(fns)-1]
}
