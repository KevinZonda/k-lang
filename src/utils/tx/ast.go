package tx

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func IsStdoutAsExpected(f func(), expected string) error {
	CaptureStdout()
	f()
	output := StopCaptureStdout()

	fmt.Println(output)
	if output != expected {
		return fmt.Errorf("The code result:\n%s\nSHA256:%s\nExpected:\n%s\nSHA256:%s\n",
			output, Sha256(output),
			expected, Sha256(expected))
	}
	return nil
}

func Sha256(o string) string {
	h := sha256.New()
	h.Write([]byte(o))
	o = hex.EncodeToString(h.Sum(nil))
	return o
}
