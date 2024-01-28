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

### References

- [Laws of Reflection](https://blog.golang.org/laws-of-reflection)
