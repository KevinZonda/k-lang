package external_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternalServer"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/async"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"github.com/gin-gonic/gin"
	"net"
	"testing"
)

func initSvr() *gin.Engine {
	p := httpExternalServer.NewPackage()

	p2 := httpExternalServer.NewPackage()
	p2.AppendFxWithName("add", func(x int, y int) int { return x + y })
	p.AppendPackage("p2", p2)
	return p.Engine("simple")
}

func initOldSvr(addr string) {
	p := httpExternalServer.NewFuncPack("simple")
	p.AppendFxWithName("add", func(x int, y int) int { return x + y })
	p.StartServer(addr)
}

func TestOpenStmt(t *testing.T) {
	l, _ := net.Listen("tcp", "localhost:0")
	l.Close()
	addr := l.Addr().String()
	p := initSvr()
	cancel := async.Run(func() {
		p.Run(addr)
	})
	defer async.CancelThenSleep(100, cancel)

	code := fmt.Sprintf(`
open "ext/http://%s/simple/p2" as simple
z = simple.add(1, 2)
println(z)
`, addr)
	tester.GeneralTest(false, t, code, "3\n")
}

func TestOpenStmtOld(t *testing.T) {
	addr := tester.FreeListenAddr()
	cancel := async.Run(func() {
		initOldSvr(addr)
	})
	defer async.CancelThenSleep(100, cancel)

	code := fmt.Sprintf(`
open "ext/http://%s/simple" as simple
z = simple.add(1, 2)
println(z)
`, addr)
	tester.GeneralTest(false, t, code, "3\n")
}
