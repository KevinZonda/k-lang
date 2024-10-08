package main

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternalServer"
)

func mainX() {
	p := httpExternalServer.NewFuncPack("simple")
	p.AppendFxWithName("add", func(x int, y int) int { return x + y })
	p.AppendFx(Foo)
	p.StartServer("localhost:11451")
}

func Foo(argI int, argX bool) string {
	fmt.Println("HOOOO")
	return ""
}

func main() {
	p := httpExternalServer.NewPackage()
	p.AppendFx(Foo)

	p2 := httpExternalServer.NewPackage()
	p2.AppendFxWithName("add", func(x int, y int) int { return x + y })
	p.AppendPackage("p2", p2)
	p.EngineWithNoPackName().Run("localhost:11451")
}
