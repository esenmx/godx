package set

import (
	"sync"
)

type node struct {
	val   interface{}
	left  *node
	right *node
}
type Comparator func(a interface{}, b interface{}) int
type TreeSet struct {
	root       node
	comparator Comparator
	mutex      sync.RWMutex
}

func New(compare Comparator, args ...interface{}) *TreeSet {
	return &TreeSet{comparator: compare}
}

//func (ts *TreeSet) Add(val interface{}) bool {
//	if ts.root.val == nil {
//
//	}
//}
