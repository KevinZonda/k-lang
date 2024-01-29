package idle

import (
	"bytes"
	"log"
)

func NewFakeWCloser() *FakeWrCloser {
	buf := bytes.NewBuffer([]byte{})
	w := FakeWrCloser{
		Buf: buf,
	}
	log.Println("NewFakeWCloser", w, "Buffer:", buf, "Buffer Bytes:", buf.Bytes())
	return &w
}

type FakeWrCloser struct {
	Buf *bytes.Buffer
}

func (f *FakeWrCloser) Close() error {
	return nil
}

func (f *FakeWrCloser) ReadAll() []byte {
	return f.Buf.Bytes()
}

func (f *FakeWrCloser) ReadAllString() string {
	return f.Buf.String()
}

func (f *FakeWrCloser) Write(p []byte) (int, error) {
	log.Printf("%p ", f.Buf)
	log.Println("BEFORE:", f.Buf.Bytes())
	v, e := f.Buf.Write(p)
	log.Println(f.Buf.Bytes())
	log.Println("FakeWrCloser.Write", string(p), "BUF:", f.ReadAllString())
	return v, e
}
