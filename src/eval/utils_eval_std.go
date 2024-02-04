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

func (e *Eval) SetStdOut(out io.WriteCloser) {
	e.builtin.StdOut = out
}

func (e *Eval) SetStdIn(in io.ReadCloser) {
	e.builtin.StdIn = in
}

func (e *Eval) SetStdErr(err io.WriteCloser) {
	e.builtin.StdErr = err
}

func (e *Eval) GetStdOut() io.WriteCloser {
	return e.builtin.StdOut
}

func (e *Eval) GetStdIn() io.ReadCloser {
	return e.builtin.StdIn
}

func (e *Eval) GetStdErr() io.WriteCloser {
	return e.builtin.StdErr
}

func (e *Eval) ResetStd() {
	e.builtin.StdOut = os.Stdout
	e.builtin.StdIn = os.Stdin
	e.builtin.StdErr = os.Stderr
}

//endregion
