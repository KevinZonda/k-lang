package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"io"
)

// region STD {IN, OUT, ERR}

func (e *Eval) LoadStdFromOS() {
	e.std = obj.DefaultIO()
}

func (e *Eval) SetStdOut(out io.Writer) {
	e.std.SetStdOut(out)
}

func (e *Eval) SetStdIn(in io.Reader) {
	e.std.SetStdIn(in)
}

func (e *Eval) SetStdErr(err io.Writer) {
	e.std.SetStdErr(err)
}

func (e *Eval) GetIO() obj.StdIO {
	return e.std
}

//endregion
