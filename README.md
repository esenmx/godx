# GO Data Structures

Yet another boring but `simple` and `efficient` Golang implementation of popular data structures with `thread-safety`.

#### Why this package? What's the difference between [gods](https://github.com/emirpasic/gods?) and `godx`?

Well, most of the time you'll go with `gods` package because of your niche logic will require thread-safety at higher
level. On the other hand, sometimes a `thread-safe` standard collection can do a big favor for you. That's where `godx`
participates.

Another important point is, while `gods` has only essential built-in methods, `godx`
does not directly re-implements but
references [Java/Collections](https://docs.oracle.com/javase/7/docs/api/java/util/Collections.html)
& [Dart/Collection](https://pub.dev/packages/collection) interfaces in terms of functionality where many convenient
built-in methods implemented.

#### A simple example comparison of `gods` & `godx`:

**`gods` HashSet interface**:

```go
type Set interface {
Add(elements ...interface{})
Remove(elements ...interface{})
Contains(elements ...interface{}) bool

containers.Container
// Empty() bool
// Size() int
// Clear()
// Values() []interface{}
}
```

**`godx` HashSet interface**:

```go
type HashSetInterface interface {
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
}
```

**other methods also incoming, especially after with long waited `generic` feature.*

As you see, `godx` has much more built-in methods and that may make possible to reduce your logic into an already
packaged collection in some cases, which is big win for your productivity.

### Progress Tracker:

- [X] HashSet

#### Others are coming...