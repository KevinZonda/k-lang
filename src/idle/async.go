package idle

import (
	"context"
)

func AsyncFunc(f func()) func() {
	ctx, cancel := context.WithCancel(context.Background())
	go doWork(ctx, f)
	return cancel
}

func doWork(ctx context.Context, f func()) {
	select {
	case <-ctx.Done():
		return
	default:
		if f != nil {
			f()
		}
	}
}
