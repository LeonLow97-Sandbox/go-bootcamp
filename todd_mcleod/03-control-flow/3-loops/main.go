package main

import "fmt"

func main() {
	basicForLoop()
	nestedLoop()
	forStatement()
	continueKeyword()
	asciiLoop()
}

func asciiLoop() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}

func continueKeyword() {
	x := 83 / 40
	y := 83 % 40   // remainder (use modulus "%")
	fmt.Println(x, y)

	// Print Even Numbers
	x = 0
	for {
		x++
		if x > 20 {break}
		if x % 2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}

func forStatement() {
	x := 1
	// like a "while" statement
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done", x)

	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("Done2", y)
}

func nestedLoop() {
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
}

func basicForLoop() {
	// for init; condition; post {}
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}
}