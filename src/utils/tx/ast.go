package tx

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func IsStdoutAsExpected(allowPanic bool, f func(), expected string) (e error) {
	CaptureStdout()
	defer func() {
		if !allowPanic {
			if rec := recover(); rec != nil {
				fmt.Println("Recovered in tester", rec)
				if IsCapturing() {
					StopCaptureStdout()
				}
				fmt.Println("Content Captured:\n", lastCall)
				e = fmt.Errorf("code panic! %v", rec)
			}
		}
		if IsCapturing() {
			output := StopCaptureStdout()
			fmt.Println("Content Captured:\n", output)
		}

	}()
	f()
	output := StopCaptureStdout()
	fmt.Println("Content Captured:\n", output)

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
