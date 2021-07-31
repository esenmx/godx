package sets

import (
	"fmt"
	"github.com/softronaut/godx/pkg"
	"strings"
	"sync"
)

type HashSet struct {
	hash  map[interface{}]void
	mutex sync.RWMutex
}

type HashSetInterface interface {
	pkg.Interface
	Add(interface{}) bool
	AddAll(other ...interface{})
	Any(predicate func(interface{}) bool) bool
	Clear()
	Contains(element interface{}) bool
	ContainsAll(...interface{}) bool
	Difference(*HashSet) *HashSet
	Every(func(interface{}) bool) bool
	ForEach(func(interface{}))
	Intersection(*HashSet) *HashSet
	Join(string) string
	Map(func(interface{}) interface{}) []interface{}
	Remove(interface{}) bool
	RemoveAll(...interface{})
	RetainAll(set *HashSet)
	Size() int
	ToArray() []interface{}
	Union(*HashSet) *HashSet
	Where(func(interface{}) bool) *HashSet
}

func NewHashSet(elements ...interface{}) *HashSet {
	hash := make(map[interface{}]void, len(elements))
	for _, v := range elements {
		hash[v] = void{}
	}
	return &HashSet{hash: hash}
}

func assertHashSetInterface() {
	var _ HashSetInterface = (*HashSet)(nil)
}

func (hs *HashSet) Add(element interface{}) bool {
	hs.mutex.Lock()
	defer func() {
		hs.hash[element] = void{}
		hs.mutex.Unlock()
	}()
	_, ok := hs.hash[element]
	return !ok
}

func (hs *HashSet) AddAll(elements ...interface{}) {
	hs.mutex.Lock()
	for _, v := range elements {
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

func (hs *HashSet) Contains(element interface{}) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	_, ok := hs.hash[element]
	return ok
}

func (hs *HashSet) ContainsAll(elements ...interface{}) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	for _, v := range elements {
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

func (hs *HashSet) IsNotEmpty() bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash) > 0
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

func (hs *HashSet) Map(fn func(interface{}) interface{}) []interface{} {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	array := make([]interface{}, len(hs.hash))
	i := 0
	for k := range hs.hash {
		array[i] = fn(k)
		i++
	}
	return array
}

func (hs *HashSet) Remove(element interface{}) bool {
	hs.mutex.Lock()
	defer delete(hs.hash, element)
	defer hs.mutex.Unlock()
	_, ok := hs.hash[element]
	return ok
}

func (hs *HashSet) RemoveAll(elements ...interface{}) {
	hs.mutex.Lock()
	for _, v := range elements {
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

func (hs *HashSet) Size() int {
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
