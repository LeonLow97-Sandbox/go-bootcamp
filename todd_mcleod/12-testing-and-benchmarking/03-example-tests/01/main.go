package main

import (
	"fmt"

	"github.com/LeonLow97/12-testing-and-benchmarking/03-example-tests/01/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
