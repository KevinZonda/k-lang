package gtks

import (
	"errors"
	"io"
	"sync"
)

type GtkCodeEditorPipeWriter struct {
	buf *CodeEditor
}

func (w *GtkCodeEditorPipeWriter) Write(p []byte) (n int, err error) {
	if w.buf == nil {
		return 0, errors.New("GtkCodeEditorPipeWriter is closed")
	}

	w.buf.AppendEnd(string(p))
	return len(p), nil
}

func (ce *CodeEditor) WriterPipe(l sync.Locker) io.WriteCloser {
	return &GtkCodeEditorPipeWriter{
		buf: ce,
	}
}

func (w *GtkCodeEditorPipeWriter) Close() error {
	w.buf = nil
	return nil
}
