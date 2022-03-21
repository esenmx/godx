package maps

type Entry[K comparable, V any] struct {
	key   K
	value V
}
