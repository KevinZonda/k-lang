package memory

import (
	"fmt"
	"strings"
)

func (t *Layer) String() string {
	return t.ToString("")
}

func (t *Layer) ToString(prefix string) string {
	w := &strings.Builder{}
	for pair := t.m.Newest(); pair != nil; pair = pair.Prev() {
		fmt.Fprintf(w, "%s%s: %v\n", prefix, pair.Key, pair.Value)
	}
	return w.String()
}

func (t *Memory) String() string {
	return t.ToString(fmt.Sprintf("%p", t))
}

func (t *Memory) ToString(addr string) string {
	w := &strings.Builder{}
	fmt.Fprintln(w, "***********************************")
	fmt.Fprintln(w, "V-MEM STACK @", addr)
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
