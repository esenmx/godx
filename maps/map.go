package maps

import (
	"sync"
)

type Entry[K comparable, V any] struct {
	key   K
	value V
}

func (e *Entry[K, V]) Key() K   { return e.key }
func (e *Entry[K, V]) Value() V { return e.value }

type Map[K comparable, V comparable] struct {
	hash  map[K]V
	mutex sync.RWMutex
}

func NewMap[K comparable, V comparable]() *Map[K, V] {
	return &Map[K, V]{hash: make(map[K]V)}
}

func NewMapFromEntries[K comparable, V comparable](entries ...Entry[K, V]) *Map[K, V] {
	hash := make(map[K]V, len(entries))
	for _, v := range entries {
		hash[v.key] = v.value
	}
	return &Map[K, V]{hash: hash}
}

//func NewMapFromArray[K comparable, V comparable](array []V, key pkg.Compute[K], value pkg.Compute[V]) *Map[K, V] {
//	hash := make(map[K]V, len(array))
//	for _, v := range array {
//		hash[key(v)] = value(v)
//	}
//	return &Map[K, V]{hash: hash}
//}

func (hs *Map[K, V]) IsEmpty() bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash) == 0
}

func (hs *Map[K, V]) IsNotEmpty() bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash) > 0
}

func (hs *Map[K, V]) AddAll(other *Map[K, V]) {
	hs.mutex.Lock()
	other.mutex.RLock()
	for k, v := range other.hash {
		hs.hash[k] = v
	}
	hs.mutex.Unlock()
	other.mutex.RUnlock()
}

func (hs *Map[K, V]) AddEntries(entries []Entry[K, V]) {
	hs.mutex.Lock()
	for _, v := range entries {
		hs.hash[v.key] = v.value
	}
	hs.mutex.Unlock()
}

func (hs *Map[K, V]) Clear() {
	hs.mutex.Lock()
	hs.hash = make(map[K]V)
	hs.mutex.Unlock()
}

func (hs *Map[K, V]) ContainsKey(key K) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	_, ok := hs.hash[key]
	return ok
}

func (hs *Map[K, V]) ContainsValue(value V) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	var ok bool
	for _, v := range hs.hash {
		if v == value {
			ok = true
			break
		}
	}
	return ok

}

func (hs *Map[K, V]) Entries() []Entry[K, V] {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	entries := make([]Entry[K, V], len(hs.hash))
	i := 0
	for k, v := range hs.hash {
		entries[i] = Entry[K, V]{key: k, value: v}
		i++
	}
	return entries
}

func (hs *Map[K, V]) ForEach(do func(key K, value V)) {
	hs.mutex.RLock()
	for k, v := range hs.hash {
		do(k, v)
	}
	hs.mutex.RUnlock()
}

func (hs *Map[K, V]) Keys() []K {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	keys := make([]K, len(hs.hash))
	i := 0
	for k := range hs.hash {
		keys[i] = k
		i++
	}
	return keys
}

func (hs *Map[K, V]) Put(key K, value V) V {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	old := hs.hash[key]
	hs.hash[key] = value
	return old
}

func (hs *Map[K, V]) PutIfAbsent(key K, put func() V) *V {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	if v, ok := hs.hash[key]; ok {
		return &v
	}
	v := put()
	hs.hash[key] = v
	return nil
}

func (hs *Map[K, V]) Remove(key K) bool {
	hs.mutex.Lock()
	defer hs.mutex.Unlock()
	_, ok := hs.hash[key]
	delete(hs.hash, key)
	return ok
}

func (hs *Map[K, V]) RemoveWhere(predicate func(key K, value V) bool) {
	hs.mutex.Lock()
	for k, v := range hs.hash {
		if predicate(k, v) {
			delete(hs.hash, k)
		}
	}
	hs.mutex.Unlock()
}

func (hs *Map[K, V]) Size() int {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash)
}

func (hs *Map[K, V]) Values() []V {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	values := make([]V, len(hs.hash))
	i := 0
	for _, v := range hs.hash {
		values[i] = v
		i++
	}
	return values
}

type MapInterface[K comparable, V comparable] interface {
	AddAll(other *Map[K, V])
	AddEntries(entries []Entry[K, V])
	Clear()
	ContainsKey(key K) bool
	ContainsValue(value V) bool
	Entries() []Entry[K, V]
	ForEach(do func(key K, value V))
	IsEmpty() bool
	IsNotEmpty() bool
	Keys() []K
	Put(key K, value V) V
	PutIfAbsent(key K, put func() V) *V
	Remove(key K) (exists bool)
	RemoveWhere(predicate func(key K, value V) bool)
	Size() int
	Values() (values []V)
}

func assertHashMapInterface[K comparable, V comparable]() {
	var _ MapInterface[K, V] = (*Map[K, V])(nil)
}
