package sets

import (
	"testing"

	"github.com/esenmx/godx/mock"
	"github.com/stretchr/testify/require"
)

func TestMethods(t *testing.T) {
	t.Run("NewSet", func(t *testing.T) {
		s := NewSet[int]()
		require.True(t, s.IsEmpty())
		s = NewSet[int](1, 2, 3)
		require.False(t, s.IsEmpty())
		require.Equal(t, 3, s.Size())
	})
	t.Run("Clear", func(t *testing.T) {
		s := NewSet(1, 2, 3)
		s.Clear()
		require.True(t, s.IsEmpty())
	})
	t.Run("Add/Remove/Contains", func(t *testing.T) {
		s := NewSet[int]()
		require.True(t, s.Add(42))
		require.False(t, s.Add(42))
		require.True(t, s.Contains(42))
		require.True(t, s.Add(1))
		require.Equal(t, 2, s.Size())
		require.True(t, s.Remove(1))
		require.False(t, s.Remove(1))
		require.Equal(t, 1, s.Size())
		require.True(t, s.Contains(42))
	})
	t.Run("AddAll/RemoveAll/ContainsAll/RetainAll/RemoveWhere", func(t *testing.T) {
		s := NewSet[int]()
		s.AddAll(42, 21)
		require.True(t, s.ContainsAll(42, 21))
		s.AddAll(84, 42, 21)
		require.Equal(t, 3, s.Size())
		s.RemoveAll(1, 2, 3)
		require.Equal(t, 3, s.Size())
		s.RemoveAll(42, 21, 1)
		require.Equal(t, 1, s.Size())
		// todo
		s.AddAll(21, 42)
		s.RemoveWhere(func(i int) bool { return i%2 == 0 })
		require.Equal(t, 1, s.Size())
		require.True(t, s.Contains(21))
	})
	t.Run("Any/Every/ForEach/Where", func(t *testing.T) {
		size := 100
		s := NewSet[int](mock.OrderedIntArray(size)...)
		require.Equal(t, size, s.Size())
		require.True(t, s.Any(func(i int) bool { return i%2 == 0 }))
		require.False(t, s.Any(func(i int) bool { return i > size-1 }))
		require.True(t, s.Every(func(i int) bool { return i >= 0 }))
		require.False(t, s.Every(func(i int) bool { return i < size-1 }))
		l := 0
		s.ForEach(func(i int) { l += i })
		require.Equal(t, (size-1)*size/2, l)
		require.Equal(t, size/2, s.Where(func(i int) bool { return i%2 == 0 }).Size())
	})
	t.Run("ToArray/Difference/Intersection/Union", func(t *testing.T) {
		s1 := NewSet[int](mock.OrderedIntArray(3)...)
		require.ElementsMatch(t, []int{0, 1, 2}, s1.ToArray())
		s2 := NewSet[int](mock.OrderedIntArray(5)[1:]...)
		require.ElementsMatch(t, []int{0}, s1.Difference(s2).ToArray())
		require.ElementsMatch(t, []int{3, 4}, s2.Difference(s1).ToArray())
		require.ElementsMatch(t, []int{1, 2}, s1.Intersection(s2).ToArray())
		require.ElementsMatch(t, []int{0, 1, 2, 3, 4}, s1.Union(s2).ToArray())
	})
	//t.Run("Map/Reduce", func(t *testing.T) {
	//	es := mock.RandomElements(100)
	//	set := NewSet[mock.Element](es...)
	//	require.Equal(t, 100, set.Size())
	//	as := set.Map[int](func(e mock.Element) int {
	//		return e.(mock.Element).A
	//	})
	//	require.Equal(t, 100, len(as))
	//	for _, v := range as {
	//		require.True(t, set.Any(func(e int) bool {
	//			return v == e.(mock.Element).A
	//		}))
	//	}
	//})
	t.Run("String/Join", func(t *testing.T) {
		s := NewSet[int]()
		require.Equal(t, "Set{}", s.String())
		s = NewSet[int](42)
		require.Equal(t, "Set{42}", s.String())
		s = NewSet[int](mock.OrderedIntArray(3)...)
		require.Equal(t, len("Set{0, 1, 2}"), len(s.String()))
		require.Equal(t, len("0 - 1 - 2"), len(s.Join(" - ")))
	})
}

//
// Mock
//
