package TMem

import (
	"fmt"
	"strings"
)

func (t *Layer) String() string {
	return t.ToString("")
}

func (t *Layer) ToString(prefix string) string {
	w := &strings.Builder{}
	for k, v := range t.m {
		fmt.Fprintf(w, "%s%s: %v\n", prefix, k, v)
	}
	return w.String()
}

func (t *Memory) String() string {
	return t.ToString()
}

func (t *Memory) ToString() string {
	w := &strings.Builder{}
	fmt.Fprintln(w, "***********************************")
	fmt.Fprintln(w, "TYPE STACK")
	fmt.Fprintln(w, "***********************************")
	if t.Empty() {
		fmt.Fprintln(w, "<EMPTY>")
		fmt.Fprintln(w, "***********************************")
		return w.String()
	}
	for i := len(t.q) - 1; i >= 0; i-- {
		fmt.Fprint(w, "LEVEL [", i, "]")
		if i == 0 {
			fmt.Fprint(w, " GLOBAL")
		}
		if i == len(t.q)-1 {
			fmt.Fprint(w, " TOP")
		}
		if t.q[i].Protect {
			fmt.Fprintln(w, " PROTECTED")
		} else {
			fmt.Fprintln(w)
		}
		fmt.Fprint(w, t.q[i].ToString("  -> "))
	}
	fmt.Fprintln(w, "***********************************")
	return w.String()
}
