package idle

import (
	"github.com/gotk3/gotk3/gtk"
	"sync/atomic"
)

type Counter struct {
	count atomic.Int32
}

func (c *Counter) GetCount() int32 {
	return c.count.Load()
}

func (c *Counter) IncrementCount() {
	c.count.Add(1)
}

func (c *Counter) Decrease() {
	c.count.Add(-1)
	if c.GetCount() <= 0 {
		gtk.MainQuit()
	}
}

var lifecycle = &Counter{}
