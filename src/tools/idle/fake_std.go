package idle

import (
	"bytes"
	"errors"
	"io"
	"time"
)

type FakeStdIn struct {
	buf   *bytes.Buffer
	close bool
}

func NewFakeStdIn() *FakeStdIn {
	return &FakeStdIn{
		buf:   bytes.NewBuffer([]byte{}),
		close: false,
	}
}
func (f *FakeStdIn) Read(p []byte) (n int, err error) {
	if f.close {
		return 0, io.EOF
	}
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
	if f.close {
		return 0, io.EOF
	}
	return f.buf.Write(p)
}

func (f *FakeStdIn) Close() error {
	f.close = true
	return nil
}
