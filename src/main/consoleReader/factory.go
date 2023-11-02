package consoleReader

import (
	"bufio"
	"fmt"
	"os"
)

func New(prompt string) (IConsoleReader, error) {
	return NewSimpleConsoleReader(prompt), nil
	// return readline.New(prompt)
}

type IConsoleReader interface {
	SetPrompt(prompt string)
	Readline() (string, error)
	Close() error
}

type SimpleConsoleReader struct {
	prompt  string
	scanner *bufio.Reader
}

func NewSimpleConsoleReader(prompt string) *SimpleConsoleReader {
	return &SimpleConsoleReader{
		prompt:  prompt,
		scanner: bufio.NewReader(os.Stdin),
	}
}

func (s *SimpleConsoleReader) Close() error {
	return nil
}

func (s *SimpleConsoleReader) SetPrompt(prompt string) {
	s.prompt = prompt
}

func (s *SimpleConsoleReader) Readline() (string, error) {
	fmt.Print(s.prompt)
	return s.scanner.ReadString('\n')
}
