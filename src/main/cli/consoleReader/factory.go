package consoleReader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func New(prompt string) (IConsoleReader, error) {
	return NewSimpleConsoleReader(prompt), nil
	//return readline.New(prompt)
}

func MultipleLine(rl IConsoleReader, line string, defaultPrompt string, multiPrompt string) (string, error) {
	var err error
	buffer := ""
	for strings.HasSuffix(line, "\\") {
		buffer += strings.TrimSuffix(line, "\\") + "\n"
		rl.SetPrompt(multiPrompt)
		line, err = rl.Readline()
		if err != nil {
			return "", err
		}
	}
	buffer += line
	rl.SetPrompt(defaultPrompt)
	return buffer, err
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
	v, e := s.scanner.ReadString('\n')
	if e != nil {
		return v, e
	}
	v = strings.TrimRight(v, "\r\n"+string(rune(13)))
	return v, e
}
