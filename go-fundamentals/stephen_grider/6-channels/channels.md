## Go Routine

- Engine that executes code line by line.
- Blocking call : takes some time to execute the function.

### How do we Fetch multiple requests at the same time?

- launch additional go routines using keyword `go` before calling the function.

```
	for _, link := range links {
		go checkLink(link)
	}
```

- The main routine then creates multiple go routines when there is a blocking call.
- `go`: creates a new thread go routine. (**only used in front of function calls**)

## Concurrency vs Parrellelism

- Concurrency: can have multiple threads executing code. If one thread blocks, another one is picked up and worked on.
- Parrellelism: multiple threads executed at the exact same time. Requires multiple CPU's.

## Running Program

- Main routine created when we launched our program.
- Child routine created by the 'go' keyword.

## Channels

- Channels are used to **communicate between the different go routines**. Otherwise, the `main` routine might end first.
- Data shared between the routines in the same channel must be of the same type. E.g., make channel of type string.

## Sending Data with Channels

|           code           |                                             description                                             |
| :----------------------: | :-------------------------------------------------------------------------------------------------: |
|      `channel <- 5`      |                                Send the value '5' into this channel                                 |
|  `myNumber <- channel`   | Wait for a value to be send into the channel. <br> When we get one, assign the value to 'myNumber'. |
| `fmt.Println(<-channel)` |     Wait for a value to be sent into the channel. <br> When we get one, log it out immediately,     |

## Creating an infite loop

```
	for {
		go checkLink(<-c, c)
	}
```

- The above code is equivalent to the one below (cleaner):

```
	// Creating an infinite loop for repeating routines
	for l := range c {
		go checkLink(l, c)
	}
```

## Integrate a 'Pause' before fetching another link

|           Code           |                                   Description                                   |
| :----------------------: | :-----------------------------------------------------------------------------: |
| `func Sleep(d Duration)` |         Sleep pauses the current goroutine for at least the duration d          |
|       `d Duration`       | Contains `Nanosecond`, `Microsecond`, `Millisecond`, `Second`, `Minute`, `Hour` |

## Function Literals

- Function literals are like Python lambda functions

```
	for l := range c {
		go func() {
			// pause for 5 seconds
			time.Sleep(5 * time.Second)
			checkLink(l,c)
		}	()
	}
```

## Additional Notes

- Never share variables between main routine and go routines.
- Only share channels between main routine and go routines.

## Quiz

1. Which of the following best describes what a go routine is?
- A separate line of code execution that can be used to handle blocking code.

2. What's the purpose of a channel?
- For communication between go routines.

3. Take a look at the following program. Are there any issues with it?
```
package main
 
import (
 "fmt"
)
 
func main() {
 greeting := "Hi There!"
 
 go (func() {
     fmt.Println(greeting) 
 })()
}
```

- The `greeting` variable is referenced from diretly in the go routine, which might lead to issues if we eventually start to change the value of greeting.
- The program will likely exit before the `fmt.Println` function has an opportunity to actually print anything out to the terminal; this might not be the intent of the program. 

4. Is there any issue with the following code?
```
package main
 
func main() {
 c := make(chan string)
 c <- []byte("Hi there!")
}
```

- Yes, the channel is expecting values of type string, but we are passing in a value of type byte slice, which is not technically a string.

5. Is there any issue with the following code?
```
package main
 
func main() {
     c := make(chan string)
     c <- "Hi there!"
}
```

- The syntax of this program is OK, but the program will never exit because it will wait for something to receive the value we're passing into the channel.


