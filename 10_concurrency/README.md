# Notes on Concurrency

Anonymous functions have a number of features which make them useful, two of which we're using above.

- Firstly, they can be executed at the same time that they're declared - this is what the () at the end of the anonymous function is doing.

- Secondly they maintain access to the lexical scope in which they are defined - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

Look out for loop closure issues when working with goroutines.

Race condition, is a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over.

Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.

We can declare anonymous structs -

```go
type result struct {
	string
	bool
}

// Accessed as -
r.string
r.bool
```

Send statement, the <- operator, takes a channel on the left and a value on the right:

```go
// Send statement
resultChannel <- result{u, wc(u)}
```

Receive expression uses the <- operator, but with the two operands now reversed: the channel is now on the right and the variable that we're assigning to is on the left:

```go
// Receive expression
r := <-resultChannel
```

By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time.

We have parallelized the part of the code that we wanted to make faster, **while making sure that the part that cannot happen in parallel still happens linearly**.

Where 'work' is making the tests pass, 'right' is refactoring the code, and 'fast' is optimizing the code to make it, for example, run quickly.
