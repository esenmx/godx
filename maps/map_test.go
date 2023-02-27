package maps

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var entries = []Entry[int, string]{{1, "1"}, {2, "2"}, {3, "3"}}

func TestHashMap(t *testing.T) {
	t.Run("Core/Put/Remove/ContainsKey/ContainsValue", func(t *testing.T) {
		hs := NewMapFromEntries[int, string](Entry[int, string]{1, "a"})
		require.Equal(t, 1, hs.Size())
		require.True(t, hs.ContainsKey(1))
		require.False(t, hs.ContainsKey(42))
		require.Equal(t, hs.hash[1], "a")
		require.True(t, hs.IsNotEmpty())
		require.False(t, hs.IsEmpty())
		hs.Put(2, "b")
		require.Equal(t, 2, hs.Size())
		require.True(t, hs.ContainsKey(2))
		require.Equal(t, hs.hash[2], "b")
		hs.Remove(1)
		require.Equal(t, 1, hs.Size())
		hs.Clear()
		require.True(t, hs.IsEmpty())
		require.False(t, hs.IsNotEmpty())
	})
	//t.Run("Entries/Keys/Values", func(t *testing.T) {
	//	hs := NewMapFromArray[int, string](
	//		[]int{1, 2, 3},
	//		func(element int) (result int) { return element },
	//		func(element int) (result string) { return strconv.Itoa(element) },
	//	)
	//	require.ElementsMatch(t, hs.Entries(), entries)
	//	require.ElementsMatch(t, []int{1, 2, 3}, hs.Keys())
	//	require.ElementsMatch(t, []string{"1", "2", "3"}, hs.Values())
	//})
	t.Run("AddAll/AddEntries/ForEach/PutIfAbsent/RemoveWhere", func(t *testing.T) {
		hs := NewMap[int, string]()
		hs.AddEntries(entries)
		require.Equal(t, 3, hs.Size())
		var sum int
		hs.ForEach(func(key int, value string) {
			sum += key
			v, err := strconv.Atoi(value)
			require.NoError(t, err)
			sum += v
		})
		require.Equal(t, 12, sum)
		v := hs.PutIfAbsent(1, func() string { return "123" })
		require.Equal(t, "1", *v)
		v = hs.PutIfAbsent(4, func() string { return "5" })
		require.Nil(t, v)
		hs.RemoveWhere(func(key int, value string) bool {
			i, err := strconv.Atoi(value)
			require.NoError(t, err)
			return key == i
		})
		require.Equal(t, 1, hs.Size())
		require.True(t, hs.ContainsKey(4))
		nhs := NewMapFromEntries(entries...)
		hs.AddAll(nhs)
		require.Equal(t, 4, hs.Size())
		require.Subset(t, hs.Entries(), entries)
	})

	t.Run("RemoveAll", func(t *testing.T) {
		base := map[int]string{1: "", 2: "", 3: ""}
		hs := NewMapWithValues[int, string](base)
		hs.RemoveAll()
		require.Equal(t, base, hs.hash)
		hs.RemoveAll(1, 3)
		require.Equal(t, map[int]string{2: ""}, hs.hash)
	})
}
