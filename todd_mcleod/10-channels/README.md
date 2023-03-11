# Golang — Understanding channel, buffer, blocking, deadlock and happy groutines.

I was so confused to understand behaviior of Golang channels, buffer, blocking, deadlocking and groutines.

I read Go by Example topics.

- [Go by Example: Channels](https://gobyexample.com/channels)
- [Go by Example: Channel Buffering](https://gobyexample.com/channel-buffering)

But I got questions.

- In the first place, what is channel?
- How Channel buffering work?
- What is deadlock of channel?
- Why main proccess needs channel buffering to use it?
- Why groutine does not raise deadlock error without channel buffering?

I tried to bridge between understanding of this topics.

## This topic tells you my understanding

If you found something strange, you can review it and send comment anytime.

## What is channel?

Golang channel makes goroutines can communicate each other.
Through channel, goroutines send or receive messages (values).

\* In this topic, I name "messaging" as "sending and receiving values".

I say, channel is like a communication hub, pipe, or bridge. 

## Understand goroutine

Before undarstand channels, let's undarstand goroutines.

Main func can spawn other goroutines, but "main" it self is one groutine.

```golang
package main

import "fmt"
import "runtime"

func main() {
	// Goroutine num includes main processing
	fmt.Println(runtime.NumGoroutine()) // 1

	// Spawn two goroutines
	go func() {}()
	go func() {}()

	// Total three goroutines run
	fmt.Println(runtime.NumGoroutine()) // 3
}
```

I know, Channel does not know defferencies between them.
Even it was a main groutine or spawned groutine, Channel handles them equally.

## Send and receive

A goroutine can send message to channel.
And, when channel receives message from goroutine, channel try to send that message "immediately" to other some goroutines.

```golang
package main

import "fmt"

func main() {
	messages := make(chan string)

	// Send message
	go func() { messages <- "Hello" }()

	// Receive message
	fmt.Println(<-messages) // Hello
}
```

## Blocking

Blocking happens when ...

For example

- A goroutine try to receive message
- But channel is empty. 
- And other some groutine is running

In this case, message receiving will be blocked and "Receiver" groutine waits it.
Because while other groutine is runnning, he hope someone send message to channel. ( Not surrender )

### Blocking example : Sleeping sender

```golang
package main

import "fmt"
import "time"

func main() {
	messages := make(chan string)

	// In spawned goroutine
	//
	// Send message to channel
	// But before do it sleep for a while
	go func() {
		time.Sleep(1000 * time.Millisecond)
		messages <- "Hello"
	}()

	// In main goroutine
	//
	// Receive message from channel
	// Message appears after spawned goroutine awaked from sleeping
	// Channel is empty until other goroutine send message
	// So this receiving will be blocked for a while
	fmt.Println(<-messages) // Hello
}
```

### Blocking example : Sleeping receiver

```golang
package main

import "fmt"
import "time"

func main() {
	messages := make(chan string)

	// Receiver
	go func() {
		fmt.Println("Receiver : I am waiting for your message.")
		msg := <-messages
		fmt.Println("Receiver : I got a mail.")
		fmt.Println(msg)
	}()

	// Sender
	time.Sleep(2000 * time.Millisecond)
	messages <- "Message : Do you like go langage?"

	// Wait spawned goroutine process
	time.Sleep(1000 * time.Millisecond)
}
```

## Deadlock

When deadlock happens? see this beautifule explaining.

>A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.

[Go: Deadlock | Programming.Guide](https://programming.guide/go/detect-deadlock.html)

For example:

- When one goroutine try to receive message from channel
- And channel is empty
- And no other goroutine running.

In this case, "Receiver" goroutine never receive message, so it surrenders messaging, and deadlock happens.

### Deadlock example. No sender

```golang
package main

func main() {
	messages := make(chan string)

	// Do nothing spawned goroutine
	go func() {}()

	// A groutine ( main groutine ) trying to send message to channel
	// But no other groutine runnning
	// And channel has no buffers
	// So it raises deadlock error
	messages <- "I wanna tell you." // fatal error: all goroutines are asleep - deadlock!
}
```

### Deadlock example. No receiver

```golang
package main

func main() {
	messages := make(chan string)

	// Do nothing spawned goroutine
	go func() {}()

	// A groutine ( main groutine ) trying to receive message from channel
	// But channel has no messages, it is empty.
	// And no other groutine running. ( means no "Sender" exists )
	// So channel will be deadlocking
	<-messages // fatal error: all goroutines are asleep - deadlock!
}
```

### Blocking example : No groutines procceeds

```golang
package main

import "fmt"

// Two groutines are running but deadlock happens

// Output example
//
// Try to receive message
// Try to receive message
// fatal error: all goroutines are asleep - deadlock!

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("Try to receive message") // Printing
		<-messages                            // Blocking
		fmt.Println("Receive message")        // Never reached
	}()

	fmt.Println("Try to receive message") // Printing
	<-messages                            // Blocking
	fmt.Println("Receive message")        // Never reached

}
```

## Addtional Questions

### Question. Why deadlock does not happen in spawned groutine ?

I think that reason is ....

Because "main goroutine" always running.
So until main groutine finishes, it means, one or more groutine running anytime.

When main groutine finishes, it is end of process, means end of all groutines.
So even if spawned groutine meets blocking, does not meet deadlock error.

### What is groutine's "asleep"

Now finding.

## How to get script examples

Download Zip or git clone.

```
git clone https://gist.github.com/YumaInaura/8d52e73dac7dc361745bf568c3c4ba37
```