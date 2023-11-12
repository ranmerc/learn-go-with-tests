# Notes on Maps

Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the [].

```go
  map[string][string]
```

Errors can be converted to a string with the `.Error()` method.

We can return multiple items from a function

```go
func (d Dictionary) Search(word string) (string, error) {
  return d[word], nil
}
```

A map lookup can return 2 values. The second value is a boolean which indicates if the key was found successfully.

```go
func (d Dictionary) Add(word, definition string) {
  d[word] = definition
}
```

Adding to a map is also similar to an array. You just need to specify a key and set it equal to a value.

An interesting property of maps is that you can modify them without passing as an address to it (e.g `&myMap`)

When you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

A gotcha with maps is that they can be a `nil` value. A `nil` map behaves like an empty map when reading, but attempts to write to a `nil` map will cause a **runtime panic**.

Therefore, you should never initialize an empty map variable, initialize with an empty map or use `make`

```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

We can group together declaration as such -

```go
var (
  ErrNotFound   = errors.New("could not find the word you were looking for")
  ErrWordExists = errors.New("cannot add word because it already exists")
)
```

Any type with an Error() string method fulfils the error interface.

Make errors immutable and more reusable by creating a type for it -

```go
type DictionaryErr string

func (e DictionaryErr) Error() string {
  return string(e)
}
```

then declaring errors as -

```go
const ErrNotFound = DictionaryErr('not found')
```

Go has a built-in function delete that works on maps. It takes two arguments. The first is the map and the second is the key to be removed.

```go
delete(d, word)
```

## References

- [If a map isn’t a reference variable, what is it?](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)
- [Constant errors](https://dave.cheney.net/2016/04/07/constant-errors)
