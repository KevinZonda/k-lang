package obj

type Table map[string]any

type TableMap map[string]Table

type TableStack struct {
	q []Table
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

func (t *TableStack) Get(key string) (any, bool) {
	for i := len(t.q) - 1; i >= 0; i-- {
		if v, ok := t.q[i][key]; ok {
			return v, true
		}
	}
	return nil, false
}

func (t *TableStack) Set(key string, val any) {
	if len(t.q) == 0 {
		v := make(Table)
		v[key] = val
		t.q = append(t.q, v)
	} else {
		for i := len(t.q) - 1; i >= 0; i-- {
			if _, ok := t.q[i][key]; ok {
				t.q[i][key] = val
				return
			}
		}
		t.q[len(t.q)-1][key] = val
	}
}

func (t *TableStack) SetAtTop(key string, val any) {
	if t.Empty() {
		t.PushEmpty()
	}
	t.q[len(t.q)-1][key] = val
}

func (t *TableStack) GetAtTop(key string) (any, bool) {
	if t.Empty() {
		return nil, false
	}
	v, ok := t.q[len(t.q)-1][key]
	return v, ok
}
