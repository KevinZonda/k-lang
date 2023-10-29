package parserHelper

import (
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/tree"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/visitor"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/parser"
	"github.com/antlr4-go/antlr/v4"
)

// WrappedParser this is self defined parser
type WrappedParser struct {
	lexer  *parser.V2Lexer
	parser *parser.V2Parser
	errors []string
}

func (p *WrappedParser) Ast() tree.Ast {
	return p.
		parser.
		Program().
		Accept(visitor.New()).(tree.Ast)
}

// region Getter & Setter

func (p *WrappedParser) Errors() []string {
	return p.errors
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
	return &WrappedParser{
		lexer:  l,
		parser: parser.NewV2Parser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)),
	}
}

//endregion
