package obj

type StackImpl[T any] struct {
	q []map[string]T
}

func (t *StackImpl[T]) Raw() []map[string]T {
	return t.q
}

func (t *StackImpl[T]) Push(n map[string]T) {
	t.q = append(t.q, n)
}

func (t *StackImpl[T]) PushEmpty() {
	t.q = append(t.q, make(map[string]T))
}

func (t *StackImpl[T]) Pop() map[string]T {
	n := t.q[len(t.q)-1]
	t.q = t.q[:len(t.q)-1]
	return n
}

func (t *StackImpl[T]) Peek() map[string]T {
	return t.q[len(t.q)-1]
}

func (t *StackImpl[T]) Len() int {
	return len(t.q)
}

func (t *StackImpl[T]) Empty() bool {
	return len(t.q) == 0
}

func (t *StackImpl[T]) Get(key string) (T, bool) {
	for i := len(t.q) - 1; i >= 0; i-- {
		if v, ok := t.q[i][key]; ok {
			return v, true
		}
	}
	var r T
	return r, false
}

func (t *StackImpl[T]) Set(key string, val T) {
	if len(t.q) == 0 {
		v := make(map[string]T)
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
