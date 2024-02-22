package idle

import (
	"context"
	"sync"
)

func AsyncFunc(l sync.Locker, f func()) func() {
	ctx, cancel := context.WithCancel(context.Background())
	go doWork(ctx, l, f)
	return cancel
}

func doWork(ctx context.Context, l sync.Locker, f func()) {
	select {
	case <-ctx.Done():
		l.Lock()
		defer l.Unlock()

		return
	default:
		if f != nil {
			f()
		}
	}
}
