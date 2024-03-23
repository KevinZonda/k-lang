package async

import "context"

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
