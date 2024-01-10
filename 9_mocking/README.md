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

When writing tests if you're not confident that your tests are giving you sufficient confidence, just break it.

People normally get in to a bad state when they don't listen to their tests and are not respecting the refactoring stage.

If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should listen to that bad feeling and think about your code. Usually it is a sign of

- The thing you are testing is having to do too many things (because it has too many dependencies to mock)
  - Break the module apart so it does less
- Its dependencies are too fine-grained
  - Think about how you can consolidate some of these dependencies into one meaningful module
- Your test is too concerned with implementation details
  - Favour testing expected behaviour rather than the implementation

Normally a lot of mocking points to bad abstraction in your code.

More often than not poor test code is a result of bad design or put more nicely, **well-designed code is easy to test**.

Rules to follow -

1. Do not focus on testing implementation details too much.
2. Refactoring should not cause major changes in your testing code.
3. Avoid testing internals of a module and spend time more on actual public behaviour.
4. Using spies causes tight coupling between implementation and tests.

Mocking helps in -

1. Creating faster feedback loops - being able to test without waiting for async processes to finish.
2. Enables to test important areas of our code.
3. Depending on actual async services can make tests fragile.

**Always be mindful about the value of your tests and what impact they would have in future refactoring.**

## References

1. [TestDouble by Martin Flower](https://martinfowler.com/bliki/TestDouble.html)
