package main

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/externel/httpExternelServer"
)

func main() {
	p := httpExternelServer.NewFuncPack("simple")
	p.AppendFxWithName("add", func(x int, y int) int { return x + y })
	p.AppendFx(Foo)
	p.StartServer("localhost:11451")
}

func Foo(argI int, argX bool) string {
	fmt.Println("HOOOO")
	return ""
}
