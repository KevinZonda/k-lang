package compressor

import (
	"bytes"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"github.com/andybalholm/brotli"
	"github.com/brimdata/zed/zson"
	"io"
)

const (
	magicStart = "[<"
	magicEnd   = ">]"
)

func serialise(a tree.Ast) string {
	m := zson.NewMarshaler()
	m.Decorate(zson.StyleSimple)
	val, e := m.Marshal(a)
	if e != nil {
		panic(e)
	}
	return val
}

func deserialise(zson string) tree.Ast {
	um := NewUnmarshal()
	var a tree.Ast
	e := um.Unmarshal(zson, &a)
	if e != nil {
		panic(e)
	}
	return a
}

func br(bs []byte) []byte {
	var b bytes.Buffer
	w := brotli.NewWriterLevel(&b, brotli.BestCompression)
	w.Write(bs)
	w.Close()
	return b.Bytes()
}

func unbr(bs []byte) []byte {
	bs, e := io.ReadAll(brotli.NewReader(bytes.NewReader(bs)))
	if e != nil {
		panic(e)
	}
	return bs
}
func Compress(a tree.Ast) []byte {
	x := br([]byte(serialise(a)))
	return append([]byte(magicStart),
		append(x, []byte(magicEnd)...)...)
}

func Decompress(bs []byte) tree.Ast {
	begin := bytes.Index(bs, []byte(magicStart))
	end := bytes.Index(bs, []byte(magicEnd))
	if begin == -1 || end == -1 {
		panic("Not a compressed file")
	}
	bs = bs[begin+len(magicStart) : end]
	return deserialise(string(unbr(bs)))
}

func IsCompressed(bs []byte) bool {
	begin := bytes.Index(bs, []byte(magicStart))
	end := bytes.Index(bs, []byte(magicEnd))
	return begin != -1 && end != -1
}
