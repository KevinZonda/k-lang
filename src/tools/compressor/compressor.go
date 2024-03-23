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

func serialise(a tree.Ast) (string, error) {
	m := zson.NewMarshaler()
	m.Decorate(zson.StyleSimple)
	return m.Marshal(a)
}

func deserialise(zson string) (tree.Ast, error) {
	um := NewUnmarshal()
	var a tree.Ast
	e := um.Unmarshal(zson, &a)
	return a, e
}

func br(bs []byte) []byte {
	var b bytes.Buffer
	w := brotli.NewWriterLevel(&b, brotli.BestCompression)
	w.Write(bs)
	w.Close()
	return b.Bytes()
}

func unbr(bs []byte) ([]byte, error) {
	return io.ReadAll(brotli.NewReader(bytes.NewReader(bs)))
}
func Compress(a tree.Ast) ([]byte, CompressorError) {
	s, err := serialise(a)
	if err != nil {
		return nil, ErrSerialiseFailed
	}
	x := br([]byte(s))
	return append([]byte(magicStart),
		append(x, []byte(magicEnd)...)...), nil
}

func Decompress(bs []byte) (tree.Ast, CompressorError) {
	begin := bytes.Index(bs, []byte(magicStart))
	end := bytes.Index(bs, []byte(magicEnd))
	if begin == -1 || end == -1 {
		return nil, ErrNotCompressed
	}
	bs = bs[begin+len(magicStart) : end]

	unbred, err := unbr(bs)
	if err != nil {
		return nil, ErrDecompressFailed
	}

	ast, err := deserialise(string(unbred))
	if err != nil {
		return nil, ErrDeserialiseFailed
	}
	return ast, nil
}

func IsCompressed(bs []byte) bool {
	begin := bytes.Index(bs, []byte(magicStart))
	end := bytes.Index(bs, []byte(magicEnd))
	return begin != -1 && end != -1
}
