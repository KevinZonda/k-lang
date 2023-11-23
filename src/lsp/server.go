package lsp

import (
	"github.com/tliron/commonlog"
	_ "github.com/tliron/commonlog/simple"
	"github.com/tliron/glsp/server"
)

type Server struct {
	*server.Server
}

func NewLsp() *Server {
	commonlog.Configure(2, nil)
	srv := &Server{
		Server: server.NewServer(&handler, "K", true),
	}
	return srv
}
