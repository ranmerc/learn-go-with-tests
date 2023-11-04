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
