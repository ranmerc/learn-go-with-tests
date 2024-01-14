# Notes on Select

In the mocking and dependency injection chapters, we covered how ideally we don't want to be relying on external services to test our code because they can be

- Slow
- Flaky
- Can't test edge cases

Package called `net/http/httptest` which enables users to easily create a mock HTTP server.

We wrap `http.HandlerFunc` in `httptest.NewServer` which finds an open port to listen on and then you can close it when you're done with your test.
