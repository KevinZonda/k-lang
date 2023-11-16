package tx

import (
	"bytes"
	"io"
	"os"
)

var out *os.File
var r *os.File
var w *os.File
var lastCall string

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

func IsCapturing() bool {
	return w != nil
}

func StopCaptureStdout() string {
	if w == nil {
		return ""
	}
	w.Close()
	w = nil
	os.Stdout = out
	var buf bytes.Buffer
	io.Copy(&buf, r)
	lastCall = buf.String()
	return buf.String()
}
