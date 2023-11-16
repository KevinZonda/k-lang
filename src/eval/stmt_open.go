package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"path"
	"path/filepath"
	"strings"
)

func (e *Eval) openFile(s string) (abs, content string, ok bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", "", false
	}
	paths := []string{path.Join(e.basePath, s), s}
	for _, p := range paths {
		//{
		//	a, e := filepath.Abs(p)
		//	fmt.Println("TRY OPEN ->", a, e)
		//}
		if iox.ExistFile(p) {
			txt, err := iox.ReadAllText(p)
			abs, _ = filepath.Abs(p)
			return abs, txt, err == nil
		}
	}
	return "", "", false
}

func (e *Eval) EvalOpenStmt(n *node.OpenStmt) {
	abs, txt, ok := e.openFile(n.Path)
	if !ok {
		panic("File not found: " + n.Path)
		return
	}

	if _, ok = openedFiles[abs]; ok {
		return
	}

	ast, errs := parserHelper.Ast(txt)
	if len(errs) > 0 {
		parserHelper.PrintAllCodeErrors(errs)
		panic("Parse failed: " + n.Path)
		return
	}
	_e := New(ast, n.Path)
	openedFiles[abs] = _e
	if n.As != "" {
		e.opened[n.As] = _e
	} else {
		base := filepath.Base(abs)
		sb := strings.Builder{}
		for _, c := range base {
			if c != '_' && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
				sb.WriteRune('_')
			} else {
				sb.WriteRune(c)
			}
		}
		e.opened[sb.String()] = _e
		sb.Reset()
	}
	_e.run()
	// e.new(ast).run()
}
