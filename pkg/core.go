package pkg

type Comparator[K any] func(a, b K) int
type Compute[V any] func(element any) (result V)
