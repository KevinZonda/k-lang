package main

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/lsp"
)

func main() {
	srv := lsp.NewLsp()
	srv.RunTCP("127.0.0.1:11451")
}
