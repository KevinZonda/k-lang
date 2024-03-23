package idle_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/idle"
	"testing"
)

func TestLifecycle(t *testing.T) {
	c := idle.Counter{}
	c.IncrementCount()
	if c.GetCount() != 1 {
		t.Error("Expected 1, got ", c.GetCount())
	}
	c.Decrease()
	if c.GetCount() != 0 {
		t.Error("Expected 0, got ", c.GetCount())
	}
}
