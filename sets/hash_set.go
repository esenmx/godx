package sets

import (
	"fmt"
	"strings"
	"sync"
)

type void struct{}

type Set[T comparable] struct {
	hash  map[T]void
	mutex sync.RWMutex
}

func NewSet[T comparable](elements ...T) *Set[T] {
	hash := make(map[T]void, len(elements))
	for _, v := range elements {
		hash[v] = void{}
	}
	return &Set[T]{hash: hash}
}

func (r *Set[T]) Add(element T) bool {
	r.mutex.Lock()
	defer func() {
		r.hash[element] = void{}
		r.mutex.Unlock()
	}()

	_, ok := r.hash[element]
	return !ok
}

func (r *Set[T]) AddAll(elements ...T) {
	r.mutex.Lock()
	for _, v := range elements {
		r.hash[v] = void{}
	}
	r.mutex.Unlock()
}

func (r *Set[T]) Any(fn func(T) bool) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for k := range r.hash {
		if fn(k) {
			return true
		}
	}
	return false
}

func (r *Set[T]) Array() []T {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	array := make([]T, len(r.hash))
	i := 0
	for k := range r.hash {
		array[i] = k
		i++
	}
	return array
}

func (r *Set[T]) Clear() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.hash = make(map[T]void)
}

func (r *Set[T]) Contains(element T) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_, ok := r.hash[element]
	return ok
}

func (r *Set[T]) ContainsAll(elements ...T) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, v := range elements {
		if _, ok := r.hash[v]; !ok {
			return false
		}
	}
	return true
}

func (r *Set[T]) Difference(o *Set[T]) *Set[T] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	o.mutex.RLock()
	defer o.mutex.RUnlock()

	diff := NewSet[T]()
	for k := range r.hash {
		if _, ok := o.hash[k]; !ok {
			diff.hash[k] = void{}
		}
	}
	return diff
}

func (r *Set[T]) Every(fn func(T) bool) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for k := range r.hash {
		if !fn(k) {
			return false
		}
	}
	return true
}

func (r *Set[T]) ForEach(fn func(T)) {
	r.mutex.RLock()
	for k := range r.hash {
		fn(k)
	}
	r.mutex.RUnlock()
}

func (r *Set[T]) Intersection(other *Set[T]) *Set[T] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	other.mutex.RLock()
	defer other.mutex.RUnlock()

	intersection := NewSet[T]()
	if len(r.hash) > len(other.hash) {
		for k := range other.hash {
			if _, ok := r.hash[k]; ok {
				intersection.hash[k] = void{}
			}
		}
	} else {
		for k := range r.hash {
			if _, ok := other.hash[k]; ok {
				intersection.hash[k] = void{}
			}
		}
	}
	return intersection
}

func (r *Set[T]) IsEmpty() bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash) == 0
}

func (r *Set[T]) IsNotEmpty() bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash) > 0
}

func (r *Set[T]) Join(separator string) string {
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

//func (r *Set[T]) Map[M any](fn func(T) M) []M {
//	r.mutex.RLock()
//	defer r.mutex.RUnlock()
//	array := make([]M, len(r.hash))
//	i := 0
//	for k := range r.hash {
//		array[i] = fn(k)
//		i++
//	}
//	return array
//}

func (r *Set[T]) Remove(element T) bool {
	r.mutex.Lock()
	defer func() { delete(r.hash, element) }()
	defer r.mutex.Unlock()
	_, ok := r.hash[element]
	return ok
}

func (r *Set[T]) RemoveAll(elements ...T) {
	r.mutex.Lock()
	for _, v := range elements {
		delete(r.hash, v)
	}
	r.mutex.Unlock()
}

func (r *Set[T]) RemoveWhere(test func(T) bool) {
	r.mutex.Lock()
	for k, _ := range r.hash {
		if test(k) {
			delete(r.hash, k)
		}
	}
	r.mutex.Unlock()
}

func (r *Set[T]) RetainAll(other *Set[T]) {
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

func (r *Set[T]) Size() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash)
}

func (r *Set[T]) ToArray() []T {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	array := make([]T, len(r.hash))
	i := 0
	for k := range r.hash {
		array[i] = k
		i++
	}
	return array
}

func (r *Set[T]) Union(other *Set[T]) *Set[T] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	other.mutex.RLock()
	defer other.mutex.RUnlock()

	union := NewSet[T]()
	for k := range r.hash {
		union.hash[k] = void{}
	}
	for k := range other.hash {
		union.hash[k] = void{}
	}
	return union
}

func (r *Set[T]) Where(fn func(T) bool) *Set[T] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	where := NewSet[T]()
	for k := range r.hash {
		if fn(k) {
			where.hash[k] = void{}
		}
	}
	return where
}

func (r *Set[T]) String() string {
	return fmt.Sprintf("Set{%s}", r.Join(", "))
}

type HashSetInterface[T comparable] interface {
	Array() []T
	Add(T) bool
	AddAll(other ...T)
	Any(predicate func(T) bool) bool
	Clear()
	Contains(element T) bool
	ContainsAll(...T) bool
	Difference(*Set[T]) *Set[T]
	Every(func(T) bool) bool
	ForEach(func(T))
	Intersection(*Set[T]) *Set[T]
	IsEmpty() bool
	IsNotEmpty() bool
	Join(string) string
	//Map(func(T) T) []T
	Remove(T) bool
	RemoveAll(...T)
	RetainAll(set *Set[T])
	Size() int
	ToArray() []T
	Union(*Set[T]) *Set[T]
	Where(func(T) bool) *Set[T]
}

func assertHashSetInterface[T comparable]() {
	var _ HashSetInterface[T] = (*Set[T])(nil)
}
