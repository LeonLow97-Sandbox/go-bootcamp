# Understanding Golang Context

- [Video Link](https://www.youtube.com/watch?v=kaZOXRqFPCw&ab_channel=AnthonyGG)

## Managing requests and cancellations in Go Applications

- Context is a powerful tool in Go for managing requests and cancellations in Go applications.
- It provides a mechanism for carrying request-scoped values such as user authentication information, request deadline, and cancellation signals, across API boundaries and between processes.
- The context package in Go allows passing a `context` to programs.
- Context can be used to pass in a timeout, deadline, or a channel to indicate stop working and return.
- In production-grade systems, it is a good practice to have timeouts for web requests or system commands to prevent <a href="#backup-request">backup of requests</a> and avoid degraded performance.
- A timeout or deadline context can come in handy in cases where APIs are running slow or not responding, preventing a cascading effect on other requests.

### What is Context in Go?

- Context is a package in Go that provides a way to carry request-scoped values across API boundaries and between processes.
- Context is designed to be used in cases where a request flows through multiple API handlers, and we need to pass some information between them.
- Context can be used to manage timeouts and cancellations for requests.

### How to create a new Context?

- The `context.Background()` function returns a new empty Context, which can be used as the parent of all other Contexts in a request.
- The `context.WithValue(parent Context, key interface{}, value interface{})` function creates a new Context that is a child of the parent Context and carries the given key-value pair.
  - In the event of logging errors in your application, can store the key as the request ID of the user and the value as the error that has occurred.
  - Thus, we can know which request is causing the error.
  - Useful in large, distributed systems where many requests are processed concurrently.
  - `context.WithValue()` is not got global or long-lived values, so don't store database connections or other resources that should be managed by a connection pool.

### How to use Context to carry request-scoped values?

- Use the `context.WithValue` function to add request-scoped values to a Context.
- For example, we can add a user authentication token to a Context using the key "authToken" and the value of the token.

```go
ctx := context.WithValue(parentContext, "authToken", "token123")

val := ctx.Value("authToken") // val is of type interface{} any "token123"
```

### How to use Context to manage timeouts and cancellations?

- Use context to manage timeouts and cancellations by creating a new Context with a `context.WithTimeout(parent Context, timeout time.Duration)` or `context.WithDeadline(parent Context, deadline time.Time)` function.
- `context.WithTimeout`: function creates a new Context with a timeout, which will automatically cancel the Context and all its children after the given duration.
- `context.WithDeadline`: function creates a new Context with a deadline, which will automatically cancel the Context and all its children at the given time.
- `context.WithCancel(parent Context)`: function creates a new Context with a cancellation function that can be used to create a new Context with a cancellation function that can be used to manually cancel the Context and all its children.

```go
ctx, cancel := context.WithTimeout(parentContext, time.Second * 10)
defer cancel() // always call cancel function to release resources

ctx, cancel := context.WithDeadline(parentContext, time.Now().Add(time.Second * 10))
defer cancel() // always call cancel function to release resources

ctx, cancel := context.WithCancel(parentContext)
defer cancel() // always call cancel function to release resources
```

### Best practices for using Context in Go Applications

- Always create a new Context for each request and use it to carry request-scoped values and manage timeouts and cancellations.
- Use Context to pass request-scoped values between API handlers instead of global variables or other shared resources.
- Always pass Context as the first argument to API handlers and make sure to propagate it to any other functions that need it.
- Always call the cancel function returned by Context functions when the request is finished to release any resources associated with the Context.

<h2 id="backup-request">Backup Request</h2>

- In a system that receives a large number of requests, a backup of requests occurs when the system is unable to process incoming requests quickly enough, causing them to accumulate and waiting in line for processing.
- This can happen for various reasons such as network latency, system overload, or limited resources.
- When requests are backed up, the system may become unresponsive, leading to increased latency and reduced performance.
- In such cases, it is essential to have measures in place, such as timeouts and cancellations using the context package, to prevent backups and ensure that the system remains responsive to incoming requests.
