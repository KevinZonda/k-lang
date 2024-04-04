package obj

import (
	"bytes"
	"io"
	"os"
)

func DefaultIO() StdIO {
	return StdIO{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}
}

type StdIO struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

func (i *StdIO) SetStdIn(in io.Reader) {
	i.in = in
}

func (i *StdIO) SetStdOut(out io.Writer) {
	i.out = out
}

func (i *StdIO) SetStdErr(err io.Writer) {
	i.err = err
}

type eofReader struct{}

func (eofReader) Read(_ []byte) (n int, err error) {
	return 0, io.EOF
}

func (i StdIO) GetStdIn() io.Reader {
	if i.in != nil {
		return i.in
	}
	if os.Stdin != nil {
		return os.Stdin
	}
	return eofReader{}
}

func (i StdIO) GetStdOut() io.Writer {
	if i.out != nil {
		return i.out
	}
	if os.Stdout != nil {
		return os.Stdout
	}
	if os.Stderr != nil {
		return os.Stderr
	}
	return bytes.NewBuffer(nil)
}

func (i StdIO) GetStdErr() io.Writer {
	if i.err != nil {
		return i.err
	}
	if os.Stderr != nil {
		return os.Stderr
	}
	if os.Stdout != nil {
		return os.Stdout
	}
	return bytes.NewBuffer(nil)
}
