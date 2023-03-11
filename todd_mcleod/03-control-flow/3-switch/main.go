package main

import "fmt"

func main() {
	// no switch expression: so the expression is default to Boolean true
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		// fallthrough Keyword: allows the switch statement to run all the way down
		fallthrough
	case (4 == 4):
		fmt.Println("also true")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		fallthrough
	default:
		fmt.Println("this is default")
	}

	n := "Leon"
	switch n {
	case "Jie Wei":
		fmt.Println("Not this.")
	case "Leon", "Stupid":   // can put multiple cases in one
		fmt.Println("This Leon")
	default:
		fmt.Println("Unknown")
	}
}