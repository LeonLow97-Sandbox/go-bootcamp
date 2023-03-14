# Materials Link

- [Context Timeout](https://www.sohamkamani.com/golang/context-cancellation-and-values/)

## `WithTimeout` vs `WithDeadline`

- Any application that needs to maintain an SLA (service level agreement) for the maximum duration of a request, should use time based cancellation.

```go
// The context will be cancelled after 3 seconds
// If it needs to be cancelled earlier, the `cancel` function can be used
ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)

// Setting a context deadline is similar to setting a timeout, except
// you specify a time when you want the context to cancel, rather than a duration.
// Here, the context will be cancelled on 2009-11-10 23:00:00
ctx, cancel := context.WithDeadline(ctx, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
```