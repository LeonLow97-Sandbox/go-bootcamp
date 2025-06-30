package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func sumNumbers[T Number](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int32{1, 2, 3, 4, 5}
	numbers3 := []int64{1, 2, 3, 4, 5}
	numbers4 := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	numbers5 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(sumNumbers(numbers1))
	fmt.Println(sumNumbers(numbers2))
	fmt.Println(sumNumbers(numbers3))
	fmt.Println(sumNumbers(numbers4))
	fmt.Println(sumNumbers(numbers5))
}
