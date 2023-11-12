package tx

import (
	"bytes"
	"io"
	"os"
)

var out *os.File
var r *os.File
var w *os.File

func CaptureStdout() {
	if out == nil {
		out = os.Stdout
	}
	if w != nil {
		panic("STDOUT IS CAPTURING")
	}
	r, w, _ = os.Pipe()
	os.Stdout = w
}

func StopCaptureStdout() string {
	if w == nil {
		return ""
	}
	w.Close()
	os.Stdout = out
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
