package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type Table map[string]*obj.Object

func NewObjectTable() *TableStack {
	return &TableStack{q: []Table{make(Table)}}
}

type TableStack struct {
	q []Table
}

func (t *TableStack) Raw() []Table {
	return t.q
}

func (t *TableStack) Push(n Table) {
	t.q = append(t.q, n)
}

func (t *TableStack) PushEmpty() {
	t.q = append(t.q, make(Table))
}

func (t *TableStack) Pop() Table {
	n := t.q[len(t.q)-1]
	t.q = t.q[:len(t.q)-1]
	return n
}

func (t *TableStack) Peek() Table {
	if t.Empty() {
		return nil
	}
	return t.q[len(t.q)-1]
}

func (t *TableStack) HasKeyAtTop(key string) bool {
	if t.Empty() {
		return false
	}
	_, ok := t.q[len(t.q)-1][key]
	return ok
}

func (t *TableStack) RemoveKeyAtTop(key string) {
	if t.Empty() {
		return
	}
	delete(t.q[len(t.q)-1], key)
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
		for k, v := range t.q[i] {
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
		v, ok := t.q[0][key]
		return v, ok
	}
	//ts := []Table{t.q[len(t.q)-1], t.q[0]}
	//for _, _t := range ts {
	//	if v, ok := _t[key]; ok {
	//		return v, true
	//	}
	//}
	for i := len(t.q) - 1; i >= 0; i-- {
		if v, ok := t.q[i][key]; ok {
			return v, true
		}
	}
	return nil, false
}

func (t *TableStack) Set(key string, val any) {
	o := cons(val)
	if len(t.q) == 0 {
		v := make(Table)
		v[key] = o
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

	for i := len(t.q) - 1; i >= 0; i-- {
		if _, ok := t.q[i][key]; ok {
			t.q[i][key] = o
			return
		}
	}
	t.q[len(t.q)-1][key] = o
}

func (t *TableStack) SetAtTop(key string, val any) {
	o := cons(val)
	if t.Empty() {
		t.PushEmpty()
	}
	t.q[len(t.q)-1][key] = o
}

func (t *TableStack) GetAtTop(key string) (any, bool) {
	if t.Empty() {
		return nil, false
	}
	v, ok := t.q[len(t.q)-1][key]
	if !ok {
		return nil, false
	}
	return v.Val, true
}
