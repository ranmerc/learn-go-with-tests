# Notes for Hello World

## Test Rules -

- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t \*testing.T

## TDD Discipline

- Write a test
- Make the compiler pass, there shouldn't be any compiler error
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

## Importance of TDD Discipline

- Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
- Writing the smallest amount of code to make it pass so we know we have working software
- Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with

`t.Helper()` is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper

We can group constants in a block instead of declaring them each on their own line. It's a good idea to use a line between sets of related constants for readability.

```go
const (
    french  = "French"
    spanish = "Spanish"
    .
    .
    .
)
```

In our function signature we have made a named return value (**prefix string**).

- This will create a variable called prefix in your function.
- It will be assigned the "zero" value. This depends on the type, for example ints are 0 and for strings it is "".
- You can return whatever it's set to by just calling return rather than return prefix.
- This will display in the Go Doc for your function so it can make the intent of your code clearer.

The function here name starts with a lowercase letter. In Go, public functions start with a capital letter and private ones start with a lowercase.

```go
func greetingPrefix(language string) (prefix string) {
 switch language {
 case french:
    prefix = frenchHelloPrefix
 case spanish:
    prefix = spanishHelloPrefix
 default:
    prefix = englishHelloPrefix
 }

 return
}
```
