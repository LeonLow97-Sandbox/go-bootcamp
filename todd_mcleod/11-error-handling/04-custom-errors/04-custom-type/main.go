package main

import (
	"fmt"
	"log"
)

type MathError struct {
	lat  string
	long string
	err  error
}

// error interface implements Error() method
func (n MathError) Error() string {
	return fmt.Sprintf("math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("math redux: square root of a negative number: %v", f)
		return 0, MathError{"50.2289 N", "99.4566 W", nme}
	}
	return 42, nil
}
