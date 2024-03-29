package async

import (
	"context"
	"time"
)

func Run(f func()) func() {
	ctx, cancel := context.WithCancel(context.Background())
	go do(ctx, f)
	return cancel
}

func do(ctx context.Context, f func()) {
	select {
	case <-ctx.Done():
		return
	default:
		if f != nil {
			f()
		}
	}
}

func CancelThenSleep(sleepMs int, cancel ...func()) {
	for _, c := range cancel {
		c()
	}
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
}
