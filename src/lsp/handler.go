package lsp

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
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

func (d *doc) pushDiagnoses(notify glsp.NotifyFunc) {
	if d == nil {
		return
	}
	ast, errs := parserHelper.Ast(d.Text)
	d.Asl = astToASL(ast)
	dia := make([]protocol.Diagnostic, 0)
	diagnoseE := protocol.DiagnosticSeverityError
	for _, err := range errs {
		dia = append(dia, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      protocol.UInteger(err.Line() - 1),
					Character: protocol.UInteger(err.Col()),
				},
				End: protocol.Position{
					Line:      protocol.UInteger(err.EndLine() - 1),
					Character: protocol.UInteger(err.EndCol()),
				},
			},
			Message:  err.Error(),
			Severity: &diagnoseE,
		})
	}
	notify(
		protocol.ServerTextDocumentPublishDiagnostics,
		protocol.PublishDiagnosticsParams{
			URI:         d.Uri,
			Diagnostics: dia,
		})

}
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
		TextDocumentDidChange: func(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
			d := docStore.getDoc(params.TextDocument.URI)
			if d == nil {
				return nil
			}
			d.applyChanges(params.ContentChanges)
			d.pushDiagnoses(context.Notify)
			return nil
		},
		TextDocumentDidClose: func(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
			docStore.rmDoc(params.TextDocument.URI)
			return nil
		},
		TextDocumentDidSave: func(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
			docStore.
				loadDoc(params.TextDocument.URI).
				pushDiagnoses(context.Notify)
			return nil
		},
		TextDocumentDidOpen: func(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
			docStore.
				loadDoc(params.TextDocument.URI).
				pushDiagnoses(context.Notify)
			return nil
		},
	}
	handler.SetInitialized(true)
}
