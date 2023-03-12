package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	// ctx := context.Background()
	ctx := context.WithValue(context.Background(), "foo", "bar")

	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ", val)
	fmt.Println("Took: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	val := ctx.Value("foo")
	fmt.Println("val:", val)

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel() // cancel the context to prevent context leakage

	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		// cancels the channel after timeout has happened (in this case after 200 milliseconds)
		case <-ctx.Done():
			return 0, fmt.Errorf("Fetching data from third party took too long. Timeout.")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

// Third party APIs might be slow which creates latency in our application
func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 500)
	// time.Sleep(time.Millisecond * 150)
	return 888, nil
}
