package external_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternalServer"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/async"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

const TEST_ADDR = "localhost:11451"

func initSvr() *httpExternalServer.FuncPack {
	p := httpExternalServer.NewFuncPack("simple")
	p.AppendFxWithName("add", func(x int, y int) int { return x + y })
	return p
}

func TestOpenStmt(t *testing.T) {
	p := initSvr()
	cancel := async.AsyncFunc(func() {
		p.StartServer(TEST_ADDR)
	})
	defer cancel()

	code := `
open "ext/http://localhost:11451/simple" as simple
z = simple.add(1, 2)
println(z)
`
	tester.GeneralTest(false, t, code, "3\n")
}
