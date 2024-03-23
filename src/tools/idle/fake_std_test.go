package idle_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/idle"
	"testing"
	"time"
)

func TestFakeStd(t *testing.T) {
	std := idle.NewFakeStdIn()
	std.Write([]byte("Hello"))
	bs := make([]byte, 5)
	n, err := std.Read(bs)
	tester.Assert(t, n, 5)
	tester.Assert(t, err, nil)
	tester.Assert(t, string(bs), "Hello")

	std.Close()
	_, err = std.Read(bs)
	tester.Assert(t, err.Error(), "EOF")
	_, err = std.Write([]byte("Hello"))
	tester.Assert(t, err.Error(), "EOF")
}

func TestFakeStdHolding(t *testing.T) {
	std := idle.NewFakeStdIn()

	bs := make([]byte, 5)
	go func() {
		time.Sleep(100 * time.Millisecond)
		std.Write([]byte("Hello"))
		time.Sleep(100 * time.Millisecond)
	}()
	std.Read(bs)
	tester.Assert(t, string(bs), "Hello")
}
