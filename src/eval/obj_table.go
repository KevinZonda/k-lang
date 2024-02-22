package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type MemLayer struct {
	m       map[string]*obj.Object
	Protect bool
}

func (t *MemLayer) Get(key string) (*obj.Object, bool) {
	v, ok := t.m[key]
	return v, ok
}

func (t *MemLayer) Len() int {
	return len(t.m)
}

func (t *MemLayer) Empty() bool {
	return t.Len() == 0
}

func (t *MemLayer) Set(key string, val *obj.Object) {
	t.m[key] = val
}

func (t *MemLayer) SetValue(key string, val any) {
	t.m[key] = cons(val)
}

func newTable(protect bool) *MemLayer {
	return &MemLayer{m: make(map[string]*obj.Object), Protect: protect}
}

func NewMemory() *Memory {
	return &Memory{q: []*MemLayer{newTable(true)}}
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
	//ts := []MemLayer{t.q[len(t.q)-1], t.q[0]}
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
	o := cons(val)
	if t.Empty() {
		v := newTable(false)
		v.Set(key, o)
		t.q = append(t.q, v)
		return
	}

	// FIXME: Could cause var not found
	//ts := []MemLayer{t.q[len(t.q)-1], t.q[0]}
	//for _, _t := range ts {
	//	if _, ok := _t[key]; ok {
	//		_t[key] = val
	//		return
	//	}
	//}

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

func (t *Memory) Top() *MemLayer {
	if t.Empty() {
		t.PushEmpty(true)
	}
	return t.q[len(t.q)-1]
}

func (t *MemLayer) Has(key string) bool {
	if t == nil {
		return false
	}
	_, ok := t.m[key]
	return ok
}

func (t *MemLayer) GetValue(key string) (any, bool) {
	o, ok := t.Get(key)
	if ok && o != nil {
		return o.Val, ok
	}
	return nil, ok
}

func (t *MemLayer) Remove(key string) {
	delete(t.m, key)
}
