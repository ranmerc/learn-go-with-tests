# Notes on Pointers and Errors

## Pointers

Pointers let us point to some values and then let us change them.
We get the pointer (memory address) of something by placing an `&` character at the beginning of the symbol.

No need to deference pointer like this before using them,

```go
func (w *Wallet) Balance() int {
  return (*w).balance
}
```

These are _struct pointers_ and they are automatically dereferenced.

```go
func (w *Wallet) Balance() int {
  return w.balance
}
```

By convention you should keep your method receiver types the same for consistency.

Go lets you create new types from existing ones. The syntax is

```go
  type MyName OriginalType
```

`Stringer` interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints.

## Errors

In Go, if you want to indicate an error it is idiomatic for your function to return an err for the caller to check and act on.

`nil` is synonymous with `null` from other programming languages. If you see a function that takes arguments or returns values that are interfaces, they can be nillable.

Like `null` if you try to access a value that is `nil` it will throw a runtime panic. This is bad! You should make sure that you check for nils.

`errors.New` creates a new error with a message of your choosing.

`t.Fatal` which will stop the test if it is called.

The `var` keyword allows us to define values global to the package.

`errcheck` is useful to check for unchecked errors.

## References

- [Stringer](https://pkg.go.dev/fmt#Stringer)
- [errcheck](https://github.com/kisielk/errcheck)
- [Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
