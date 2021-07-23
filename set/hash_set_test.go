package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMethods(t *testing.T) {
	t.Run("NewHashSet", func(t *testing.T) {
		s := NewHashSet()
		assert.True(t, s.IsEmpty())
		s = NewHashSet(nil, true, 2, "3")
		assert.False(t, s.IsEmpty())
		assert.Equal(t, 4, s.Size())
	})
	t.Run("Clear", func(t *testing.T) {
		s := NewHashSet(1, 2, 3)
		s.Clear()
		assert.True(t, s.IsEmpty())
	})
	t.Run("Add/Remove/Contains", func(t *testing.T) {
		s := NewHashSet()
		assert.True(t, s.Add("foo"))
		assert.False(t, s.Add("foo"))
		assert.True(t, s.Contains("foo"))
		assert.True(t, s.Add(42))
		assert.True(t, s.Contains(42))
		assert.Equal(t, 2, s.Size())
		assert.True(t, s.Remove("foo"))
		assert.False(t, s.Remove("foo"))
		assert.Equal(t, 1, s.Size())
		assert.True(t, s.Contains(42))
	})
	t.Run("AddAll/RemoveAll/ContainsAll/RetainAll", func(t *testing.T) {
		s := NewHashSet()
		s.AddAll(42, "foo")
		assert.True(t, s.ContainsAll(42, "foo"))
		s.AddAll(42, "foo", "bar")
		assert.Equal(t, 3, s.Size())
		s.RemoveAll(1, 2, 3)
		assert.Equal(t, 3, s.Size())
		s.RemoveAll(42, "bar", "baz")
		assert.Equal(t, 1, s.Size())
		// todo
	})
	t.Run("Any/Every/ForEach/Where", func(t *testing.T) {
		size := 100
		s := NewHashSet(orderedIntArray(size)...)
		assert.Equal(t, size, s.Size())
		assert.True(t, s.Any(func(i interface{}) bool {
			v, ok := i.(int)
			return ok && v%2 == 0
		}))
		assert.False(t, s.Any(func(i interface{}) bool {
			return i.(int) > size-1
		}))
		assert.True(t, s.Every(func(i interface{}) bool {
			_, ok := i.(int)
			return ok
		}))
		assert.False(t, s.Every(func(i interface{}) bool { return i.(int) < size-1 }))
		l := 0
		s.ForEach(func(i interface{}) {
			l += i.(int)
		})
		assert.Equal(t, (size-1)*size/2, l)
		assert.Equal(t, size/2, s.Where(func(i interface{}) bool {
			v, _ := i.(int)
			return v%2 == 0
		}).Size())
	})
	t.Run("ToArray/Difference/Intersection/Union", func(t *testing.T) {
		s1 := NewHashSet(orderedIntArray(3)...)
		assert.ElementsMatch(t, []interface{}{0, 1, 2}, s1.ToArray())
		s2 := NewHashSet(orderedIntArray(5)[1:]...)
		assert.ElementsMatch(t, []interface{}{0}, s1.Difference(s2).ToArray())
		assert.ElementsMatch(t, []interface{}{3, 4}, s2.Difference(s1).ToArray())
		assert.ElementsMatch(t, []interface{}{1, 2}, s1.Intersection(s2).ToArray())
		assert.ElementsMatch(t, []interface{}{0, 1, 2, 3, 4}, s1.Union(s2).ToArray())
	})
	t.Run("String/Join", func(t *testing.T) {
		s := NewHashSet()
		assert.Equal(t, "HashSet{}", s.String())
		s = NewHashSet(42)
		assert.Equal(t, "HashSet{42}", s.String())
		s = NewHashSet(orderedIntArray(3)...)
		assert.Equal(t, len("HashSet{0, 1, 2}"), len(s.String()))
		assert.Equal(t, len("0 - 1 - 2"), len(s.Join(" - ")))
	})
}

func orderedIntArray(size int) []interface{} {
	elements := make([]interface{}, size)
	for i := 0; i < size; i++ {
		elements[i] = i
	}
	return elements
}
