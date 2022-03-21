package sets

import (
	"fmt"
	"github.com/softronaut/godx/pkg"
	"strings"
	"sync"
)

type HashSet struct {
	hash  map[interface{}]struct{}
	mutex sync.RWMutex
}

func NewHashSet(elements ...interface{}) *HashSet {
	hash := make(map[interface{}]struct{}, len(elements))
	for _, v := range elements {
		hash[v] = struct{}{}
	}
	return &HashSet{hash: hash}
}

func assertHashSetInterface() {
	var _ HashSetInterface = (*HashSet)(nil)
}

func (r *HashSet) Add(element interface{}) bool {
	r.mutex.Lock()
	defer func() {
		r.hash[element] = struct{}{}
		r.mutex.Unlock()
	}()

	_, ok := r.hash[element]
	return !ok
}

func (r *HashSet) AddAll(elements ...interface{}) {
	r.mutex.Lock()
	for _, v := range elements {
		r.hash[v] = struct{}{}
	}
	r.mutex.Unlock()
}

func (r *HashSet) Any(fn func(interface{}) bool) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for k := range r.hash {
		if fn(k) {
			return true
		}
	}
	return false
}

func (r *HashSet) Clear() {
	r.mutex.Lock()
	r.hash = make(map[interface{}]struct{})
	r.mutex.Unlock()
}

func (r *HashSet) Contains(element interface{}) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, ok := r.hash[element]
	return ok
}

func (r *HashSet) ContainsAll(elements ...interface{}) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, v := range elements {
		if _, ok := r.hash[v]; !ok {
			return false
		}
	}
	return true
}

func (r *HashSet) Difference(o *HashSet) *HashSet {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	o.mutex.RLock()
	defer o.mutex.RUnlock()

	diff := NewHashSet()
	for k := range r.hash {
		if _, ok := o.hash[k]; !ok {
			diff.hash[k] = struct{}{}
		}
	}
	return diff
}

func (r *HashSet) Every(fn func(interface{}) bool) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for k := range r.hash {
		if !fn(k) {
			return false
		}
	}
	return true
}

func (r *HashSet) ForEach(fn func(interface{})) {
	r.mutex.RLock()
	for k := range r.hash {
		fn(k)
	}
	r.mutex.RUnlock()
}

func (r *HashSet) Intersection(other *HashSet) *HashSet {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	other.mutex.RLock()
	defer other.mutex.RUnlock()

	intersection := NewHashSet()
	if len(r.hash) > len(other.hash) {
		for k := range other.hash {
			if _, ok := r.hash[k]; ok {
				intersection.hash[k] = struct{}{}
			}
		}
	} else {
		for k := range r.hash {
			if _, ok := other.hash[k]; ok {
				intersection.hash[k] = struct{}{}
			}
		}
	}
	return intersection
}

func (r *HashSet) IsEmpty() bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash) == 0
}

func (r *HashSet) IsNotEmpty() bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash) > 0
}

func (r *HashSet) Join(separator string) string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var builder strings.Builder
	switch len(r.hash) {
	case 0:
		return ""
	case 1:
		for k := range r.hash {
			builder.WriteString(fmt.Sprintf("%v", k))
		}
		return builder.String()
	default:
		for k := range r.hash {
			builder.WriteString(fmt.Sprintf("%v%s", k, separator))
		}
		str := builder.String()
		return str[:len(str)-len(separator)]
	}
}

func (r *HashSet) Map(fn func(interface{}) interface{}) []interface{} {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	array := make([]interface{}, len(r.hash))
	i := 0
	for k := range r.hash {
		array[i] = fn(k)
		i++
	}
	return array
}

func (r *HashSet) Remove(element interface{}) bool {
	r.mutex.Lock()
	defer delete(r.hash, element)
	defer r.mutex.Unlock()
	_, ok := r.hash[element]
	return ok
}

func (r *HashSet) RemoveAll(elements ...interface{}) {
	r.mutex.Lock()
	for _, v := range elements {
		delete(r.hash, v)
	}
	r.mutex.Unlock()
}

func (r *HashSet) RetainAll(other *HashSet) {
	r.mutex.Lock()
	other.mutex.RLock()
	for k := range r.hash {
		if _, ok := other.hash[k]; !ok {
			delete(r.hash, k)
		}
	}
	r.mutex.Unlock()
	other.mutex.RUnlock()
}

func (r *HashSet) Size() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash)
}

func (r *HashSet) ToArray() []interface{} {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	array := make([]interface{}, len(r.hash))
	i := 0
	for k := range r.hash {
		array[i] = k
		i++
	}
	return array
}

func (r *HashSet) Union(other *HashSet) *HashSet {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	other.mutex.RLock()
	defer other.mutex.RUnlock()

	union := NewHashSet()
	for k := range r.hash {
		union.hash[k] = struct{}{}
	}
	for k := range other.hash {
		union.hash[k] = struct{}{}
	}
	return union
}

func (r *HashSet) Where(fn func(interface{}) bool) *HashSet {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	where := NewHashSet()
	for k := range r.hash {
		if fn(k) {
			where.hash[k] = struct{}{}
		}
	}
	return where
}

func (r *HashSet) String() string {
	return fmt.Sprintf("HashSet{%s}", r.Join(", "))
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
