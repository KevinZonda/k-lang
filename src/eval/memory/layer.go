package memory

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Layer struct {
	m       *orderedmap.OrderedMap[string, *obj.Object]
	Protect bool
}

func (t *Layer) Get(key string) (*obj.Object, bool) {
	v, ok := t.m.Get(key)
	return v, ok
}

func (t *Layer) Len() int {
	return t.m.Len()
}

func (t *Layer) Empty() bool {
	return t.Len() == 0
}

func (t *Layer) Set(key string, val *obj.Object) {
	t.m.Set(key, val)
}

func (t *Layer) SetValue(key string, val any) {
	t.m.Set(key, obj.Construct(val))
}

func NewLayer(protect bool) *Layer {
	return &Layer{m: orderedmap.New[string, *obj.Object](), Protect: protect}
}

func (t *Layer) Copy() *Layer {
	n := NewLayer(t.Protect)
	for pair := t.m.Oldest(); pair != nil; pair = pair.Next() {
		n.Set(pair.Key, pair.Value)
	}
	return n
}

func NewMemory() *Memory {
	return &Memory{q: []*Layer{NewLayer(true)}}
}

func (t *Layer) WithObj(key string, val *obj.Object) *Layer {
	t.Set(key, val)
	return t
}

func (t *Layer) Has(key string) bool {
	if t == nil {
		return false
	}
	_, ok := t.m.Get(key)
	return ok
}

func (t *Layer) HasAny(key ...string) bool {
	if t == nil {
		return false
	}
	for _, k := range key {
		if t.Has(k) {
			return true
		}
	}
	return false

}

func (t *Layer) GetValue(key string) (any, bool) {
	o, ok := t.Get(key)
	if ok && o != nil {
		return o.Value(), ok
	}
	return nil, ok
}

func (t *Layer) Remove(key string) {
	t.m.Delete(key)
}

func (t *Layer) Raw() *orderedmap.OrderedMap[string, *obj.Object] {
	return t.m
}

func (t *Layer) NoReserved() {
	if t == nil {
		return
	}
	t.Protect = false
	t.m.Delete(reserved.Break)
	t.m.Delete(reserved.Continue)
	t.m.Delete(reserved.Return)
	t.m.Delete(reserved.Loop)
}
