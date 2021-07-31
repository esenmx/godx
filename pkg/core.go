package pkg

type Comparator func(a, b interface{}) int
type Compute func(element interface{}) (result interface{})

type Interface interface {
	Clear()
	IsEmpty() bool
	IsNotEmpty() bool
	// Map()
	Size() int
}
