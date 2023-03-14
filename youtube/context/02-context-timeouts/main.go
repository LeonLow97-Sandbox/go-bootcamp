package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a new context
	// With a deadline of 100 milliseconds
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)

	// Make a request, that will call the google homepage
	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	res, err := client.Do(req)

	// If the request failed, log to STDOUT
	// Request failed: Get "http://google.com": context deadline exceeded
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	} else {
		cancel()
	}

	// Print the status code if the request succeeds
	// Response received, status code:  200
	fmt.Println("Response received, status code: ", res.StatusCode)
}
