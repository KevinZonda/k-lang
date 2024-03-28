package httpExternal_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternal"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternalServer"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/async"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestServer(t *testing.T) {
	p := initSvr()
	addr := tester.FreeListenAddr()
	cancel := async.AsyncFunc(func() {
		p.StartServer(addr)
	})
	defer async.CancelThenSpeep(100, cancel)

	l := &httpExternal.Library{EndPoint: "http://" + addr + "/simple"}
	rst := l.InvokeFunc("add", 1, 2)
	fmt.Println(rst)
	x := external.NewExternalLibrary(l, "simple")
	fx := x.GetFunc("add")
	fmt.Println(fx.EvalBinary([]any{1, 2}))
	fx = x.GetFunc("hi")
	fmt.Println(fx.EvalBinary([]any{}))

}

func initSvr() *httpExternalServer.FuncPack {
	p := httpExternalServer.NewFuncPack("simple")
	p.AppendFxWithName("add", func(x int, y int) int { return x + y })
	p.AppendFxWithName("hi", func() { fmt.Println("hi from svr") })
	return p
}
