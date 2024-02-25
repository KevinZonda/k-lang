package eval

import (
	"io"
	"os"
)

type IO struct {
	StdIn  io.Reader
	StdOut io.Writer
	StdErr io.Writer
}

func (i IO) GetStdout() io.Writer {
	return i.StdOut
}

func (i IO) GetStdin() io.Reader {
	return i.StdIn
}

func (i IO) GetStderr() io.Writer {
	return i.StdErr
}

func NewIO() IO {
	return IO{
		StdIn:  os.Stdin,
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}
}
