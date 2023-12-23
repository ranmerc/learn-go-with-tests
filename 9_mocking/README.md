# Notes on Mocking

When building programs, always take an iterative, test-driven approach.

We make sure we take the smallest steps we can to have useful software.

We don't want to spend a long time with code that will theoretically work after some hacking because that's often how developers fall down rabbit holes. **It's an important skill to be able to slice up requirements as small as you can so you can have working software.**

Whenever we want to write somewhere, io.Writer is the de-facto way of capturing that as an interface in Go.

- In main we will send to `os.Stdout` so our users see the countdown printed to the terminal.
- In test we will send to `bytes.Buffer` so our tests can capture what data is being generated.

**Take a thin slice of functionality and make it work end-to-end, backed by tests.**
