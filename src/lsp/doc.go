package lsp

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
	"sync"
)

type doc struct {
	lock sync.Mutex
	Uri  protocol.DocumentUri
	Text string
}

var docMap = sync.Map{}

func getDoc(uri protocol.DocumentUri) *doc {
	if d, ok := docMap.Load(uri); ok {
		return d.(*doc)
	}
	return nil
}

func rmDoc(uri protocol.DocumentUri) {
	docMap.Delete(uri)
}

func setDoc(d *doc) {
	docMap.Store(d.Uri, d)
}

func (d *doc) applyChanges(changes []interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	for _, change := range changes {
		switch c := change.(type) {
		case protocol.TextDocumentContentChangeEvent:
			begin, end := c.Range.IndexesIn(d.Text)
			if begin == -1 || end == -1 ||
				begin > end ||
				begin < 0 ||
				end > len(d.Text) {
				continue
			}
			d.Text = d.Text[:begin] + c.Text + d.Text[end:]
		case protocol.TextDocumentContentChangeEventWhole:
			d.Text = c.Text
		}
	}
}
