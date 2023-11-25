package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"path"
	"path/filepath"
	"strings"
)

func (e *Eval) fineFile(s string) (abs string, ok bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", false
	}
	paths := []string{path.Join(e.basePath, s), s}
	for _, p := range paths {
		//{
		//	a, e := filepath.Abs(p)
		//	fmt.Println("TRY OPEN ->", a, e)
		//}
		if iox.ExistFile(p) {
			var err error
			abs, err = filepath.Abs(p)
			return abs, err == nil
		}
	}
	return "", false
}

func (e *Eval) loadBuiltInLibrary(name, as string) (ok bool) {
	lib := builtin.GetLibrary(name)
	if lib == nil {
		return false
	}

	if as != "" {
		e.objTable.Set(normaliseName(as), lib)
	} else {
		e.objTable.Set(name, lib)
	}
	return true
}

func (e *Eval) EvalOpenStmt(n *node.OpenStmt) {
	if e.loadBuiltInLibrary(n.Path, n.As) {
		return
	}

	abs, ok := e.fineFile(n.Path)
	if !ok {
		panic("File not found: " + n.Path)
		return
	}

	var openedEval *Eval

	if openedEval, ok = openedFiles[abs]; !ok {
		txt, err := iox.ReadAllText(abs)
		if err != nil {
			panic(fmt.Sprintln("Error reading file: ", abs, err))
		}

		ast, errs := parserHelper.Ast(txt)
		if len(errs) > 0 {
			parserHelper.PrintAllCodeErrors(errs)
			panic("Parse failed: " + n.Path)
		}
		openedEval = New(ast, n.Path)
		openedFiles[abs] = openedEval
		openedEval.run()
	}

	if n.As != "" {
		e.objTable.Set(n.As, openedEval)
	} else {
		base := filepath.Base(abs)
		e.objTable.Set(normaliseName(base), openedEval)
	}
}

func normaliseName(n string) string {
	ns := strings.Split(n, "/")
	n = ns[len(ns)-1]
	sb := strings.Builder{}
	for _, c := range n {
		if c != '_' && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			sb.WriteRune('_')
		} else {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
