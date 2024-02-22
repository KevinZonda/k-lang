package eval

import (
	"fmt"
	"io"
	"strings"
)

type Memory struct {
	q []*MemLayer
}

func (t *Memory) Bottom() *MemLayer {
	if t.Empty() {
		return nil
	}
	return t.q[0]
}

func (t *Memory) Raw() []*MemLayer {
	return t.q
}

func (t *Memory) Push(n *MemLayer) {
	t.q = append(t.q, n)
}

func (t *Memory) PushEmpty(protect bool) *MemLayer {
	tbl := newTable(protect)
	t.q = append(t.q, tbl)
	return tbl
}

func (t *Memory) Pop() *MemLayer {
	n := t.q[len(t.q)-1]
	t.q = t.q[:len(t.q)-1]
	return n
}

func (t *Memory) Peek() *MemLayer {
	if t.Empty() {
		return nil
	}
	return t.q[len(t.q)-1]
}

func (t *Memory) Len() int {
	return len(t.q)
}

func (t *Memory) Empty() bool {
	return len(t.q) == 0
}

func (t *Memory) Println(w io.Writer, addr string) {
	fmt.Fprintln(w, t.ToString(addr))
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
		for k, v := range t.q[i].m {
			fmt.Fprintf(w, "  -> %s: %v\n", k, v)
		}
	}
	fmt.Fprintln(w, "***********************************")
	return w.String()
}
