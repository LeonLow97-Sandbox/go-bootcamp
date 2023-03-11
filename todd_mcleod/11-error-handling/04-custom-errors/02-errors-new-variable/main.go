package main

import (
	"errors"
	"fmt"
	"log"
)

var Err = errors.New("square root of negative number")

func main() {

	fmt.Printf("%T\n", Err)

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// // func New(text string) error)
		return 0, Err
	}
	return 42, nil
}
