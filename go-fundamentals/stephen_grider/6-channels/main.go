package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://epomaker.com",
	}

	// creating a channel
	c := make(chan string)

	for _, link := range links {
		// create a 'go' routine
		go checkLink(link, c) // pass channel 'c' into checkLink function
	}

	// Creating an infinite loop for repeating routines
	// dont pass time.sleep here because we want the main routine to always be awake
	for l := range c {
		// go checkLink(l, c)

		// using a function literal (like Python lambda)
		// receive as an argument (go routine now has a copy of 'l')
		go func(link string) {
			// pause for 5 seconds
			time.Sleep(5 * time.Second)
			checkLink(link, c) // main routine and child routine are always looking at the same variable, same location in memory. ('l' is defined in the outer scope.)
			// never reference the same variable in different routines.
			// main routine reference variable keeps changing, but child routine still looks at the copy of the same link.
		}(l) // function is invoked immediately (calling itself)
		// pass in 'l' to the function literal.
	}

	// // using a for loop for the channels
	// for i := 0; i< len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
}

// 2nd argument will be the channel
func checkLink(link string, c chan string) {
	// Make a GET request to the link (blocking code --> pass execution to another go routine)
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
