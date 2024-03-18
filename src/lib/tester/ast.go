package tester

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func IsStdoutAsExpected(allowPanic bool, f func(), expected string) (e error) {
	output := Stdout(allowPanic, f)

	if output != expected {
		return fmt.Errorf("The code result:\n%s\nSHA256:%s\nExpected:\n%s\nSHA256:%s\n",
			output, Sha256(output),
			expected, Sha256(expected))
	}
	return nil
}

func Stdout(allowPanic bool, f func()) (output string) {
	CaptureStdout()
	defer func() {
		if !allowPanic {
			if rec := recover(); rec != nil {
				fmt.Println("Recovered in tester", rec)
				if IsCapturing() {
					StopCaptureStdout()
				}
				fmt.Println("Content Captured:\n" + lastCall)
				output = lastCall
			}
		}
		if IsCapturing() {
			output = StopCaptureStdout()
			fmt.Println("Content Captured:\n" + output)
		}

	}()
	f()
	output = StopCaptureStdout()
	fmt.Println("Content Captured:\n" + output)

	return output
}

func Sha256(o string) string {
	h := sha256.New()
	h.Write([]byte(o))
	o = hex.EncodeToString(h.Sum(nil))
	return o
}
