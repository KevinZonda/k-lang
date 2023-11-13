package parserHelper

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/visitor"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
)

// WrappedParser this is self defined parser
type WrappedParser struct {
	lexer  *parser.V2Lexer
	parser *parser.V2Parser
	lexE   *errorListener
	parE   *errorListener
}

func (p *WrappedParser) Ast() tree.Ast {
	return p.
		parser.
		Program().
		Accept(visitor.New()).(tree.Ast)
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
		lexE: &errorListener{},
		parE: &errorListener{},
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

func Ast(s string) (tree.Ast, []CodeError) {
	p := FromString(s)
	ast := p.Ast()
	return ast, p.Errors()
}
