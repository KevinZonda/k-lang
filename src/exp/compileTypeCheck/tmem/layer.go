package TMem

import (
	. "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/compileTypeCheck/common"
)

type Layer struct {
	m       map[string]T
	Protect bool
}

func (t *Layer) Get(key string) (T, bool) {
	v, ok := t.m[key]
	return v, ok
}

func (t *Layer) Len() int {
	return len(t.m)
}

func (t *Layer) Empty() bool {
	return t.Len() == 0
}

func (t *Layer) Set(key string, val T) {
	t.m[key] = val
}

func newLayer(protect bool) *Layer {
	return &Layer{m: make(map[string]T), Protect: protect}
}

func NewMemory() *Memory {
	return &Memory{q: []*Layer{newLayer(true)}}
}

func (t *Layer) Has(key string) bool {
	if t == nil {
		return false
	}
	_, ok := t.m[key]
	return ok
}

func (t *Layer) Remove(key string) {
	delete(t.m, key)
}
