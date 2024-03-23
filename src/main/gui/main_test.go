package main

import (
	"os"
	"testing"
)

func TestMainProg(m *testing.T) {
	os.Args = nil
	startW()
}
