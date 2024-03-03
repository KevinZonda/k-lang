package TMem

import (
	"fmt"
	. "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/compileTypeCheck/common"
	"io"
)

type Memory struct {
	q []*Layer
}

func (t *Memory) Bottom() *Layer {
	if t.Empty() {
		return nil
	}
	return t.q[0]
}

func (t *Memory) Raw() []*Layer {
	return t.q
}

func (t *Memory) Push(n *Layer) {
	t.q = append(t.q, n)
}

func (t *Memory) PushEmpty(protect bool) *Layer {
	tbl := newLayer(protect)
	t.Push(tbl)
	return tbl
}

func (t *Memory) Pop() *Layer {
	n := t.q[len(t.q)-1]
	t.q = t.q[:len(t.q)-1]
	return n
}

func (t *Memory) Len() int {
	return len(t.q)
}

func (t *Memory) Empty() bool {
	return len(t.q) == 0
}

func (t *Memory) Println(w io.Writer) {
	fmt.Fprintln(w)
}

// Get gets the value from the table stack
func (t *Memory) Get(key string) (T, bool) {
	// Get only valid on 2 tables
	// Top
	//  |
	//  |
	// Bottom (oldest)
	// t.Println()
	if t.Empty() {
		return T{}, false
	}
	if t.Len() == 1 {
		v, ok := t.q[0].Get(key)
		return v, ok
	}
	//ts := []Layer{t.q[len(t.q)-1], t.q[0]}
	//for _, _t := range ts {
	//	if v, ok := _t[key]; ok {
	//		return v, true
	//	}
	//}
	for i := t.Len() - 1; i >= 0; i-- {
		stack := t.q[i]
		if v, ok := stack.Get(key); ok {
			return v, true
		}
		if stack.Protect {
			if t.Len() >= 1 {
				stack = t.q[0]
				if v, ok := stack.Get(key); ok {
					return v, true
				}
			}
			return T{}, false
		}
	}
	return T{}, false
}

// Set overwrites the value if it exists in the table stack
// otherwise it sets the value at the top of the stack
func (t *Memory) Set(key string, _t T) {
	if t.Empty() {
		l := newLayer(true)
		l.Set(key, _t)
		t.q = append(t.q, l)
		return
	}

	for i := t.Len() - 1; i >= 0; i-- {
		stack := t.q[i]
		if _, ok := stack.Get(key); ok {
			t.q[i].Set(key, _t)
			return
		}
		if stack.Protect {
			if t.Len() >= 1 {
				stack = t.q[0]
				if _, ok := stack.Get(key); ok {
					stack.Set(key, _t)
					return
				}
			}
			break
		}
	}
	t.q[len(t.q)-1].Set(key, _t)
}

func (t *Memory) Top() *Layer {
	if t.Empty() {
		return nil
		// t.PushEmpty(true)
	}
	return t.q[len(t.q)-1]
}
