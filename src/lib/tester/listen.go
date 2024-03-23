package tester

import "net"

const TEST_ADDR = "localhost:11451"

func FreeListenAddr() string {
	l, _ := net.Listen("tcp", "localhost:0")
	l.Close()
	return l.Addr().String()
}
