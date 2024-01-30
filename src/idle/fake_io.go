package idle

import (
	"bufio"
	"bytes"
)

func NewFakeWCloser() *FakeWrCloser {
	buf := &bytes.Buffer{}
	w := FakeWrCloser{
		Buf:    buf,
		Writer: bufio.NewWriter(buf),
	}
	return &w
}

type FakeWrCloser struct {
	*bufio.Writer
	Buf *bytes.Buffer
}

func (f *FakeWrCloser) Close() error {
	return nil
}

func (f *FakeWrCloser) ReadAll() []byte {
	f.Flush()
	return f.Buf.Bytes()
}

func (f *FakeWrCloser) ReadAllString() string {
	f.Flush()
	return f.Buf.String()
}
