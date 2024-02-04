package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const K_HOME = "K_HOME"

func (e *Eval) findFile(s string) (abs string, ok bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", false
	}
	paths := []string{path.Join(e.basePath, s)}
	if khome := os.Getenv(K_HOME); khome != "" {
		paths = append(paths, path.Join(khome, s))
	}
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
		name = as
	}
	e.objTable.Set(normaliseName(name), lib)
	return true
}

func (e *Eval) EvalOpenStmt(n *node.OpenStmt) {
	e.currentToken = n.GetToken()
	if e.loadBuiltInLibrary(n.Path, n.As) {
		return
	}

	abs, ok := e.findFile(n.Path)
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
		parserHelper.IfErrorsPrintAndPanic(errs)

		openedEval = New(ast, n.Path)
		openedFiles[abs] = openedEval
		openedEval.runWithBreak(reserved.Return)
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
	if strings.HasSuffix(n, ".k") {
		n = n[:len(n)-2]
	}
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
