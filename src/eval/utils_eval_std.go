package eval

import (
	"io"
	"os"
)

// region STD {IN, OUT, ERR}

func (e *Eval) LoadStdFromOS() {
	e.builtin.StdOut = os.Stdout
	e.builtin.StdIn = os.Stdin
	e.builtin.StdErr = os.Stderr
}

func (e *Eval) SetStdOut(out io.Writer) {
	e.builtin.StdOut = out
}

func (e *Eval) SetStdIn(in io.Reader) {
	e.builtin.StdIn = in
}

func (e *Eval) SetStdErr(err io.Writer) {
	e.builtin.StdErr = err
}

func (e *Eval) GetStdOut() io.Writer {
	return e.builtin.StdOut
}

func (e *Eval) GetStdIn() io.Reader {
	return e.builtin.StdIn
}

func (e *Eval) GetStdErr() io.Writer {
	return e.builtin.StdErr
}

func (e *Eval) ResetStd() {
	e.builtin.StdOut = os.Stdout
	e.builtin.StdIn = os.Stdin
	e.builtin.StdErr = os.Stderr
}

//endregion
