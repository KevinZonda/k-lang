package memory

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
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

func (t *Memory) PushMem(m *Memory) {
	if m == nil || len(m.q) == 0 {
		return
	}
	t.q = append(t.q, m.q...)
}

func (t *Memory) PushEmpty(protect bool) *Layer {
	tbl := NewLayer(protect)
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
	return t == nil || len(t.q) == 0
}

func (t *Memory) Println(w io.Writer, addr string) {
	fmt.Fprintln(w, t.ToString(addr))
}

// Get gets the value from the table stack
func (t *Memory) Get(key string) (*obj.Object, bool) {
	// Get only valid on 2 tables
	// Top
	//  |
	//  |
	// Bottom (oldest)
	// t.Println()
	if t.Empty() {
		return nil, false
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
			return nil, false
		}
	}
	return nil, false
}

// Set overwrites the value if it exists in the table stack
// otherwise it sets the value at the top of the stack
func (t *Memory) Set(key string, val any) {
	o := obj.Construct(val)
	if t.Empty() {
		t.q = append(t.q, NewLayer(true).WithObj(key, o))
		return
	}

	for i := t.Len() - 1; i >= 0; i-- {
		stack := t.q[i]
		if _, ok := stack.Get(key); ok {
			t.q[i].Set(key, o)
			return
		}
		if stack.Protect {
			if t.Len() >= 1 {
				stack = t.q[0]
				if _, ok := stack.Get(key); ok {
					stack.Set(key, o)
					return
				}
			}
			break
		}
	}
	t.q[len(t.q)-1].Set(key, o)
}

func (t *Memory) Top() *Layer {
	if t.Empty() {
		return nil
		// t.PushEmpty(true)
	}
	return t.q[len(t.q)-1]
}

func (t *Memory) Squeeze() *Layer {
	if t.Empty() {
		return nil
	}
	if t.Len() == 1 {
		return t.q[0].Copy()
	}
	l := NewLayer(false)
	for i := t.Len() - 1; i > 0; i-- {
		for k, v := range t.q[i].m {
			l.Set(k, v)
		}
	}
	return l
}
