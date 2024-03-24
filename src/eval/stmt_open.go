package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
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

func (e *Eval) loadExternalLibrary(name, as string) bool {
	if !strings.HasPrefix(name, "ext/") {
		return false
	}
	name = name[4:]
	lib := external.GetExternalLibrary(name)
	if lib == nil {
		panic("Library not found: " + name)
	}
	if as == "" {
		as = normaliseName(name)
	}
	e.memory.Top().SetValue(as, lib)
	return true
}

func (e *Eval) loadBuiltInFeature(name string) bool {
	if !strings.HasPrefix(name, "feat/") {
		return false
	}
	val := true
	if strings.HasSuffix(name, ":on") || strings.HasSuffix(name, ":true") {
		name = name[:strings.LastIndex(name, ":")]
	} else if strings.HasSuffix(name, ":off") || strings.HasSuffix(name, ":false") {
		name = name[:strings.LastIndex(name, ":")]
		val = false
	}
	switch name {
	case "feat/staticType":
		e.FeatStaticType = val
	case "feat/verbose":
		e.FeatVerbose = val
	case "feat/unknownVarNil":
		e.FeatUnknownVarNil = val
	case "feat/refAll":
		e.FeatRefAll = val
	case "feat/visualise", "feat/visualize":
		e.visualise = val
	default:
		return false
	}
	return true
}

func (e *Eval) loadBuiltInLibrary(name, as string) (ok bool) {
	lib := builtin.GetLibrary(name)
	if lib == nil {
		return false
	}

	if as != "" {
		name = as
	}
	e.memory.Top().SetValue(normaliseName(name), lib)
	return true
}

func (e *Eval) EvalOpenStmt(n *node.OpenStmt) {
	e.currentToken = n.GetToken()

	if e.loadBuiltInFeature(n.Path) {
		return
	}

	if e.loadBuiltInLibrary(n.Path, n.As) {
		return
	}

	if e.loadExternalLibrary(n.Path, n.As) {
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
			panic(fmt.Sprint("Error reading file: ", abs, err))
		}

		ast, errs := parserHelper.Ast(txt)
		errs.PanicIfError()

		openedEval = New().WithBasePath(n.Path)
		openedFiles[abs] = openedEval
		openedEval.runAst(ast, reserved.Return)
	}

	name := ""
	if n.As != "" {
		name = n.As
	} else {
		name = normaliseName(filepath.Base(abs))
	}
	e.memory.Top().Set(name, obj.NewImmutableObj(obj.EvalObj, openedEval))
}

func normaliseName(n string) string {
	ns := strings.Split(n, "/")
	n = ns[len(ns)-1]
	if strings.HasSuffix(n, ".k") {
		n = n[:len(n)-2]
	}
	sb := strings.Builder{}
	for i, c := range n {
		if i == 0 && (c > '0' && c < '9') {
			sb.WriteRune('_')
			continue
		}
		if c != '_' && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
			sb.WriteRune('_')
		} else {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
