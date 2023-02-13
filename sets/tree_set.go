package sets

//import (
//	"github.com/esenmx/godx/pkg"
//	"sync"
//)
//
//type node struct {
//	left  *node
//	right *node
//}
//type TreeSet struct {
//	hash       map[interface{}]*node
//	root       *node
//	comparator pkg.Comparator
//	mutex      sync.RWMutex
//}
//
//func NewTreeSet(compare pkg.Comparator, args ...interface{}) *TreeSet {
//	set := &TreeSet{comparator: compare}
//
//	return set
//}
//
////func (ts *TreeSet) Add(arg interface{}) bool {
////	ts.mutex.Lock()
////	defer ts.mutex.Unlock()
////	if ts.root == nil {
////		ts.root = &node{}
////		ts.hash[arg] = ts.root
////		return true
////	}
////	if _, ok := ts.hash[arg]; ok {
////		return false
////	}
////
////}
