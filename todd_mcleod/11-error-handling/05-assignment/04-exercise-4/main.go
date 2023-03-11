package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("more coffee needed")
		// e := fmt.Errorf("More coffee needed %v", f)
		return 0, sqrtError{"50.2289 N", "99.4565W", e}
	}
	return 42, nil
}
