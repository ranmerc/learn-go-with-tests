# Notes on Mocking

When building programs, always take an iterative, test-driven approach.

We make sure we take the smallest steps we can to have useful software.

We don't want to spend a long time with code that will theoretically work after some hacking because that's often how developers fall down rabbit holes. **It's an important skill to be able to slice up requirements as small as you can so you can have working software.**

Whenever we want to write somewhere, io.Writer is the de-facto way of capturing that as an interface in Go.

- In main we will send to `os.Stdout` so our users see the countdown printed to the terminal.
- In test we will send to `bytes.Buffer` so our tests can capture what data is being generated.

**Take a thin slice of functionality and make it work end-to-end, backed by tests.**

We can also use backticks for strings and include new lines.

To sleep for 1 seconds.

```go
  time.Sleep(1 * time.Second)
```

## Mocking

We have a dependency on Sleeping which we need to extract so we can then control it in our tests.

If we can mock `time.Sleep` we can use dependency injection to use it instead of a "real" `time.Sleep` and then we can spy on the calls to make assertions on them.

_Spies_ are a kind of mock which can record how a dependency is used. They can record the arguments sent in, how many times it has been called, etc.

When writing tests if you're not confident that your tests are giving you sufficient confidence, just break it
