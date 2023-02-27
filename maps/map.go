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

func (r *Map[K, V]) Get(key K) (V, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	v, ok := r.hash[key]
	return v, ok
}

func (r *Map[K, V]) Set(key K, value V) {
	r.mutex.Lock()
	r.hash[key] = value
	r.mutex.Unlock()
}

func (r *Map[K, V]) IsEmpty() bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash) == 0
}

func (r *Map[K, V]) IsNotEmpty() bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash) > 0
}

func (r *Map[K, V]) AddAll(other *Map[K, V]) {
	r.mutex.Lock()
	other.mutex.RLock()
	for k, v := range other.hash {
		r.hash[k] = v
	}
	r.mutex.Unlock()
	other.mutex.RUnlock()
}

func (r *Map[K, V]) AddEntries(entries []Entry[K, V]) {
	r.mutex.Lock()
	for _, v := range entries {
		r.hash[v.key] = v.value
	}
	r.mutex.Unlock()
}

func (r *Map[K, V]) Clear() {
	r.mutex.Lock()
	r.hash = make(map[K]V)
	r.mutex.Unlock()
}

func (r *Map[K, V]) ContainsKey(key K) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	_, ok := r.hash[key]
	return ok
}

func (r *Map[K, V]) ContainsValue(value V) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	var ok bool
	for _, v := range r.hash {
		if v == value {
			ok = true
			break
		}
	}
	return ok

}

func (r *Map[K, V]) Entries() []Entry[K, V] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	entries := make([]Entry[K, V], len(r.hash))
	i := 0
	for k, v := range r.hash {
		entries[i] = Entry[K, V]{key: k, value: v}
		i++
	}
	return entries
}

func (r *Map[K, V]) ForEach(do func(key K, value V)) {
	r.mutex.RLock()
	for k, v := range r.hash {
		do(k, v)
	}
	r.mutex.RUnlock()
}

func (r *Map[K, V]) Keys() []K {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	keys := make([]K, len(r.hash))
	i := 0
	for k := range r.hash {
		keys[i] = k
		i++
	}
	return keys
}

func (r *Map[K, V]) Put(key K, value V) V {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	old := r.hash[key]
	r.hash[key] = value
	return old
}

func (r *Map[K, V]) PutIfAbsent(key K, put func() V) *V {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if v, ok := r.hash[key]; ok {
		return &v
	}
	v := put()
	r.hash[key] = v
	return nil
}

func (r *Map[K, V]) Remove(key K) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	_, ok := r.hash[key]
	delete(r.hash, key)
	return ok
}

func (r *Map[K, V]) RemoveWhere(predicate func(key K, value V) bool) {
	r.mutex.Lock()
	for k, v := range r.hash {
		if predicate(k, v) {
			delete(r.hash, k)
		}
	}
	r.mutex.Unlock()
}

func (r *Map[K, V]) Size() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return len(r.hash)
}

func (r *Map[K, V]) Values() []V {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	values := make([]V, len(r.hash))
	i := 0
	for _, v := range r.hash {
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
	Get(key K) (V, bool)
	Set(key K, value V)
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
