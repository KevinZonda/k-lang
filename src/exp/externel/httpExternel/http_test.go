package httpExternel_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/externel/httpExternel"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/externel/httpExternelServer"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/async"
	"testing"
)

const TEST_ADDR = "localhost:11451"

func TestServer(t *testing.T) {
	p := initSvr()
	cancel := async.AsyncFunc(func() {
		p.StartServer(TEST_ADDR)
	})
	defer cancel()

	l := &httpExternel.Library{EndPoint: "http://" + TEST_ADDR + "/simple"}
	rst := l.InvokeFunc("add", 1, 2)
	fmt.Println(rst)

}

func initSvr() *httpExternelServer.FuncPack {
	p := httpExternelServer.NewFuncPack("simple")
	p.AppendFxWithName("add", func(x int, y int) int { return x + y })
	return p
}
