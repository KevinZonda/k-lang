package parserHelper

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type CodeError interface {
	error
	Line() int
	Col() int
}

type SyntaxError struct {
	line, column int
	msg          string
}

func (s SyntaxError) Error() string {
	return fmt.Sprintf("line %d:%d -> %s", s.line, s.column, s.msg)
}

func (s SyntaxError) Line() int {
	return s.line
}

func (s SyntaxError) Col() int {
	return s.column
}

type errorListener struct {
	*antlr.DefaultErrorListener
	Errors []CodeError
}

func (c *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	synE := SyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	}
	c.Errors = append(c.Errors, synE)
}

func PrintAllCodeErrors(errs []CodeError) {
	fmt.Println("Parsing Error(s):")
	for i, e := range errs {
		fmt.Println("[", i, "] :", e)
	}
}
