package async

import (
	"context"
	"time"
)

func AsyncFunc(f func()) func() {
	ctx, cancel := context.WithCancel(context.Background())
	go DoWork(ctx, f)
	return cancel
}

func DoWork(ctx context.Context, f func()) {
	select {
	case <-ctx.Done():
		return
	default:
		if f != nil {
			f()
		}
	}
}

func CancelThenSpeep(sleepMs int, cancel ...func()) {
	for _, c := range cancel {
		c()
	}
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
}
