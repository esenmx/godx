# GO Data Structures

Data structures implemented with `golang` with optional `thread-safety`.

### What's the difference between [gods](https://github.com/emirpasic/gods?) and `godx`?

Simple, optional `thread-safe` data structures types and wider interfaces.

Interface
references: [Java/Collections](https://docs.oracle.com/javase/7/docs/api/java/util/Collections.html)
& [Dart/Collection](https://pub.dev/packages/collection)

An example comparison:

**`gods` HashSet interface**:

```go
Add(elements ...interface{})
Remove(elements ...interface{})
Contains(elements ...interface{}) bool
Empty() bool
Size() int
Clear()
Values() []interface{}
```

**`godx` HashSet interface**:

```go
Add(interface{}) bool
AddAll(...interface{})
Any(func (interface{}) bool) bool
Clear()
Contains(interface{}) bool
ContainsAll(...interface{}) bool
Difference(*HashSet) *HashSet
Every(func (interface{}) bool) bool
ForEach(func (interface{}))
Intersection(*HashSet) *HashSet
IsEmpty() bool
Join(string) string
Remove(interface{}) bool
RemoveAll(...interface{})
RetainAll(set *HashSet)
Size() interface{}
ToArray() []interface{}
Union(*HashSet) *HashSet
Where(func (interface{}) bool) *HashSet
```

**other methods also incoming, especially with long awaited `generics` feature.*

### Progress Tracker:

- [X] HashSet
- [ ] TreeSet
