package external_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternalServer"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/async"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"github.com/gin-gonic/gin"
	"testing"
)

const TEST_ADDR = "localhost:11451"

func initSvr() *gin.Engine {
	p := httpExternalServer.NewPackage()

	p2 := httpExternalServer.NewPackage()
	p2.AppendFxWithName("add", func(x int, y int) int { return x + y })
	p.AppendPackage("p2", p2)
	return p.Engine("simple")
}

func TestOpenStmt(t *testing.T) {
	p := initSvr()
	cancel := async.AsyncFunc(func() {
		p.Run(TEST_ADDR)
	})
	defer cancel()

	code := `
open "ext/http://localhost:11451/simple/p2" as simple
z = simple.add(1, 2)
println(z)
`
	tester.GeneralTest(false, t, code, "3\n")
}
