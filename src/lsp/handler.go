package lsp

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"strings"
	"sync"
)

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()
	ver := "0.0.1"
	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    "LSP for K",
			Version: &ver,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func completionItemResolve(context *glsp.Context, params *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	fmt.Println("completionItemResolve ->", *params)
	return nil, nil
}

func completion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	// params.TextDocument.URI
	return nil, nil
}

var handler protocol.Handler
var initLock sync.Mutex

func initHandler() {
	initLock.Lock()
	defer initLock.Unlock()
	if handler.IsInitialized() {
		return
	}
	handler = protocol.Handler{
		Initialize:  initialize,
		Initialized: initialized,
		Shutdown:    shutdown,
		SetTrace:    setTrace,

		//CompletionItemResolve:  completionItemResolve,

		TextDocumentCompletion: completion,
		TextDocumentDidSave: func(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
			path := string(params.TextDocument.URI)
			if strings.HasPrefix(path, "file://") {
				path = path[7:]
			}
			txt, _ := iox.ReadAllText(path)
			_, errs := parserHelper.Ast(txt)
			dia := make([]protocol.Diagnostic, 0)
			diagnoseE := protocol.DiagnosticSeverityError
			for _, err := range errs {
				dia = append(dia, protocol.Diagnostic{
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      protocol.UInteger(err.Line() - 1),
							Character: protocol.UInteger(err.Col()) - 1,
						},
						End: protocol.Position{
							Line:      protocol.UInteger(err.EndLine() - 1),
							Character: protocol.UInteger(err.EndCol() - 1),
						},
					},
					Message:  err.Error(),
					Severity: &diagnoseE,
				})
			}
			context.Notify(
				protocol.ServerTextDocumentPublishDiagnostics,
				protocol.PublishDiagnosticsParams{
					URI:         params.TextDocument.URI,
					Diagnostics: dia,
				})

			return nil
		},
	}
	handler.SetInitialized(true)
}
