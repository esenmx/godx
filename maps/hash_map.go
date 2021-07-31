package maps

import (
	"github.com/softronaut/godx/pkg"
	"sync"
)

type HashMapInterface interface {
	MapInterface
}

func NewHashMap() *HashMap {
	return &HashMap{hash: make(map[interface{}]interface{})}
}

func NewHashMapFromEntries(entries ...Entry) *HashMap {
	hash := make(map[interface{}]interface{}, len(entries))
	for _, v := range entries {
		hash[v.key] = v.value
	}
	return &HashMap{hash: hash}
}

func NewHashMapFromArray(array []interface{}, key pkg.Compute, value pkg.Compute) *HashMap {
	hash := make(map[interface{}]interface{}, len(array))
	for _, v := range array {
		hash[key(v)] = value(v)
	}
	return &HashMap{hash: hash}
}

func assertHashMapInterface() {
	var _ HashMapInterface = (*HashMap)(nil)
}

type HashMap struct {
	hash  map[interface{}]interface{}
	mutex sync.RWMutex
}

func (hs *HashMap) IsEmpty() bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash) == 0
}

func (hs *HashMap) IsNotEmpty() bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash) > 0
}

func (hs *HashMap) AddAll(other *HashMap) {
	hs.mutex.Lock()
	other.mutex.RLock()
	for k, v := range other.hash {
		hs.hash[k] = v
	}
	hs.mutex.Unlock()
	other.mutex.RUnlock()
}

func (hs *HashMap) AddEntries(entries []Entry) {
	hs.mutex.Lock()
	for _, v := range entries {
		hs.hash[v.key] = v.value
	}
	hs.mutex.Unlock()
}

func (hs *HashMap) Clear() {
	hs.mutex.Lock()
	hs.hash = make(map[interface{}]interface{})
	hs.mutex.Unlock()
}

func (hs *HashMap) ContainsKey(key interface{}) bool {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	_, ok := hs.hash[key]
	return ok
}

func (hs *HashMap) ContainsValue(value interface{}) bool {
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

func (hs *HashMap) Entries() []Entry {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	entries := make([]Entry, len(hs.hash))
	i := 0
	for k, v := range hs.hash {
		entries[i] = Entry{key: k, value: v}
		i++
	}
	return entries
}

func (hs *HashMap) ForEach(do func(key interface{}, value interface{})) {
	hs.mutex.RLock()
	for k, v := range hs.hash {
		do(k, v)
	}
	hs.mutex.RUnlock()
}

func (hs *HashMap) Keys() []interface{} {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	keys := make([]interface{}, len(hs.hash))
	i := 0
	for k := range hs.hash {
		keys[i] = k
		i++
	}
	return keys
}

func (hs *HashMap) Put(key interface{}, value interface{}) interface{} {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	old := hs.hash[key]
	hs.hash[key] = value
	return old
}

func (hs *HashMap) PutIfAbsent(key interface{}, put func() interface{}) interface{} {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	if v, ok := hs.hash[key]; ok {
		return v
	}
	hs.hash[key] = put()
	return nil
}

func (hs *HashMap) Remove(key interface{}) bool {
	hs.mutex.Lock()
	defer hs.mutex.Unlock()
	_, ok := hs.hash[key]
	delete(hs.hash, key)
	return ok
}

func (hs *HashMap) RemoveWhere(predicate func(key interface{}, value interface{}) bool) {
	hs.mutex.Lock()
	for k, v := range hs.hash {
		if predicate(k, v) {
			delete(hs.hash, k)
		}
	}
	hs.mutex.Unlock()
}

func (hs *HashMap) Size() int {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	return len(hs.hash)
}

func (hs *HashMap) Values() []interface{} {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()
	values := make([]interface{}, len(hs.hash))
	i := 0
	for _, v := range hs.hash {
		values[i] = v
		i++
	}
	return values
}
