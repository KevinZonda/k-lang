package parserHelper

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/visitor"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
	"strings"
)

// WrappedParser this is self defined parser
type WrappedParser struct {
	lexer  *parser.V2Lexer
	parser *parser.V2Parser
	lexE   *ErrorListener
	parE   *ErrorListener
}

func (p *WrappedParser) Ast() tree.Ast {
	v := visitor.New()
	a := p.
		parser.
		Program().
		Accept(v).(tree.Ast)

	for _, e := range v.Errs {
		msg := "Related Code: " + e.Text
		if e.Msg != "" {
			msg = e.Msg + "\n" + msg
		}

		p.parE.Errors = append(p.parE.Errors,
			SyntaxError{
				line:    e.Line,
				column:  e.Column,
				endLine: e.EndLine,
				endCol:  e.EndColumn,
				msg:     msg,
			})
	}

	return a
}

// region Getter & Setter

func (p *WrappedParser) LexerErrors() []CodeError {
	return p.lexE.Errors
}

func (p *WrappedParser) ParserErrors() []CodeError {
	return p.parE.Errors
}

func (p *WrappedParser) Errors() []CodeError {
	return append(p.LexerErrors(), p.ParserErrors()...)
}

func (p *WrappedParser) Lexer() *parser.V2Lexer {
	return p.lexer
}

func (p *WrappedParser) Parser() *parser.V2Parser {
	return p.parser
}

//endregion

// region Constructor

func FromString(s string) *WrappedParser {
	return FromLexer(parser.NewV2Lexer(antlr.NewInputStream(s)))
}

func FromLexer(l *parser.V2Lexer) *WrappedParser {
	w := &WrappedParser{
		lexE: &ErrorListener{},
		parE: &ErrorListener{},
	}

	l.RemoveErrorListeners()
	l.AddErrorListener(w.lexE)

	w.lexer = l

	p := parser.NewV2Parser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel))

	p.RemoveErrorListeners()
	p.AddErrorListener(w.parE)
	w.parser = p

	return w
}

//endregion

func Ast(s string) (tree.Ast, CodeErrors) {
	p := FromString(s)
	ast := p.Ast()
	return ast, p.Errors()
}

type CodeErrors []CodeError

func (c CodeErrors) String() string {
	sb := &strings.Builder{}
	sb.WriteString("Parser Error(s):\n")
	for idx, err := range c {
		sb.WriteString("[" + strconv.Itoa(idx) + "] ")
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}
	fmt.Fprintln(sb, "Count:", len(c))
	return sb.String()
}

func (c CodeErrors) PanicIfError() {
	if len(c) > 0 {
		panic(c.String())
	}
}
