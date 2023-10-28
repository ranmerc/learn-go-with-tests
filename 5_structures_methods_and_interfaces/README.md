# Notes on Structures, Methods and Interfaces

## Structures

We can declare Structures as -

```go
type Rectangle struct {
  Width float64
  Height float64
}

Rectangle{10.0, 10.0}
```

## Methods

Go does not have function overloading. We can

- Create a new package having function with same name
- Define methods on types

A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

Methods are very similar to functions but they are called by invoking them on an instance of a particular type.

`func (receiverName ReceiverType) MethodName(args)`

```go
type Rectangle struct {
  Width  float64
  Height float64
}

func (r Rectangle) Area() float64 {
  return 0
}
```

It is a convention in Go to have the receiver variable be the first letter of the type.

```go
r Rectangle
```

## Interfaces

Interfaces are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.

In Go **interface resolution is implicit**. If the type you pass in matches what the interface is asking for, it will compile. No need to do `My type Foo implements interface Bar`.

By declaring an interface, the we can create functions that are decoupled from the concrete types and only concerns itself with the required methods on interface.

**Table driven tests** are useful when you want to build a list of test cases that can be tested in the same manner.

Table driven tests can be a great item in your toolbox, but be sure that you have a need for the extra noise in the tests. They are a great fit when you wish to test various implementations of an interface, or if the data being passed in to a function has lots of different requirements that need testing.

In Test-Driven Development by Example Kent Beck refactors some tests to a point and asserts:

> The test speaks to us more clearly, as if it were an assertion of truth, not a sequence of operations

```go
func TestArea(t *testing.T) {

  areaTests := []struct {
    name    string
    shape   Shape
    hasArea float64
  }{
    {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
    {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
    {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
  }

  for _, tt := range areaTests {
    // using tt.name from the case to use it as the `t.Run` test name
    t.Run(tt.name, func(t *testing.T) {
      got := tt.shape.Area()
      if got != tt.hasArea {
        t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
      }
    })

  }
```

You can run specific tests within your table with `go test -run TestArea/Rectangle`

## References

- [Parametric Polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism)
