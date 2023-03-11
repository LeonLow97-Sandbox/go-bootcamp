package main

import "fmt"

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	// exercise7()
	// exercise8()
	exercise9()
}

func exercise9() {
	// Create a program that uses a switch statement with the switch expression specified
	// as "Skiing"
	switch "Skiing" {
	case "Leon":
		fmt.Println("Leon")
	case "Skiing":
		fmt.Println("Skiing")
	default:
		fmt.Println("None of the above")
	}
}

func exercise8() {
	// Create a program that uses a switch statement with no switch expression specified
	// expression defaults to Boolean true
	switch {
	case (2 == 2) :
		fmt.Println("This is true")
	case false:
		fmt.Println("False")
	}
}

func exercise7() {
	// Create a program that uses "else if" and "else"
	n := 1
	if (n == 60) {
		fmt.Println("n is 60")
	} else if (n == 59) {
		fmt.Println("n is 59")
	} else {
		fmt.Println("n is not 59 or 60.")
	}
}

func exercise6() {
	// Create a program that shows the "if statement" in action
	n := 60
	if (n == 60) {
		fmt.Println("n is 60")
	}
}

func exercise5() {
	// Print out the remainder (modulus) which is found for each number
	// between 10 and 100 when it is divided by 4.
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v\t%v\n", i, i%4)
	} 
}

func exercise4() {
	// Create a for loop to print out the years you have been alive
	// Using for {}
	by := 1997
	for {
		if by > 2022 {
			break
		}
		fmt.Println(by)
		by++
	}
}

func exercise3() {
	// Create a for loop to print out the years you have been alive
	// Using for condition {}
	by := 1997   // by is birth year
	for by <= 2022 {
		fmt.Println(by)
		by++
	}
}

func exercise2() {
	// Print every rune code point of the uppercase alphabet 3 times.
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func exercise1() {
	// print every number from 1 to 100
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}