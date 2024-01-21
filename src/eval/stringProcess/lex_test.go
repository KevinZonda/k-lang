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
	fmt.Println(lexer.Next())
	fmt.Println(lexer.HasNext())
}

func TestLex2(t *testing.T) {
	s := "Hello X\nWorld {"
	lexer := &stringProcess.Lexer{}
	lexer.SetRaw(s)
	lexer.SetMode(stringProcess.ModeVar)
	fmt.Println(lexer.Next())
	fmt.Println(lexer.HasNext())
}

func TestLex3(t *testing.T) {
	s := "Hello X\nWorld {VAR} {{VAR}}"
	lexer := &stringProcess.Lexer{}
	lexer.SetRaw(s)
	lexer.SetMode(stringProcess.ModeVar)

	for lexer.HasNext() {
		kind, token := lexer.Next()
		fmt.Println("Kind:", kind, "Token:", token)
	}
}
