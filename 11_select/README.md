# Notes on Select

In the mocking and dependency injection chapters, we covered how ideally we don't want to be relying on external services to test our code because they can be

- Slow
- Flaky
- Can't test edge cases

## httptest

Package called `net/http/httptest` which enables users to easily create a mock HTTP server.

We wrap `http.HandlerFunc` in `httptest.NewServer` which finds an open port to listen on and then you can close it when you're done with your test.

## defer

By prefixing a function call with defer it will now call that function at the end of the containing function.
Sometimes you will need to clean up resources, such as closing a file or in our case closing a server so that it does not continue to listen to a port.
You want this to execute at the end of the function, but keep the instruction near where you created the server for the benefit of future readers of the code.

## Select

If something can be done without waiting for the previous then we can use concurrency.

**`chan struct{}` is the smallest data type available from a memory perspective**

Always use `make` to create channels. Since for channels the zero value is nil and sending to it with <- will block it forever because you cannot send to nil channels.

`select` allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.

### References

- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
