package maps

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

var entries = []Entry{{1, "1"}, {2, "2"}, {3, "3"}}

func TestHashMap(t *testing.T) {
	t.Run("Core/Put/Remove/ContainsKey/ContainsValue", func(t *testing.T) {
		hs := NewMapFromEntries(Entry{1, "a"})
		require.Equal(t, 1, hs.Size())
		require.True(t, hs.ContainsKey(1))
		require.False(t, hs.ContainsKey(42))
		require.True(t, hs.ContainsValue("a"))
		require.False(t, hs.ContainsValue("foo"))
		require.True(t, hs.IsNotEmpty())
		require.False(t, hs.IsEmpty())
		hs.Put(2, "b")
		require.Equal(t, 2, hs.Size())
		require.True(t, hs.ContainsKey(2))
		require.True(t, hs.ContainsValue("b"))
		hs.Remove(1)
		require.Equal(t, 1, hs.Size())
		hs.Clear()
		require.True(t, hs.IsEmpty())
		require.False(t, hs.IsNotEmpty())
	})
	t.Run("Entries/Keys/Values", func(t *testing.T) {
		hs := NewMapFromArray(
			[]interface{}{1, 2, 3},
			func(element interface{}) (result interface{}) { return element },
			func(element interface{}) (result interface{}) { return strconv.Itoa(element.(int)) },
		)
		require.ElementsMatch(t, hs.Entries(), entries)
		require.ElementsMatch(t, []int{1, 2, 3}, hs.Keys())
		require.ElementsMatch(t, []string{"1", "2", "3"}, hs.Values())
	})
	t.Run("AddAll/AddEntries/ForEach/PutIfAbsent/RemoveWhere", func(t *testing.T) {
		hs := NewMap()
		hs.AddEntries(entries)
		require.Equal(t, 3, hs.Size())
		var sum int
		hs.ForEach(func(key interface{}, value interface{}) {
			sum += key.(int)
			v, err := strconv.Atoi(value.(string))
			require.NoError(t, err)
			sum += v
		})
		require.Equal(t, 12, sum)
		v := hs.PutIfAbsent(1, func() interface{} { return "123" })
		require.Equal(t, "1", v)
		v = hs.PutIfAbsent(4, func() interface{} { return "5" })
		require.Nil(t, v)
		hs.RemoveWhere(func(key interface{}, value interface{}) bool {
			i, err := strconv.Atoi(value.(string))
			require.NoError(t, err)
			return key.(int) == i
		})
		require.Equal(t, 1, hs.Size())
		require.True(t, hs.ContainsKey(4))
		nhs := NewMapFromEntries(entries...)
		hs.AddAll(nhs)
		require.Equal(t, 4, hs.Size())
		require.Subset(t, hs.Entries(), entries)
	})
}
