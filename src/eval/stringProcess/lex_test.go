package stringProcess_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/stringProcess"
	"testing"
)

func TestLex(t *testing.T) {
	s := "Hello X\nWorld"
	lexer := &stringProcess.Lexer{}
	lexer.SetRaw(s)
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.HasNext())
}
