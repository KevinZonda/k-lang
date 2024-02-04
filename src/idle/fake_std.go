package idle

import (
	"bytes"
	"errors"
	"io"
	"time"
)

type FakeStdIn struct {
	buf *bytes.Buffer
}

func (f *FakeStdIn) Read(p []byte) (n int, err error) {
	n, err = f.buf.Read(p)
	if err == nil {
		return n, nil
	}
	if errors.Is(err, io.EOF) {
		time.Sleep(100 * time.Millisecond)
		return f.Read(p)
	}
	return n, err
}

func (f *FakeStdIn) Write(p []byte) (n int, err error) {
	return f.buf.Write(p)
}
