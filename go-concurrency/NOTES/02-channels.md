# Content

# Channels

- Channels provide a mechanism to **synchronize and communicate between goroutines**.
- Channels help avoid data races by allowing goroutines to share memory by communicating, not by sharing variables directly.
- Channels behave like **FIFO (First In, First Out)** queues internally.
- Value Transfer and Ownership
  - Sending a value to a channel can be seen as releasing ownership.
  - Receiving a value from a channel is like acquiring ownership.
  - Ownership here is logical, not enforced by the language.
- Channels are **first-class types** in Go, requiring no imports. Unlike `sync.Mutex`, which requires importing the `sync` package.
- Channels help write safer code (avoid data races) without manual locking, though they don't prevent all bad concurrency patterns.
- Go also provides traditional tools (`sync`, `sync/atomic`) which may be better suited in specific performance-critical or low-level use cases.

## Channel Types and Values

- Like arrays, slices and maps, a channel has an element type `T`. It only transfers values of that type.
- Channel direction types:
  - `chan T`: Bidirectional channel, can both send and receive values.
  - `chan<- T`: Send-only channel, can only send values, cannot receive.
  - `<-chan T`: Receive-only channel, can only receive values, cannot send.
- Type conversions:
  - Bidirectional channels (`chan T`) can be implicitly converted to send-only (`chan<- T`) or receive-only (`<-chan T`) channels.
  - Conversion from send-only to receive-only or vice versa is not allowed.
  - Direction arrows `<-` in channel types act as modifiers.
- Channel Capacity:
  - Channels have a capacity, defined when created.
  - **Unbuffered channels**: capacity is zero; **synchronous** send/receive.
  - **Buffered channels**: capacity > 0; **asynchronous** send/receive up to capacity.
- Zero value and creation:
  - The zero value of a channel is `nil` (no channel).
    - `var ch chan int`. Any operation like send or receive on a `nil` channel blocks forever.
  - Non-nil channel must be created with `make`
    - `make (chan int, 10)` creates a buffered channel of `int` with capacity 10.
    - `make (chan int)` creates an unbuffered channel of int with zero capacity.
    - The capacity argument is optional; default is zero (unbuffered).

## Channel Value Comparisons

- All channel types are comparable (e.g., `ch1 == ch2` is valid).
- Two channels are equal if they refer to the same underlying channel. (see code below)

```go
ch1 := make(chan int)
ch2 := make(chan int)
fmt.Println(ch1 == ch2) // false, two separate channels

ch3 := ch1
fmt.Println(ch1 == ch3) // true, both refer to the same channel
```

- Channel values are multi-part (internal pointer-like structure).
  - Internally, a Go channel value is **not the actual channel** itself but a reference (like a pointer) to a runtime structure that holds:
    - The buffer (if any)
    - Queue of senders/receivers
    - Synchronization metadata
  - Think of the channel value as a handle or reference to a Go runtime object that lives elsewhere in memory.
- `nil` is the zero value of channel types.
  - Just like a pointer, map or slice, when you declare a channel without initializing it, the value is `nil`.
  - Sending/receiving from a `nil` channel **blocks forever**.
  - `ch == nil` checks if a channel has been initialized.

```go
var ch chan int
fmt.Println(ch == nil) // true
```

## Go Channel Operations

1. Closing a Channel

- Syntax: `close(ch)`
- Only channels that are not receive-only can be closed.
- After closing, you **cannot send** more values to the channel.
- Receivers can still receive **remaining buffered values**, or get the **zero value** once the buffer is empty.
- `close()` is typically used by senders to signal no more values will be sent.

2. Sending to a Channel

- Syntax: `ch <- v`, sends value `v` into channel `ch`.
- `v` must be assignable to the channel's element type.
- Channel must not be **receive-only**.
- Synchronized operation (safe for concurrent use).

3. Receiving from a Channel

- Syntax (**single-value receive**): `v = <-ch`
- Syntax (**multi-value receive**): `v, ok = <-ch`
  - `ok` is a boolean indicating whether the value was received **before the channel was closed**.
- Cannot be used on a **send-only** channel.
- Synchronized operation.

4. Channel Capacity

- Syntax: `cap(ch)`
- Returns the capacity of the channel buffer as an `int`.
- For a `nil` channel, returns 0.
- Rarely used in practice but helpful in debugging or profiling.

5. Channel Length

- Syntax: `len(ch)`
- Returns the number of values **currently in the buffer** (sent but not yet received).
- For a `nil` channel, returns 0.
- Rarely used in practice but helpful in debugging or profiling.

### Thread-Safety and Synchronization

- All channel operations (`send`, `receive`, `close`) are **synchronized** and concurrency-safe.
- However, regular value assignments (e.g., assigning a received value to a variable) are not synchronized.

### Detailed Explanation for Channel Operations

- Channel Categories
  1. `nil` channels: `var ch chan int` (declared but not initialized)
  2. **Closed channels**: explicitly closed with `close(ch)`
  3. **Not-Closed Non-Nil channels**: normal channels that are open and usable.

Behaviors of different channels

| Operation | Nil Channel   | Closed Channel    | Not-Closed Non-Nil Channel        |
| --------- | ------------- | ----------------- | --------------------------------- |
| Close     | panic         | panic             | succeed to close `(C)`            |
| Send      | block forever | panic             | block or succeed to send `(B)`    |
| Receive   | block forever | never block `(D)` | block or succeed to receive `(A)` |

## Internal Structure of a Channel

Each Go channel consists of the following **3 internal FIFO queues**:

1. Receiving Goroutine Queue

- Stores goroutines blocked on receive (`<-ch`).
- Implemented as a **linked list** with no size limit.
- Goroutines **wait here until a value is available to receive**.

2. Sending Goroutine Queue

- Stores goroutines blocked on send (`ch <- val`).
- Also a **linked list** with no size limit.
- Stores both the goroutine and the **value (or its address)** being sent.
- Waits until **a receiver is available** (or buffer space if channel is buffered).

3. Value Buffer Queue

- A **circular queue** holding actual values.
- **Size = channel capacity** (`make(chan T, capacity)`)
- Values are of the **element type** of the channel.
- Behavior:
  - **Full** when number of elements = capacity.
  - **Empty** when no elements are stored.
  - **Unbuffered channels (capacity = 0)** are always both full and empty.
    - value buffer queue is `nil` for unbuffered channels.

Summary Table:

| Queue Type                | Characteristics                          | Triggered When                       |
| ------------------------- | ---------------------------------------- | ------------------------------------ |
| Receiving Goroutine Queue | Unbounded FIFO, stores waiting receivers | Receive happens, but no value yet    |
| Sending Goroutine Queue   | Unbounded FIFO, stores waiting senders   | Send happens, but no receiver/buffer |
| Value Buffer Queue        | Bounded circular buffer (capacity size)  | Buffered channel only; stores values |

```
                 ┌────────────────────────────────────────┐
                 │               Channel                  │
                 └────────────────────────────────────────┘
                        ▲             ▲            ▲
                        │             │            │
        ┌───────────────┘             │            └────────────┐
        │                             │                         │
        ▼                             ▼                         ▼
┌────────────────┐        ┌────────────────────┐     ┌────────────────────┐
│ Receiver Queue │◄───────┤   Value Buffer Q   ├────►| Sender Queue       │
│  (goroutines)  │        │ (circular buffer)  │     │ (goroutine+value)  │
└────────────────┘        └────────────────────┘     └────────────────────┘
        ▲                       ▲   ▲   ▲   ▲                ▲
        │                       │   │   │   │                │
        │                      Val Val Val Val               │
        │                       ▼   ▼   ▼   ▼                │
        │                 (0 ≤ len ≤ cap)                    │
        └────────────────────────────────────────────────────┘
                          Protected by internal lock
```

- Mutex for Synchronization
  - Every channel has an **internal mutex lock**.
  - Ensures **synchronization and thread safety** during send/receive/close operations.
  - Prevents **data races** when multiple goroutines access the channel concurrently.

### Internal runtime of Channel, `hchan`

```go
// Found in src/runtime/chan.go
type hchan struct {
	qcount   uint           // No. of elements current in the channel (queue count)
	dataqsiz uint           // Capacity of the channel buffer

	buf      unsafe.Pointer // Points to the Value Buffer queue storing elements (circular queue)
    recvq    waitq  // Queue of goroutines waiting for receive (Receiving Goroutine Queue)
	sendq    waitq  // Queue of goroutines waiting to send (Sending Goroutine Queue)

	sendx    uint   // Index where the next element will be sent (write pointer)
	recvx    uint   // Index where the next element will be received (read pointer)

	lock mutex // Internal lock protecting the entire structure
}

// Simple linked list of goroutines (represented as `sudog` structs) waiting on the channel
type waitq struct {
	first *sudog // first goroutine waiting
	last  *sudog // last goroutine waiting
}
```

To understand this further, let's look at 3 examples.

| Code Version             | Sender | Receiver  | Deadlock? | Why?                               |
| ------------------------ | ------ | --------- | --------- | ---------------------------------- |
| Unbuffered, 1 goroutine  | Main   | Main      | ✅ YES    | Main blocks on send, can't receive |
| Unbuffered, 2 goroutines | Main   | Goroutine | ❌ NO     | Send/receive match directly        |
| Buffered (size 1)        | Main   | Main      | ❌ NO     | Buffer holds value temporarily     |

#### Example 1: Unbuffered channel with 1 goroutine (🔴 Deadlock)

```go
func main() {
    ch := make(chan int)
    ch <- 1     // <-- blocks forever
    <-ch
}
```

---

Problem:

- On `ch <- 1`, the main goroutine tries to send `1` into the channel.
- But since **no other goroutine is ready to receive**, the main goroutine **blocks**.
- Because the same (and only) goroutine is blocked on the `ch <- 1` line, it **never reaches the next line** (`<-ch`) to receive the value.
- --> **Deadlock** occurs.

---

Internal Mechanics (Unbuffered Channel):

1. `make(chan int)` creates a **zero-capacity channel**.
   - Internally:
     - `dataqsiz = 0` (buffer size)
     - No values can be stored.
     - Send and receive **must happen at the same time**.
     - `buf`: ❌ (buffer queue doesn't exist because capacity = 0)
   - For `ch := make(chan int), channel created with:
     - `buf = nil`
     - `recvq = []`
     - `sendq = []`
2. `ch <- 1`:
   - This is a **send operation**.
   - Go's runtime checks:
     - Buffer? ❌ (unbuffered)
     - Any goroutines in `recvq` (receivers waiting)? ❌
   - So the current goroutine (`main`) is added to the `sendq` and goes into a **blocked** state.
3. `main` is the only goroutine running.
   - It's now blocked, so it never reaches the next line (`<-ch`) to receive the value.
   - Since **no other goroutines exists to receive, nothing can unblock `main`**.
   - --> **DEADLOCK**
   - At runtime, GO detects this during execution and prints: `fatal error: all goroutines are asleep - deadlock!`

---

#### Example 2: Unbuffered channel with 2 goroutines (✅ Fixed Version)

```go
func main() {
    ch := make(chan int)

    go func() {
        <-ch // Receiver is now running in a separate goroutine
    }()

    ch <- 1 // Sender in main, will succeed since receiver is ready
}
```

---

Behind the Scenes

1. `make(chan int)` creates a **zero-capacity channel**.
2. `go func() { <-ch }()`:
   - Starts a **new goroutine that tries to receive**.
   - Since there's no value yet, it is added to `recvq` and blocks, waiting for a sender.
3. Back in `main`, we call `ch <- 1`:
   - Runtime sees a **receiver is waiting in `recvq`**.
   - Value is handed directly from sender to the waiting goroutine receiver.
     - The receiving goroutine is then removed from `recvq`
     - The receiving goroutine is marked as ready to run (scheduled to continue execution).
   - Receiver is **unblocked**, continues to run.
   - Sender (`main`) also continues.
   - ✅ NO DEADLOCK 🎉🎉🎉🎉🎉

---

#### Example 3: Buffered Channel (Size 1)

```go
func main() {
    ch := make(chan int, 1)
    ch <- 1    // succeeds — there's room in the buffer
    <-ch       // receives from the buffer
}
```

1. `make(chan int, 1)` creates a **buffered channel** with `dataqsiz = 1`.
   - `buf`: size 1
2. `ch <- 1`:
   - Buffer has space (`qcount = 0 < dataqsiz = 1`), so value is added to buffer.
     - `buf` is not full --> put value in `buf[0]`, set `qcount = 1`, `sendx` pointer moves forward.
   - Sender continues immediately (no blocking).
3. `<- ch`:
   - Value is taken from the buffer.
     - `buf` is not empty --> remove value from `buf[0]`, `qcount = 0`.
   - Receiver continues immediately (no blocking).

- ✅ Both operations are non-blocking!

# References

- [Channels in Go](https://go101.org/article/channel.html)
