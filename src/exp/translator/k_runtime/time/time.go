package time

import "time"

func Sleep(i int) {
	time.Sleep(time.Duration(i) * time.Millisecond)
}

func Unix() int64 {
	return time.Now().Unix()
}

func UnixNano() int64 {
	return time.Now().UnixNano()
}

func Now() string {
	return time.Now().String()
}
