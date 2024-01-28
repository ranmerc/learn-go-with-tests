# Notes on Reflection

Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming.

`any` is an alias for `interface{}` which means any type.

In go `switch` cases don't fallthrough by default. You need to opt in using `fallthrough`

Or you can match multiple values by clubbing them together in `case` statement.

```go
switch(x):
case 1, 2:
  fmt.Printf("One or two")
```

Iterating over channels -

```go
for v, ok := value.Recv(); ok; v, ok = value.Recv() {
  walkValue(v)
}
```

Send in base type `nil` places where you are expected to send something, but have nothing to send.

## Corresponding functions on `reflect.Value` for different data types

### String

- `String()` value as string.

### Struct

- `NumField()` number of fields present in struct.
- `Field(i)` reflect value of field at index.

### Array/Slice

- `Len()` returns length.
- `Index()` reflect value at index

### Map

- `MapKeys()` slice containing all keys of the map.
- `MapIndex()` reflect values present at index.

### Channel

- `Recv()` receive from channel.
- `Send()` send to channel.

### Function

- `Call()` calls with reflect value arguments and returns reflect values result

## References

- [Laws of Reflection](https://blog.golang.org/laws-of-reflection)
