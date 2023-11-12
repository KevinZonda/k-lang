package tx

import (
	"fmt"
)

func IsStdoutAsExpected(f func(), expected string) error {
	CaptureStdout()
	f()
	output := StopCaptureStdout()

	fmt.Println(output)
	if output != expected {
		return fmt.Errorf("The code result:\n%sExpected:\n%s\n", output, expected)
	}
	return nil
}
