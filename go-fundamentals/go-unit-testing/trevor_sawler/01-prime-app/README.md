# Learning Outcomes

- Table Test
- Test for output to the console window.

## Testing Commands

- `go test .`
  - runs tests for Go packages.
- `go test -v .`
  - runs tests for Go packages and prints verbose output.

## Testing function naming convention

- Test function should be named using the pattern `Test_Xxx` Where `Xxx` is the name of the function being tested.
- Include a comment explaining what the test is testing.

```go
func Test_Add(t *testing.T) {
    // Test that the Add function returns the correct result
    result := Add(1, 2)
    if result != 3 {
        t.Errorf("Add(1, 2) = %d; want 3", result)
    }
}
```

## Test Coverage

- `go test -cover .`
  - runs tests for Go packages and also reports on the coverage of the tested packages.
- `go test -coverprofile=coverage.out`
  - runs tests for Go packages and also generates a coverage profile.
- `go tool cover -html=coverage.out`
  - converts a coverage profile to HTML and opens it in a web browser.
  - red: not covered
  - green: covered

## Running Single Tests

- `go test -run Test_isPrime`
- `go test -v -run Test_isPrime`

## Running Groups of Tests (Test Suites)

- `go test -v -run Test_alpha`
- Need to rename the test functions

```
=== RUN   Test_alpha_isPrime
--- PASS: Test_alpha_isPrime (0.00s)
=== RUN   Test_alpha_prompt
--- PASS: Test_alpha_prompt (0.00s)
PASS
ok      primeapp        0.163s
```


