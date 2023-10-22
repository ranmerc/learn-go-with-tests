# Notes on Arrays and Slices

Arrays have a fixed capacity which you define when you declare the variable.

We can initialize an array in two ways:

```go
  numbers := [5]int{1, 2, 3, 4, 5}

  numbers := [...]int{1, 2, 3, 4, 5}
```

## Slices

The syntax is very similar to arrays, you just omit the size when declaring them
`mySlice := []int{1,2,3}` rather than `myArray := [3]int{1,2,3}`

It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. **Every test has a cost**.

`make(type, n)` allows you to n slices of type `type`.

```go
  func SumAll(numbersToSum ...[]int) []int {
    lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)
    .
    .
  }
```

The append function which takes a slice and a new value, then returns a new slice with all the items in it.

```go
  func SumAll(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
      sums = append(sums, Sum(numbers))
    }
    .
    .
  }
```

The syntax to slice a slice is `slice[low:high]`. If you omit the value on one of the sides of the : it captures everything to that side of it.

For example, the expression b[1:4] creates a slice including elements 1 through 3 of b (the indices of the resulting slice will be 0 through 2).

Slicing does not copy the slice’s data. It creates a new slice value that points to the original array. Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice.

`len(slice)` to get length of slice. The length is the number of elements referred to by the slice.

`cap(slice)` to get capacity of slice. The capacity is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).

Go has support for first class functions.

## References

- [Go by Example - Variadic functions](https://gobyexample.com/variadic-functions)
- [reflect.DeepEqual](https://pkg.go.dev/reflect#DeepEqual)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
