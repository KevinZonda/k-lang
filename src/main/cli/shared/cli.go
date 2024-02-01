package shared

import (
	"flag"
)

var Input = flag.String("input", "", "Input file")
var Output = flag.String("output", "", "Output file")

func GetStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
