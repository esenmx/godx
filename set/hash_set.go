package set

import (
	"fmt"
	"strings"
	"sync"
)

type void struct{}

type HashSet struct {
	hash  map[interface{}]void
	mutex sync.RWMutex
}

type HashSetInterface interface {
	Add(interface{}) bool
	AddAll(...interface{})
	Any(func(interface{}) bool) bool
	Clear()
	Contains(interface{}) bool
	ContainsAll(...interface{}) bool
	Difference(*HashSet) *HashSet
	Every(func(interface{}) bool) bool
	ForEach(func(interface{}))
	Intersection(*HashSet) *HashSet
	IsEmpty() bool
	Join(string) string
	Remove(interface{}) bool
	RemoveAll(...interface{})
	RetainAll(set *HashSet)
	Size() interface{}
	ToArray() []interface{}
	Union(*HashSet) *HashSet
	Where(func(interface{}) bool) *HashSet
}

func assertHashSetInterface() {
	var _ HashSetInterface = (*HashSet)(nil)
}

func NewHashSet(args ...interface{}) *HashSet {
	elements := make(map[interface{}]void, len(args))
	for _, v := range args {
		elements[v] = void{}
	}
	return &HashSet{hash: elements}
}

func (hs *HashSet) Add(arg interface{}) bool {
	hs.mutex.Lock()
	defer func() {
		hs.hash[arg] = void{}
		hs.mutex.Unlock()
	}()
	_, ok := hs.hash[arg]
	return !ok
}

func (hs *HashSet) AddAll(args ...interface{}) {
	hs.mutex.Lock()
	for _, v := range args {
		hs.hash[v] = void{}
	}
	hs.mutex.Unlock()
}

func (hs *HashSet) Any(fn func(interface{}) bool) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	for k := range hs.hash {
		if fn(k) {
			return true
		}
	}
	return false
}

func (hs *HashSet) Clear() {
	hs.mutex.Lock()
	hs.hash = make(map[interface{}]void)
	hs.mutex.Unlock()
}

func (hs *HashSet) Contains(arg interface{}) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	_, ok := hs.hash[arg]
	return ok
}

func (hs *HashSet) ContainsAll(args ...interface{}) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	for _, v := range args {
		if _, ok := hs.hash[v]; !ok {
			return false
		}
	}
	return true
}

func (hs *HashSet) Difference(o *HashSet) *HashSet {
	hs.mutex.RLock()
	o.mutex.RLock()
	defer hs.mutex.RUnlock()
	defer o.mutex.RUnlock()
	diff := NewHashSet()
	for k := range hs.hash {
		if _, ok := o.hash[k]; !ok {
			diff.hash[k] = void{}
		}
	}
	return diff
}

func (hs *HashSet) Every(fn func(interface{}) bool) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	for k := range hs.hash {
		if !fn(k) {
			return false
		}
	}
	return true
}

func (hs *HashSet) ForEach(fn func(interface{})) {
	hs.mutex.RLock()
	for k := range hs.hash {
		fn(k)
	}
	hs.mutex.RUnlock()
}

func (hs *HashSet) Intersection(other *HashSet) *HashSet {
	hs.mutex.RLock()
	other.mutex.RLock()
	defer hs.mutex.RUnlock()
	defer other.mutex.RUnlock()
	intersection := NewHashSet()
	if len(hs.hash) > len(other.hash) {
		for k := range other.hash {
			if _, ok := hs.hash[k]; ok {
				intersection.hash[k] = void{}
			}
		}
	} else {
		for k := range hs.hash {
			if _, ok := other.hash[k]; ok {
				intersection.hash[k] = void{}
			}
		}
	}
	return intersection
}

func (hs *HashSet) IsEmpty() bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash) == 0
}

func (hs *HashSet) Join(separator string) string {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	var builder strings.Builder
	switch len(hs.hash) {
	case 0:
		return ""
	case 1:
		for k := range hs.hash {
			builder.WriteString(fmt.Sprintf("%v", k))
		}
		return builder.String()
	default:
		for k := range hs.hash {
			builder.WriteString(fmt.Sprintf("%v%s", k, separator))
		}
		str := builder.String()
		return str[:len(str)-len(separator)]
	}
}

func (hs *HashSet) Remove(arg interface{}) bool {
	hs.mutex.Lock()
	defer delete(hs.hash, arg)
	defer hs.mutex.Unlock()
	_, ok := hs.hash[arg]
	return ok
}

func (hs *HashSet) RemoveAll(args ...interface{}) {
	hs.mutex.Lock()
	for _, v := range args {
		delete(hs.hash, v)
	}
	hs.mutex.Unlock()
}

func (hs *HashSet) RetainAll(other *HashSet) {
	hs.mutex.Lock()
	other.mutex.RLock()
	for k := range hs.hash {
		if _, ok := other.hash[k]; !ok {
			delete(hs.hash, k)
		}
	}
	hs.mutex.Unlock()
	other.mutex.RUnlock()
}

func (hs *HashSet) Size() interface{} {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash)
}

func (hs *HashSet) ToArray() []interface{} {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	array := make([]interface{}, len(hs.hash))
	i := 0
	for k := range hs.hash {
		array[i] = k
		i++
	}
	return array
}

func (hs *HashSet) Union(other *HashSet) *HashSet {
	hs.mutex.RLock()
	other.mutex.RLock()
	defer hs.mutex.RUnlock()
	defer other.mutex.RUnlock()
	union := NewHashSet()
	for k := range hs.hash {
		union.hash[k] = void{}
	}
	for k := range other.hash {
		union.hash[k] = void{}
	}
	return union
}

func (hs *HashSet) Where(fn func(interface{}) bool) *HashSet {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	where := NewHashSet()
	for k := range hs.hash {
		if fn(k) {
			where.hash[k] = void{}
		}
	}
	return where
}

func (hs *HashSet) String() string {
	return fmt.Sprintf("HashSet{%s}", hs.Join(", "))
}
