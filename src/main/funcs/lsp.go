package funcs

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lsp"

func Lsp(addr string) {
	if addr == "" {
		addr = "127.0.0.1:11451"
	}
	srv := lsp.NewLsp()
	err := srv.RunTCP(addr)
	if err != nil {
		panic(err)
	}
}
