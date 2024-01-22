package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type Table struct {
	m       map[string]*obj.Object
	Protect bool
}

func (t *Table) Get(key string) (*obj.Object, bool) {
	v, ok := t.m[key]
	return v, ok
}

func (t *Table) Len() int {
	return len(t.m)
}

func (t *Table) Empty() bool {
	return t.Len() == 0
}

func (t *Table) Set(key string, val *obj.Object) {
	t.m[key] = val
}

func newTable(protect bool) *Table {
	return &Table{m: make(map[string]*obj.Object), Protect: protect}
}

func NewObjectTable() *TableStack {
	return &TableStack{q: []*Table{newTable(true)}}
}

type TableStack struct {
	q []*Table
}

func (t *TableStack) Raw() []*Table {
	return t.q
}

func (t *TableStack) Push(n *Table) {
	t.q = append(t.q, n)
}

func (t *TableStack) PushEmpty(protect bool) {
	t.q = append(t.q, newTable(protect))
}

func (t *TableStack) Pop() *Table {
	n := t.q[len(t.q)-1]
	t.q = t.q[:len(t.q)-1]
	return n
}

func (t *TableStack) Peek() *Table {
	if t.Empty() {
		return nil
	}
	return t.q[len(t.q)-1]
}

func (t *TableStack) HasKeyAtTop(key string) bool {
	if t.Empty() {
		return false
	}
	_, ok := t.q[len(t.q)-1].Get(key)
	return ok
}

func (t *TableStack) RemoveKeyAtTop(key string) {
	if t.Empty() {
		return
	}
	delete(t.q[len(t.q)-1].m, key)
}

func (t *TableStack) Len() int {
	return len(t.q)
}

func (t *TableStack) Empty() bool {
	return len(t.q) == 0
}

func (t *TableStack) Println() {
	if t.Empty() {
		fmt.Println("<EMPTY>")
		return
	}
	for i := len(t.q) - 1; i >= 0; i-- {
		fmt.Println("TABLE [", i, "]")
		for k, v := range t.q[i].m {
			printId(i)
			fmt.Printf("%s: %v\n", k, v)
		}
	}
}

func printId(id int) {
	for i := 0; i < id; i++ {
		fmt.Print("  ")
	}
	fmt.Print("->")
}

func getFromObjTable[T any](t *TableStack, key string) (T, bool) {
	v, ok := t.Get(key)
	if !ok {
		var t T
		return t, false
	}
	tV, ok := v.Val.(T)
	return tV, ok
}

// Get gets the value from the table stack
func (t *TableStack) Get(key string) (*obj.Object, bool) {
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
	//ts := []Table{t.q[len(t.q)-1], t.q[0]}
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
			return nil, false
		}
	}
	return nil, false
}

// Set overwrites the value if it exists in the table stack
// otherwise it sets the value at the top of the stack
func (t *TableStack) Set(key string, val any) {
	o := cons(val)
	if t.Empty() {
		v := newTable(false)
		v.Set(key, o)
		t.q = append(t.q, v)
		return
	}

	// FIXME: Could cause var not found
	//ts := []Table{t.q[len(t.q)-1], t.q[0]}
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
			break
		}
	}
	t.q[len(t.q)-1].Set(key, o)
}

// SetAtTop always sets the value at the top of the stack
func (t *TableStack) SetAtTop(key string, val any) {
	o := cons(val)
	if t.Empty() {
		t.PushEmpty(false)
	}
	t.q[len(t.q)-1].Set(key, o)
}

// GetAtTop always gets the value at the top of the stack
func (t *TableStack) GetAtTop(key string) (any, bool) {
	if t.Empty() {
		return nil, false
	}
	v, ok := t.q[len(t.q)-1].Get(key)
	if !ok {
		return nil, false
	}
	return v.Val, true
}
