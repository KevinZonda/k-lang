package lsp

import (
	"github.com/KevinZonda/GoX/pkg/iox"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"log"
	"net/url"
	"runtime"
	"strings"
	"sync"
)

type doc struct {
	lock sync.Mutex
	Uri  protocol.DocumentUri
	Text string
	Asl  *ASL
}

type docStoreType struct {
	docMap sync.Map
}

var docStore = &docStoreType{
	docMap: sync.Map{},
}

func (s *docStoreType) getDoc(uri protocol.DocumentUri) *doc {
	if d, ok := s.docMap.Load(uri); ok {
		return d.(*doc)
	}
	return nil
}

func (s *docStoreType) rmDoc(uri protocol.DocumentUri) {
	s.docMap.Delete(uri)
}

func (s *docStoreType) setDoc(d *doc) {
	s.docMap.Store(d.Uri, d)
}

func (s *docStoreType) loadDoc(uri protocol.DocumentUri) *doc {
	path := string(uri)
	if strings.HasPrefix(path, "file://") {
		path = path[7:]
	}
	path, err := url.QueryUnescape(path)
	if err != nil {
		log.Println(path, err)
		return nil
	}
	if runtime.GOOS == "windows" && strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	txt, err := iox.ReadAllText(path)
	if err != nil {
		return nil
	}
	d := &doc{
		Uri:  uri,
		Text: txt,
	}
	docStore.setDoc(d)
	return d
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
