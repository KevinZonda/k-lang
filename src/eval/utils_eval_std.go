package eval

import (
	"io"
	"os"
)

// region STD {IN, OUT, ERR}

func (e *Eval) LoadStdFromOS() {
	e.std.StdOut = os.Stdout
	e.std.StdIn = os.Stdin
	e.std.StdErr = os.Stderr
}

func (e *Eval) SetStdOut(out io.Writer) {
	e.std.StdOut = out
}

func (e *Eval) SetStdIn(in io.Reader) {
	e.std.StdIn = in
}

func (e *Eval) SetStdErr(err io.Writer) {
	e.std.StdErr = err
}

func (e *Eval) GetStdOut() io.Writer {
	return e.std.StdOut
}

func (e *Eval) GetStdIn() io.Reader {
	return e.std.StdIn
}

func (e *Eval) GetStdErr() io.Writer {
	return e.std.StdErr
}

func (e *Eval) ResetStd() {
	e.std.StdOut = os.Stdout
	e.std.StdIn = os.Stdin
	e.std.StdErr = os.Stderr
}

//endregion
