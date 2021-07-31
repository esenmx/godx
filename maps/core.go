package maps

import (
	"github.com/softronaut/godx/pkg"
)

type Entry struct {
	key   interface{}
	value interface{}
}

type MapInterface interface {
	pkg.Interface
	AddAll(other *HashMap)
	AddEntries(entries []Entry)
	ContainsKey(key interface{}) bool
	ContainsValue(value interface{}) bool
	Entries() []Entry
	ForEach(do func(key interface{}, value interface{}))
	Keys() []interface{}
	Put(key interface{}, value interface{}) interface{}
	PutIfAbsent(key interface{}, put func() interface{}) interface{}
	Remove(key interface{}) (exists bool)
	RemoveWhere(predicate func(key interface{}, value interface{}) bool)
	Values() (values []interface{})
}
