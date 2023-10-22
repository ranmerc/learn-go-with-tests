# Notes for Iteration

## Looping

Go only has `for` as a looping construct. It sort of works as overloaded term for all the other looping constructs -

```go
  // like while
  for i > 10 {
    fmt.Println(i)
    i += 1
  }

  // like for
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  // infinite loop
  for {
    fmt.Println()
  }

```

`break` and `continue` work as expected.

## Benchmarking

The `testing.B` gives you access to the cryptically named `b.N`.

When the benchmark code is executed, it runs ` b.N`` times and measures how long it takes.

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.

To run the benchmarks do `go test -bench=.`

## References

- [Go by Example - variables](https://gobyexample.com/variables)
- [Go by Example - for](https://gobyexample.com/for)
