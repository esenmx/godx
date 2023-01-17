package mock

import "math/rand"

type Element struct {
	A int
	B int
}

func RandomElements(size int) []Element {
	array := make([]Element, size)
	for i := 0; i < len(array); i++ {
		array[i] = Element{A: rand.Int()}
	}
	return array
}

func OrderedIntArray(size int) []int {
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = i
	}
	return elements
}
